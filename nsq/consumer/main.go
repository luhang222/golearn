package main

import (
	"crypto/tls"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

type MsgHandler struct{}

var count int64

func (*MsgHandler) HandleMessage(msg *nsq.Message) error {
	count++
	fmt.Println(count)
	return nil
}

func main() {
	cfg := nsq.NewConfig()
	cfg.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	cfg.TlsV1 = true
	cfg.TlsConfig.Certificates = make([]tls.Certificate, 1)
	path, _ := os.Getwd()
	certFile := path + "/nsq/cert.pem1"
	keyFile := path + "/nsq/key.pem1"
	cfg.TlsConfig.Certificates[0], _ = tls.LoadX509KeyPair(certFile, keyFile)
	cfg.MaxInFlight = 2
	consumer, err := nsq.NewConsumer("MYTEST", "TESTch1", cfg)
	if nil != err {
		log.Println(err)
		return
	}
	consumer.AddHandler(&MsgHandler{})
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if nil != err {
		log.Println(err)
		return
	}
	select {}
}
