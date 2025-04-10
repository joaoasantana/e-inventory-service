FROM ubuntu:latest
LABEL authors="jvsantana"

ENTRYPOINT ["top", "-b"]
