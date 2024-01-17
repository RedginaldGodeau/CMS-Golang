package contactform

import (
	"cms/internal/models/contactform"
	"cms/internal/services/mailer"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var Post gin.HandlerFunc = func(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		return
	}

	var data = contactform.ContactFormModel{
		Username: c.PostForm("Username"),
		Email:    c.PostForm("Email"),
		Message:  c.PostForm("Message"),
		ClientIp: c.ClientIP(),
		SendAt:   time.Now().Unix(),
	}

	if len(data.Username) < 2 || len(data.Email) < 3 || len(data.Message) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Les données entrées ne sont pas correcte"})
		return
	}

	err = data.Create()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Les services sont actuelement en panne"})
		return
	}

	var builded = mailer.Mail{
		From:     viper.GetString("server.mailer"),
		To:       "redginaldgodeau@proton.me",
		Cc:       "",
		Subject:  "Message de " + data.Username,
		Template: "contact.html",
		Data:     data,
	}

	err = builded.SendMail()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Une erreur est survenue lors de l'envoi du mail"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Votre message à bien été envoyé"})
}
