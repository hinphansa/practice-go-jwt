package services

import (
	"errors"
	"fmt"
	"github/Hiinnn/practice-go/models"
	"regexp"
	"unicode"
)

// Signup -> Create new user
func Signup(user *models.User) (err error) {
	isMatch, _ := regexp.MatchString("[0-9a-zA-Z]{4,12}", user.Username)
	if !isMatch {
		return errors.New(fmt.Sprintln("invalid username", user.Username, user.Password))
	}

	isMatch, _ = regexp.MatchString("(?=.*[a-z])", user.Password)
	if !isMatch {
		return errors.New(fmt.Sprintln("invalid password", user.Username, user.Password))
	}

	// err = config.DB.Create(user).Error
	// if err != nil {
	// 	return err
	// }

	return nil
}

func validatePassword(password string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}
