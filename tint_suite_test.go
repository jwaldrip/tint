package tint_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tint Suite")
}
