package main

import (
	"cms/internal/config"
	"cms/internal/models/auth"
	"cms/internal/privates/password_generator"
	"fmt"
	"github.com/manifoldco/promptui"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		println(err.Error())
		return
	}

	var args = os.Args
	if len(args) < 2 {
		println("Pas assez d'arguments: [Username]")
		return
	}

	fmt.Println("Etes vous sur de vouloir créer le compte pour", args[1])
	if !yesNo() {
		return
	}

	username := args[1]
	pwd := password_generator.GeneratePassword(12, 3, 3, 1)

	fmt.Println("Votre compte à bien été créer")
	fmt.Println("Utilisteur:", username)
	fmt.Println("Password:", pwd)

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		println(err.Error())
		return
	}

	var account = auth.UserModel{
		Username: username,
		Password: string(hash),
	}
	err = account.Create()
	if err != nil {
		println(err.Error())
		return
	}
}

func yesNo() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
