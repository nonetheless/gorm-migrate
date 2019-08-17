package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var letterSeed = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

var seed = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

// generate a string by seed
// match the k8s pattern: [a-z]([-a-z0-9]*[a-z0-9])?
func GeneratorX(length int) string {
	if length <= 0 {
		return ""
	}

	prefix := doGenerate(1, letterSeed)
	random := doGenerate(length-1, seed)

	return fmt.Sprintf("%s%s", prefix, random)
}

func doGenerate(length int, seed []byte) string {
	sc := len(seed)

	buffer := new(bytes.Buffer)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		idx := r.Intn(sc)
		buffer.WriteByte(seed[idx])
	}

	return buffer.String()
}
