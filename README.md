# An opinionated template for Go projects

Fully isolated from other Go environments with [Go Modules](https://github.com/golang/go/wiki/Modules).

## setup

0. Bootstrap your project:  
`curl -L -s https://github.com/blalor/go-template/archive/master.tar.gz | tar -xz --strip-components=1`
1. Write tests
2. Write code
3. Profit

## `make` targets

* `all` (default) -- runs tests and builds `stage/$(NAME)`
* `rpm` -- builds an RPM with [fpm](https://github.com/jordansissel/fpm/)

## features

* [go-flags](https://github.com/jessevdk/go-flags) for command-line option parsing
* [logrus](https://github.com/sirupsen/logrus) for logging (because `log` is too simplistic)
* [ginkgo](https://github.com/onsi/ginkgo/ginkgo) for BDD tests
* [testify](https://github.com/stretchr/testify) for mocks, which isn't great, but is better than nothing

## credits

* [Embeddable version string](http://technosophos.com/2014/06/11/compile-time-string-in-go.html)
* [Capturing panic output to log file](https://code.google.com/p/go/issues/detail?id=325)
