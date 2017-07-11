package models

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"fmt"
	// "github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
	"y2bmp3/models/mymongo"
)

type Video struct {
	// https://www.youtube.com/watch?v=XKqWnOtbSr8
	Id    string `bson:"_id"			json:"_id,omitempty"`
	Title string `bson:"title"			json:"name,omitempty"`
	Path  string `bson:"path"			json:"path,omitempty"`
	// CreateTime time.Time `bson:"create_time"	json:"create_time,omitempty"`
}

// Insert insert a document to collection.
func (v *Video) Insert() (code int, err error) {
	mConn := mymongo.Conn()
	defer mConn.Clone()

	c := mConn.DB("").C("videos")
	err = c.Insert(v)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}

	return
}

// FindByID query a document according to input id.
func (v *Video) FindById(id string) (code int, err error) {
	mConn := mymongo.Conn()
	defer mConn.Clone()

	c := mConn.DB("").C("videos")
	err = c.FindId(id).One(v)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = ErrNotFound
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}

	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Download from Youtube, and convert it to MP3.
func Download(id string) (title string, path string, err error) {
	path = fmt.Sprintf("/app/data/videos/%s", time.Now().Format("200601/02"))
	pathExists, _ := PathExists(path)
	if !pathExists {
		// fmt.Printf("%s not exist.\n", path)
		mkdirErr := os.MkdirAll(path, 0777)
		if mkdirErr != nil {
			fmt.Println("mkdir Err :" + mkdirErr.Error())
		}
	}

	cmdStr := fmt.Sprintf("cd %s && youtube-dl -c --no-warnings -x --audio-format mp3 https://www.youtube.com/watch?v=%s | grep .mp3", path, id)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return nil, nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return nil, nil, err
	}

	// Handle Stdout
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return nil, nil, err
	}
	s := strings.Split(string(bytes), ": ")
	_, mp3 := s[0], s[1]
	title = mp3
	return
}
