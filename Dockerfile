FROM alpine

# configurate
ENV OS=linux
ENV ARCH=arm64

# app name
ENV APP_NAME=${PWD##*/}-${OS}-${ARCH}

# copy bin
COPY ./bin/${APP_NAME} /opt/

# setup
WORKDIR /opt
EXPOSE 8080
CMD ["/opt/${APP_NAME}"]
