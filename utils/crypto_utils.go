package utils

import (
	"hash/crc32"
)

func CalculateEmailHash(email string) uint32 {
	return crc32.ChecksumIEEE([]byte(email))
}


func DecryptMock(cipherText string) string {
	return cipherText
}