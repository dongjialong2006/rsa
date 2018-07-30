package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type Reader struct {
	num   int
	value []byte
	seeds map[string][]byte
}

func NewReader(value []byte) *Reader {
	r := &Reader{
		value: value,
		seeds: make(map[string][]byte),
	}

	if 0 != len(value) {
		r.init(value)
	}

	r.num = 0
	return r
}

func (r *Reader) Read(p []byte) (n int, err error) {
	copy(p, seeds[fmt.Sprintf("rsa%d", r.num)])
	r.num++
	return len(p), nil
}

func GenerateKey(key string, bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	var err error = nil
	var private *rsa.PrivateKey = nil
	if "" != key {
		private, err = rsa.GenerateKey(NewReader([]byte(key)), bits)
		if err != nil {
			return nil, nil, err
		}
	} else {
		private, err = rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			return nil, nil, err
		}
	}

	return private, &private.PublicKey, nil
}

func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

func EncodePublicKey(public *rsa.PublicKey) ([]byte, error) {
	publicBytes, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{
		Bytes: publicBytes,
		Type:  "PUBLIC KEY",
	}), nil
}

func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(publicKey), nil
}

func MakeSSHKeyPair(id string, length int) (string, string, error) {
	pkey, pubkey, err := GenerateKey(id, length)
	if err != nil {
		return "", "", err
	}

	pub, err := EncodeSSHKey(pubkey)
	if err != nil {
		return "", "", err
	}

	return string(EncodePrivateKey(pkey)), string(pub), nil
}
