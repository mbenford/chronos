package chronos_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChronos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chronos Suite")
}
