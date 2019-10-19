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

func (r Result) IsError() bool {
	return r.err != nil
}

func (r Result) IsOk() bool {
	return !r.IsError()
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
