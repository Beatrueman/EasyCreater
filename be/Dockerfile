# 编译环境
FROM golang:latest as builder
WORKDIR /EasyCreater

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o main .

FROM debian:latest

RUN rm -rf /etc/apt/sources.list.d/debian.sources \ 
    && echo "deb http://mirrors.aliyun.com/debian stable main contrib non-free" > /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/debian stable-updates main contrib non-free" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/debian-security stable-security main contrib non-free" >> /etc/apt/sources.list \
    && apt-get clean \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates netcat-traditional \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /EasyCreater/main .
COPY start.sh /start.sh

RUN chmod +x /start.sh

EXPOSE 8888

CMD ["/start.sh"]
