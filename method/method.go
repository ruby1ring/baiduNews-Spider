package method

import (
	"fmt"
	"spider/model"

	"github.com/gocolly/colly/v2"
)
type FetchMethod struct{}

//热点要闻
func (f FetchMethod)Fetch_activte(N model.BaiduNewsSlice){
	c1:=colly.NewCollector()
	newsCategory:="pane-news"
	c1.OnHTML("div[id='pane-news']",func(e *colly.HTMLElement){
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
		
	})

	c1.Visit(model.ActiveURL)
}

//获取国际信息
func (f FetchMethod)Fetch_World(N model.BaiduNewsSlice){
	newsCategory:="world"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='guojie'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.Visit(model.InternetURL)

}
//获取娱乐新闻
func (f FetchMethod)Fetch_Entertainment(N model.BaiduNewsSlice){
	newsCategory:="Entertainment"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='yule'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='yule'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.EntertainmentURL)

}

//获取体育新闻
func (f FetchMethod)Fetch_Sports(N model.BaiduNewsSlice){
	newsCategory:="Sports"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='tiyu'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='tiyu'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.SportsURL)

}

//获取财经新闻
func (f FetchMethod)Fetch_Finance(N model.BaiduNewsSlice){
	newsCategory:="Finance"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='caijing'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.FinanceURL)

}

//获取科技新闻
func (f FetchMethod)Fetch_Technology(N model.BaiduNewsSlice){
	newsCategory:="Technology"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='col-tech'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='col-tech'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.TechnologyURL)

}

//获取军事新闻
func (f FetchMethod)Fetch_Military(N model.BaiduNewsSlice){
	newsCategory:="Military"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='junshi'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='col-tech'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.MilitaryURL)

}

//获取互联网新闻
func (f FetchMethod)Fetch_Internet(N model.BaiduNewsSlice){
	newsCategory:="Internet"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='hulianwang'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='hulianwang'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.InternetURL)

}

//获取探索新闻
func (f FetchMethod)Fetch_Discovery(N model.BaiduNewsSlice){
	newsCategory:="Discovery"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='col-discovery'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.DiscoveryURL)

}

//获取女人新闻
func (f FetchMethod)Fetch_Lady(N model.BaiduNewsSlice){
	newsCategory:="Lady"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='col-lady'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='col-lady'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.LadyURL)

}

//获取健康新闻
func (f FetchMethod)Fetch_Health(N model.BaiduNewsSlice){
	newsCategory:="Military"
	c1:=colly.NewCollector()
	c1.OnHTML("div[id='col-healthy'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})

	c1.OnHTML("div[id='col-healthy'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				fmt.Println(url,text)
				N.AddNews(newsCategory,url,text)
			})
		})
	})
	c1.Visit(model.HealthyURL)

}


