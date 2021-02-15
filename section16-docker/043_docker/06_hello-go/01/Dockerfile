# Hello-go
FROM golang:1.15.8
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

# MAINTAINER instruction allows you to set the Author field of the generated images.
MAINTAINER yarn_woe@hotmail.com