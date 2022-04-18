package utils

import "echo-starter/internal/utils/secrets"

func GenerateHandle() string {
	token, _ := secrets.NewToken(32)
	handle := token.Base62()
	return handle
}
