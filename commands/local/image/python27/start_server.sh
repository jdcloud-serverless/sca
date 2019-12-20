#!/bin/bash

# utf-8
export LC_ALL=en_US.utf8

# user
groupadd -g 1000 function
useradd -u 1000 -g 1000 function

#start wsgi
/usr/local/bin/uwsgi --ini /function/wsgi_server.ini
