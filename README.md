<img src="./assets/golang.png" height="80" width="200">   

# Go DDD

[![Go](https://github.com/fabianogoes/go-ddd/actions/workflows/go.yml/badge.svg)](https://github.com/fabianogoes/go-ddd/actions/workflows/go.yml)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fabianogoes/go-ddd)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcusolsson/goddd)](https://goreportcard.com/report/github.com/fabianogoes/go-ddd)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

Este projeto visa demonstrar uma ideia de implementação dos padrões de design tático do Domain Driven Design em Go usando influências da arquitetura [DIP](https://www.martinfowler.com/articles/dipInTheWild.html) e [Ports and Adapters](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)).

> :heavy_exclamation_mark: DIP - Dependency Inversion Principle

## Domínio de exemplo para ser desenvolvido

![Domain](./assets/domain.png)

## Dependências

- [UUID - Google UUID](https://pkg.go.dev/github.com/google/uuid)
- [Assert - Testfy](https://pkg.go.dev/github.com/stretchr/testify/assert)
- [Log - Logrus](https://github.com/Sirupsen/logrus)
- [ORM - GORM](https://gorm.io/index.html)
- [Database - Postgres](https://www.postgresql.org/docs/current/index.html)

## Referências

- [Using Enums (and Enum Types) in Golang](https://www.sohamkamani.com/golang/enums/)
- [Logging in Go: Choosing a System and Using it](https://www.honeybadger.io/blog/golang-logging/)

## Test

> Para rodar todos os testes unitários

```shell
 go test -covermode=count -coverprofile=count.out -v ./...
```
