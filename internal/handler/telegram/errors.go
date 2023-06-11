package telegram

import (
	"errors"
	"fmt"
)

const (
	ErrorInitBot tgError = "error authorize in telegram %s"
)

type tgError string

func (t *TgBot) fmtError(msg tgError, params ...any) error {
	errMsg := fmt.Sprintf(string(msg), params...)
	t.Log.Error().Msg(errMsg)

	return errors.New(errMsg)
}

func (t *TgBot) fmtDebug(msg tgError, params ...any) error {
	errMsg := fmt.Sprintf(string(msg), params...)
	t.Log.Debug().Msg(errMsg)

	return errors.New(errMsg)
}
