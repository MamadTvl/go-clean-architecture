package bcrypt

import bcryptlib "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) GenerateHash(password string) (string, error) {
	bytes, err := bcryptlib.GenerateFromPassword([]byte(password), 4)
	if err == nil {
		return string(bytes), nil
	}
	return "", err
}

func (b *Bcrypt) Compare(hashedPassword string, password string) bool {
	err := bcryptlib.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
