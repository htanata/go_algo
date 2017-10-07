package partition_test

import (
	. "github.com/htanata/go_algo/partition"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Partition", func() {
	It("partitions a list by 1", func() {
		Expect(Partition([]int{1, 2, 3}, 1)).To(Equal(
			[][]int{
				[]int{1, 2, 3},
			},
		))
	})

	It("partitions a list by 2", func() {
		Expect(Partition([]int{1, 2}, 2)).To(Equal(
			[][]int{
				[]int{1},
				[]int{2},
			},
		))
	})

	It("partitions a list by 3", func() {
		Expect(Partition([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)).To(Equal(
			[][]int{
				[]int{1, 2, 3, 4, 5},
				[]int{6, 7},
				[]int{8, 9},
			},
		))
	})
})
