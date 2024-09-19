package common

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

func Url(ep string) (*url.URL, error) {
	if !strings.Contains(ep, ":") {
		return nil, errors.New("invalid endpoint, endpoint format=<ip>:<port>")
	}

	if !strings.HasPrefix(ep, "http") {
		ep = "http://" + ep
	}
	u, err := url.Parse(ep)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Request(method string, u *url.URL, path string, token, user, passwd string, body []byte, header http.Header) (*http.Request, error) {
	var bodyReader io.Reader
	if body == nil {
		bodyReader = nil
	} else {
		bodyReader = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, u.JoinPath(path).String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if header != nil {
		req.Header = header
	}

	// Bearer token
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
		return req, nil
	}

	// basic auth
	if user != "" && passwd != "" {
		req.SetBasicAuth(user, passwd)
		return req, nil
	}

	return nil, errors.New("must provide token or user and passwd")
}

func Ep(cmd *cobra.Command) (ip, port, token string, errs error) {
	ip, err := cmd.Flags().GetString("http.ip")
	if err != nil {
		errs = errors.Join(errs, err)
	}

	port, err = cmd.Flags().GetString("http.port")
	if err != nil {
		errs = errors.Join(errs, err)
	}

	token, err = cmd.Flags().GetString("http.token")
	if err != nil {
		errs = errors.Join(errs, err)
	}

	return

}
