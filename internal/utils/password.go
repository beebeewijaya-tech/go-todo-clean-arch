package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword can be use to hashing a password and generate some hash
// eg. Input "secret" will return "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK"
func (u *Utils) HashPassword(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}

	return string(b)
}

// VerifyPassword will return error if the comparation fails, or not match
func (u *Utils) VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
