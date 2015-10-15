#!/bin/bash

KEYDIR="keys"

DIR=$(pwd `dirname ..` | sed 's#/# #g' | awk '{print $2}')
DIRPATH="${KEYDIR}/${DIR}"
rm -rf ${DIRPATH}
mkdir -pv ${DIRPATH}
openssl req \
  -nodes \
  -x509 \
  -newkey rsa:2048 \
  -keyout ${DIRPATH}/key.pem \
  -out ${DIRPATH}/cert.pem \
  -days 365
