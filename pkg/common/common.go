package common

import (
	"errors"
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
