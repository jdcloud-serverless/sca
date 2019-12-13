#!/bin/bash

language="python27"
tag="local"

docker build -t ${language}:${tag} .
