# Sample GO project to generate test coverage and publish to Coveralls.io

## prerequisites

* [Setup GO and add go/bin to $PATH](https://www.contributing.md/2019/04/08/setup-go-on-mac/)
* `go get -u github.com/ory/go-acc`

## Generate Coverage

    go-acc ./...

## See Coverage

    go tool cover -html=coverage.txt