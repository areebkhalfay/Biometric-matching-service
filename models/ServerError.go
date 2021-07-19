package models

import (
	"fmt"
)

type ServerError struct {

}

func (e ServerError) Error(error string, code int) {
	fmt.Sprintf("%d Internal Server Error %s", error, code)
}