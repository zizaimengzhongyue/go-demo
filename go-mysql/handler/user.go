package handler

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
)

type User struct {
	canal.DummyEventHandler
}

func (this *User) OnRow(e *canal.RowsEvent) error {
	switch e.Action {
	case canal.InsertAction:
		fmt.Println("user 插入数据")
	case canal.UpdateAction:
		fmt.Println("user 更新数据")
	case canal.DeleteAction:
		fmt.Println("user 删除数据")
	}
	return nil
}
