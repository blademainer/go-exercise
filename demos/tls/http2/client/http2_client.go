package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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
	// file, e := ioutil.ReadFile(tlsConfig.CertFile)
	// if e != nil {
	// 	msg := fmt.Sprintf("Error when read file: %s error: %s \n", tlsConfig.CertFile, e.Error())
	// 	panic(msg)
	// }
	// Logger.Infof("Read cert file: %s \n", file)
	// roots := x509.NewCertPool()
	// ok := roots.AppendCertsFromPEM([]byte(file))
	// if !ok {
	// 	panic("failed to parse root certificate")
	// }
	// c := &tls.Config{RootCAs: roots}
	clientCertPool := x509.NewCertPool()
	caCertPath := "demos/tls/key/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
	}
	clientCertPool.AppendCertsFromPEM(caCrt)

	//certBytes, e := x509.ParseCertificate(file)
	//certificate, e := tls.LoadX509KeyPair("demos/tls/key/client.crt", "demos/tls/key/client.key")
	//if e != nil {
	//	fmt.Printf("Error when load certificate！error: %s \n", e.Error())
	//	return nil, e
	//}
	//certificates := []tls.Certificate{certificate}
	//c := &tls.Config{
	//	Certificates: certificates,
	//	//ClientAuth:   tls.RequireAndVerifyClientCert,
	//	RootCAs: clientCertPool,
	//	//InsecureSkipVerify:true,
	//}
	// certBytes, e := x509.ParseCertificate(file)
	// certificate, e := tls.LoadX509KeyPair("demos/tls/key/client.crt", "demos/tls/key/client.key")
	// if e != nil {
	// 	fmt.Printf("Error when load certificate！error: %s \n", e.Error())
	// 	return nil, e
	// }
	// certificates := []tls.Certificate{certificate}
	// c := &tls.Config{
	// 	Certificates: certificates,
	// 	//ClientAuth:   tls.RequireAndVerifyClientCert,
	// 	RootCAs: clientCertPool,
	// 	//InsecureSkipVerify:true,
	// }
	c, err := newTLSCofig(caCertPath, "demos/tls/key/client.crt", "demos/tls/key/client.key")
	if err != nil{
		panic(err)
	}
	//c := &tls.Config{InsecureSkipVerify: true}
	// c := &tls.Config{InsecureSkipVerify: true}

	tr := &http2.Transport{
		AllowHTTP: true,
		// DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		//	cfg.Certificates = []tls.Certificate(certificates)
		//	return net.Dial(network, addr)
		// },
		TLSClientConfig: c,
	}
	cli := &http.Client{Transport: tr, Timeout: HTTP_TIMEOUT}
	return cli, nil
}

// newTLSCofig only skip hostname verification
func newTLSCofig(caCrtFile, certFile, keyFile string) (*tls.Config, error) {
	caCrt, err := ioutil.ReadFile(caCrtFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCrt); !ok {
		return nil, errors.New("Failed to append CA cert as trusted cert")
	}
	cliCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{cliCrt},
		InsecureSkipVerify: true, // Not actually skipping, we check the cert in VerifyPeerCertificate
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			// Code copy/pasted and adapted from
			// https://github.com/golang/go/blob/81555cb4f3521b53f9de4ce15f64b77cc9df61b9/src/crypto/tls/handshake_client.go#L327-L344, but adapted to skip the hostname verification.
			// See https://github.com/golang/go/issues/21971#issuecomment-412836078.

			// If this is the first handshake on a connection, process and
			// (optionally) verify the server's certificates.
			certs := make([]*x509.Certificate, len(rawCerts))
			for i, asn1Data := range rawCerts {
				cert, err := x509.ParseCertificate(asn1Data)
				if err != nil {
					return errors.New("bitbox/electrum: failed to parse certificate from server: " + err.Error())
				}
				certs[i] = cert
			}

			opts := x509.VerifyOptions{
				Roots:         caCertPool,
				CurrentTime:   time.Now(),
				DNSName:       "", // <- skip hostname verification
				Intermediates: x509.NewCertPool(),
			}

			for i, cert := range certs {
				if i == 0 {
					continue
				}
				opts.Intermediates.AddCert(cert)
			}
			_, err := certs[0].Verify(opts)
			return err
		},
	}, nil
}


func main() {

	if client, e := InitHttp2Client(); e == nil {
		fmt.Println(client)
		respCh := make(chan []byte, 1)

		wg := sync.WaitGroup{}
		for i := 0; i < 10000; i++ {
			go func(response chan []byte) {
				wg.Add(1)
				defer wg.Done()
				// fmt.Println("new go func")
				// response <- []byte("ccc")

				reader := bytes.NewReader([]byte("hello!"))
				if resp, err := client.Post("https://127.0.0.1:8443/h2", "application/json", reader); err == nil {
					body := make([]byte, 1024)
					if n, err2 := resp.Body.Read(body); err2 == nil {
						data := body[:n]
						// fmt.Println("data: ", string(data))
						response <- data
					} else {
						fmt.Printf("Error to send data: %s \n", err2.Error())
					}
					// if not close, may produce memory leak
					// resp.Body.Close()
				} else {
					fmt.Printf("Error to send data: %s \n", err.Error())
				}
			}(respCh)
		}

		go func() {
			for {
				select {
				case resp := <-respCh:
					fmt.Printf("received: %s \n", string(resp))
				}
			}
		}()
		wg.Wait()
	} else {
		fmt.Printf("Error to init client: %s \n", e.Error())
	}

}
