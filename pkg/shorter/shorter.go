package shorter

import (
	"crypto/sha256"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func URLToID(url string) string {
	hash := sha256.Sum256([]byte(url))
	intValue := new(big.Int).SetBytes(hash[:])
	return base62Encode(intValue)[:10]
}

func base62Encode(number *big.Int) string {
	base := big.NewInt(int64(len(charset)))
	zero := big.NewInt(0)
	isNegative := number.Sign() < 0

	if isNegative {
		number = new(big.Int).Neg(number)
	}

	var encoded string
	for number.Cmp(zero) > 0 {
		remainder := new(big.Int)
		number.DivMod(number, base, remainder)
		encoded = string(charset[remainder.Int64()]) + encoded
	}

	if isNegative {
		encoded = "-" + encoded
	}

	return encoded
}
