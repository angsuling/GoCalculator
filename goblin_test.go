package calc

import (
	"math"
	"testing"

	. "github.com/franela/goblin"
)

func TestCalc(t *testing.T) {
	gob := Goblin(t)
	gob.Describe("Calc File Test", func() {
		gob.It("should add two numbers ", func() {
			gob.Assert(Add(1, 2)).Equal(3)
		})

		gob.It("should subtract two numbers", func() {
			gob.Assert(Subtract(1, 2)).Equal(-1)
		})

		gob.It("should multiply two numbers", func() {
			gob.Assert(Multiply(1, 3)).Equal(3)
		})

		gob.It("should divide two numbers", func() {
			gob.Assert(Divide(10, 2)).Equal(float64(5))
		})
	})
}
