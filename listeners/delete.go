package listeners

import (
	"context"

	configs "github.com/crowdeco/skeleton/configs"
	events "github.com/crowdeco/skeleton/events"
	handlers "github.com/crowdeco/skeleton/handlers"
	elastic "github.com/olivere/elastic/v7"
)

type Delete struct {
	Context       context.Context
	Elasticsearch *elastic.Client
}

func (d *Delete) Handle(event interface{}) {
	e := event.(*events.ModelEvent)

	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("id", e.Id()))
	result, _ := d.Elasticsearch.Search().Index(e.Service()).Query(query).Do(d.Context)
	for _, hit := range result.Hits.Hits {
		d.Elasticsearch.Delete().Index(e.Service()).Id(hit.Id).Do(d.Context)
	}
}

func (d *Delete) Listen() string {
	return handlers.AFTER_DELETE_EVENT
}

func (d *Delete) Priority() int {
	return configs.HIGEST_PRIORITY + 1
}
