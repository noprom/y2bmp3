package controllers

import (
	"encoding/json"
	"time"
	"y2bmp3/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

type VideoController struct {
	BaseController
}

var bm, err = cache.NewCache("redis", `{"key":"collectionName","conn":"redis:6379","dbNum":"0","password":""}`)

// @Title Convert
// @Description Convert and Download Youtube MP3
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Video.Id
// @Failure 403 body is empty
func (c *VideoController) Convert() {
	id := c.GetString("v")
	beego.Debug("Convert video id: ", id)
	exist, _ := models.VideoExist(id)
	if !exist {
		beego.Debug("Video does not exist on Youtube.")
		c.ApiReturn(&ApiResult{400, "Video does not exist on Youtube.", nil})
	}
	// Get Video from Redis
	cacheKey := "v_" + id
	cacheExist := bm.IsExist(cacheKey)
	if cacheExist {
		result := bm.Get(cacheKey)
		if result == nil {
			beego.Error("Cache get nil")
		} else {
			cacheValue := string(result.([]uint8))
			beego.Debug("Cache hit: " + cacheValue)
			v := models.Video{}
			json.Unmarshal([]byte(cacheValue), &v)
			c.ApiReturn(&ApiResult{200, "success", v})
			// c.Data["json"] = v
			// c.ServeJSON()
			// return
		}
	} else {
		// Find the Data from Mongo
		video := models.Video{}
		if code, err := video.FindById(id); err != nil {
			beego.Error("FindVideoById: ", err)
			if code == models.ErrNotFound {
				// Download Video
				title, path, err := models.DownloadVideo(id)

				if err != nil {
					beego.Error("download error: ", err)
					c.Data["json"] = models.NewErrorInfo(ErrDownload)
					c.ServeJSON()
					return
				}

				v := models.Video{
					Id:         id,
					Title:      title,
					Path:       path,
					CreateTime: time.Now(),
				}
				beego.Debug("Download video info: ", &v)
				// Insert Into MongoDB
				if code, err := v.Insert(); err != nil {
					beego.Error("InsertVideo:", err)
					if code == models.ErrDupRows {
						c.Data["json"] = models.NewErrorInfo(ErrDupUser)
					} else {
						c.Data["json"] = models.NewErrorInfo(ErrDatabase)
					}
					c.ServeJSON()
					return
				}
				c.Data["json"] = v
				if b, err := json.Marshal(v); err == nil {
					beego.Debug("Cache to redis: ", string(b))
					bm.Put(cacheKey, string(b), 24*365*time.Hour)
				}
			} else {
				c.Data["json"] = video
			}
			c.ServeJSON()
			return
		}

		if b, err := json.Marshal(video); err == nil {
			beego.Debug("Cache to redis: ", string(b))
			bm.Put(cacheKey, string(b), 24*365*time.Hour)
		}
		beego.Debug("Read from Mongo, Video Info: ", &video)
		c.Data["json"] = video
		c.ServeJSON()
	}
}
