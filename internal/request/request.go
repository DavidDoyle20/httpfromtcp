package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request, err := io.ReadAll(reader)
	if err != nil {
		err = fmt.Errorf("error reading from reader: %s", err)
		return nil, err
	}

	line := strings.Split(string(request), "\r\n")[0]
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		err = fmt.Errorf("number of arguments in request line is invalid")
		return nil, err
	}
	
	method := parts[0]
	target := parts[1]
	version := parts[2]

	if !isAllUpper(method) {
		err = fmt.Errorf("http method is not valid")
		return nil, err
	}
	if version != "HTTP/1.1" {
		err = fmt.Errorf("http version not supported")
		return nil, err
	}
	version = strings.Split(version, "/")[1]

	requestLine := RequestLine{
		HttpVersion: version,
		RequestTarget: target,
		Method: method,
	}
	req := Request {
		RequestLine: requestLine,
	}
	return &req, nil

}

func isAllUpper(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) || !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}