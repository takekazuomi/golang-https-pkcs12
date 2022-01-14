SHELL=bash

help:           ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

crt: server.crt client.crt server.pfx

client.crt:	## Create Client certificate
	openssl req \
		-x509 \
		-nodes \
		-newkey rsa:4096 \
		-keyout client.key \
		-out client.crt \
		-days 3650 \
		-subj "/CN=localhost"

ca.crt: 	## Create CA certificate
	openssl req \
		-newkey rsa:4096 \
		-nodes \
		-days 3650 \
		-x509 \
		-keyout ca.key \
		-out ca.crt \
		-subj "/CN=DemoCA"

server.csr:	## Generate a server Certificate Signing Request
	openssl req \
		-newkey rsa:4096 \
		-nodes \
		-keyout server.key \
		-out server.csr \
		-subj "/CN=localhost"

server.crt: 	## Create a server Certificate
server.crt: ca.crt server.csr
	openssl x509 \
		-req \
		-days 365 \
		-sha256 \
		-in server.csr \
		-CA ca.crt \
		-CAkey ca.key \
		-CAcreateserial \
		-out server.crt \
		-extfile <(echo "subjectAltName = DNS:localhost, IP:127.0.0.1")


server.pfx: server.crt server.key
	openssl pkcs12 -export -inkey server.key -in server.crt -out server.pfx -password pass:


dump-server: 	## Dump a server Certificate
	openssl x509 -in server.crt -text -noout

clean:		## clean
	-@rm -f client.*
	-@rm -f server.*
	-@rm -f ca.*
