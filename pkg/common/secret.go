package common

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(serverPublicKey *rsa.PublicKey, clientPublicKey *rsa.PublicKey, code string) []byte {

	hash := sha256.New()
	hash.Write(clientPublicKey.N.Bytes())
	//hash.Write([]byte(strconv.Itoa(clientPublicKey.E)))
	hash.Write([]byte{1, 0, 1})

	hash.Write(serverPublicKey.N.Bytes())
	//hash.Write([]byte(strconv.Itoa(serverPublicKey.E)))
	hash.Write([]byte{1, 0, 1})

	codeEnd, _ := hex.DecodeString(code)
	hash.Write(codeEnd)

	return hash.Sum(nil)
}
