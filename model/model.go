package model

type BaiduNews struct {
	NewsCategory string
	Url          string
	Text         string
	Sorce        float64
}
type BaiduNewsSlice []BaiduNews

func (N BaiduNewsSlice) AddNews(newsCategory string, urls []string, texts []string) {
	for i := 0; i < len(urls); i++ {
		new := BaiduNews{NewsCategory: newsCategory, Url: urls[i], Text: texts[i]}
		N = append(N, new)
	}
}

const (
	ActiveURL        = "https://news.baidu.com/"
	WorldURL         = "https://news.baidu.com/widget?id=InternationalNews"
	EntertainmentURL = "https://news.baidu.com/widget?id=EnterNews"
	SportsURL        = "https://news.baidu.com/widget?id=SportNews"
	FinanceURL       = "https://news.baidu.com/widget?id=FinanceNews"
	TechnologyURL    = "https://news.baidu.com/widget?id=TechNews"
	MilitaryURL      = "https://news.baidu.com/widget?id=MilitaryNews"
	InternetURL      = "https://news.baidu.com/widget?id=InternetNews"
	DiscoveryURL     = "https://news.baidu.com/widget?id=DiscoveryNews"
	LadyURL          = "https://news.baidu.com/widget?id=LadyNews"
	HealthyURL       = "https://news.baidu.com/widget?id=HealthNews"
)
