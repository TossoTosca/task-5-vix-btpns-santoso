package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword digunakan untuk melakukan proses hashing pada password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash digunakan untuk melakukan verifikasi antara password yang dimasukkan dengan password yang sudah di-hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
