package errs

import (
	"fmt"
)

func NewErrIndexOutOfRange(length, index int) error {
	return fmt.Errorf("ekit:下标查出范围:长度 %d,下标 %d", length, index)
}
