#!/bin/bash

language="jdcloudchina/sca"
tag="python36"

docker build -t ${language}:${tag} .