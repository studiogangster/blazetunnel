package main

import (
	"log"
	"os"

	"github.com/hako/branca"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Secretkey = getEnv("SECRET_KEY", "supersecretkeyyoushouldnotcommit")

func encrypt(data string) string {
	b := branca.NewBranca(Secretkey) // This key must be exactly 32 bytes long.
	// Encode String to Branca Token.
	token, err := b.EncodeToString(data)
	log.Println("Encryptinh", Secretkey, token, err)
	if err != nil {
		log.Println("Encryptinh Failure", err)
		return ""
	}

	log.Println("Encryptinh Success", token)
	return token

}

func decrypt(data string) string {
	b := branca.NewBranca(Secretkey) // This key must be exactly 32 bytes long.
	// Encode String to Branca Token.
	message, err := b.DecodeToString(data)
	log.Println("Decryption", Secretkey, data, message, err)
	if err != nil {
		return ""
	}

	return message

}

func main() {

	encrypt("oot:root:localhost")
	// decrypt("hello")
	decrypt("hGoJec54nntTDwD2zkd04kv5foTeTWwPOFpRxjOXTWaHAgcEjbRmMOSQXAq0MGdrezUxvXtHWoSQ5g8itLYY8L")

}
