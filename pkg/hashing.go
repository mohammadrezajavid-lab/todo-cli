package pkg

import "crypto/sha256"

func HashPassword(password string) []uint8 {

	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return bs
}
