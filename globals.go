package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20poly1305"
)

type DbConfig struct {
	NamespaceUUID uuid.UUID `mapstructure:"NamespaceUUID"`
}

const (
	dataDir = "data"
)

var (
	encryptionKey = make([]byte, chacha20poly1305.KeySize)
	aead, _       = chacha20poly1305.NewX(encryptionKey)
	dbConfig      = DbConfig{}
	configPath    = fmt.Sprintf("%s/config.guilidb", dataDir)
)
