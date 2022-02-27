# pair-storing
simple in-memory database

[![Build Status](https://travis-ci.org/matmust/pair-storing.svg?branch=master)](https://travis-ci.org/matmust/pair-storing)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/matmust/pair-storing)
[![Go Report Card](https://goreportcard.com/badge/github.com/matmust/pair-storing)](https://goreportcard.com/report/github.com/matmust/pair-storing)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
![stability-unstable](https://img.shields.io/badge/stability-unstable-yellow.svg)

This project aims to:

- Demonstrate how the tactical design patterns from Domain Driven Design using the hexagonal architecture. 

## Running the application

Start the application on port 9000 (or whatever the `PORT` variable is set to).


```
cd /cmd/server

go run main.go
```

### Docker

You can also run the application using Docker.

```
# Build
docker build --tag pair-storing .

# Start application
ddocker run -d --name pair-storing -p 9000:9000  pair-storing
```

... or if you're using Docker Compose:

```
docker-compose up
```

## Try it!

A RESTful API provides the following endpoints:

* `GET /v1/pairs/:key`: returns the value of key if exists.
* `PUT /v1/pairs/:key`: sets value of key if non-exists. Otherwise replace the value with the given value
* `DELETE /v1/pairs`: deletes all key-value pairs

## Project Layout

The application uses the following project layout:
 
```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── mock                 mock repository for the purpose of testing
│
└── pkg                  public application and library code
    ├── flushing         flushing data service
    ├── getting          getting data service
    ├── http             includes the versions of api (rest,rpc..)
    │   └── rest         rest api implementations
    ├── storage          
    │   ├── inmem        provides in-memory implementations of all the domain repositories.
    │   ├── json         json file storage for the purpose of backup  
    ├── setting          setting data service
    ├── storage          storage for backup purposes
    └── validator        helpers for validation
```


## Contributing

If you want to fork the repository, follow these step to avoid having to rewrite the import paths.

```shell
go get github.com/matmust/pair-storing
cd $GOPATH/src/github.com/matmust/pair-storing
git remote add fork git://github.com:<yourname>/pair-storing.git

# commit your changes

git push fork
```

For more information, read [this](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html).

## Additional resources

### For watching

- [Building an Enterprise Service in Go](https://www.youtube.com/watch?v=twcDf_Y2gXY) at Golang UK Conference 2016

### Related projects

Also, if you want to learn more about Domain Driven Design, I encourage you to take a look at the [Domain Driven Design](http://www.amazon.com/Domain-Driven-Design-Tackling-Complexity-Software/dp/0321125215) book by Eric Evans.