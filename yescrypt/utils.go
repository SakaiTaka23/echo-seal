package yescrypt

import "errors"

const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var errEncodeFailedTooBig = errors.New("encode number must be in 1~63")

func encode64Int(value int) (string, error) {
	if value < 0 || value >= 64 {
		return "", errEncodeFailedTooBig
	}

	return string(itoa64[value]), nil
}

func encode64(src []byte) []byte {
	//nolint: mnd
	dst := make([]byte, 0, (len(src)*8+5)/6)

	for i := 0; i < len(src); {
		value, bits := uint32(0), 0
		for ; bits < 24 && i < len(src); bits += 8 {
			value |= uint32(src[i]) << bits
			i++
		}

		for ; bits > 0; bits -= 6 {
			dst = append(dst, itoa64[value&0x3f])
			value >>= 6
		}
	}

	return dst
}
