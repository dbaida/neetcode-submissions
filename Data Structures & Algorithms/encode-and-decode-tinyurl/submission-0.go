import (
	"math/big"
)

type Solution struct {
	data map[string]string
}
func Constructor() Solution {
	return Solution{
		data: make(map[string]string),
	}
}
func (this *Solution) Encode(longUrl string) string {
	encoded := encodeString(longUrl)
	this.data[encoded] = longUrl

	return encoded
}
func (this *Solution) Decode(shortUrl string) string {
	return this.data[shortUrl]
}

const base62Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// EncodeString encodes an arbitrary input string into a Base62 string.
func encodeString(input string) string {
	bytes := []byte(input)
	if len(bytes) == 0 {
		return ""
	}

	// Count leading zero bytes to preserve them
	var leadingZeros int
	for _, b := range bytes {
		if b == 0 {
			leadingZeros++
		} else {
			break
		}
	}

	n := new(big.Int).SetBytes(bytes)
	base := big.NewInt(62)
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result []byte
	for n.Cmp(zero) > 0 {
		n.DivMod(n, base, mod)
		result = append(result, base62Alphabet[mod.Int64()])
	}

	// Reverse the encoded characters
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Prepend encoded leading zero bytes
	prefix := strings.Repeat(string(base62Alphabet[0]), leadingZeros)
	return prefix + string(result)
}

// DecodeString decodes a Base62 string back to the original string.
func decodeString(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	// Count leading '0' characters
	var leadingZeros int
	for i := 0; i < len(input); i++ {
		if input[i] == base62Alphabet[0] {
			leadingZeros++
		} else {
			break
		}
	}

	base := big.NewInt(62)
	n := big.NewInt(0)

	for i := 0; i < len(input); i++ {
		idx := strings.IndexByte(base62Alphabet, input[i])
		if idx == -1 {
			return "", fmt.Errorf("invalid Base62 character %q at index %d", input[i], i)
		}
		n.Mul(n, base)
		n.Add(n, big.NewInt(int64(idx)))
	}

	decodedBytes := n.Bytes()

	// Prepend preserved zero bytes
	if leadingZeros > 0 {
		padding := make([]byte, leadingZeros)
		decodedBytes = append(padding, decodedBytes...)
	}

	return string(decodedBytes), nil
}
