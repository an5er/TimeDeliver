package util

import "golang.org/x/crypto/bcrypt"

func BcryptE(src string) []byte {
	dst, _ := bcrypt.GenerateFromPassword([]byte(src), 14)
	return dst
}

func BcryptC(passwdHash, passwdHashBuf string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(passwdHashBuf))
	if err == nil {
		return true
	}
	return false
}
