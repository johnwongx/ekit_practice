package sqlx

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptColumn_Value(t *testing.T) {
	testCases := []struct {
		name    string
		c       EncryptColumn[any]
		wantVal driver.Value
		wantErr error
	}{
		{
			name: "int",
			c: EncryptColumn[any]{
				Val:   1,
				Valid: true,
				Key:   "",
			},
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.c.Value()
			assert.Equal(t, err, tc.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, res, tc.wantVal)
		})
	}
}

func TestEncryptColumn_aesEncrypt(t *testing.T) {
	var val int64 = 1
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, val)
	if err != nil {
		return
	}
	b := buffer.Bytes()
	fmt.Println(b)
}
