#!/bin/bash

language="jdcloudchina/sca"
tag="python27"

docker build -t ${language}:${tag} .
