[![Coverage Status](https://coveralls.io/repos/github/magrathealabs/jarvis/badge.svg)](https://coveralls.io/github/magrathealabs/jarvis)
[![Build Status](https://travis-ci.org/magrathealabs/jarvis.svg?branch=master)](https://travis-ci.org/magrathealabs/jarvis)

# Jarvis :robot:

Apenas um sistema bastante inteligente para controlar o nosso dia-a-dia de trabalho físico e remoto.

## Configuração

Servidor WEB:
```sh
    $WEB_HOST=0.0.0.0
    $WEB_PORT=3000
```

Postgres:
```sh
    $POSTGRES_HOST=localhost
    $POSTGRES_USER=postgres
    $POSTGRES_PASS=postgres
    $POSTGRES_DBNAME=test
    $POSTGRES_PORT=5432
```

Graphite:
```sh
    $METRIC_HOST=localhost
    $METRIC_PORT=2003
```

## Dev

Utilizamos Makefile & Docker & Docker-Compose para automatizar o fluxo de trabalho:

```
 - make up: Executa toda a infraestrutura com o código fonte local para executar o ambiente de dev.
 - make dev: Sobe somente os serviços de dependência para executar testes unitários no VisualCode.
 - make test: Executa toda a suite de testes. Utilizamos a infraestrutura Docker para executar os testes.
```

Golang 1.12:

```sh
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

    gvm install go1.12 -B
    gvm use go1.12
```

GOPATH:

```sh
    mkdir -p $GOPATH/src/github.com/magrathealabs/

    cd $GOPATH/src/github.com/magrathealabs/
    git clone git@github.com:magrathealabs/jarvis.git
```

Utilizamos o Govendor:

```sh
    go get -u github.com/kardianos/govendor
```

Golint-CI:

```sh
    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
    golangci-lint run
```

