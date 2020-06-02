package util

import (
	"github.com/satori/go.uuid"
	"strings"
)

func GetV1() (string, error) {

	u, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func GetV4() (string, error) {

	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

//GetV1 去掉 "-"
func GetV1Trim() (string, error) {

	u, err := uuid.NewV1()
	if err != nil {
		return "", err
	}


	return strings.ReplaceAll(u.String(),"-",""),nil
}

//GetV4 去掉 "-"
func GetV4Trim() (string, error) {

	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(u.String(),"-",""),nil
}
