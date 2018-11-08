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

$ dignore serviceA serviceB
serviceA/some_rules
serviceB/node_modules
serviceC
serviceD
serviceE

$ dignore serviceA serviceB > .dockerignore
$ docker build -f serviceA/Dockerfile .
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT
