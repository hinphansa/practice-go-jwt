package services

import (
	"errors"
	"github/Hiinnn/practice-go/config"
	"github/Hiinnn/practice-go/models"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
)

// var secretKey = []byte(os.Open("secret_key.txt"))

/* -------------------------------------------------------------------------- */
/*                               Public Function                              */
/* -------------------------------------------------------------------------- */

// Register -> Create new user
func Register(user *models.User) error {
	isMatch := validateUsername(user.Username)
	if !isMatch {
		return errors.New("invalid username")
	}

	isMatch = validatePassword(user.Password)
	if !isMatch {
		return errors.New("invalid password")
	}

	if result := config.DB.FirstOrCreate(user); result.RowsAffected == 0 {
		return errors.New("username exists")
	}

	return nil
}

// Signin -> Check valid username & password, Create JWT token
func Signin(user *models.User) (string, error) {
	var queryUser models.User

	config.DB.Find(&queryUser, "username = ?", user.Username)
	if queryUser.Username == user.Username && queryUser.Password == user.Password {
		token, err := createJWTToken(user.Username)
		return token, err
	}

	return "", errors.New("Invalid Username or Password")
}

// GetUSerData -> Check token is valid
func GetUSerData(tokenString string) (interface{}, error) {
	var user models.User
	claims, err := parseJWTToken(tokenString)
	if err != nil {
		return "", err
	}

	result := config.DB.Find(&user, "username = ?", claims.(jwt.MapClaims)["aud"])
	if result.Error != nil {
		return "", result.Error
	}

	return user, nil
}

/* -------------------------------------------------------------------------- */
/*                              Private Function                              */
/* -------------------------------------------------------------------------- */
func createJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "hiinnn"                              //isseur
	claims["aud"] = username                              //audience
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // expiration time

	t, err := token.SignedString([]byte("secretkeytxt"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func parseJWTToken(token string) (interface{}, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkeytxt"), nil
	})

	if err != nil {
		return "", err
	}

	return t.Claims, err
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

func validateUsername(username string) bool {
	for _, char := range username {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return false
		}
	}

	return true
}
