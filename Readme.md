# jubo-hw
## Introduction
For this project, I used gin as my framework, mysql as my database, docker containerization, and gin-swagger as the API file generation mechanism.

## How to deploy
1. first of all, clone a `.env` file from `example.env` and setup the value.
```
cp example.env .env
```

2. change `docker-compose.yml` environment and port.
```
MYSQL_ROOT_PASSWORD: example <- change the to secret password
...
ports:
  - 80:8787 <- change 80 port to the port you want to deploy services
```

3. docker compose!
```
docker compose up -d
```

## develop
```
docker compose up -f docker-compose.dev.yml up -d
```

## swagger(api docs)
After deploy, you can link /swagger/index.html to use swagger check and test api.

example: http://103.122.116.222/swagger/index.html

![image](https://user-images.githubusercontent.com/48404862/231569409-50b0478f-877f-4146-8e6c-76d865418387.png)

