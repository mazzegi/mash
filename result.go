package mash

type Result struct {
	ctx  string
	text string
	err  error
}

func NewResult(ctx string, text string, err error) Result {
	return Result{
		ctx:  ctx,
		text: text,
		err:  err,
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

func (r Result) Text() string {
	return r.text
}

func (r Result) ErrorText() string {
	return r.err.Error()
}

func (r Result) Error() error {
	return r.err
}
