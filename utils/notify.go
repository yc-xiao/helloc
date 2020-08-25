package utils

/*
　短信跟邮件
*/

import (
	"Helloc/confs"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandNum() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(8999) + 1000
	return strconv.Itoa(n)
}

func SendPhoneCode(phoneNumber, code string) bool {
	return true
	key, secret := confs.Cfg["SMS_KEY"], confs.Cfg["SMS_SECRET"]
	signName, TemplateCode := confs.Cfg["SIGN_NAME"], confs.Cfg["TEMPLATE_CODE"]
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", key, secret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumber
	request.SignName = signName
	request.TemplateCode = TemplateCode
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

	response, err := client.SendSms(request)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if response.Code == "OK" {
		return true
	}else{
		fmt.Printf("response is %#v\n", response)
		log.Println(response.Message)
		return false
	}
	return true
}

func DefaultSendEmailCode(email, code string) bool {
	receiver := []string{email}
	subject := "Helloc-邮件通知"
	body := "邮箱验证码:" + code
	if err := SendMail(receiver, subject, body); err != nil {
		return false
	}
	return true
}

func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "xxx@sample.cn",
		"pass": "r4r3St*****7a7Uk",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "XX官方")) //这种方式可以添加别名，即“XX官方”
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)

	if err != nil {
		log.Println("邮件发送失败,", err)
	}
	return err
}