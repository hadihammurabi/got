package got

func Ok(data interface{}) Result {
	return Result{data: data}
}

func Err(err error) Result {
	return Result{err: err}
}
