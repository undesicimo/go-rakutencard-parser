package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/resend/resend-go/v2"
)

var endpoint string = "https://login.account.rakuten.com/sso/authorize?client_id=rakuten_card_enavi_web&redirect_uri=https://www.rakuten-card.co.jp/e-navi/auth/login.xhtml&scope=openid%20profile&response_type=code&prompt=login#/sign_in"

func main() {
	// Login to Rakuten Card
	browser := rod.New().MustConnect()
	defer browser.MustClose()
	RakutenPage := browser.MustPage(endpoint).MustWaitStable()
	fmt.Println("Login Page loaded, Inputing User ID")
	RakutenPage.MustElement("input[id='user_id']").MustInput("dehaanchillax13")

	fmt.Println("going next")
	RakutenPage.MustElementX("//*[@id='cta001']").MustClick()
	fmt.Println("Password Inputted, going next")
	RakutenPage.MustElement("input[id='password_current']").MustInput("Vongola26//")
	fmt.Println("Inputting password again")
	RakutenPage.MustElementX("//*[@id='cta011']").MustClick()
	RakutenPage.MustWaitStable().MustElement("input[id='password_current']").MustInput("Vongola26//")
	RakutenPage.MustElementX("//*[@id='cta011']").MustClick()
	fmt.Println("Logged in")
	RakutenPage.MustWaitStable().MustElement("select[id='cardChangeForm:cardtype']").MustSelect("楽天カード（Visa）")
	fmt.Println("Getting the amount")
	amount, e := RakutenPage.MustWaitStable().MustElementX("//*[@id='js-rd-billInfo-amount_show']/span").Text()
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	fmt.Println(amount)
	SendEmail(amount)
}

func SendEmail(amount string) {
	apiKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	resend.NewClient(apiKey)
}
