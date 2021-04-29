package handler

import (
	"fmt"

	"github.com/go-mysql-org/go-mysql/canal"
)

type Blog struct {
	canal.DummyEventHandler
}

func (this *Blog) OnRow(e *canal.RowsEvent) error {
	switch e.Action {
	case canal.InsertAction:
		fmt.Println("blog 插入数据")
	case canal.UpdateAction:
		fmt.Println("blog 更新数据")
	case canal.DeleteAction:
		fmt.Println("blog 删除数据")
	}
	return nil
}
