# goslayer

## One statement intro

GoSLayer is a tool that helps you to create a golang project in seconds.

* layered base on a [standard architecture layout](https://github.com/golang-standards/project-layout) ([for chinese](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md))
* followed by [Package-Oriented-Design guideline](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)
* creating a runnable Rest ful Web Service

## Usage

1. `$ go get -u github.com/danceyoung/goslayer`
2. cd any directory where you will creating your project: `$ goslayer`
3. Enter your project name (here is `goapp`) and then choose a web framework(gin or http handler, default is gin)
4. cd `goapp`: `go mod init goapp;go run ./cmd/myapp `
5. Open another terminal, run `curl --location --request GET 'http://127.0.0.1:8080/goslayer/events'` , response data will output.

Creation process and your project structure might look like this:

```
GoSLayer is a tool that helps you to create a golang project in seconds.

• layered base on a standard architecture layout
• followed by Package-Oriented-Design guideline
link: https://github.com/danceyoung/goslayer

Please enter your project name: goapp
Please choose a web framework,
(1) use gin, (2) use handler buildin: 1

Creating your go project with GIN
The go project is created successfully.
goapp
├── cmd/
│   └── myapp/
│       └── router/
│           └── handler/
│           └── router.go
│       └── main.go
├── internal/
│   └── myapp/
│       └── event/
├── └── pkg/
│       └── middleware/
```

## References

This tool is based on material taken from the following posts.

* [Ardan Labs: Package-Oriented-Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html),
* [Github: golang standard project layout](https://github.com/golang-standards/project-layout),
* [Microsoft: Design Fundamentals - Layout Application Guideline](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/ee658109(v=pandp.10))
* [Go面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

todo:

redis

logger

unit test, integration test
