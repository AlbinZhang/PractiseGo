/**
 * @author [zhangfeng]
 * @email [312024054@qq.com]
 * @create date 2017-10-26 09:00:34
 * @modify date 2017-10-26 09:00:34
 * @desc [description]
 */
package controllers

// URLController ...
type URLController struct {
	BaseController
}

// Get ...
func (c *URLController) Get() {
	//c.Ctx.WriteString("This is UrlController:Get")
	c.success("This url sucessful method.")
}
