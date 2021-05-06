#!/usr/bin/env bash
cur_script_dir="$(dirname "$0")"

openssl genrsa -out ${cur_script_dir}/ca.key 2048
openssl req -new -x509 -key ${cur_script_dir}/ca.key -days 1095 -out ${cur_script_dir}/ca.crt -subj "/C=CN/ST=GD/L=SZ/O=BLADEMAINER/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"

openssl x509 -enddate -noout -in ${cur_script_dir}/ca.crt

openssl genrsa -out ${cur_script_dir}/server.pem 2048
openssl rsa -in ${cur_script_dir}/server.pem -out ${cur_script_dir}/server.key
openssl req -new -key ${cur_script_dir}/server.pem -out server.csr -subj "/C=CN/ST=GD/L=SZ/O=BLADEMAINER/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"
openssl x509 -req -sha256 -in server.csr -days 1095 -CA ${cur_script_dir}/ca.crt -CAkey ${cur_script_dir}/ca.key -CAcreateserial -out ${cur_script_dir}/server.crt


openssl genrsa -out ${cur_script_dir}/client.pem 2048
openssl rsa -in ${cur_script_dir}/client.pem -out ${cur_script_dir}/client.key
openssl req -new -key ${cur_script_dir}/client.pem -out ${cur_script_dir}/client.csr -subj "/C=CN/ST=GD/L=SZ/O=BLADEMAINER/OU=IT/CN=localhost/emailAddress=blademainer@gmail.com"
openssl x509 -req -sha256 -in ${cur_script_dir}/client.csr -days 1095 -CA ${cur_script_dir}/ca.crt -CAkey ${cur_script_dir}/ca.key -CAcreateserial -out ${cur_script_dir}/client.crt
#openssl pkcs12 -export -clcerts -in ${cur_script_dir}/client.crt -inkey ${cur_script_dir}/client.key -out client.pfx

