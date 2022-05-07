package main

import (
	"spider/method"
	"spider/model"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)



func main() {
	
	go fetchNews()
	web()
}

func web(){
	gin.SetMode(gin.ReleaseMode)
	r:=gin.Default()
	r.GET("/AllNews",func(ctx *gin.Context) {
		ctx.JSON(200,model.GetAllBadNews())
	})
	r.GET("/News/:Category",func(ctx *gin.Context) {
		c:=ctx.Param("Category")
		ctx.JSON(200,model.GetSingleCategoryBadNews(c))
	})

	r.Run(":8080")
}


func fetchNews(){
	model.NewBaiduNewsSlice()
	model.NewBadBaiduNewsSlice()
	DoSyncTaskCronJob()
	
}

//定时1分钟判断新闻是否更新并抓取
func DoSyncTaskCronJob(){
	c := cron.New()
	spec:="00 * * * * ?"
	f:=method.FetchMethod{}
	c.AddFunc(spec,f.Fetch_activte)
	c.AddFunc(spec,f.Fetch_Discovery)
	c.AddFunc(spec,f.Fetch_Entertainment)
	c.AddFunc(spec,f.Fetch_Finance)
	c.AddFunc(spec,f.Fetch_Health)
	c.AddFunc(spec,f.Fetch_Internet)
	c.AddFunc(spec,f.Fetch_Lady)
	c.AddFunc(spec,f.Fetch_Military)
	c.AddFunc(spec,f.Fetch_World)
	c.AddFunc(spec,f.Fetch_Sports)
	c.AddFunc(spec,f.Fetch_Technology)
	c.Start()

}



