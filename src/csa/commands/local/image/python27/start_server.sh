#!/bin/bash

# utf-8
export LC_ALL=en_US.utf8

# user
chmod -R 777 /tmp
groupadd -g 1000 function
useradd -d /tmp -u 1000 -g 1000 function
chown -R function:function /tmp

#modify wsgi config
sed -i -e "/listen.*$/c\listen = 128" /function/wsgi_server.ini 

#start wsgi
/usr/local/bin/uwsgi --ini /function/wsgi_server.ini
