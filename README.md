# connectnew

`connectnew` checks if &connect.Request and &connect.Response are used.

## Install

```shell
$ go install github.com/cloverrose/connectnew/cmd/connectnew@latest
```

### Or Build from source

```shell
$ make build
```

### Or Install via aqua

https://aquaproj.github.io/

## Usage

```shell
$ go vet -vettool=`which connectnew` ./...
```
