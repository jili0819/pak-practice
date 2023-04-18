default:
	mkdir ./tls_config
	# ca
	openssl genrsa -out ./tls_config/ca.key 2048
	openssl req -new -x509 -days 3650 \
		-subj "/C=GB/L=China/O=gobook/CN=github.com" \
		-key ./tls_config/ca.key -out ./tls_config/ca.crt
 
	# server
	openssl genrsa \
		-out ./tls_config/server.key 2048
	openssl req -new \
		-subj "/C=GB/L=China/O=server/CN=server.io" \
		-key ./tls_config/server.key \
		-out ./tls_config/server.csr
	openssl x509 -req -sha256 \
		-CA ./tls_config/ca.crt -CAkey ./tls_config/ca.key -CAcreateserial -days 3650 \
		-in ./tls_config/server.csr \
		-out ./tls_config/server.crt
 
	# client
	openssl genrsa \
		-out ./tls_config/client.key 2048
	openssl req -new \
		-subj "/C=GB/L=China/O=client/CN=client.io" \
		-key ./tls_config/client.key \
		-out ./tls_config/client.csr
	openssl x509 -req -sha256 \
		-CA ./tls_config/ca.crt -CAkey ./tls_config/ca.key -CAcreateserial -days 3650 \
		-in ./tls_config/client.csr \
		-out ./tls_config/client.crt
 
clean:
	rm -rf ./tls_config
