package main

import (
	"fmt"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)


type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!\n")
}

func main() {

	//Router = gin.Default()
	//
	//Router.POST("/h2", func(context *gin.Context) {
	//	if bytes, e := context.GetRawData(); e == nil {
	//		fmt.Println(string(bytes))
	//		context.Writer.WriteString("Good: " + string(bytes))
	//	}
	//})

	pool := x509.NewCertPool()
	caCertPath := "demos/tls/key/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	//初始化一个server 实例。
	s := &http.Server{
		//设置宿主机的ip地址，并且端口号为8081
		Addr:    ":8443",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,

		},
	}
	err = s.ListenAndServeTLS("demos/tls/key/server.crt", "demos/tls/key/server.key")

	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}


	//// support http2
	//if err := Router.RunTLS(Config.Listen, Config.Tls.CertFile, Config.Tls.KeyFile); err != nil {
	//	panic(err)
	//}
}


