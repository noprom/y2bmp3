package models

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "time"

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
