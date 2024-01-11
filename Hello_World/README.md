# Установка Go на Ubuntu

## 1. Обновить систему

```bash
    sudo apt-get update && sudo apt-get upgrade -y
```

## 2. Скачать дистрибутив (поменять версию на актуальную)

```bash
    wget https://dl.google.com/go/go1.21.6.linux-amd64.tar.gz
```

```bash
    sudo tar -xvf go1.21.6.linux-amd64.tar.gz
```

```bash
    sudo mv go /usr/local
```

Находясь в пользовательской директории (/home/user/):

```bash
    nano .bashrc
```

```
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

## 3. Проверить версию go

```bash
    go version
```

## 4. Запустить программу hello world

```bash
    go run hello_world.go
```
