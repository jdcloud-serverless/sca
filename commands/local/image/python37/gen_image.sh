#!/bin/bash

language="jdccloudserverless/sca"
tag="python37"

docker build -t ${language}:${tag} .