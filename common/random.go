package common

import (
	"crypto/rand"
	"encoding/base64"
)

// GenSalt tạo một salt an toàn, không thể đoán được.
// - sizeBytes: số byte ngẫu nhiên (>= 16 khuyến nghị).
// Trả về một chuỗi Base64 URL-safe (Raw, không padding).
func GenSalt(sizeBytes int) (string, error) {
	if sizeBytes <= 0 {
		sizeBytes = 32
	}

	b := make([]byte, sizeBytes)
	// crypto/rand.Read đảm bảo entropy cryptographic-grade
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Base64 URL-safe, không có padding (thích hợp lưu vào DB / URL)
	salt := base64.RawURLEncoding.EncodeToString(b)
	return salt, nil
}
