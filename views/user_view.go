package views

import (
	"io"
	"log"
	"sqli/initializers"
)

func LoginRender(file io.Writer, errorMsgs ...error) {
	var errorMsg string
	if len(errorMsgs) > 0 {
		errorMsg = errorMsgs[0].Error()
	}
	err := initializers.Template.ExecuteTemplate(file, "login.html", errorMsg)
	if err != nil {
		log.Printf("error occured while executing template, %v\n", err)
	}

}

func ChangePasswordRender(file io.Writer, hasErrorMsg bool) {
	err := initializers.Template.ExecuteTemplate(file, "change_password.html", hasErrorMsg)
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
