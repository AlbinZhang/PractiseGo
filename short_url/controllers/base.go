/**
 * @author [zhangfeng]
 * @email [312024054@qq.com]
 * @create date 2017-10-26 08:54:20
 * @modify date 2017-10-26 08:54:20
 * @desc [description]
 */

package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

func (bc *BaseController) optput(res string) {
	bc.Ctx.WriteString(res)
}

func (bc *BaseController) success(res interface{}) {
	bc.Data["json"] = map[string]interface{}{
		"sucessful": true,
		"result":    res,
	}

	bc.ServeJSON()
}

func (bc *BaseController) error(errorMsg string) {
	bc.Data["json"] = map[string]interface{}{
		"sucessful": false,
		"message":   errorMsg,
	}
}
