package main

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/pkcs12"
)

func main() {
	// 公開鍵、秘密鍵を含んだpkcs12形式のpfxファイルを読む
	data, err := ioutil.ReadFile("cert/server.pfx")
	if err != nil {
		log.Fatal(err)
	}

	// https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20220112180741-5e0467b6c7ce/pkcs12#example-ToPEM
	// PEMに変換する。パスワードは空白文字列
	blocks, err := pkcs12.ToPEM(data, "")
	if err != nil {
		log.Fatal(err)
	}

	// 公開鍵、秘密鍵と２つの証明書が入っているので、個別に切り出す
	// map のkeyには、それぞれのブロック名を使う
	pems := map[string][]byte{}

	for _, b := range blocks {
		log.Printf("%s", b.Type)
		pems[b.Type] = pem.EncodeToMemory(b)
	}

	// X509KeyPairを作って、tls.Configにセットし、TLS Serverのlistenに入る
	// https://pkg.go.dev/crypto/tls#X509KeyPair
	cert, err := tls.X509KeyPair(pems["CERTIFICATE"], pems["PRIVATE KEY"])
	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	srv := &http.Server{
		Addr:         ":9081",
		Handler:      &handler{},
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello\n"))
}
