module github.com/blalor/go-template

require (
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.4.3
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/vektra/mockery v0.0.0-20181123154057-e78b021dcbb5
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 // indirect
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 // indirect
	golang.org/x/sys v0.0.0-20191218084908-4a24b4065292 // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

// looks like a couple of issues in 1.12 related to go.mod parsing:
// * https://github.com/go-resty/resty/issues/230#issuecomment-467911306
replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0

// * https://github.com/golang/lint/issues/436#issuecomment-468450096
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

go 1.13
