package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

// newTLSConfig only skip hostname verification
func newTLSConfig(caCrtFile, certFile, keyFile string) (*tls.Config, error) {
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
	config, err := newTLSConfig(
		"demos/tls/key/ca.crt", "demos/tls/key/server.crt", "demos/tls/key/server.key",
	)
	if err != nil {
		panic(err.Error())
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8080", config)
	if err != nil {
		panic(err.Error())
	}

	_, err = conn.Write([]byte("hello tls server"))
	if err != nil {
		panic(err.Error())
	}
	err = conn.CloseWrite()
	if err != nil {
		panic(err.Error())
	}

	all, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("reply: %v\n", string(all))
}
