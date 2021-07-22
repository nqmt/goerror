package goerror

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type GoError struct {
	TransactionId string `json:"transactionId"`
	Status        int    `json:"status"`
	Code          string `json:"code"`
	Msg           string `json:"msg"`
	cause         string
}

func (e GoError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func (e *GoError) WithCause(cause error) *GoError {
	e.cause = cause.Error()
	return e
}

func (e *GoError) GetCause() string {
	return fmt.Sprintf("'%s'", e.cause)
}

func (e *GoError) SetTransactionId(txId string) {
	e.TransactionId = txId
}


func EchoErrorReturn(err error, c echo.Context, tx string) {
	err.(*GoError).SetTransactionId(tx)
	log.Println("CAUSE:", err.(*GoError).GetCause())
	if err = c.JSON(err.(*GoError).Status, err.(*GoError)); err != nil {
		panic(err)
	}
}

// 4xx
func DefineBadRequest(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusBadRequest,
		Code:   code,
		Msg:    msg,
	}
}

func DefineUnauthorized(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    msg,
	}
}

func DefineForbidden(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusForbidden,
		Code:   code,
		Msg:    msg,
	}
}

func DefineNotFound(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusNotFound,
		Code:   code,
		Msg:    msg,
	}
}

func DefineGone(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusGone,
		Code:   code,
		Msg:    msg,
	}
}

// 5xx
func DefineInternalServerError(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusInternalServerError,
		Code:   code,
		Msg:    msg,
	}
}

func DefineNotImplemented(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusNotImplemented,
		Code:   code,
		Msg:    msg,
	}
}

func DefineBadGateway(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusBadGateway,
		Code:   code,
		Msg:    msg,
	}
}

func DefineServiceUnavailable(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusServiceUnavailable,
		Code:   code,
		Msg:    msg,
	}
}

func DefineGatewayTimeout(code, msg string) *GoError {
	return &GoError{
		Status: http.StatusGatewayTimeout,
		Code:   code,
		Msg:    msg,
	}
}
