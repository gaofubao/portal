package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/mdp/qrterminal/v3"
	"github.com/wechaty/go-wechaty/wechaty"
	filebox "github.com/wechaty/go-wechaty/wechaty-puppet/file-box"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

func main() {
	var bot = wechaty.NewWechaty()

	bot.OnScan(onScan).OnLogin(func(ctx *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnMessage(onMessage).OnLogout(func(ctx *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	bot.DaemonStart()
}

func onMessage(ctx *wechaty.Context, message *user.Message) {
	log.Println(message)

	if message.Self() {
		log.Println("Message discarded because its outgoing")
	}

	if message.Age() > 2*60*time.Second {
		log.Println("Message discarded because its TOO OLD(than 2 minutes)")
	}

	if message.Type() != schemas.MessageTypeText || message.Text() != "#ding" {
		log.Println("Message discarded because it does not match 'ding'")
		return
	}

	// 1. reply 'dong'
	_, err := message.Say("dong")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("REPLY: dong")

	// 2. reply image(qrcode image)
	fileBox := filebox.FromUrl("https://wechaty.github.io/wechaty/images/bot-qr-code.png")
	_, err = message.Say(fileBox)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("REPLY: %s\n", fileBox)
}

func onScan(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
	if status == schemas.ScanStatusWaiting || status == schemas.ScanStatusTimeout {
		qrterminal.GenerateHalfBlock(qrCode, qrterminal.L, os.Stdout)

		qrcodeImageUrl := fmt.Sprintf("https://wechaty.js.org/qrcode/%s", url.QueryEscape(qrCode))
		fmt.Printf("onScan: %s - %s\n", status, qrcodeImageUrl)
		return
	}
	fmt.Printf("onScan: %s\n", status)
}
