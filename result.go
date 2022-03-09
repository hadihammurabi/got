package got

type Result struct {
	data interface{}
	err  error
}

func (r Result) Data() interface{} {
	return r.data
}

func (r Result) HasErr() bool {
	return r.err != nil
}

func (r Result) Err() error {
	return r.err
}
