tls-config-name := tls-config
tls-config := $(shell pwd)/$(tls-config-name)
.PHONY: build clean run proto tls
default:
	echo $(tls-config)
	mkdir $(tls-config-name)
	# ca
	openssl genrsa -out $(tls-config)/ca.key 2048
	openssl req -new -x509 -days 3650 \
		-subj "/C=GB/L=China/O=gobook/CN=github.com" \
		-key $(tls-config)/ca.key -out $(tls-config)/ca.crt
 
	# server
	openssl genrsa \
		-out $(tls-config)/server.key 2048
	openssl req -new \
		-subj "/C=GB/L=China/O=server/CN=server.io" \
		-key $(tls-config)/server.key \
		-out $(tls-config)/server.csr
	openssl x509 -req -sha256 \
		-CA $(tls-config)/ca.crt -CAkey $(tls-config)/ca.key -CAcreateserial -days 3650 \
		-in $(tls-config)/server.csr \
		-out $(tls-config)/server.crt
 
	# client
	openssl genrsa \
		-out $(tls-config)/client.key 2048
	openssl req -new \
		-subj "/C=GB/L=China/O=client/CN=client.io" \
		-key $(tls-config)/client.key \
		-out $(tls-config)/client.csr
	openssl x509 -req -sha256 \
		-CA $(tls-config)/ca.crt -CAkey $(tls-config)/ca.key -CAcreateserial -days 3650 \
		-in $(tls-config)/client.csr \
		-out $(tls-config)/client.crt
 
clean:
	rm -rf $(tls-config)

