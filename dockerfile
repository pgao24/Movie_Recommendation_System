FROM golang:1.23.3-bookworm

RUN apt update &&  apt install vim -y
