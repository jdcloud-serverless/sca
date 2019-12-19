#!/bin/bash

language="python27"
tag="local"

cp -rf ../../../../../../images/language/${language}/server ./
sed -i -e "/processes.*$/c\processes = ${process_num}" -e "/uid.*$/c\uid = ${uid}" -e "/gid.*$/c\gid = ${gid}" ${WSGI_ROOT}/function/wsgi_server.ini 
docker build -t ${language}:${tag} .
rm -rf ./server
