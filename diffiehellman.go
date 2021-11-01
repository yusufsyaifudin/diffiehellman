package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a private key greater than 2 and less than p.
func PrivateKey() *big.Int {
	prime, _ := rand.Prime(rand.Reader, 1024)
	return prime
}

// PublicKey generates a public key.
// Alice = g ** private mod |p|
// Bob   = g ** private mod |p|
func PublicKey(private, p *big.Int, g int64) (A *big.Int) {
	A = new(big.Int).Exp(big.NewInt(g), private, p)
	return
}

// NewPair creates a key pair using prime numbers p and g.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey()
	public = PublicKey(private, p, g)
	return
}

// SecretKey creates the secret key used for encryption.
// secret = BobPrivateKey   ** AlicePublicKey mod p
// or
// secret = AlicePrivateKey ** BobPublicKey   mod p
// Both operation will result equal value.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
