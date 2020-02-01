package main

import (
	"net/http"
	"testing"
)

func Test_handler(t *testing.T) {
	/*
		type args struct {
			writer  http.ResponseWriter
			request *http.Request
		}
	*/
	tests := []struct {
		name string
		args *args
	}{
		// TODO: Add test cases.
		//	{ "test01", NewArgs(http.ResponseWriter, new(http.Request))},
		//	{ "test02", new(args)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler(tt.args.writer, tt.args.request)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

type args struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewArgs(writer http.ResponseWriter, request *http.Request) *args {
	m := new(args)
	m.writer = writer
	m.request = request
	return m
}
