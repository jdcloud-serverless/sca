#!/bin/bash

language="jdcloudchina/sca"
tag="python37"

docker build -t ${language}:${tag} .