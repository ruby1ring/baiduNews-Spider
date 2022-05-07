package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BaiduNews struct{
	NewsCategory string  `json:"newsCategory,omitempty"`
	NewsContext  []Context `json:"newsContext,omitempty"`
}

type Context struct{
	Url          string  `json:"url,omitempty"`
	Text         string  `json:"text,omitempty"`
	Sorce        float64  `json:"result,float64"`
}

type NewsTextEmotion struct{
	Result float64  `json:"result,float64"`
}

var BaiduNewsSlice []BaiduNews
var BadBaiduNewsSlice []BaiduNews

func NewBaiduNewsSlice(){
	BaiduNewsSlice=make([]BaiduNews,11)
	BaiduNewsSlice[CategoryDict[PaneNews]].NewsCategory=PaneNews
	BaiduNewsSlice[CategoryDict[World]].NewsCategory=World
	BaiduNewsSlice[CategoryDict[Entertainment]].NewsCategory=Entertainment
	BaiduNewsSlice[CategoryDict[Sports]].NewsCategory=Sports
	BaiduNewsSlice[CategoryDict[Finance]].NewsCategory=Finance
	BaiduNewsSlice[CategoryDict[Technology]].NewsCategory=Technology
	BaiduNewsSlice[CategoryDict[Military]].NewsCategory=Military
	BaiduNewsSlice[CategoryDict[Internet]].NewsCategory=Internet
	BaiduNewsSlice[CategoryDict[Discovery]].NewsCategory=Discovery
	BaiduNewsSlice[CategoryDict[Lady]].NewsCategory=Lady
	BaiduNewsSlice[CategoryDict[Healthy]].NewsCategory=Healthy
	
}

func NewBadBaiduNewsSlice(){
	BadBaiduNewsSlice=make([]BaiduNews,11)
	BadBaiduNewsSlice[CategoryDict[PaneNews]].NewsCategory=PaneNews
	BadBaiduNewsSlice[CategoryDict[World]].NewsCategory=World
	BadBaiduNewsSlice[CategoryDict[Entertainment]].NewsCategory=Entertainment
	BadBaiduNewsSlice[CategoryDict[Sports]].NewsCategory=Sports
	BadBaiduNewsSlice[CategoryDict[Finance]].NewsCategory=Finance
	BadBaiduNewsSlice[CategoryDict[Technology]].NewsCategory=Technology
	BadBaiduNewsSlice[CategoryDict[Military]].NewsCategory=Military
	BadBaiduNewsSlice[CategoryDict[Internet]].NewsCategory=Internet
	BadBaiduNewsSlice[CategoryDict[Discovery]].NewsCategory=Discovery
	BadBaiduNewsSlice[CategoryDict[Lady]].NewsCategory=Lady
	BadBaiduNewsSlice[CategoryDict[Healthy]].NewsCategory=Healthy
	
}

func AddNews(newsCategory string, urls *[]string, texts *[]string) {
	//如果是重复新闻则不更新
	if len(BaiduNewsSlice[CategoryDict[newsCategory]].NewsContext)>0 && (BaiduNewsSlice[CategoryDict[newsCategory]].NewsContext[0].Url==(*urls)[0]){
		return
	}
	newContextSlice:=make([]Context,len(*urls))
	for i := 0; i < len(*urls); i++ {
		newContextSlice[i].Url=(*urls)[i]
		newContextSlice[i].Text=(*texts)[i]
	}
	BaiduNewsSlice[CategoryDict[newsCategory]].NewsContext=newContextSlice
	GetSocre(BaiduNewsSlice[CategoryDict[newsCategory]])
	addBadNews(newsCategory,BaiduNewsSlice[CategoryDict[newsCategory]])
} 

func GetSocre(baiduNews BaiduNews) {
	result:=NewsTextEmotion{}
	for i:=0;i<len(baiduNews.NewsContext);i++ {
		resp, err := http.Post(EmotionUrl,"application/x-www-form-urlencoded",strings.NewReader(baiduNews.NewsContext[i].Text))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body,err:=ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(body,&result)
		baiduNews.NewsContext[i].Sorce=result.Result
		log.Printf("func: GetSocre(), Text: %s,result: %f",baiduNews.NewsContext[i].Text,baiduNews.NewsContext[i].Sorce)
	}
}

func addBadNews(Category string,baiduNews BaiduNews){
	newSlice:=BaiduNews{NewsCategory: Category}
	for i:=0;i<len(baiduNews.NewsContext);i++{
		if (baiduNews.NewsContext[i].Sorce)<(-0.5){
			newContext:=Context{baiduNews.NewsContext[i].Url,baiduNews.NewsContext[i].Text,baiduNews.NewsContext[i].Sorce}
			newSlice.NewsContext=append(newSlice.NewsContext, newContext)
			log.Printf("func: addBadNews(), Url: %s,result: %f" ,baiduNews.NewsContext[i].Url,baiduNews.NewsContext[i].Sorce)
		}
	}
	BadBaiduNewsSlice[CategoryDict[Category]]=newSlice
}

func GetAllBadNews()[]BaiduNews{
	return BadBaiduNewsSlice
}

func GetSingleCategoryBadNews(Category string)BaiduNews{
	return BadBaiduNewsSlice[CategoryDict[Category]]
}





const (
	//百度新闻不同类别的URL
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
	//情感分析URL
	EmotionUrl ="http://baobianapi.pullword.com:9091/get.php"
)

const(
	PaneNews ="Pane-news"
	World = "World"
	Entertainment = "Entertainment"
	Sports ="Sports"
	Finance ="Finance"
	Technology ="Technology"
	Military ="Military"
	Internet ="Internet"
	Discovery ="Discovery"
	Lady ="Lady"
	Healthy ="Healthy"
	AllNews ="All"
)


var CategoryDict = map[string]int{   PaneNews: 0,  World: 1,  Entertainment: 2,  Sports : 3,  Finance: 4,  Technology  : 5,  Military  : 6,  Internet : 7,  Discovery  : 8,  Lady  : 9,  Healthy   : 10}


