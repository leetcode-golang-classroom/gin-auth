# gin-auth

This repository is to use gin to implementation authentication logic

## package install

```shell
go get -u golang.org/x/crypto/bcrypt 
```

## mysql setup

```
services:
  db: 
    image: mysql:8
    container_name: mysql
    restart: always
    environment:
      MYSQL_DATABASE: "${MYSQL_DATABASE}"
      MYSQL_USER: "${MYSQL_USER}"
      MYSQL_PASSWORD: "${MYSQL_PASSWORD}"
      MYSQL_ROOT_PASSWORD: "${MYSQL_ROOT_PASSWORD}"
    ports:
      - "${PORT}:${PORT}"
    expose:
      - "${PORT}"
    volumes:
      - ./data:/var/lib/mysql
    logging:
      driver: "json-file"
      options: 
        max-size: "1k"
        max-file: "3"
```