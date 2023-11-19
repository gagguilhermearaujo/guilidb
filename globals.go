package main

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/chacha20poly1305"
)

type DbConfig struct {
	NamespaceUUID uuid.UUID
}

const (
	dataDir          = "data"
	configCollection = "#guilidb"
	configKey        = "config"
)

var (
	encryptionKey = make([]byte, chacha20poly1305.KeySize)
	aead, _       = chacha20poly1305.NewX(encryptionKey)
	dbConfig      = DbConfig{NamespaceUUID: uuid.NameSpaceDNS}
)
