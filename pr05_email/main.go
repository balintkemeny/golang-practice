package main

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
)

func getEnvVars(keyName string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("An error has occurred while reading in config")
		return ""
	}

	value, ok := viper.Get(keyName).(string)

	if !ok {
		fmt.Println("An error has occurred while reading specified key")
		return ""
	}

	return value
}

func main() {
	envFrom := getEnvVars("FROM")
	envTo := getEnvVars("TO")
	envKey := getEnvVars("SENDGRID_KEY")

	from := mail.NewEmail("Blob", envFrom)
	to := mail.NewEmail("Bálint Kemény", envTo)
	subject := "I hope this mail sender works"
	plainTextContent := "I've sent this mail using Sendgrid and Go."

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, plainTextContent)
	client := sendgrid.NewSendClient(envKey)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
