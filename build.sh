#!/bin/bash

git pull origin master
rm gamification
dep ensure
go build -o gamification main.go
