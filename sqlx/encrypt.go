package sqlx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type EncryptColumn[T any] struct {
	Val   T
	Valid bool
	Key   string
}

var errInvalid = errors.New("ekit EncryptColumn invalid")
var errKeyLengthInvalid = errors.New("ekit EncryptColumn key only support 16/24/32 byte")

func (e *EncryptColumn[T]) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, errInvalid
	}
	if len(e.Key) != 16 && len(e.Key) != 24 && len(e.Key) != 32 {
		return nil, errKeyLengthInvalid
	}

	var val any = e.Val
	var b []byte
	var err error
	switch valT := val.(type) {
	case string:
		b = []byte(valT)
	case []byte:
		b = valT
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64,
		float32, float64:
		buffer := new(bytes.Buffer)
		err = binary.Write(buffer, binary.BigEndian, val)
		b = buffer.Bytes()
	case int:
		tmp := int64(valT)
		buffer := new(bytes.Buffer)
		err = binary.Write(buffer, binary.BigEndian, tmp)
		b = buffer.Bytes()
	case uint:
		tmp := uint64(valT)
		buffer := new(bytes.Buffer)
		err = binary.Write(buffer, binary.BigEndian, tmp)
		b = buffer.Bytes()
	default:
		b, err = json.Marshal(val)
	}
	if err != nil {
		return nil, err
	}
	return e.aesEncrypt(b)
}

func (e *EncryptColumn[T]) Scan(src any) error {
	var err error
	var b []byte
	switch value := src.(type) {
	case []byte:
		b, err = e.aesDecrypt(value)
	case string:
		b, err = e.aesDecrypt([]byte(value))
		if err != nil {
			return nil
		}
	default:
		return fmt.Errorf("ekit:EncryptColumn.Scan not support type %v", src)
	}
	if err != nil {
		return err
	}
	err = e.setValAfterDecrypt(b)
	e.Valid = err == nil
	return err
}

func (e *EncryptColumn[T]) setValAfterDecrypt(deEncrypt []byte) error {
	var val any = &e.Val
	var err error
	switch valT := val.(type) {
	case *string:
		*valT = string(deEncrypt)
	case *[]byte:
		*valT = deEncrypt
	case *int8, *int16, *int32, *int64, *uint8, *uint16, *uint32, *uint64,
		*float32, *float64:
		reader := bytes.NewReader(deEncrypt)
		err = binary.Read(reader, binary.BigEndian, valT)
	case *int:
		tmp := new(int64)
		reader := bytes.NewReader(deEncrypt)
		err = binary.Read(reader, binary.BigEndian, tmp)
		*valT = int(*tmp)
	case *uint:
		tmp := new(uint64)
		reader := bytes.NewReader(deEncrypt)
		err = binary.Read(reader, binary.BigEndian, tmp)
		*valT = uint(*tmp)
	default:
		err = json.Unmarshal(deEncrypt, &e.Val)
	}
	return err
}

func (e *EncryptColumn[T]) aesDecrypt(src []byte) ([]byte, error) {
	newCipher, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(newCipher)
	if err != nil {
		return nil, err
	}
	nonce, cipherData := src[:gcm.NonceSize()], src[gcm.NonceSize():]
	return gcm.Open(nil, nonce, cipherData, nil)
}

func (e *EncryptColumn[T]) aesEncrypt(b []byte) (driver.Value, error) {
	newCliper, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(newCliper)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encrypted := gcm.Seal(nonce, nonce, b, nil)
	return encrypted, nil
}
