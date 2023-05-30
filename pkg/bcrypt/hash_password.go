package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func CheckPasswordHash(password, hashedByte string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedByte), []byte(password))
	return err == nil
}
