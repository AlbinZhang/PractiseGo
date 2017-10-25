/**
 * @author [zhangfeng]
 * @email [312024054@qq.com]
 * @create date 2017-10-22 07:51:55
 * @modify date 2017-10-22 07:51:55
 * @desc [description]
 */

package main

import (
	"flag"

	"github.com/astaxie/beego"
)

var gPort = flag.String("port", "8080", "web listen port")

func main() {
	beego.Run(":" + *gPort)
}

/*
func main() {
	flag.Parse()

	router := mux.NewMuxHandler()

	router.Handle("/hello/golang/", &BaseHandler{})
	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))
	})

	log.Fatalln(http.ListenAndServe(":"+*gPort, router))
	//log.Fatalln(http.ListenAndServe(":"+*gPort, &BaseHandler{}))
}

// BaseHandler ...
type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a => ", req.URL.Query().Get("a"))

	resp.Write([]byte("hello golang"))
}
*/
