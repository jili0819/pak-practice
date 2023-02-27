default:
	# server
	openssl genrsa -out ./tls_config/server.key 2048
	openssl req -new -x509 -sha256 \
         -key ./tls_config/server.key \
         -out ./tls_config/server.pem -days 3650
clean:
	-rm ./tls_config/*
