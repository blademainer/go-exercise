package tls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

// NewSecureTLSConfig only skip hostname verification
func NewSecureTLSConfig(caCrtFile, certFile, keyFile string) (*tls.Config, error) {
	caCrt, err := ioutil.ReadFile(caCrtFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCrt); !ok {
		return nil, errors.New("failed to append CA cert as trusted cert")
	}
	cliCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	c := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{cliCrt},
	}
	return c, nil
}

// NewServerTLSConfig create server tls
func NewServerTLSConfig(caCrtFile, certFile, keyFile string) (*tls.Config, error) {
	tc := &tls.Config{}

	pool := x509.NewCertPool()
	caCertPath := caCrtFile

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)

	tc.ClientCAs = pool

	if certFile != "" {
		cliCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		tc.Certificates = []tls.Certificate{cliCrt}
	}
	return tc, nil
}

// NewClientTLSConfig only skip hostname verification
func NewClientTLSConfig(caCrtFile, certFile, keyFile string) (*tls.Config, error) {
	tc := &tls.Config{}

	if caCrtFile != "" {
		caCrt, err := ioutil.ReadFile(caCrtFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCrt); !ok {
			return nil, errors.New("failed to append CA cert as trusted cert")
		}

		tc.RootCAs = caCertPool
		tc.ClientAuth = tls.RequireAndVerifyClientCert
		tc.InsecureSkipVerify = true // Not actually skipping, we check the cert in VerifyPeerCertificate
		tc.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
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
		}
	}

	if certFile != "" {
		cliCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		tc.Certificates = []tls.Certificate{cliCrt}
	}

	return tc, nil
}
