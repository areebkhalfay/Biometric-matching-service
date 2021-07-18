package models

import "fmt"

type RequestError struct {

}

func(e *RequestError) Error() string {
	return fmt.Sprintf("Unable to decode image data as PNG.")
}