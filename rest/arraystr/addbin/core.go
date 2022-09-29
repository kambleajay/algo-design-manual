package addbin

import (
	"math/big"
)

func addBinary(a string, b string) string {
	x, y, answer, carry, zero := new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	x.SetString(a, 2)
	y.SetString(b, 2)
	zero.SetString("0", 2)
	for y.Cmp(zero) != 0 {
		answer.Xor(x, y)
		carry.And(x, y).Lsh(carry, 1)
		x.Set(answer)
		y.Set(carry)
	}
	return x.Text(2)
}
