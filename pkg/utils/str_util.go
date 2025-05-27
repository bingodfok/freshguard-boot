package utils

import "crypto/rand"

func GenNumberString(n int) (string, error) {
	if n <= 0 {
		return "", nil
	}
	const digits = "0123456789"
	b := make([]byte, n)
	_, err := rand.Read(b) // 生成随机字节
	if err != nil {
		return "", err
	}
	for i := 0; i < n; i++ {
		b[i] = digits[b[i]%10] // 转换为数字字符
	}
	return string(b), nil
}
