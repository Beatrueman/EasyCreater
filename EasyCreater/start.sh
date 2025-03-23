#!/bin/bash

printf "\n"
printf "  _____                 ____                _\n"
printf " | ____|__ _ ___ _   _ / ___|_ __ ___  __ _| |_ ___ _ __\n"
printf " |  _| / _\` / __| | | | |   | '__/ _ \\/ _\` | __/ _ \\ '__|\n"
printf " | |__| (_| \\__ \\ |_| | |___| | |  __/ (_| | ||  __/ |\n"
printf " |_____\\__,_|___/\\__, |\\____|_|  \\___|\\__,_|\\__\\___|_|\n"
printf "                 |___/\n"

echo "======开始部署EasyCreater======"
# 检查 node20, golang1.23, nginx 是否安装
if ! command -v node &>/dev/null || ! node -v | grep -q "v20"; then
    echo "Node.js 20 未安装，请安装 Node.js 20 后重试。"
    exit 1
fi

if ! command -v go &>/dev/null || ! go version | grep -q "go1.23"; then
    echo "Golang 1.23 未安装，请安装 Golang 1.23 后重试。"
    exit 1
fi

if ! command -v nginx &>/dev/null; then
    echo "Nginx 未安装，请安装 Nginx 后重试。"
    exit 1
fi

# 交互式获取配置信息
echo "请输入通义千问 token:"
read TOKEN
echo "请输入阿里云 OSS Endpoint:"
read OSS_ENDPOINT
echo "请输入阿里云 OSS AccessKey:"
read OSS_ACCESSKEY
echo "请输入阿里云 OSS SecretKey:"
read OSS_SECRETKEY
echo "请输入阿里云 OSS BucketName:"
read OSS_BUCKETNAME
echo "请输入管理员用户名:"
read ADMIN_USERNAME
echo "请输入管理员密码:"
read -s ADMIN_PASSWORD
echo "请输入管理员邮箱:"
read ADMIN_EMAIL
echo "请输入管理员手机号:"
read ADMIN_PHONE
echo "请输入 MySQL 地址:"
read MYSQL_HOST
echo "请输入 MySQL 端口:"
read MYSQL_PORT
echo "请输入 MySQL 数据库名:"
read MYSQL_DATABASE
echo "请输入 MySQL 用户名:"
read MYSQL_USER
echo "请输入 MySQL 密码:"
read -s MYSQL_PASSWORD

# 写入配置文件
mkdir -p be/config
cat > be/config/config.yaml <<EOL
# 通义千问 token
token: "$TOKEN"

# 阿里云 oss 存储凭据
OSS:
  Endpoint: "$OSS_ENDPOINT"
  AccessKey: "$OSS_ACCESSKEY"
  SecretKey: "$OSS_SECRETKEY"
  BucketName: "$OSS_BUCKETNAME"

# 管理员账户
username: "$ADMIN_USERNAME"
password: "$ADMIN_PASSWORD"
email: "$ADMIN_EMAIL"
phone: "$ADMIN_PHONE"

# MySQL数据库配置
MySQL:
  host: "$MYSQL_HOST"
  port: $MYSQL_PORT
  database: "$MYSQL_DATABASE"
  user: "$MYSQL_USER"
  password: "$MYSQL_PASSWORD"
EOL

echo "配置文件已写入 be/config/config.yaml"

# 修改 nginx 配置
if [ -f "fe/nginx.conf" ]; then
    cp fe/nginx.conf /etc/nginx/conf.d/
    sed -i 's|proxy_pass .*;|proxy_pass http://localhost:8888;|' fe/nginx.conf
    sed -i 's|root .*|root /var/www/html;|' /etc/nginx/conf.d/nginx.conf
    echo "Nginx 配置文件已修改并移动至 /etc/nginx/conf.d"
else
    echo "错误: fe/nginx.conf 文件不存在！"
    exit 1
fi

# 进入前端目录并安装依赖
cd fe || exit 1
npm install
npm run build

# 移动构建产物至 Nginx HTML 目录
if [ -d "dist" ]; then
    mv dist/* /var/www/html/
    echo "前端构建完成，文件已移动至 /var/www/html"
else
    echo "错误: 构建目录 dist 不存在！"
    exit 1
fi

# 检查数据库连接
until nc -z -v -w30 "$MYSQL_HOST" 3306; do
    echo "等待 MySQL 数据库启动..."
    sleep 10
done
echo "MySQL 数据库连接成功"


# 运行后端服务
cd ../be || exit 1
nohup go run main.go &> /dev/null &

printf "\n"
printf "  _____                 ____                _\n"
printf " | ____|__ _ ___ _   _ / ___|_ __ ___  __ _| |_ ___ _ __\n"
printf " |  _| / _\` / __| | | | |   | '__/ _ \\/ _\` | __/ _ \\ '__|\n"
printf " | |__| (_| \\__ \\ |_| | |___| | |  __/ (_| | ||  __/ |\n"
printf " |_____\\__,_|___/\\__, |\\____|_|  \\___|\\__,_|\\__\\___|_|\n"
printf "                 |___/\n"

