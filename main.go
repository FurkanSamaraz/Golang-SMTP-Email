package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	//Gönderen mail adresi.
	//Sender e-mail address.
	// !!! Gönderen mail adresinin "Daha az güvenli uygulama erişimi" ni açınız.
	from := "samaraz545454@gmail.com"

	//gonderen mail sifresi.
	//sender mail password.
	password := "Samaraz172754"

	//Alıcı mail adresi.
	//Recipient e-mail address.
	to := []string{"samaraz4545@gmail.com"}

	//SMTP ayarları.
	//SMTP settings.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Gönderilecek mesaj
	//Message to be sent.
	message := []byte("Test ediliyor...")

	//Kimlik doğrulama
	//Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	//E-mail Gönderme işlemi.
	//E-mail Sending process.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println("hata var")
		os.Exit(1)
	}

	fmt.Println("E-mail Gönderildi.")

}
