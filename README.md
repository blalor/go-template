# A template project for the way I like to write Go

Fully isolated from other Go environments with
[gvp](https://github.com/pote/gvp) and manages dependencies with
[gpm](https://github.com/pote/gpm).  These seem to be a reasonable alternative
until Go gets its act together regarding dependency management.

## targets

* `all` (default) -- runs tests and builds `stage/$(NAME)`
* `rpm` -- builds an RPM with [fpm](https://github.com/jordansissel/fpm/)

## setup

1. Set `NAME` and `PKG_PATH` in `Makefile`.
2. Write tests
3. Write code
4. Profit

## credits

* Embeddable version: http://technosophos.com/2014/06/11/compile-time-string-in-go.html
* Capturing panic output to log file: https://code.google.com/p/go/issues/detail?id=325
