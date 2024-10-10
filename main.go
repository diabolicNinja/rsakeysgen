package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/* rsa.GenerateKey() =>
x509.MarshalPKIXPublicKey() =>
pem.Encode() */

func main() {
	/* generate keys pair */
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publickey := &privatekey.PublicKey

	/* write private key to files */
	privatekeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	privatekeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privatekeyBytes,
	}
	privatekeyPEM, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(privatekeyPEM, privatekeyBlock)
	if err != nil {
		panic(err)
	}

	/* write pub key to file */
	publickeyBytes := x509.MarshalPKCS1PublicKey(publickey)
	publickeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publickeyBytes,
	}
	publickeyPEM, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(publickeyPEM, publickeyBlock)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
