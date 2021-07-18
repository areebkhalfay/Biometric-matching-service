package models

import "fmt"

type ServerError struct {

}

func(e *ServerError) Error() string {
	return fmt.Sprintf("Unable to decode image data as PNG.")
}
