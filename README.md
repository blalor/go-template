# An opinionated template for Go projects

Fully isolated from other Go environments with
[gvp](https://github.com/pote/gvp) and manages dependencies with
[gpm](https://github.com/pote/gpm).  These seem to be a reasonable alternative
until Go gets its act together regarding dependency management.

## setup

0. Bootstrap your project:  
`curl -L -s https://github.com/blalor/go-template/archive/master.tar.gz | tar -xz --strip-components=1`
1. Set `NAME` and `PKG_PATH` in `Makefile`.
2. Write tests
3. Write code
4. Profit

## `make` targets

* `all` (default) -- runs tests and builds `stage/$(NAME)`
* `rpm` -- builds an RPM with [fpm](https://github.com/jordansissel/fpm/)

## features

* [go-flags](https://github.com/jessevdk/go-flags) for command-line option parsing
* [logrus](https://github.com/Sirupsen/logrus) for logging (because `log` is too simplistic)
* [ginkgo](https://github.com/onsi/ginkgo/ginkgo) for BDD tests
* [testify](https://github.com/stretchr/testify) for mocks, which isn't great, but is better than nothing

## credits

* [Embeddable version string](http://technosophos.com/2014/06/11/compile-time-string-in-go.html)
* [Capturing panic output to log file](https://code.google.com/p/go/issues/detail?id=325)
