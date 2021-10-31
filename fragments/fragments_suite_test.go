package fragments_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFragments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fragments Suite")
}
