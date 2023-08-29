package hasher

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type StringService interface {
	HashPassword(password string) (string, error)
	CheckHashPassword(hashPassword string, password string) error
	ConvertStringToInt(value string) (int, error)
	Sscanf(str string, format string, v ...interface{}) (int, error)
	RandomStringGenerator(length int) string
}

type stringService struct{}

func NewStringService() StringService {
	return &stringService{}
}

func (s *stringService) HashPassword(password string) (string, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), err
}

func (s *stringService) CheckHashPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (s *stringService) RandomStringGenerator(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	//length := 8
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

func (s *stringService) ConvertStringToInt(value string) (int, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *stringService) Sscanf(str string, format string, a ...interface{}) (int, error) {
	n, err := fmt.Sscanf(str, format, a...)
	return n, err
}
