package strstr

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

func randomInt(min, max int, rng *rand.Rand) int64 {
	n := rng.Intn(max-min) + min
	return int64(n)
}

func randomPrime(min, max int, rng *rand.Rand) int64 {
	var bigN *big.Int
	for {
		n := randomInt(min, max, rng)
		bigN = big.NewInt(n)
		if bigN.ProbablyPrime(20) {
			break
		}
	}
	return bigN.Int64()
}

func makeRng() *rand.Rand {
	seed := time.Now().UTC().UnixNano()
	return rand.New(rand.NewSource(seed))
}

func (rk *RabinKarp) computeRM() int64 {
	var RM int64 = 1
	for i := 1; i <= rk.M-1; i++ {
		RM = (rk.R * RM) % rk.Q
	}
	return RM
}

func (rk *RabinKarp) hash(key string) int64 {
	var h int64
	for i := 0; i < rk.M; i++ {
		h = (rk.R*h + int64(key[i])) % rk.Q
	}
	return h
}

func (rk *RabinKarp) search(txt string) int {
	N := len(txt)
	txtHash := rk.hash(txt)
	if rk.patHash == txtHash {
		return 0
	}
	for i := rk.M; i < N; i++ {
		txtHash = (txtHash + rk.Q - rk.RM*int64(txt[i-rk.M])%rk.Q) % rk.Q
		txtHash = (txtHash*rk.R + int64(txt[i])) % rk.Q
		if rk.patHash == txtHash {
			return i - rk.M + 1
		}
	}
	return -1
}

type RabinKarp struct {
	pat     string
	patHash int64
	M       int
	Q       int64
	R       int64
	RM      int64
}

func NewRabinKarp(needle string) *RabinKarp {
	rng := makeRng()
	pat := needle
	M := len(pat)
	Q := randomPrime(math.MinInt32, math.MaxInt32, rng)
	var R int64 = 256
	rk := RabinKarp{pat: pat, M: M, Q: Q, R: R}
	rk.RM = rk.computeRM()
	rk.patHash = rk.hash(pat)
	return &rk
}

func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	rk := NewRabinKarp(needle)
	return rk.search(haystack)
}

func strStr(haystack string, needle string) int {
	return strStr2(haystack, needle)
}
