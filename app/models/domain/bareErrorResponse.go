package domain

import (
	"strings"
)

type BaseErrorMessage struct {
	ErrorCode    string `json:"error-code"`
	ErrorMessage string `json:"error-message"`
}

func JSONError(error string) (errJSON *BaseErrorMessage) {
	mainMsg := strings.Split(error, ":")
	part1 := strings.Split(mainMsg[0], " ")

	errCode, errMessage := part1[1], mainMsg[1]

	errorMsg := &BaseErrorMessage{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}

	return errorMsg
}
