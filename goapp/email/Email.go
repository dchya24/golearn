package email

import (
	"encoding/json"
	"log"
	"net/http"

	"dchya24/golearn/utils"

	"github.com/spf13/viper"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	viper.SetConfigName("../config.yaml")

	err := viper.ReadInConfig()

	if err != nil {
		response := utils.Response{
			Status:  "failed",
			Message: "Failed to send Email",
		}

		log.Println(err)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}

	var key string = viper.GetString("sendgrid_token")
	from := mail.NewEmail("No Reply", "no-reply@kasumi.co.jp") // Change to your verified sender
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Cahya Dinar Prastyo", "cahyadinar241@gmail.com") // Change to your recipient
	plainTextContent := "Hello World"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)

	_, err = client.Send(message)

	if err != nil {
		response := utils.Response{
			Status:  "failed",
			Message: "Failed to send Email",
		}

		log.Println(err)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	} else {
		response := utils.Response{
			Status:  "success",
			Message: "Success Send Email",
		}

		log.Println(err)
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
