package mash

type Result struct {
	ctx   string
	value string
	err   error
}

func NewResult(ctx string, value string, err error) Result {
	return Result{
		ctx:   ctx,
		value: value,
		err:   err,
	}
}

var NoResult = NewResult("", "", nil)

func (r Result) Failed() bool {
	return r.err != nil
}

func (r Result) Ok() bool {
	return !r.Failed()
}

func (r Result) Context() string {
	return r.ctx
}

func (r Result) Value() string {
	return r.value
}

func (r Result) ErrorText() string {
	return r.err.Error()
}

func (r Result) Error() error {
	return r.err
}
