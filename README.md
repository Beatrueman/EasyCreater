# EasyCreater（易创）—— 一站式简历生成器

## 简介

**EasyCreater——让简历制作更高效、更智能！**

### **产品简介**

> *EasyCreater，让简历制作更轻松，让求职之路更顺畅！ 🚀✨*

  EasyCreater 是一款在线简历制作工具，旨在帮助用户专注于内容创作，而无需操心排版和格式，让简历制作更加高效便捷。

🚀 **高度自定义**

- 提供极高的自由度，用户可以自由调整简历侧栏颜色、宽度、头像形状等，打造个性化简历。

🤖 **AI 智能辅助**

- 内置 AI 大模型，无论是优化表达、润色内容，还是解决写作难题，都可以随时向 AI 寻求帮助，让简历更专业、更出色。

📝 **自动保存草稿**

- 系统会自动保存用户的编辑内容，防止数据丢失，确保写作过程更安心顺畅。

💡 **灵感词汇推荐**

- 编辑界面侧边栏提供丰富的灵感词汇和参考语句，助力用户精准表达，提升简历亮点。

🌍 **简历广场分享**

- 用户可将简历分享至简历广场，与他人交流、学习，实现互帮互助，共同提升求职竞争力。

📌 **求职导航**

- 提供就业经验荟萃专区，整合知名面试刷题网站的快捷入口，并精选简历撰写必备词句，助力高效求职。

🖨 **一键导出 PDF**

- 支持简历导出为 PDF，方便打印或在线投递，确保格式美观、排版整齐，让求职更高效。

📥 **多格式简历导入**

- 支持导入 PDF、Word 及图片格式的简历，并提供云端存储，便捷管理已有简历。无论何时何地，用户都能随时访问、编辑和优化自己的简历，确保求职更高效、更从容。

![image-20250302190157153](https://gitee.com/beatrueman/images/raw/master/20250420130900061.png)

**首页**

![image-20250302190233547](https://gitee.com/beatrueman/images/raw/master/img/202503021902627.png)

**简历制作编辑**

![image-20250302191608349](https://gitee.com/beatrueman/images/raw/master/20250420130859730.png)

**我的简历**

![image-20250314234945073](https://gitee.com/beatrueman/images/raw/master/img/202503142349206.png)

**简历广场**

![image-20250302191650384](https://gitee.com/beatrueman/images/raw/master/20250420130859583.png)

**通过EasyCreater所生成的简历**

![image-20250302192331269](https://gitee.com/beatrueman/images/raw/master/img/202503021923331.png)

**AI 简历内容生成**

![image-20250314161044624](https://gitee.com/beatrueman/images/raw/master/img/202503141610734.png)

**AI 简历润色**

![image-20250314161150083](https://gitee.com/beatrueman/images/raw/master/20250420130901880.png)

**灵感词汇**

![image-20250418190011480](https://gitee.com/beatrueman/images/raw/master/20250420130901621.png)

## 技术架构与部署支持

> *EasyCreater，助力高效求职，提供强大技术支撑！ 🚀*

**EasyCreater 采用前后端分离架构开发，支持多种灵活的部署方式，包括 Linux 服务器、Docker Compose 及 Kubernetes，确保高效稳定的运行。**

💻 **前端技术栈**

- **Vue 3 + Element Plus** —— 提供流畅的用户交互体验，界面简洁美观，操作便捷高效。

⚙️ **后端技术栈**

- **Go（Gin 框架 + Gorm）** —— 高性能 Web 框架 Gin，结合 Gorm ORM，高效处理业务逻辑，提供稳定可靠的 API 服务。

🗄 **数据库**

- **MySQL** —— 采用 MySQL 作为核心数据库，支持高并发数据存储，保障简历数据的稳定存储与管理。

☁ **对象存储**

- **阿里云 OSS** —— 负责存储简历缩略图，确保数据安全、访问高效。

🧠 **AI 智能支持**

- **阿里云通义千问 API** —— 提供 AI 内容优化、智能润色、写作建议等功能，让简历更加专业出色。

🚀 **多种部署方式**

- **Linux 单机部署** —— 适用于轻量级服务器，简单易维护。
- **Docker Compose** —— 通过容器化管理，快速部署，降低环境依赖。
- **Kubernetes（K8s）** —— 适用于大规模分布式部署，实现弹性扩展与高可用性。

## 部署

建议使用docker-compose部署，方便快捷

### Linux单机部署

本项目提供了一键启动脚本`start.sh`

**环境准备**

- Linux
- Golang1.23
- Node.js 20
- Nginx
- MySQL
- [通义大模型_企业拥抱 AI 时代首选-阿里云](https://www.aliyun.com/product/tongyi)
- [阿里云 对象存储OSS_](https://www.aliyun.com/activity/purchase/storage?utm_content=se_1020490232)

***如何开始？***

```
chmod +x start.sh
./start.sh
```

### Docker compose一键部署

本项目支持 `docker-compose` 一键部署

**注意**

- 运行前需要填写 `be/config/config.yaml`，其中`MySQL.host`请填写 `mysql`，如果要使用外部MySQL，请修改相关配置。
- 按照如下配置，项目启动后，前端运行在`8080`端口，后端运行在`8888`端口。如果启动后提示端口占用，请自行修改端口`ports`

```
version: '3'

services:
  mysql:
    image: mysql:latest
    # 确保与config.yaml中填写的信息一致
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: demo
      MYSQL_PASSWORD: 123456
    ports:
      - "3306:3306"
    networks:
      - EasyCreater-network

  frontend:
    build:
      context: ./fe
    ports:
      - "8080:80"
    depends_on:
      - backend
    networks:
      - EasyCreater-network

  backend:
    build:
      context: ./be
    ports:
      - "8888:8888"
    volumes:
      - ./be/config/config.yaml:/config/config.yaml
    depends_on:
      - mysql
    networks:
      - EasyCreater-network

networks:
  EasyCreater-network:
    driver: bridge
```

***如何开始？***

```
docker-compose up -d
```

### Kubernetes部署

如需使用 Kubernetes 部署，请 clone **cloud分支**

**注意**

- 部署前需要填写 `be/config/config.yaml`，其中`MySQL.host`请填写 `easycreater-mysql`，不建议使用外部数据库。

***如何开始？***

首先在项目根目录下，创建 Secret

```
kubectl create secret generic easycreater-secret-config --from-file=./be/config/config.yaml
```

然后执行

```
kubectl apply -f deploy/service.yaml
kubectl apply -f deploy/deployment.yaml
```

如有域名解析需求，请先部署 **Traefk**，然后修改 ingressroute.yaml，最后执行

```
kubectl apply -f deploy/ingressroute.yaml
```

### CI 自动化流水线部署支持

添加了用于打包和推送镜像的**Github Action**

使用时在Settings >> Secrets and varibles >> Actions中添加secrets 

`REGISTRY_USERNAME`和`REGISTRY_PASSWORD`

![image-20240915002948963](https://gitee.com/beatrueman/images/raw/master/img/202503111459911.png)

如果要推送到类似Harbor的自建仓库，请添加varibles

`IMAGE_REGISTRY_SERVICE`：默认为docker.io

`IMAGE_FE_REPOSITORY`：默认为beatrueman/easycreater-fe

`IMAGE_BE_REPOSITORY`：默认为beatrueman/easycreater-be

推送时请指定**tag**，格式为`v1.0.0`，用于指定镜像版本

```
git tag v1.0.0
git push origin v1.0.0
```

或者手动指定**tag**
