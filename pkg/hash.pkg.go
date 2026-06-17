package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Memory  uint32
	Time    uint32
	Threads uint8
	KeyLen  uint32
	SaltLen uint32
}

func NewHashConfig() *HashConfig {
	h := &HashConfig{}
	// based on OWASP min recommendation (May 2023)
	h.Memory = 64 * 1024 // 64 MiB
	h.Time = 2
	h.Threads = 1
	h.KeyLen = 32
	h.SaltLen = 16
	return h
}

func (h *HashConfig) genSalt() []byte {
	salt := make([]byte, h.SaltLen)
	_, _ = rand.Read(salt)
	return salt
}

func (h *HashConfig) GenHash(pwd string) string {
	salt := h.genSalt()
	hash := argon2.IDKey([]byte(pwd), salt, h.Time, h.Memory, h.Threads, h.KeyLen)

	version := argon2.Version
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	// format: $argon2id$v=$m=,t=,p=$salt$hash
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", version, h.Memory, h.Time, h.Threads, encodedSalt, encodedHash)
}

func (h *HashConfig) Compare(pwd string, hashedPwd string) error {
	splitHash := strings.Split(hashedPwd, "$")
	if len(splitHash) != 6 {
		return errors.New("invalid Hash")
	}
	if splitHash[1] != "argon2id" {
		return errors.New("not argon2id hash")
	}

	var version int
	if _, err := fmt.Sscanf(splitHash[2], "v=%d", &version); err != nil || version != argon2.Version {
		return errors.New("wrong argon2id version used")
	}

	var memory, time uint32
	var threads uint8
	if _, err := fmt.Sscanf(splitHash[3], "m=%d,t=%d,p=%d", &memory, &time, &threads); err != nil {
		return errors.New("wrong scan syntax")
	}

	salt, err := base64.RawStdEncoding.DecodeString(splitHash[4])
	if err != nil {
		return errors.New("failed to decode salt")
	}
	hash, err := base64.RawStdEncoding.DecodeString(splitHash[5])
	if err != nil {
		return errors.New("failed to decode hash")
	}

	newHash := argon2.IDKey([]byte(pwd), salt, time, memory, threads, uint32(len(hash)))

	if subtle.ConstantTimeCompare(hash, newHash) == 0 {
		return errors.New("wrong password")
	}
	return nil
}
