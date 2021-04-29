package main

import (
	"github.com/go-mysql-org/go-mysql/canal"

	"github.com/zizaimengzhongyue/go-demo/go-mysql/handler"
)

var user = []canal.EventHandler{}
var blog = []canal.EventHandler{}

type Bootstrap struct {
	canal.DummyEventHandler
}

func (this *Bootstrap) OnRow(e *canal.RowsEvent) error {
	switch e.Table.String() {
	case "test.user":
		return this.notifyUser(e)
	case "test.blog":
		return this.notifyBlog(e)
	}
	return nil
}

func (this *Bootstrap) notifyUser(e *canal.RowsEvent) error {
	for _, v := range user {
		v.OnRow(e)
	}
	return nil
}

func (this *Bootstrap) notifyBlog(e *canal.RowsEvent) error {
	for _, v := range blog {
		v.OnRow(e)
	}
	return nil
}

func init() {
	user = append(user, &handler.User{})
	blog = append(blog, &handler.Blog{})
}
