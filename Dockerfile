FROM ubuntu:latest
LABEL authors="gocam"

ENTRYPOINT ["top", "-b"]