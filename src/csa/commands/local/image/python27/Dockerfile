FROM centos:7

ARG version="2.7.16"

#install python
ENV PYTHON_VERSION $version
RUN yum -y install wget gcc* make unzip epel-release openssl openssl-devel python-devel ncurses-devel.x86_64 bzip2-devel sqlite-devel zlib \
    && wget https://function-inuse.oss.cn-north-1.jcloudcs.com/Python-${PYTHON_VERSION}.tgz \
    && tar xvf Python-${PYTHON_VERSION}.tgz \
    && cd Python-${PYTHON_VERSION} \
    && ./configure --prefix=/usr/local \
    && make \
    && make altinstall \
    && cd - \
    && rm -rf Python-${PYTHON_VERSION}* \
    && rm -rf /usr/bin/python2 \
    && rm -rf /usr/bin/python \
    && ln -s /usr/local/bin/python2.7 /usr/bin/python2 \
    && ln -s /usr/local/bin/python2.7 /usr/bin/python \
    && ln -s /usr/bin/python2.7 /usr/bin/python2.7.5 \
    && sed -i '1c #!/usr/bin/python2.7.5' /usr/bin/yum \
    && sed -i '1c #! /usr/bin/python2.7.5' /usr/libexec/urlgrabber-ext-down \
    && wget https://files.pythonhosted.org/packages/5f/ad/1fde06877a8d7d5c9b60eff7de2d452f639916ae1d48f0b8f97bf97e570a/distribute-0.7.3.zip \
    && unzip distribute-0.7.3.zip \
    && cd distribute-0.7.3 \
    && python setup.py install \
    && cd - \
    && rm -rf distribute-0.7.3* \
    && curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
    && python get-pip.py \
    && rm -rf get-pip.py
RUN pip install uwsgi boto3 psutil jdcloud-wsgi-wrapper \
    && mkdir -p /function /dev/shm /proc \
    && touch /etc/fstab \
    && echo 'none /dev/shm tmpfs rw,nosuid,nodev,noexec 0' > /etc/fstab \
    && yum -y erase gcc make wget unzip \
    && yum clean all \
    && rm -rf /var/cache/yum \
    && mkdir -p /function

#copy wsgi server
COPY server/* /function/

#copy start command
COPY start_server.sh /function/

#entrypoint
ENTRYPOINT ["/function/start_server.sh"]