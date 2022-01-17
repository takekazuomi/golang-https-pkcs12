RG=omi01kv-rg
LOCATION=eastus2
KEYVAULT_NAME=omi01kv01

run: cert/server.pfx
	go run main.go

cert/server.pfx:
	cd cert; make crt
curl:
	curl -v --cacert cert/ca.crt https://localhost:9081/

s_client:
	echo "GET /" | openssl s_client -showcerts -CAfile cert/ca.crt -connect localhost:9081


.kv-init:
	az group create --name $(RG) --location $(LOCATION)
	az keyvault create --resource-group $(RG) --name $(KEYVAULT_NAME)
	touch .kv-init

kv-import: .kv-init
	az keyvault certificate import --file cert/server.pfx --name server --vault-name $(KEYVAULT_NAME)

clean:
	cd cert; make clean
	@rm .kv-init
