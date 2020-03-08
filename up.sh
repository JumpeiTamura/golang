#!/bin/bash

docker-compose up -d
alias go='docker container exec golang_dev_1 go'
alias sh='docker container exec -it golang_dev_1 bash'
alias stop='docker container stop golang_dev_1'
