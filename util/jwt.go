package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"io"
	"time"
)

// KeyJWT jwt key
var KeyJWT = []byte("hihi")
var keyAES = "hihi"
var expireHour = time.Duration(24)
var jwtValidate = validator.New()

// JwtData data
type JwtData struct {
	ID int    `validate:"required"`
	IP string `validate:"required"`
}

// JwtCreate create new jwt token
func JwtCreate(data JwtData) (string, error) {

	// validate struct
	if err := jwtValidate.Struct(data); err != nil {
		return "", err
	}
	// data to json
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	// encrypt json
	dataEncrypted, err := encrypt(dataJSON, keyAES)
	if err != nil {
		return "", err
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"data": dataEncrypted,
		"exp":  time.Now().Add(time.Hour * expireHour).Unix(),
	}
	// claims := token.Claims.(jwt.MapClaims)
	// claims["data"] = dataEncrypted
	// claims["exp"] = time.Now().Add(time.Hour * expireHour).Unix()

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString(KeyJWT)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// JwtVerify verify data
func JwtVerify(tokenString string, ip string) (JwtData, error) {

	// parse jwt token
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return KeyJWT, nil
	})
	if err != nil {
		return JwtData{}, err
	}
	if !token.Valid {
		return JwtData{}, jwt.NewValidationError("not validated", jwt.ValidationErrorMalformed)
	}

	// get encrypted Data from claim
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return JwtData{}, jwt.NewValidationError("not validated", jwt.ValidationErrorMalformed)
	}
	dataEncrypted, ok := claims["data"].(string)
	if !ok {
		return JwtData{}, jwt.NewValidationError("not validated", jwt.ValidationErrorMalformed)
	}

	// decrypt to json
	dataJSON, err := decrypt(dataEncrypted, keyAES)
	if err != nil {
		return JwtData{}, err
	}
	// json to struct
	var data JwtData
	err = json.Unmarshal([]byte(dataJSON), &data)
	if err != nil {
		return JwtData{}, err
	}
	// validate struct
	if err := jwtValidate.Struct(data); err != nil {
		return JwtData{}, err
	}
	// check user IP
	if data.IP != ip {
		return JwtData{}, err
	}

	return data, nil

}

func createHash(key string) string {
	hash := md5.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// Encrypt encrypt string
func encrypt(data []byte, passphrase string) (string, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypt string
func decrypt(data string, passphrase string) (string, error) {
	dataByte, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
