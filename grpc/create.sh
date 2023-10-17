#!/bin/bash
# https://github.com/grpc/grpc-go/blob/master/examples/data/x509/create.sh
if [ -d "./cert" ]; then
    echo "cert exist, skip create"
else
    echo "cert folder not exist, create"
    mkdir cert
    chmod 700 ./cert
fi

# Create the server CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout ./cert/ca_key.pem                                  \
  -out ./cert/ca_cert.pem                                    \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca                                 \
  -sha256

# Create the client CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout ./cert/client_ca_key.pem                           \
  -out ./cert/client_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca                                 \
  -sha256

# Generate a server cert.
openssl genrsa -out ./cert/server_key.pem 4096
openssl req -new                                    \
  -key ./cert/server_key.pem                               \
  -days 3650                                        \
  -out ./cert/server_csr.pem                               \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server1/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \
  -in ./cert/server_csr.pem        \
  -CAkey ./cert/ca_key.pem         \
  -CA ./cert/ca_cert.pem           \
  -days 3650                \
  -set_serial 1000          \
  -out ./cert/server_cert.pem      \
  -extfile ./openssl.cnf    \
  -extensions test_server   \
  -sha256
openssl verify -verbose -CAfile ./cert/ca_cert.pem  ./cert/server_cert.pem

# Generate a client cert.
openssl genrsa -out ./cert/client_key.pem 4096
openssl req -new                                    \
  -key ./cert/client_key.pem                               \
  -days 3650                                        \
  -out ./cert/client_csr.pem                               \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \
  -in ./cert/client_csr.pem        \
  -CAkey ./cert/client_ca_key.pem  \
  -CA ./cert/client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out ./cert/client_cert.pem      \
  -extfile ./openssl.cnf    \
  -extensions test_client   \
  -sha256
openssl verify -verbose -CAfile ./cert/client_ca_cert.pem  ./cert/client_cert.pem

rm ./cert/*_csr.pem