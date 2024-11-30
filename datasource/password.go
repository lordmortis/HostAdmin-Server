package datasource

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/crypto/argon2"
	"gopkg.in/errgo.v2/errors"
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

func argonValidate(hash []byte, password string) error {
	stringHash := string(hash)
	params, salt, hash, err := decodeHash(stringHash)
	if err != nil {
		return err
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return nil
	}
	return errors.New("invalid password")
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("the encoded hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version of argon2")
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
