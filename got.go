package got

import "errors"

func Ok(data interface{}) Result {
	return Result{data: data}
}

func Err(err string) Result {
	return Result{err: errors.New(err)}
}
