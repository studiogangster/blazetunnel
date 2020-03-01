package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

func generateTLSConfig() *tls.Config {
	return generateTLSConfigFallback()
}

func generateTLSConfigFallback() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
}

func genRandomString() string {

	return os.Getenv("SERVICE_NAME")

}

// XOR is used to get the XOR of 2 byte arrays
func XOR(in []byte, with []byte) []byte {
	out := make([]byte, len(in))
	for i := range in {
		// Circular
		out[i] = in[i] ^ with[i%len(with)]
	}
	return out
}
