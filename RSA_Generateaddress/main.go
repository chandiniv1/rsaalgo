package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	privateKey, publicKey, err := generateKeys()
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	address := generateAddress(publicKey)

	fmt.Printf("Private key (hex): %x\n", privateKey.D.Bytes())
	fmt.Printf("Public key (hex): %x\n", publicKey.N.Bytes())
	fmt.Printf("Address (hex): %s\n", address)
	//fmt.Println("length of address:", len(address))
}

func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {

	//generated private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	//generates public key from the generated private key
	return privateKey, &privateKey.PublicKey, nil //similar to ((rsa.privatekey).publickey)
}

//generates address using public key

func generateAddress(publicKey *rsa.PublicKey) string {
	publicKeyBytes := publicKey.N.Bytes()
	hash := sha256.Sum256(publicKeyBytes)
	address := fmt.Sprintf("%x", hash)

	return "0x" + address[:30]
}
