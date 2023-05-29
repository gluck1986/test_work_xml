package datasource

import (
	"gluck1986/test_work_xml/internal/datasource/repository"
	"gluck1986/test_work_xml/internal/infrastructure"
)

// DataSources data source dependency container
type DataSources struct {
	Repositories  *repository.Repositories
	ParserFactory ISdnParserFactory
}

// NewSources constructor
func NewSources(infDep *infrastructure.Infrastructure) *DataSources {

	return &DataSources{
		Repositories:  repository.NewRepositories(infDep.Db.Db, infDep.Log),
		ParserFactory: NewSdnXMLParserFactory(infDep.Config, infDep.Log),
	}
}
