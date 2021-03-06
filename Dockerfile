FROM debian:stable-slim

RUN apt-get update --fix-missing && apt-get -uy upgrade
RUN apt-get -y install ca-certificates dnsutils && update-ca-certificates

ADD dbcontroller /dbcontroller

ENTRYPOINT ["/dbcontroller"]