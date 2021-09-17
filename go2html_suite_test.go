package go2html_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGo2html(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go2html Suite")
}
