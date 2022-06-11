package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}

func CheckPassword(password string, hashpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
}
