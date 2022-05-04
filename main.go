package main

import (
	"spider/method"
	"spider/model"
)

func main() {
	err:=fetch_AllNews()
	if err!=nil{
		
	}
}

func fetch_AllNews()error{
	fetchNews := make(model.BaiduNewsSlice,0,300)
	f:=method.FetchMethod{}
	f.Fetch_activte(fetchNews)
	f.Fetch_Discovery(fetchNews)
	f.Fetch_Entertainment(fetchNews)
	f.Fetch_Finance(fetchNews)
	f.Fetch_Health(fetchNews)
	f.Fetch_Internet(fetchNews)
	f.Fetch_Lady(fetchNews)
	f.Fetch_Military(fetchNews)
	f.Fetch_World(fetchNews)
	f.Fetch_Sports(fetchNews)
	f.Fetch_Technology(fetchNews)
	return nil
}
