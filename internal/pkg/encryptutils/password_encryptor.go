package encryptutils

import "golang.org/x/crypto/bcrypt"

type PasswordEncryptor interface {
	Hash(password string) (string, error)
	Check(password, hash string) bool
}

type bcryptPasswordEncryptor struct {
	cost int
}

func NewBcryptPasswordEncryptor(cost int) *bcryptPasswordEncryptor {
	return &bcryptPasswordEncryptor{
		cost: cost,
	}
}

func (e *bcryptPasswordEncryptor) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (e *bcryptPasswordEncryptor) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
