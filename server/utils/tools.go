package utils

import (
	"admire-avatar/entities"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetJWTClaims(tokenData string, key []byte) (*entities.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenData, &entities.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in token parsing")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entities.UserClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

type GeneratedImage struct {
	Phrase string   `json:"phrase"`
	Tags   []string `json:"tags"`
}

func DownloadFile(data GeneratedImage) (string, error) {
	body, err := json.Marshal(data)

	resp, err := http.Post("http://192.168.43.7:8000/images/", "application/json", bytes.NewBuffer(body))
	if err == nil && resp.StatusCode == 404 {
		fmt.Println("Image not found")
		return "", err
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	filename := uuid.New().String() + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".png"
	out, err := os.Create("files/temporary/" + filename)
	if err != nil {
		return "", err
	}
	defer func() {
		err = out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	return filename, err
}
