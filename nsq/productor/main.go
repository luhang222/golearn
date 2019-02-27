package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//produce("192.168.33.10:4150","node1", &w)
		produceHttps("192.168.33.10:4152", "node1", &w)
	}()
	//go func() {
	//	produce("127.0.0.1:4140","node2", &w)
	//}()
	w.Wait()
}

func produce(addr string, node string, w *sync.WaitGroup) {
	defer w.Done()
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
	p, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(1 * time.Second)
		p.Publish("hello", []byte(node+":"+time.Now().String()))
	}
}

func produceHttps(addr string, node string, w *sync.WaitGroup) {
	defer w.Done()
	url := "https://" + addr + "/pub?topic=" + "hello"
	fmt.Println("url:", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(node+":"+time.Now().String())))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	path, _ := os.Getwd()
	certFile := path + "/nsq/cert.pem1"
	keyFile := path + "/nsq/key.pem1"
	transport.TLSClientConfig.Certificates = make([]tls.Certificate, 1)
	transport.TLSClientConfig.Certificates[0], _ = tls.LoadX509KeyPair(certFile, keyFile)
	client.Transport = transport
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
