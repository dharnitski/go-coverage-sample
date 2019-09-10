# Sample GO project to generate test coverage and publish to Coveralls.io

[![CircleCI](https://circleci.com/gh/dharnitski/go-coverage-sample.svg?style=svg)](https://circleci.com/gh/dharnitski/go-coverage-sample)
[![Coverage Status](https://coveralls.io/repos/github/dharnitski/go-coverage-sample/badge.svg?branch=master)](https://coveralls.io/github/dharnitski/go-coverage-sample?branch=master)

## prerequisites

* [Setup GO and add go/bin to $PATH](https://blog.dharnitski.com/2019/04/08/setup-go-on-mac/)

## Generate Coverage

    go test -coverpkg ./... ./... -coverprofile=coverage.txt

## See Coverage

    go tool cover -html=coverage.txt