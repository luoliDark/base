package httputil

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// request 请求参数转 Map
// list 值，转为逗号拼接都分割。
func RequestPramsToMapString(ctx *gin.Context) map[string]string {
	if ctx == nil || ctx.Request == nil {
		return make(map[string]string)
	}
	params := make(map[string]string, len(ctx.Request.PostForm))
	for key, val := range ctx.Request.PostForm {
		params[key] = strings.Join(val, ",")
	}
	return params
}
