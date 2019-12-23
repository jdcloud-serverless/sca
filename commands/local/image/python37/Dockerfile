FROM centos:7

ARG version="3.7.3"

#install python
ENV PYTHON_VERSION $version
RUN yum -y install wget openssl-devel gcc make git python3-devel libffi-devel sqlite-devel \
    && wget https://function-inuse.oss.cn-north-1.jcloudcs.com/Python-${PYTHON_VERSION}.tgz \
    && tar xvf Python-${PYTHON_VERSION}.tgz \
    && cd Python-${PYTHON_VERSION} \
    && ./configure --prefix=/usr/local --enable-shared --enable-loadable-sqlite-extensions \
    && make \
    && make altinstall \
    && cp libpython3.7m.so.1.0 /usr/lib64/ \
    && cd - \
    && rm -rf Python-${PYTHON_VERSION}* \
    && rm -f /usr/bin/python \
    && rm -f /usr/bin/python3 \
    && ln -s /usr/local/bin/python3.7 /usr/bin/python \
    && ln -s /usr/local/bin/python3.7 /usr/bin/python3 \
    && sed -i '1c #!/usr/bin/python2' /usr/bin/yum \
    && sed -i '1c #! /usr/bin/python2' /usr/libexec/urlgrabber-ext-down \
    && pip3.7 install --upgrade pip \
    && pip3.7 install uwsgi boto3 psutil pysqlite3 jdcloud-wsgi-wrapper \
    && mkdir -p /function /dev/shm /proc \
    && touch /etc/fstab \
    && echo 'none /dev/shm tmpfs rw,nosuid,nodev,noexec 0' > /etc/fstab \
    && yum -y erase gcc make git wget \
    && yum clean all \
    && rm -rf /var/cache/yum
    
#copy wsgi server
COPY server/* /function/

#copy start command
COPY start_server.sh /function/

#entrypoint
ENTRYPOINT ["/function/start_server.sh"]