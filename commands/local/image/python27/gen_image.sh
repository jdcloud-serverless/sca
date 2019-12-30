#!/bin/bash

language="jdccloudserverless/sca"
tag="python27"

docker build -t ${language}:${tag} .
