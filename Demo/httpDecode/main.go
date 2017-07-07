package main

import (
	fm "fmt"
	"net/url"
	rt "runtime"
)

func main() {
	l, err := url.ParseQuery("tos=312024054%40qq.com&subject=%5BP0%5D%5BOK%5D%5BmonitorServer%5D%5B%5D%5Bcpu%E8%BF%87%E8%BD%BD+all%28%231%29+cpu.idle++100%3C100%5D%5BO1+2017-07-05+18%3A40%3A00%5D&content=OK%0D%0AP0%0D%0AEndpoint%3AmonitorServer%0D%0AMetric%3Acpu.idle%0D%0ATags%3A%0D%0Aall%28%231%29%3A+100%3C100%0D%0ANote%3Acpu%E8%BF%87%E8%BD%BD%0D%0AMax%3A3%2C+Current%3A1%0D%0ATimestamp%3A2017-07-05+18%3A40%3A00%0D%0Ahttp%3A%2F%2F127.0.0.1%3A8081%2Fportal%2Ftemplate%2Fview%2F1%0D%0A")
	fm.Println(l, err)

	l3, err3 := url.Parse("tos=312024054%40qq.com&subject=%5BP0%5D%5BOK%5D%5BmonitorServer%5D%5B%5D%5Bcpu%E8%BF%87%E8%BD%BD+all%28%231%29+cpu.idle++100%3C100%5D%5BO1+2017-07-05+18%3A40%3A00%5D&content=OK%0D%0AP0%0D%0AEndpoint%3AmonitorServer%0D%0AMetric%3Acpu.idle%0D%0ATags%3A%0D%0Aall%28%231%29%3A+100%3C100%0D%0ANote%3Acpu%E8%BF%87%E8%BD%BD%0D%0AMax%3A3%2C+Current%3A1%0D%0ATimestamp%3A2017-07-05+18%3A40%3A00%0D%0Ahttp%3A%2F%2F127.0.0.1%3A8081%2Fportal%2Ftemplate%2Fview%2F1%0D%0A")
	fm.Println(l3, err3)
	fm.Println(l3.Path)
	fm.Println(l3.RawQuery)
	fm.Println(l3.Query())
	fm.Println(l3.Query().Encode())

	fm.Println(l3.RequestURI())
	fm.Printf("Hello World! version : %s", rt.Version())
}
