package yescrypt_test

import (
	"echo-seal/common"
	"echo-seal/yescrypt"
	"testing"
)

func newYescrypt() *yescrypt.Yescrypt {
	saltGen := common.NewSaltGenerator(16)

	return yescrypt.New(saltGen, 11, 31)
}

func TestValidPassword(t *testing.T) {
	t.Parallel()

	ys := newYescrypt()

	hashed, err := ys.Hash("password")
	if err != nil {
		t.Error(err)
	}

	verified, err := ys.Verify("password", hashed)
	if err != nil {
		t.Error(err)
	}

	if !verified {
		t.Error("not equal")
	}
}

func TestInValidPassword(t *testing.T) {
	t.Parallel()

	ys := newYescrypt()

	hashed, err := ys.Hash("password")
	if err != nil {
		t.Error(err)
	}

	verified, err := ys.Verify("passworddd", hashed)
	if err != nil {
		t.Error(err)
	}

	if verified {
		t.Error("should not be verified")
	}
}
