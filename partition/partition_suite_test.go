package partition_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPartition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Partition Suite")
}
