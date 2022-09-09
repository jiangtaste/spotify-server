FROM alpine

# copy bin
COPY ./bin/app /opt/

# setup
WORKDIR /opt
EXPOSE 8080
CMD ["/opt/app"]
