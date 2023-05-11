package httpmiddleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

//跨域处理
func CrosMiddleware(Ctx *gin.Context) {
	origin := Ctx.Request.Header.Get("Origin")
	Ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	Ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	Ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Cookie, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie")
	Ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if Ctx.Request.Method == "OPTIONS" {
		Ctx.Writer.Header().Set("Access-Control-Max-Age", "15552000") //缓存半年
		//Ctx.Header("Cache-Control", "private, max-age=15552000")      //缓存半年   李大神说这句千万不能用
		Ctx.JSON(200, "Options Request!")
	} else if strings.Index(Ctx.Request.RequestURI, "/base/static") >= 0 {
		//静态资源
		//Ctx.Header("Cache-Control", "private, max-age=15552000") //缓存半年
	}

	Ctx.Next()

}

//登录拦截检查 注：api和定时任务rest不检查
func CheckLoginMiddleware(Ctx *gin.Context) {

	//临时不检查cookie
	Ctx.Next()

	//获取cookie中token
	//sidCookie, err := Ctx.Request.Cookie("sid")
	//if err != nil || sidCookie == nil {
	//	Ctx.Error(errors.New("cookie获取sid失败"))
	//	Ctx.Abort()
	//}
	//
	//token := sidCookie.Value
	////从redis获取用户信息
	//user, err := sso.GetUserByToken(token)
	//if err != nil {
	//	Ctx.Error(err)
	//	//	Ctx.Abort()
	//} else {
	//	//将token转为userid 写入head
	//	Ctx.Writer.Header().Set("userid", user.UserID)
	//	Ctx.Next()
	//}

}
