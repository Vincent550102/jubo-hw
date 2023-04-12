# jubo-hw

## how to deploy

### first of all, clone a `.env` file from `example.env` and setup the value
```
cp example.env .env
```

### change `docker-compose.yml` environment and port.
```
MYSQL_ROOT_PASSWORD: example <- change the to secret password
```
```
ports:
  - 80:8787 <- change 80 port to the port you want to deploy services
```

### docker compose!
```
docker compose up -d
```

## develop
```
docker compose up -f docker-compose.dev.yml up -d
```

## swagger
After deploy, you can link /swagger/index.html to use swagger check and test api.

example: http://103.122.116.222/swagger/index.html

