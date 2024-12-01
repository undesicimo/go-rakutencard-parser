package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ENDPOINT := os.Getenv("RAKUTEN_ENDPOINT")
	USERNAME := os.Getenv("RAKUTEN_USERNAME")
	PASSWORD := os.Getenv("RAKUTEN_PASSWORD")

	// Login to Rakuten Card
	browser := rod.New().MustConnect()
	defer browser.MustClose()
	RakutenPage := browser.MustPage(ENDPOINT).MustWaitStable()
	log.Println("Login Page loaded, Inputing User ID")
	RakutenPage.MustElement("input[id='user_id']").MustInput(USERNAME)

	log.Println("Inputting Password")
	RakutenPage.MustElementX("//*[@id='cta001']").MustClick()
	RakutenPage.MustElement("input[id='password_current']").MustInput(PASSWORD)
	RakutenPage.MustElementX("//*[@id='cta011']").MustClick()

	// We get redirected again to the login page
	RakutenPage.MustWaitStable().MustElement("input[id='password_current']").MustInput(PASSWORD)
	RakutenPage.MustElementX("//*[@id='cta011']").MustClick()
	log.Println("Logged in")

	// Get amount information
	RakutenPage.MustWaitStable().MustElement("select[id='cardChangeForm:cardtype']").MustSelect("楽天カード（Visa）")
	log.Println("Getting the amount")
	amount, e := RakutenPage.MustWaitStable().MustElementX("//*[@id='js-rd-billInfo-amount_show']/span").Text()
	if e != nil {
		log.Println(e)
		panic(e)
	}

	SendEmail(amount)
}

func SendEmail(amount string) {
	RESEND_API_KEY := os.Getenv("RESEND_API_KEY")
	EMAIL_RECIPIENT := os.Getenv("EMAIL_RECIPIENT")
	EMAIL_RECIPIENT_2 := os.Getenv("EMAIL_RECIPIENT_2")
	EMAIL_ORIGIN := os.Getenv("EMAIL_ORIGIN")

	client := resend.NewClient(RESEND_API_KEY)
	html := fmt.Sprintf("<b>¥%s</b>", amount)

	params := &resend.SendEmailRequest{
		From:    EMAIL_ORIGIN,
		To:      []string{EMAIL_RECIPIENT, EMAIL_RECIPIENT_2},
		Subject: "今のところ、これぐらい使ってるよ",
		Html:    html,
	}
	log.Println("Sending email...")
	_, err := client.Emails.Send(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent")
}
