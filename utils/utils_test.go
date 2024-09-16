package utils_test

import (
	"testing"

	"github.com/chlyNiklas/lou-taylor-api/utils"
)

func TestMachesAny(t *testing.T) {
	as := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	bs := []int{13, 23, 33, 43, 53, 63, 73, 83, 93, 10, 113, 123}

	if !utils.MachesAny(as, bs) {
		t.Failed()
	}
}

func TestMachesAny_false(t *testing.T) {
	as := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	bs := []int{13, 23, 33, 43, 53, 63, 73, 83, 93, 113, 123}

	if utils.MachesAny(as, bs) {
		t.Failed()
	}
}

func BenchmarkMachesAny(b *testing.B) {
	as := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	bs := []int{13, 23, 33, 43, 53, 63, 73, 83, 93, 10, 113, 123}

	for i := 0; i < b.N; i++ {

		_ = utils.MachesAny(as, bs)

	}
}
