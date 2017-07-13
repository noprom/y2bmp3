package main

import (
	_ "y2bmp3/routers"
	
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/astaxie/beego/session/redis"
)

var bm, err = cache.NewCache("redis", `{"key":"collectionName","conn":"redis:6379","dbNum":"0","password":""}`)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Debug("Gegin Run")
	bm.Put("astaxie", 1, 3600*time.Second)
	bm.Get("astaxie")
	bm.IsExist("astaxie")

	beego.SetLogger("file", `{"filename":"logs/run.log"}`)
	beego.Run()
}
