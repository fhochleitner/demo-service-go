package utils

import (
	"errors"
	"testing"
)

func TestCheckIfError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestErrorExists", args{errors.New("test error")}},
		{"TestNoError", args{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckIfError(tt.args.err)
		})
	}
}
