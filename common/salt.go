package common

import (
	"crypto/rand"
	"log"
)

const SaltMinLength = 16

type Salt struct {
	length int
}

func NewSaltGenerator(length int) Salt {
	if length < SaltMinLength {
		log.Fatalln("Salt length should be at least 16 bytes long")
	}

	return Salt{length}
}

func (s *Salt) Generate() ([]byte, error) {
	salt := make([]byte, s.length)
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}
