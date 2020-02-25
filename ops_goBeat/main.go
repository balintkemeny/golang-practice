package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

func readDotEnv(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		//Handle error
	}

	keyValue, ok := viper.Get(key).(string)

	if !ok {
		//Handle error
	}

	return keyValue
}

func pingEndpoint(endpoint string) (int, error) {
	resp, err := http.Get(endpoint)
	return resp.StatusCode, err
}

func sendErrorMail(endpoint string) {
	from := mail.NewEmail(readDotEnv("FROM_NAME"), readDotEnv("FROM_ADDR"))
	to := mail.NewEmail(readDotEnv("TO_NAME"), readDotEnv("TO_ADDR"))
	subject := "Server error message"
	plainTextContent := "Warning, your server endpoint at " + endpoint + " is unreachable."

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, plainTextContent)
	client := sendgrid.NewSendClient(readDotEnv("SENDGRID_KEY"))

	resp, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		//Handle output
		log.Println(resp)
	}
}

func main() {
	endpoint := readDotEnv("ENDPOINT")
	statusCode, err := pingEndpoint(endpoint)

	if err != nil {
		sendErrorMail(endpoint)
	} else {
		fmt.Println(statusCode)
	}
}
