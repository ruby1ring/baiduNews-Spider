package method

import (
	"log"
	"spider/model"

	"github.com/gocolly/colly/v2"
)
type FetchMethod struct{}

//热点要闻
func (f FetchMethod)Fetch_activte(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='pane-news']",func(e *colly.HTMLElement){
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.PaneNews,url,text)

			})
		})
		model.AddNews(model.PaneNews,&urlSlices,&textSlices)
		
	})

	c1.Visit(model.ActiveURL)
}

//获取国际信息
func (f FetchMethod)Fetch_World(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='guojie'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.World,url,text)
			})
		})
		model.AddNews(model.World,&urlSlices,&textSlices)
	})

	c1.Visit(model.WorldURL)

}
//获取娱乐新闻
func (f FetchMethod)Fetch_Entertainment(){

	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='yule'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Entertainment,url,text)
			})
		})
		model.AddNews(model.Entertainment,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='yule'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Entertainment,url,text)
			})
		})
		model.AddNews(model.Entertainment,&urlSlices,&textSlices)
	})
	c1.Visit(model.EntertainmentURL)

}

//获取体育新闻
func (f FetchMethod)Fetch_Sports(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='tiyu'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Sports,url,text)
			})
		})
		model.AddNews(model.Sports,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='tiyu'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Sports,url,text)
			})
		})
		model.AddNews(model.Sports,&urlSlices,&textSlices)
	})
	c1.Visit(model.SportsURL)

}

//获取财经新闻
func (f FetchMethod)Fetch_Finance(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='caijing'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Finance,url,text)
			})
		})
		model.AddNews(model.Finance,&urlSlices,&textSlices)
	})
	c1.Visit(model.FinanceURL)

}

//获取科技新闻
func (f FetchMethod)Fetch_Technology(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='col-tech'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Technology,url,text)
			})
		})
		model.AddNews(model.Sports,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='col-tech'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Technology,url,text)
			})
		})
		model.AddNews(model.Technology,&urlSlices,&textSlices)
	})
	c1.Visit(model.TechnologyURL)

}

//获取军事新闻
func (f FetchMethod)Fetch_Military(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='junshi'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Military,url,text)
			})
		})
		model.AddNews(model.Military,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='col-tech'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Military,url,text)
			})
		})
		model.AddNews(model.Military,&urlSlices,&textSlices)
	})
	c1.Visit(model.MilitaryURL)

}

//获取互联网新闻
func (f FetchMethod)Fetch_Internet(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='hulianwang'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Internet,url,text)
			})
		})
		model.AddNews(model.Internet,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='hulianwang'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Internet,url,text)
			})
		})
		model.AddNews(model.Internet,&urlSlices,&textSlices)
	})
	c1.Visit(model.InternetURL)

}

//获取探索新闻
func (f FetchMethod)Fetch_Discovery(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='col-discovery'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Lady,url,text)
			})
		})
		model.AddNews(model.Lady,&urlSlices,&textSlices)
	})
	c1.Visit(model.DiscoveryURL)

}

//获取女人新闻
func (f FetchMethod)Fetch_Lady(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='col-lady'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Lady,url,text)
			})
		})
		model.AddNews(model.Lady,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='col-lady'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Lady,url,text)
			})
		})
		model.AddNews(model.Lady,&urlSlices,&textSlices)
	})
	c1.Visit(model.LadyURL)

}

//获取健康新闻
func (f FetchMethod)Fetch_Health(){
	c1:=colly.NewCollector()
	urlSlices:=make([]string,0,30)
	textSlices:=make([]string,0,30)
	c1.OnHTML("div[id='col-healthy'] > div[class='l-left-col col-mod']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Healthy,url,text)
			})
		})
		model.AddNews(model.Healthy,&urlSlices,&textSlices)
	})

	c1.OnHTML("div[id='col-healthy'] > div[class='l-right-col']",func(e *colly.HTMLElement) {
		e.ForEach("ul",func(i int, item *colly.HTMLElement) {
			item.ForEach("li",func(i int, h *colly.HTMLElement) {
				url:=h.ChildAttrs("a[target='_blank']","href")
				text:=h.ChildTexts("a[target='_blank']")
				urlSlices = append(urlSlices, url...)
				textSlices=append(textSlices, text...)
				log.Printf("Category: %s,Url: %s ,TextContext: %s",model.Healthy,url,text)
			})
		})
		model.AddNews(model.Healthy,&urlSlices,&textSlices)
	})
	c1.Visit(model.HealthyURL)

}


