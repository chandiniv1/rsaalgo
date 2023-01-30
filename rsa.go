package main

import (
	"fmt"
)

var publickey int
var privatekey int
var n int

func gcd(e, phi int) int {
	if phi != 0 {
		return gcd(phi, e%phi)
	}
	return e

}

func encode(input string) []int64 {
	encrypted := []int64{}
	for i := 0; i < len(input); i++ {
		encrypted = append(encrypted, encryption(int64(input[i])))
	}
	return encrypted
}

func decode(encrypted []int64) string {
	str := ""
	for i := 0; i < len(encrypted); i++ {
		str = str + fmt.Sprintf("%c", decryption(encrypted[i]))
	}
	return str
}

func encryption(message int64) int64 {
	var encrypted_text int64 = 1
	for e := publickey; e > 0; e-- {
		encrypted_text *= int64(message)
	}
	encrypted_text %= int64(n)

	return encrypted_text
}

func decryption(encrypted_text int64) int64 {
	var decrypted int64 = 1
	for d := privatekey; d > 0; d-- {
		decrypted *= int64(encrypted_text)
		decrypted %= int64(n)

	}
	return int64(decrypted)
}
func init() {
	p := 3
	q := 7
	n = p * q
	phi := (p - 1) * (q - 1)
	e := 2
	// d:=((k*phi)+1)/e
	for e < phi {
		//e must be smaller than phi and coprime to phi
		if gcd(e, phi) == 1 {
			break
		} else {
			e = e + 1
		}
	}
	publickey = e
	k := 2 //some constant
	privatekey = (1 + (k * phi)) / e
}
