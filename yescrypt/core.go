package yescrypt

import (
	"echo-seal/common"
	"errors"
	"fmt"
	"strings"

	"github.com/openwall/yescrypt-go"
)

const identifier = "$y$"
const hashPart = 3

var errInvalidPrefix = fmt.Errorf("does not start with yescrypt identifier %s", identifier)
var errInvalidHash = errors.New("invalid hash")

type Yescrypt struct {
	salt    common.Salt
	setting setting
}

func New(salt common.Salt, n, r int) *Yescrypt {
	setting := setting{
		NLog2: n,
		R:     r,
	}
	setting.validate()

	return &Yescrypt{
		salt:    salt,
		setting: setting,
	}
}

func (y *Yescrypt) toSettingBin(salt []byte) ([]byte, error) {
	encodedNLog2, err := encode64Int(y.setting.NLog2)
	if err != nil {
		return []byte{}, err
	}

	encodedR, err := encode64Int(y.setting.R)
	if err != nil {
		return []byte{}, err
	}

	encodedSalt := encode64(salt)

	setting := fmt.Sprintf("%sj%s%s$%s",
		identifier, encodedNLog2,
		encodedR, encodedSalt)

	return []byte(setting), nil
}

func (y *Yescrypt) Hash(password string) (string, error) {
	salt, err := y.salt.Generate()
	if err != nil {
		return "", err
	}

	setting, err := y.toSettingBin(salt)
	if err != nil {
		return "", err
	}

	hash, err := yescrypt.Hash([]byte(password), setting)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (y *Yescrypt) Verify(password, hash string) (bool, error) {
	splittedHash, found := strings.CutPrefix(hash, identifier)
	if !found {
		return false, errInvalidPrefix
	}

	parts := strings.Split(splittedHash, "$")
	if len(parts) != hashPart {
		return false, errInvalidHash
	}

	setting := fmt.Sprintf("%s%s$%s", identifier, parts[0], parts[1])

	newHash, err := yescrypt.Hash([]byte(password), []byte(setting))
	if err != nil {
		return false, err
	}

	return string(newHash) == hash, nil
}
