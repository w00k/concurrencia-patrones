package main

import "fmt"

//consiste en definir una familia de funciones en bases separadas,
//de tal manera que una clase base las este utilizando de manera indiferente,
//las puede intercambiar sin tener que modificar el código
//es importante para principio solid

type PassweordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PassweordProtector)
}

func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PassweordProtector {
	return &PassweordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PassweordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

//función base
func (p *PassweordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PassweordProtector) {
	fmt.Printf("Hasshing using SHA for %s\n", p.passwordName)
}

type MD5 struct{}

func (MD5) Hash(p *PassweordProtector) {
	fmt.Printf("Hasshing using MD5 for %s\n", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("Francisco", "gmail password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
