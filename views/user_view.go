package views

import (
	"io"
	"log"
	"sqli/initializers"
)

func LoginRender(file io.Writer) {
	err := initializers.Template.ExecuteTemplate(file, "login.html", nil)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)
	}

}

func ChangePasswordRender(file io.Writer) {
	err := initializers.Template.ExecuteTemplate(file, "change_password.html", nil)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)
	}

}

func ForgotPasswordRender(file io.Writer, data ...any) {
	var dataObject any = nil
	if len(data) > 0{
		dataObject = data[0]
		log.Printf("dataObject: %+v\n", dataObject)
	}
	err := initializers.Template.ExecuteTemplate(file, "forgot_password.html", dataObject)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)
	}

}
