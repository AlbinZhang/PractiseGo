/**
 * @author [zhangfeng]
 * @email [312024054@qq.com]
 * @create date 2017-10-22 11:39:02
 * @modify date 2017-10-22 11:39:02
 * @desc [description]
 */

package mux

import (
	"net/http"
)

// MuxHandler ...
type MuxHandler struct {
	handlers map[string]http.Handler
}

// NewMuxHandler ...
func NewMuxHandler() *MuxHandler {
	return &MuxHandler{
		handlers: make(map[string]http.Handler),
	}
}

func (handler *MuxHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// 分发请求
	// 精确匹配
	urlPath := req.URL.Path
	if hl, ok := handler.handlers[urlPath]; ok {
		hl.ServeHTTP(resp, req)
		return
	}

	http.NotFound(resp, req)
}

// Handle ...
func (handler *MuxHandler) Handle(pattern string, hl http.Handler) {
	handler.handlers[pattern] = hl
}

// HandleFunc ...
func (handler *MuxHandler) HandleFunc(pattern string, fn func(resp http.ResponseWriter, req *http.Request)) {
	handler.Handle(pattern, http.HandlerFunc(fn))
}
