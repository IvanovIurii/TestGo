package dao

import (
	"strconv"

	"github.com/IvanovYura/restApi/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const POSTS = "posts"

func (m *PostDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		panic(err)
	}
	db = session.DB(m.Database)
}

func (m *PostDao) DropDb() {
	err := db.DropDatabase()
	if err != nil {
		panic(err)
	}
}

func (m *PostDao) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := db.C(POSTS).Find(bson.M{}).All(&posts)
	return posts, err
}

func (m *PostDao) Save(post model.Post) error {
	var results []model.Post

	err := db.C(POSTS).Find(nil).All(&results)
	if err != nil {
		panic(err)
	}

	p := &post
	p.Id = strconv.Itoa(len(results) + 1)

	return db.C(POSTS).Insert(p)
}

func (m *PostDao) Delete(post model.Post) error {
	return db.C(POSTS).Remove(&post)
}

func (m *PostDao) Update(post model.Post) error {
	return db.C(POSTS).UpdateId(post.Id, &post)
}
