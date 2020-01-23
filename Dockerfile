FROM golang:latest

LABEL maintainer="Yuhando Charlilio <yuhando74@gmail.com>"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD [ "./main" ]