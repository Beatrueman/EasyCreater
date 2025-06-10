# 编译环境
FROM golang:latest AS builder
WORKDIR /EasyCreater

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o main .

FROM debian:bookworm

RUN rm -rf /etc/apt/sources.list.d/debian.sources \
    && echo "deb http://deb.debian.org/debian stable main contrib non-free" > /etc/apt/sources.list \
    && echo "deb http://deb.debian.org/debian stable-updates main contrib non-free" >> /etc/apt/sources.list \
    && echo "deb http://security.debian.org/debian-security stable-security main contrib non-free" >> /etc/apt/sources.list \
    && apt-get clean \
    && apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates netcat-traditional \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /EasyCreater

WORKDIR /EasyCreater
COPY --from=builder /EasyCreater/main /EasyCreater/main
COPY idea.json /EasyCreater/idea.json
COPY start.sh /EasyCreater/start.sh

RUN chmod +x /EasyCreater/start.sh

EXPOSE 8888

CMD ["/EasyCreater/start.sh"]
