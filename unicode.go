package goshell_magix

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func Utf8ToUtf16LE(uft8String string) ([]byte, error) {
	runes := []rune(uft8String)
	buf := new(bytes.Buffer)

	for _, r := range runes {
		err := binary.Write(buf, binary.LittleEndian, uint16(r))
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func Utf16LEToUtf8(utf16LEBytes []byte) (string, error) {
	if len(utf16LEBytes)%2 != 0 {
		return "", errors.New("UTF-8 length must be even")
	}

	runes := make([]rune, 0, len(utf16LEBytes)/2)
	for i := 0; i < len(utf16LEBytes); i += 2 {
		codeUnit := binary.LittleEndian.Uint16(utf16LEBytes[i : i+2])
		runes = append(runes, rune(codeUnit))
	}
	return string(runes), nil
}
