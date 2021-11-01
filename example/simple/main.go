package main

import (
	"crypto/rand"
	"fmt"

	"yusufsyaifudin/diffiehellman"
)

func main() {
	// Alice and Bob both agrees of prime number and primitive number
	// Server generates this number and received by client.
	// Server must return the sessionID that valid for several duration of time.
	// This sessionID will be use later as ID to transfer generated publicKey between Alice (Server) and Bob (Client).
	var prime, _ = rand.Prime(rand.Reader, 1024)
	const primitive = 1024

	// Alice produce own private and public key
	alicePriv, alicePub := diffiehellman.NewPair(prime, primitive)

	// Bob produce own private and public key
	bobPriv, bobPub := diffiehellman.NewPair(prime, primitive)

	// alicePub transferred to Bob
	// bobPub transferred to Alice
	// Transfer using sessionID as a key.

	// Alice's private key must not be same of Bob's private
	if alicePriv.String() == bobPriv.String() {
		fmt.Println("both private key must not be the same!")
		return
	}

	// Alice's public key must not be same of Bob's public
	if alicePub.String() == bobPub.String() {
		fmt.Println("both public key must not be the same!")
		return
	}

	fmt.Println("BOB PUBLIC KEY:")
	fmt.Println(bobPub)
	fmt.Println()

	// Alice receives Bob's public key, and generate secret key using Alice private key and Bob public key.
	aliceSecretKey := diffiehellman.SecretKey(alicePriv, bobPub, prime)

	// Bob receives Alice's public key, and generate secret key using Alice public key and Bob private key.
	bobSecretKey := diffiehellman.SecretKey(bobPriv, alicePub, prime)

	// Both secret key must be the same!
	if aliceSecretKey.String() != bobSecretKey.String() {
		fmt.Println("both secret key must be the same!")
		return
	}

	fmt.Println("SECRET KEY ALICE:")
	fmt.Println(aliceSecretKey.String())

	fmt.Println()

	fmt.Println("SECRET KEY BOB:")
	fmt.Println(bobSecretKey.String())

	fmt.Println()

	// Required API:
	// 1. User get modulus (prime number) and generator (primitive number), and server's public key.
	// note that p and g is constant and public:
	// https://crypto.stackexchange.com/questions/67797/in-diffie-hellman-are-g-and-p-universal-constants-or-are-they-chosen-by-one
	// 2. User send public key and server generates shared secret key, then server encrypt data and send to user.
	// But that operation can be broken by man in the middle attack. So, consider use authenticated version
	// https://stackoverflow.com/questions/10471009/how-does-the-man-in-the-middle-attack-work-in-diffie-hellman

	// To avoid that, use https://en.wikipedia.org/wiki/Station-to-Station_protocol#Full_STS
	//
}
