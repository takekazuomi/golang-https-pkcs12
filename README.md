---
title: README
---

go httpsで、pkcs12ファイルを使う。

goにpkcs12を扱う機能が付いたのは、[crypto/x509: reading certificates from PKCS12 files #10621](https://github.com/golang/go/issues/10621) のPRから。

> It would really be best for the world if all the PKCS standards disappeared so there is a long-term cost to making it easier to use them. However, that doesn't remove the need that some people have to deal with them today.

やれやれって言う感じのことが書いてあった。

## Self-signed certificate

俗にいう、オレオレ証明書を作成し、pkcs12(pfx)ファイルに収容する。ここでは、パスワードに空文字列を使っている。

```sh
cd cert
make crt
```

## goで使う

[main.go](./main.go) を見ると最小限のことをやっている。

サーバーの起動

```sh
make run
```

別のターミナルで、接続確認する

```sh
make curl
```
