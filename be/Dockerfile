# 编译环境
FROM golang:latest as builder
WORKDIR /EasyCreater

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o main .

# 运行环境
FROM debian:stable-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates netcat-traditional\
    && update-ca-certificates

COPY --from=builder /EasyCreater/main .
COPY start.sh /start.sh

RUN chmod +x /start.sh

EXPOSE 8888

CMD ["/start.sh"]
