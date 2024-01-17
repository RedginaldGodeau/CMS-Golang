package api

import (
	"cms/internal/transports/api/homepage"
	"cms/internal/transports/api/mail/contactform"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {

	r.GET("/homepage", homepage.Get)
	r.GET("/homepage/*lang", homepage.GetByLang)
	r.POST("/homepage", homepage.Post)
	r.PATCH("/homepage", homepage.Update)
	r.DELETE("/homepage", homepage.Delete)

	r.POST("/mail/contact", contactform.Post)
}
