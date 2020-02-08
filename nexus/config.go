package main

import (
	"log"
	"os"

	"github.com/LiveSocket/bot/conv"
)

type Config struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Realm      string `json:"realm"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Key        string `json:"key"`
	Salt       string `json:"salt"`
	Iterations int    `json:"iterations"`
	KeyLen     int    `json:"keylen"`
}

func LoadConfiguration() Config {
	iterations, err := conv.ToInt(os.Getenv("NEXUS_ITERATIONS"))
	if err != nil {
		log.Fatal(err)
	}
	keyLen, err := conv.ToInt(os.Getenv("NEXUS_KEY_LENGTH"))
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		Host:       os.Getenv("NEXUS_HOST"),
		Port:       os.Getenv("NEXUS_PORT"),
		Realm:      os.Getenv("NEXUS_REALM"),
		User:       os.Getenv("NEXUS_USER"),
		Password:   os.Getenv("NEXUS_PWORD"),
		Role:       os.Getenv("NEXUS_ROLE"),
		Key:        os.Getenv("NEXUS_KEY"),
		Salt:       os.Getenv("NEXUS_SALT"),
		Iterations: iterations,
		KeyLen:     keyLen,
	}
}
