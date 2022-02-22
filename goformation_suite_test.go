package goformation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoformation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goformation Suite")
}
