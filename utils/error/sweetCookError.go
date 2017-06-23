package error

import "fmt"

type SweetCookError struct {
	Problem string
}

func (e SweetCookError) Error() string {
	return fmt.Sprintf("err: %s ", e.Problem)
}