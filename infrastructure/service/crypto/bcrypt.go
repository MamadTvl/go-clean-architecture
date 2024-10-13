package bcrypt

import bcryptlib "golang.org/x/crypto/bcrypt"

type Adapter interface {
	GenerateHash(password string) (string, error)
	Compare(hashedPassword string, password string) bool
}

type bcrypt struct{}

func NewBcrypt() Adapter {
	var adapter Adapter
	adapter = &bcrypt{}
	return adapter
}

func (b *bcrypt) GenerateHash(password string) (string, error) {
	bytes, err := bcryptlib.GenerateFromPassword([]byte(password), 4)
	if err == nil {
		return string(bytes), nil
	}
	return "", err
}

func (b *bcrypt) Compare(hashedPassword string, password string) bool {
	err := bcryptlib.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
