package service

import (
	"context"
	"fmt"
	"gluck1986/test_work_xml/internal/data_source"
	"gluck1986/test_work_xml/internal/model"
	"log"
	"sync/atomic"
	"time"
)

const batchSize = 500

type SdnSyncroniser struct {
	isIdle atomic.Bool
	log    *log.Logger
	parser data_source.ISdnParser
	repo   data_source.ISdnRepository
}

type SdnSyncroniserParams struct {
	Log    *log.Logger
	Parser data_source.ISdnParser
	Repo   data_source.ISdnRepository
}

func NewSdnSyncroniser(p *SdnSyncroniserParams) ISdnSyncroniser {
	return &SdnSyncroniser{
		isIdle: atomic.Bool{},
		log:    p.Log,
		parser: p.Parser,
		repo:   p.Repo,
	}
}

// Syncronise run sycronise, it has background tasks
func (t *SdnSyncroniser) Syncronise(ctx context.Context) error {
	if t.IsIdle() {
		return fmt.Errorf("this instance already started")
	}
	t.isIdle.Store(true)
	parsed := make(chan model.SdnParseResponse)
	batchToWrite := make(chan []model.SdnEntity)
	done := make(chan bool)
	go t.parse(ctx, parsed)
	go t.filterMapReduce(parsed, batchToWrite)
	go t.write(batchToWrite, done)

	<-done

	return nil
}

func (t *SdnSyncroniser) mapExternalEntity(input model.SdnParseResponse) model.SdnEntity {
	return model.SdnEntity{
		UID:         input.Data.UID,
		FirstName:   input.Data.FirstName,
		LastName:    input.Data.LastName,
		PublishDate: t.parseDate(input.PublishInformation.PublishDate),
	}
}

func (t *SdnSyncroniser) parseDate(srcDate string) time.Time {
	parsedTime, err := time.Parse("01/02/2006", srcDate)
	if err != nil {
		t.log.Println("Cannot parse publish date, use now", "SdnSyncroniser", "parseDate", srcDate)
		parsedTime = time.Now()
	}
	return parsedTime
}

func (t *SdnSyncroniser) IsIdle() bool {
	return t.isIdle.Load()
}

func (t *SdnSyncroniser) shouldWrite(entity model.SdnParseResponse) bool {
	return true
}

func (t *SdnSyncroniser) parse(ctx context.Context, out chan<- model.SdnParseResponse) {
	defer close(out)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, ok := t.parser.Next()
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- res:
			}
		}
	}
}

func (t *SdnSyncroniser) filterMapReduce(input <-chan model.SdnParseResponse, out chan<- []model.SdnEntity) {
	defer close(out)
	batch := make([]model.SdnEntity, 0, batchSize)

	for parseResp := range input {
		if !t.shouldWrite(parseResp) {
			continue
		}
		entity := t.mapExternalEntity(parseResp)
		batch = append(batch, entity)
		if len(batch) >= batchSize {
			out <- batch
			batch = make([]model.SdnEntity, 0, batchSize)
		}

	}
	if len(batch) > 0 {
		out <- batch
	}
}

func (t *SdnSyncroniser) write(input <-chan []model.SdnEntity, done chan<- bool) {
	for batch := range input {
		err := t.repo.WriteMany(batch)
		if err != nil {
			t.log.Println("error: cannot write batch of sdn", err)
		}
	}
	t.isIdle.Store(false)
	done <- true
	close(done)
}
