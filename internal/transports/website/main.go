package website

import (
	"cms/internal/transports/website/homepage"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	r.GET("/*lang", homepage.Index)
}
