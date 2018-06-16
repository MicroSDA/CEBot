package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
	"net/http"
	"io"
	"httpscerts"
	"fmt"
)


func startBot() {
	bot, err := tgbotapi.NewBotAPI("600819807:AAGSkqms9KCzDV0YFtyNY0utNolaxlHZ3j0")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	//log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://109.254.5.53:80/", "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/")


	go http.ListenAndServeTLS(":80", "cert.pem", "key.pem",nil)


	for update := range updates {
		log.Printf("%+v\n", update)
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,"Hello"))
	}
}


func serf(){

	httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:80")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
func main() {

	httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:80")
	bot, err := tgbotapi.NewBotAPI("600819807:AAGSkqms9KCzDV0YFtyNY0utNolaxlHZ3j0")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://109.254.5.53:80/", "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}


	updates := bot.ListenForWebhook("/")
	go http.ListenAndServeTLS(":80", "cert.pem", "key.pem", nil)


	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}