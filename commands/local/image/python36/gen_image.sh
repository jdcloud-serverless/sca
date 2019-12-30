#!/bin/bash

language="jdccloudserverless/sca"
tag="python36"

docker build -t ${language}:${tag} .