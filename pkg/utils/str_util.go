package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"time"
)

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

// GenerateUserName 生成包含小写字母的随机游客名称
func GenerateUserName() string {
	// 设置随机种子
	mrand.NewSource(time.Now().UnixNano())

	chars := make([]byte, 10)
	for i := 0; i < 10; i++ {
		switch mrand.Intn(3) {
		case 0: // 大写字母
			chars[i] = byte(mrand.Intn(26) + 65)
		case 1: // 小写字母
			chars[i] = byte(mrand.Intn(26) + 97)
		case 2: // 数字
			chars[i] = byte(mrand.Intn(10) + 48)
		}
	}
	return string(chars)
}

// GenerateNumericUserID 生成指定位数的纯数字 UserId
func GenerateNumericUserID(length int) (string, error) {
	if length < 2 {
		return "", fmt.Errorf("length must be at least 2")
	}

	firstDigit, err := randomDigitFromSet([]int{3, 4, 5, 6})
	if err != nil {
		return "", err
	}
	middleDigits := make([]byte, length-2)
	for i := range middleDigits {
		digit, err := randomDigit(0, 9)
		if err != nil {
			return "", err
		}
		middleDigits[i] = byte(digit + '0')
	}
	lastDigit, err := randomDigit(1, 9)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d%s%d", firstDigit, string(middleDigits), lastDigit), nil
}

// randomDigit 生成指定范围内的随机数字
func randomDigit(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min cannot be greater than max")
	}

	rangeSize := big.NewInt(int64(max - min + 1))

	n, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	return min + int(n.Int64()), nil
}

// randomDigitFromSet 从指定数字集合中随机选择一个
func randomDigitFromSet(set []int) (int, error) {
	if len(set) == 0 {
		return 0, fmt.Errorf("set cannot be empty")
	}

	index, err := randomDigit(0, len(set)-1)
	if err != nil {
		return 0, err
	}

	return set[index], nil
}
