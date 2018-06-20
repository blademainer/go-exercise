package main

import (
	"net/http"
	"crypto/tls"
	"golang.org/x/net/http2"
	"time"
	"fmt"
	"crypto/x509"
	"io/ioutil"
	"bytes"
)

type (
	ContentType string
	Encoder interface {
		Encode() ([]byte, error)
	}
)

const HTTP_TIMEOUT = time.Second * 60

func InitHttp2Client() (*http.Client, error) {
	//file, e := ioutil.ReadFile(tlsConfig.CertFile)
	//if e != nil {
	//	msg := fmt.Sprintf("Error when read file: %s error: %s \n", tlsConfig.CertFile, e.Error())
	//	panic(msg)
	//}
	//Logger.Infof("Read cert file: %s \n", file)
	//roots := x509.NewCertPool()
	//ok := roots.AppendCertsFromPEM([]byte(file))
	//if !ok {
	//	panic("failed to parse root certificate")
	//}
	//c := &tls.Config{RootCAs: roots}
	clientCertPool := x509.NewCertPool()
	caCertPath := "demos/tls/key/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}
	clientCertPool.AppendCertsFromPEM(caCrt)

	//certBytes, e := x509.ParseCertificate(file)
	certificate, e := tls.LoadX509KeyPair("demos/tls/key/client.crt", "demos/tls/key/client.key")
	if e != nil {
		fmt.Printf("Error when load certificate！error: %s \n", e.Error())
		return nil, e
	}
	certificates := []tls.Certificate{certificate}
	c := &tls.Config{
		Certificates: certificates,
		//ClientAuth:   tls.RequireAndVerifyClientCert,
		RootCAs:    clientCertPool,
		//InsecureSkipVerify:true,
	}
	//c := &tls.Config{InsecureSkipVerify: true}

	tr := &http2.Transport{
		AllowHTTP: true,
		//DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		//	cfg.Certificates = []tls.Certificate(certificates)
		//	return net.Dial(network, addr)
		//},
		TLSClientConfig: c,
	}
	cli := &http.Client{Transport: tr, Timeout: HTTP_TIMEOUT}
	return cli, nil
}

func main() {

	if client, e := InitHttp2Client(); e == nil {
		reader := bytes.NewReader([]byte("hello!"))
		if resp, err := client.Post("https://localhost:8443/h2", "application/json", reader); err == nil {
			body := make([]byte, 1024)
			if n, err2 := resp.Body.Read(body); err2 == nil {
				fmt.Println(string(body[:n]))
			} else {
				fmt.Printf("Error to send data: %s \n", err2.Error())
			}
		} else {
			fmt.Printf("Error to send data: %s \n", err.Error())
		}

	} else {
		fmt.Printf("Error to init client: %s \n", e.Error())
	}
}
