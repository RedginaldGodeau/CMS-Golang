package homepage

import (
	"cms/internal/models/homepage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var Get gin.HandlerFunc = func(c *gin.Context) {
	trads, err := homepage.All()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": trads,
	})
}

var GetByLang gin.HandlerFunc = func(c *gin.Context) {
	var lang = strings.Replace(c.Param("lang"), "/", "", -1)

	trad, err := homepage.New(lang)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": trad,
	})
}

var Post gin.HandlerFunc = func(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}

	var trad = homepage.HomepageModel{
		Title:                                      c.PostForm("Title"),
		Section_1_headline:                         c.PostForm("Section_1_headline"),
		Section_1_card_1_headline:                  c.PostForm("Section_1_card_1_headline"),
		Section_1_card_1_text:                      c.PostForm("Section_1_card_1_text"),
		Section_1_card_2_headline:                  c.PostForm("Section_1_card_2_headline"),
		Section_1_card_2_text:                      c.PostForm("Section_1_card_2_text"),
		Section_1_card_3_headline:                  c.PostForm("Section_1_card_3_headline"),
		Section_1_card_3_text:                      c.PostForm("Section_1_card_3_text"),
		Section_2_headline:                         c.PostForm("Section_2_headline"),
		Section_2_textblock_1:                      c.PostForm("Section_2_textblock_1"),
		Section_2_textblock_2:                      c.PostForm("Section_2_textblock_2"),
		Section_2_textblock_3:                      c.PostForm("Section_2_textblock_3"),
		Section_3_headline:                         c.PostForm("Section_3_headline"),
		Section_3_block_headline_backend:           c.PostForm("Section_3_block_headline_backend"),
		Section_3_block_text_backend:               c.PostForm("Section_3_block_text_backend"),
		Section_3_block_headline_frontend:          c.PostForm("Section_3_block_headline_frontend"),
		Section_3_block_text_frontend:              c.PostForm("Section_3_block_text_frontend"),
		Section_3_block_headline_creativedesign:    c.PostForm("Section_3_block_headline_creativedesign"),
		Section_3_block_text_creativedesign:        c.PostForm("Section_3_block_text_creativedesign"),
		Section_3_block_headline_server_management: c.PostForm("Section_3_block_headline_server_management"),
		Section_3_block_text_server_management:     c.PostForm("Section_3_block_text_server_management"),
		Section_4_headline:                         c.PostForm("Section_4_headline"),
		Section_4_textblock:                        c.PostForm("Section_4_textblock"),
		Footer_text_1:                              c.PostForm("Footer_text_1"),
		Footer_text_2:                              c.PostForm("Footer_text_2"),
	}

	err = trad.Create(c.PostForm("Lang"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Created",
	})
}

var Update gin.HandlerFunc = func(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}

	var trad = homepage.HomepageModel{
		Title:                                      c.PostForm("Title"),
		Section_1_headline:                         c.PostForm("Section_1_headline"),
		Section_1_card_1_headline:                  c.PostForm("Section_1_card_1_headline"),
		Section_1_card_1_text:                      c.PostForm("Section_1_card_1_text"),
		Section_1_card_2_headline:                  c.PostForm("Section_1_card_2_headline"),
		Section_1_card_2_text:                      c.PostForm("Section_1_card_2_text"),
		Section_1_card_3_headline:                  c.PostForm("Section_1_card_3_headline"),
		Section_1_card_3_text:                      c.PostForm("Section_1_card_3_text"),
		Section_2_headline:                         c.PostForm("Section_2_headline"),
		Section_2_textblock_1:                      c.PostForm("Section_2_textblock_1"),
		Section_2_textblock_2:                      c.PostForm("Section_2_textblock_2"),
		Section_2_textblock_3:                      c.PostForm("Section_2_textblock_3"),
		Section_3_headline:                         c.PostForm("Section_3_headline"),
		Section_3_block_headline_backend:           c.PostForm("Section_3_block_headline_backend"),
		Section_3_block_text_backend:               c.PostForm("Section_3_block_text_backend"),
		Section_3_block_headline_frontend:          c.PostForm("Section_3_block_headline_frontend"),
		Section_3_block_text_frontend:              c.PostForm("Section_3_block_text_frontend"),
		Section_3_block_headline_creativedesign:    c.PostForm("Section_3_block_headline_creativedesign"),
		Section_3_block_text_creativedesign:        c.PostForm("Section_3_block_text_creativedesign"),
		Section_3_block_headline_server_management: c.PostForm("Section_3_block_headline_server_management"),
		Section_3_block_text_server_management:     c.PostForm("Section_3_block_text_server_management"),
		Section_4_headline:                         c.PostForm("Section_4_headline"),
		Section_4_textblock:                        c.PostForm("Section_4_textblock"),
		Footer_text_1:                              c.PostForm("Footer_text_1"),
		Footer_text_2:                              c.PostForm("Footer_text_2"),
	}

	err = trad.Update(c.PostForm("lang"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success Updated",
	})
}

var Delete gin.HandlerFunc = func(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}
	lang := c.PostForm("lang")

	err = homepage.Delete(lang)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Les données sont incorrectes",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success Deleted",
	})

}
