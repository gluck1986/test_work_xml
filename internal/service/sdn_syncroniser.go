package service

import (
	"context"
	"gluck1986/test_work_xml/internal/datasource"
	"gluck1986/test_work_xml/internal/model"
	"log"
	"sync/atomic"
	"time"
)

const batchSize = 500

// SdnSyncroniser implementation of service that parse from external source sdn data and store to writer
type SdnSyncroniser struct {
	isIdle atomic.Bool
	log    *log.Logger
	parser datasource.ISdnParser
	writer ISdnWriter
	ctx    context.Context
	cache  datasource.IUIDCache
}

// SdnSyncroniserParams dependency
type SdnSyncroniserParams struct {
	Log    *log.Logger
	Writer ISdnWriter
	Cache  datasource.IUIDCache
}

// NewSdnSyncroniser constructor
func NewSdnSyncroniser(p *SdnSyncroniserParams) ISdnSyncroniser {
	return &SdnSyncroniser{
		isIdle: atomic.Bool{},
		log:    p.Log,
		writer: p.Writer,
		cache:  p.Cache,
	}
}

// Init set context and parser, use it before every run Syncronise
func (t *SdnSyncroniser) Init(ctx context.Context, parser datasource.ISdnParser) error {
	if t.IsIdle() {
		return ErrorSyncroniserAlreadyInProgress
	}
	t.parser = parser
	t.ctx = ctx
	return nil
}

// Syncronise run sycronise, it has background tasks
func (t *SdnSyncroniser) Syncronise() error {
	if t.IsIdle() {
		return ErrorSyncroniserAlreadyInProgress
	}
	t.isIdle.Store(true)
	parsed := make(chan model.SdnParseResponse)
	batchToWrite := make(chan []model.SdnEntity)
	done := make(chan bool)
	go t.parse(t.ctx, parsed)
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

// IsIdle returns true if the service has active background tasks
func (t *SdnSyncroniser) IsIdle() bool {
	return t.isIdle.Load()
}

func (t *SdnSyncroniser) shouldWrite(entity model.SdnParseResponse) bool {
	defer t.cache.Add(entity.Data.UID)
	return !t.cache.Has(entity.Data.UID)
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
		err := t.writer.WriteMany(batch)
		if err != nil {
			t.log.Println("error: cannot write batch of sdn", err)
		}
	}
	t.isIdle.Store(false)
	done <- true
	close(done)
}
