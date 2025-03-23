FROM node:20 AS build

WORKDIR /app

RUN npm config set registry https://registry.npmmirror.com

COPY package.json package-lock.json ./
RUN npm install vue-tsc --save-dev
RUN npm install

COPY . .

RUN npm run build

# 使用 Nginx 运行构建后的前端产物
FROM nginx:latest

COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

