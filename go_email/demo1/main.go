package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	// 以163邮箱为例子
	m := gomail.NewMessage()
	m.SetHeader("From", "******@163.com")
	m.SetHeader("To", "******@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")
	d := gomail.NewDialer("smtp.163.com", 465, "******@163.com", "授权码")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
