# dignore
dockerignore generator for monorepo.

## Installation
```bash
go get -u github.com/orisano/dignore
```

## How to use
```bash
$ ls -1
serviceA
serviceB
serviceC
serviceD
serviceE

$ dignore -o - serviceA serviceB
serviceA/some_rules
serviceB/node_modules
serviceC
serviceD
serviceE

$ dignore serviceA serviceB
$ docker build -f serviceA/Dockerfile .
```
or
```bash
$ docker run orisano/dignore serviceA serviceB
```

## Author
Nao Yonashiro (@orisano)

## License
MIT
