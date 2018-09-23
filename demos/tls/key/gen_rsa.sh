#!/usr/bin/env bash
openssl genrsa -out ca.key 2048
openssl req -new -x509 -key ca.key -out ca.crt -subj "/C=CN/ST=GD/L=SZ/O=XL/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"

openssl genrsa -out server.pem 2048
openssl rsa -in server.pem -out server.key
openssl req -new -key server.pem -out server.csr -subj "/C=CN/ST=GD/L=SZ/O=XL/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt


openssl genrsa -out client.pem 2048
openssl rsa -in client.pem -out client.key
openssl req -new -key client.pem -out client.csr -subj "/C=CN/ST=GD/L=SZ/O=XL/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"
openssl x509 -req -sha256 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt
#openssl pkcs12 -export -clcerts -in client.crt -inkey client.key -out client.pfx



