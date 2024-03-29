package str

import (
	"bytes"
	"strconv"
)

// Long 自定义int64类型防止传给前端精度丢失
type Long int64

func (I *Long) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, "\"")
	parseInt, err := strconv.ParseInt(BytesToString(b), 10, 64)
	if err != nil {
		return err
	}
	*I = Long(parseInt)
	return nil
}

func (I Long) MarshalJSON() ([]byte, error) {
	str := `"` + strconv.FormatInt(int64(I), 10) + `"`
	return StringToBytes(str), nil
}
