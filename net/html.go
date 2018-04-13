package net

import (
	"io"

	"github.com/Typeform/jenny/encoders"
	"github.com/Typeform/jenny/mime"
)

func init() {
	encoders.Register(mime.Type("text/html"), newHTMLEncoder)
}

func newHTMLEncoder(w io.Writer) encoders.Encoder {
	return &htmlEncoder{}
}

type htmlEncoder struct{}

func (*htmlEncoder) Encode(i interface{}) error {
	return nil
}
