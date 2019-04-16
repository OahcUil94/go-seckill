package route

import (
	"github.com/OahcUil94/go-seckill/controller"
	"github.com/astaxie/beego"
)

func init() {
	// 路由，处理的controller，处理的方法GET/POST，都支持的话，用method:handleFunc： post:handleFunc
	// SecKill是controller的方法
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}
