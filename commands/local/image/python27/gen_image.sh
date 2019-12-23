#!/bin/bash

language="sca1/python27"
tag="latest"

docker build -t ${language}:${tag} .
