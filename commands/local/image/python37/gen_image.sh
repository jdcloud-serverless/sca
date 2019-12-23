#!/bin/bash

language="sca1/python37"
tag="latest"

docker build -t ${language}:${tag} .