package submodule_test

import (
    log "github.com/sirupsen/logrus"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestSubmodule(t *testing.T) {
    // initialize logging and capture output with Ginkgo
    log.SetLevel(log.DebugLevel)
    log.SetOutput(GinkgoWriter)

    RegisterFailHandler(Fail)
    RunSpecs(t, "Submodule Suite")
}
