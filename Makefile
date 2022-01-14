run: cert/server.pfx
	go run main.go

cert/server.pfx:
	cd cert; make crt
curl:
	curl -v --cacert cert/ca.crt https://localhost:9081/

s_client:
	echo "GET /" | openssl s_client -showcerts -CAfile cert/ca.crt -connect localhost:9081
