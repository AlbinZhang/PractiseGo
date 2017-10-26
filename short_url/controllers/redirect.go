/**
 * @author [zhangfeng]
 * @email [312024054@qq.com]
 * @create date 2017-10-26 09:07:40
 * @modify date 2017-10-26 09:07:40
 * @desc [description]
 */

package controllers

// RedirectController ...
type RedirectController struct {
	BaseController
}

// Get ...
func (rc *RedirectController) Get() {
	rc.Ctx.ResponseWriter.Header().Set("Location", "/api/url")
	//rc.Ctx.ResponseWriter.WriteHeader(302)
	rc.Redirect("/api/url", 302)
}
