#!/bin/bash

language="sca1/python36"
tag="latest"

docker build -t ${language}:${tag} .