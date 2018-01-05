package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"metalApi/logger"
)

type PageProcesser struct {
}

type DelayMarket struct {
	MetalType   string `json:"metaltype"`
	Feature     int    `json:"feature"`
	NewestPrice int    `json:"newestprice"`
	PriceChange int    `json:"pricechage"`
	Inventory   int    `json:"inventory"`
}

func NewPageProcesser() *PageProcesser {
	return &PageProcesser{}
}

func (this *PageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		logger.GlobalLogReporter.Debug(p.Errormsg())
		return
	}
	var url = []string{"cu", "al", "zn", "pb", "ni", "sn", "au", "ag"}

	query := p.GetHtmlParser()
	var urls []string
	for i := 0; i < len(url); i++ {
		a := fmt.Sprintf("http://www.shfe.com.cn/statements/delaymarket_%s.html", url[i])
		urls = append(urls, a)
	}

	p.AddTargetRequests(urls, "html")
	name := query.Find("table tbody tr td").Eq(2).Text()

	p.AddField("name", name)
}

func (this *PageProcesser) Finish() {}

func main() {
	sp := spider.NewSpider(NewPageProcesser(), "Task")

	req := request.NewRequest("http://www.shfe.com.cn/statements/delaymarket_cu.html", "html", "", "GET", "", nil, nil, nil, nil)
	pageItem := sp.GetByRequest(req)
	url := pageItem.GetRequest().GetUrl()
	// println("-----------------------------------spider.Get---------------------------------")
	// println("url\t:\t" + url)
	// for name, value := range pageItem.GetAll() {
	// 	println(name + "\t:\t" + value)
	// }
	println("\n--------------------------------spider.GetAll---------------------------------")
	urls := []string{
		"http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn",
		"http://baike.baidu.com/view/383720.htm?fromtitle=html&fromid=97049&type=syn",
	}
	var reqs []*request.Request
	for _, url := range urls {
		req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
		reqs = append(reqs, req)
	}
	pageItemsArr := sp.SetThreadnum(2).GetAllByRequest(reqs)
	//pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
	for _, item := range pageItemsArr {
		url = item.GetRequest().GetUrl()
		println("url\t:\t" + url)
		fmt.Printf("item\t:\t%s\n", item.GetAll())
	}
}
