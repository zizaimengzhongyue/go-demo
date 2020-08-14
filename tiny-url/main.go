package main

import (
	"net/http"

	"github.com/zizaimengzhongyue/go-demo/tiny-url/controller"
)

func main() {
	http.HandleFunc("/", controller.Redirect)
	http.HandleFunc("/compress", controller.Compress)
	http.ListenAndServe(":8080", nil)
}
