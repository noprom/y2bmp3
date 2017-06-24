package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"y2bmp3/models"

	"github.com/astaxie/beego"
)

type VideoController struct {
	BaseController
}

// @Title Convert
// @Description Convert and Download Youtube MP3
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.Video.Id
// @Failure 403 body is empty
func (c *VideoController) Convert() {
	id := c.GetString("id")
	c.Data["json"] = models.NewNormalInfo(id)
}

// Register method.
// func (c *UserController) Register() {
// 	form := models.RegisterForm{}
// 	if err := c.ParseForm(&form); err != nil {
// 		beego.Debug("ParseRegsiterForm:", err)
// 		c.Data["json"] = models.NewErrorInfo(ErrInputData)
// 		c.ServeJSON()
// 		return
// 	}
// 	beego.Debug("ParseRegsiterForm:", &form)

// 	if err := c.VerifyForm(&form); err != nil {
// 		beego.Debug("ValidRegsiterForm:", err)
// 		c.Data["json"] = models.NewErrorInfo(ErrInputData)
// 		c.ServeJSON()
// 		return
// 	}

// 	regDate := time.Now()
// 	user, err := models.NewUser(&form, regDate)
// 	if err != nil {
// 		beego.Error("NewUser:", err)
// 		c.Data["json"] = models.NewErrorInfo(ErrSystem)
// 		c.ServeJSON()
// 		return
// 	}
// 	beego.Debug("NewUser:", user)

// 	if code, err := user.Insert(); err != nil {
// 		beego.Error("InsertUser:", err)
// 		if code == models.ErrDupRows {
// 			c.Data["json"] = models.NewErrorInfo(ErrDupUser)
// 		} else {
// 			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
// 		}
// 		c.ServeJSON()
// 		return
// 	}

// 	go models.IncTotalUserCount(regDate)

// 	c.Data["json"] = models.NewNormalInfo("Succes")
// 	c.ServeJSON()
// }
