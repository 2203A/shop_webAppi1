package tools

import (
	"SX1/shop_webAppi/userserve/cmd/global"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// elasticsearch结构体
type Elasticsearch struct {
	Host string
	Port int
}

// 创建elasticsearch实例
func NewElasticsearchClient(c Elasticsearch) *Elasticsearch {
	return &Elasticsearch{
		Host: c.Host,
		Port: c.Port,
	}
}

// 类方法接口
type ElasticsearchInterface interface {
	GetData(id string, index string) *elastic.GetResult
	GetDataList(index string, page, size int) *elastic.SearchResult
	SyncElasticsearch(data interface{}, index string) *elastic.IndexResponse
}

// 获取数据
func (e *Elasticsearch) GetData(id string, index string) *elastic.GetResult {
	data, err := global.ElasticClient.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		fmt.Println("es get data error:", err)
		return nil
	}
	return data
}

// 获取列表
func (e *Elasticsearch) GetDataList(index string, page, size int) *elastic.SearchResult {
	list, err := global.ElasticClient.Search(index).Size(size).From((page - 1) * size).
		Highlight(elastic.NewHighlight().Field("name").PreTags("<em>").PostTags("<em/>")).Do(context.Background())
	if err != nil {
		fmt.Println("es get dataList err", err)
		return nil
	}
	return list
}

// 同步数据
func (e *Elasticsearch) SyncElasticsearch(data interface{}, index string) *elastic.IndexResponse {
	do, err := global.ElasticClient.Index().Index(index).BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println("elasticsearch sync err")
		return nil
	}
	return do
}
