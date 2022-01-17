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

## keyvaultにアップロードする

ここからは、Azure を使うので、`az login --use-device-code --tenant` して置く。現在ログインしているサブスクリプションは、`az account list  -o table` で確認できる。

`az keyvault certificate import` を使ってKeyVaultにアップロードする。`{}` 内を適当に変更する。

```sh
export KEYVAULT_NAME={your key vault name}
export RG={your resource group name}

make kv-import
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

## メモ

- KeyVault から、証明書をダウンローするするときは、パスワードなし（空文字）のpkcs12になり、pkcs12がbase64 encode されて落ちてくる
- Azure Go SDKの新しいやつを使ってる。[新しいやつ](https://github.com/Azure/azure-sdk-for-go#client-new-releases)は、[Azure SDK ガイドライン](https://azure.github.io/azure-sdk/golang_introduction.html)に沿っており、再試行、ロギング、トランスポートプロトコル、認証プロトコルなどで共通の機構が使われる。ただ、機能に別にstableとbetaが混在している。（と書いてある）
  - しかし、[Azure SDK Releases Go](https://azure.github.io/azure-sdk/releases/latest/go.html)をみると、stableは１つも無い。（が、BlobはGAしたような気がする）
  - `azure-sdk-for-go/services` の下が前のリリースというやつ。こっちも、リリースが続いてる[v61.2.0](https://github.com/Azure/azure-sdk-for-go/releases/tag/v61.2.0)
  - 新しいやつのKeyVaultは、keysと、secretsしかない。[sdk/keyvault](https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/keyvault)。ここは本来、[key, secret, certificate, storage](https://docs.microsoft.com/en-us/rest/api/keyvault/)の３つがあるはず。storageはオワコンだけど。certificateは、[こっち](https://github.com/Azure/azure-sdk-for-go/issues/16768)でやってるらしい。リリースは[大分先](https://github.com/Azure/azure-sdk-for-go/milestone/55)。
- 秘密鍵入の証明書は、secretsに入り、secretsとしてダウンロードできる。ポータルでは見えないけど。
- kv単体だと、CSRに署名できないらしいことに気がついたので、ここにメモ。
  - <https://stackoverflow.com/questions/60694494/how-to-sign-csr-in-azure-key-vault-using-a-issuer-certificate>
  - <https://docs.microsoft.com/en-us/archive/blogs/kv/get-started-with-azure-key-vault-certificates#create-a-certificate-manually-and-get-signed-by-a-ca>
