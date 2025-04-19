# connectnew

`connectnew` enforces the use of constructor functions (connect.NewRequest/connect.NewResponse) instead of direct struct initialization.

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

### A. Use as go vet tool

```shell
$ go vet -vettool=`which connectnew` ./...
```

### B. Use as golangci-lint custom plugin

https://golangci-lint.run/plugins/module-plugins/

Here are reference settings

`.custom-gcl.yml`

```yaml
version: v2.1.2
name: custom-golangci-lint
destination: bin
plugins:
  - module: 'github.com/cloverrose/connectnew'
    import: 'github.com/cloverrose/connectnew'
    version: v0.1.4
```

`.golangci.yml`

config file path can be relative.

```yaml
linters-settings:
  custom:
    connectnew:
      type: module
      description: connectnew enforces the use of constructor functions (connect.NewRequest/connect.NewResponse) instead of direct struct initialization.
```
