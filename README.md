# Overview
`number-crunch` is a simple Go lang web application that listens on port 8080.
You can use this application for Kubernetes learning.


# Get started

Clone the repo and run the server.
```shell
git clone git@github.com:cloudqubes/number-crunch.git
go run number-crunch.go
```

# URL endpoints

`number-crunch` has two URL end points.

## Square root URL

Return the square root of a number.
Path: `/square-root/<number>

Example with `curl`
```shell
curl http://127.0.0.1/square-root/4
  => {"InputNumber":4,"SquareRoot":2}
```

## Cube root URL

Return the cube root of a number.
Path: `/cube-root/<number>

Example with `curl`
```shell
curl http://127.0.0.1:8080/cube-root/8
  => {"InputNumber":8,"CubeRoot":2}
```