package datasource

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"runtime"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var argonParams params

func init() {
	numCpus := runtime.NumCPU()
	if numCpus >= 256 {
		numCpus = 255
	}

	argonParams = params{
		memory:      64 * 1024,
		iterations:  2,
		parallelism: uint8(numCpus),
		saltLength:  16,
		keyLength:   32,
	}
}

func argonHash(password string) (hash []byte, err error) {
	hashString := fmt.Sprintf("{ARGON2ID}$argon2id$v=%d", argon2.Version)
	hashString += fmt.Sprintf("$m=%d,t=%d,p=%d", argonParams.memory, argonParams.iterations, argonParams.parallelism)
	salt, err := generateRandomBytes(argonParams.saltLength)
	if err != nil {
		return nil, err
	}
	hashString += fmt.Sprintf("$%s", base64.RawStdEncoding.EncodeToString(salt))

	rawHash := argon2.IDKey([]byte(password), salt, argonParams.iterations, argonParams.memory, argonParams.parallelism, argonParams.keyLength)
	hashString += fmt.Sprintf("$%s", base64.RawStdEncoding.EncodeToString(rawHash))

	return []byte(hashString), nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
