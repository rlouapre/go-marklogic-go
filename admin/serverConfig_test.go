package admin

import (
	"testing"

	handle "github.com/ryanjdew/go-marklogic-go/handle"
)

func TestXMLServerConfigResponseHandleDeserialize(t *testing.T) {
	want := `<host xmlns="http://marklogic.com/manage"><timestamp>2013-06-18T08:11:29.188561-07:00</timestamp><version>7.0</version><platform>linux</platform><edition>Essential Enterprise</edition><host-id>4808503609057420751</host-id><host-name>my-host.marklogic.com</host-name><bind-port>7999</bind-port><connect-port>7999</connect-port><foreign-bind-port>7998</foreign-bind-port><foreign-connect-port>7998</foreign-connect-port><ssl-certificate>...elided...</ssl-certificate></host>`
	result := ServerConfigResponseHandle{Format: handle.XML}
	result.Deserialize([]byte(want))
	resp := result.Get()
	if resp.Version != "7.0" {
		t.Errorf("Not equal - ServerConfigResponseHandle Version = %+v, Want = %+v", resp, want)
	}
	if resp.BindPort != 7999 {
		t.Errorf("Not equal - ServerConfigResponseHandle BindPort = %+v, Want = %+v", resp, want)
	}
}
