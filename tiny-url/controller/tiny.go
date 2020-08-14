package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zizaimengzhongyue/go-demo/tiny-url/page"
)

const permanentRedirect = 301
const temporaryRedirect = 302

type TinyURL struct {
	OriginURL string `json:"origin_url"`
	URL       string `json:"url"`
}

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   TinyURL `json:"data"`
}

var host string = "127.0.0.1:8080"

func Compress(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bts, _ := ioutil.ReadAll(body)
	req := &Request{}
	err := json.Unmarshal(bts, req)
	if err != nil {
		fmt.Println(err)
	}

	kv := page.Store(req.URL)

	res := &Response{
		Status: 0,
		Msg:    "ok",
		Data: TinyURL{
			OriginURL: kv.Value,
			URL:       genetorShortURL(kv.Key),
		},
	}
	bts, _ = json.Marshal(res)
	fmt.Fprintln(w, string(bts))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	var err error
	var res Response
	params := parsePath(r.URL.Path)
	kv, err := page.Load(params)
	if err != nil {
		res = Response{
			Status: 1,
			Msg:    "params error",
		}
		bts, _ := json.Marshal(res)
		fmt.Fprintln(w, string(bts))
		return
	}
	http.Redirect(w, r, kv.Value, temporaryRedirect)
}

func genetorShortURL(params string) string {
	return host + "/" + params
}

func parsePath(path string) string {
	bts := []byte(path)
	if len(bts) == 0 {
		return ""
	}
	if bts[0] == '/' {
		return string(bts[1:])
	}
	return string(bts)
}
