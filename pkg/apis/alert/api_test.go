package alert

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	u, err := url.Parse("http://1.2.3.4:8080/")
	if err != nil {
		t.Log("err:", err)
	}
	t.Logf("url: %+v", u)
	fmt.Println(u.JoinPath(apiAllAlertRules))
}
