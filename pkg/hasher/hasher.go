package hasher

import (
	"b2b/m/pkg/generator"
	"encoding/base64"
	"errors"
	"fmt"
	"hash/fnv"

	"github.com/gofrs/uuid"
)

type Hasher interface {
	EncodeString(value string) string
	DecodeString(value string) (string, error)
}

type hasher struct {
	prefixLen int
	gen       generator.UUIDGenerator
}

func (h *hasher) EncodeString(value string) string {
	str := h.gen.GenerateString()
	return str[:h.prefixLen] + base64.StdEncoding.EncodeToString([]byte(value))
}

func (h *hasher) DecodeString(value string) (string, error) {
	if len(value) < h.prefixLen {
		return "", errors.New("value's length is less then common prefix")
	}

	bytes, err := base64.StdEncoding.DecodeString(value[h.prefixLen:])
	return string(bytes), err
}

func NewHasher(prefixLen int) Hasher {
	return &hasher{
		prefixLen: prefixLen,
		gen:       generator.NewUUIDGenerator(uuid.NewGen()),
	}
}

func SimpleHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}
