package homepage

import (
	"cms/internal/models/homepage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var Index gin.HandlerFunc = func(c *gin.Context) {
	var lang = c.Param("lang")
	if lang == "/" {
		lang = viper.GetString("translator.default")
	}
	translate, err := homepage.New(lang)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"traduction": translate,
	})
}
