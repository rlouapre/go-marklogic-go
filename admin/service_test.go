package admin

import (
	"encoding/xml"
	"reflect"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	handle "github.com/ryanjdew/go-marklogic-go/handle"
	test "github.com/ryanjdew/go-marklogic-go/test"
	util "github.com/ryanjdew/go-marklogic-go/util"
)

var instanceAdminResponse = `<restart xmlns="http://marklogic.com/manage"><last-startup host-id="13544732455686476949">2013-04-01T10:35:19.09913-07:00</last-startup><link><kindref>timestamp</kindref><uriref>/admin/v1/timestamp</uriref></link><message>Check for new timestamp to verify host restart.</message></restart>`

func TestInstanceAdmin(t *testing.T) {
	client, server := test.AdminClient(instanceAdminResponse)
	defer server.Close()
	lastStartup, _ := time.Parse(time.RFC3339Nano, "2013-04-01T10:35:19.09913-07:00")
	lastStartupDateTimeNano, _ := util.NewDateTimeNano(lastStartup)
	want :=
		RestartResponse{
			XMLName: xml.Name{"http://marklogic.com/manage", "restart"},
			LastStartup: LastStartupElement{
				XMLName: xml.Name{"http://marklogic.com/manage", "last-startup"},
				Value:   *lastStartupDateTimeNano,
				HostId:  "13544732455686476949",
			},
			Link: LinkElement{
				XMLName: xml.Name{"http://marklogic.com/manage", "link"},
				KindRef: "timestamp",
				UriRef:  "/admin/v1/timestamp",
			},
			Message: "Check for new timestamp to verify host restart.",
		}
	// Using Basic Auth for test so initial call isn't actually made
	respHandle := RestartResponseHandle{Format: handle.XML}
	err := instanceAdmin(client, "admin", "password", "public", &respHandle)
	resp := respHandle.Get()
	if err != nil {
		t.Errorf("Error = %v", err)
	} else if resp == nil {
		t.Errorf("No response found")
	} else if !reflect.DeepEqual(want.LastStartup, resp.LastStartup) {
		t.Errorf("InstanceAdmin LastStartup = %+v, Want = %+v", spew.Sdump(resp.LastStartup), spew.Sdump(want.LastStartup))
	} else if !reflect.DeepEqual(resp.Link, want.Link) {
		t.Errorf("InstanceAdmin Link = %+v, Want = %+v", spew.Sdump(resp.Link), spew.Sdump(want.Link))
	} else if !reflect.DeepEqual(*resp, want) {
		t.Errorf("InstanceAdmin Response = %+v, Want = %+v", spew.Sdump(*resp), spew.Sdump(want))
	}
}

var initResponse = `<restart xmlns="http://marklogic.com/manage"><last-startup host-id="13544732455686476949">2013-05-15T09:01:43.019261-07:00</last-startup><link><kindref>timestamp</kindref><uriref>/admin/v1/timestamp</uriref></link><message>Check for new timestamp to verify host restart.</message></restart>`

func TestInit(t *testing.T) {
	client, server := test.AdminClient(initResponse)
	defer server.Close()
	lastStartup, _ := time.Parse(time.RFC3339Nano, "2013-05-15T09:01:43.019261-07:00")
	lastStartupDateTimeNano, _ := util.NewDateTimeNano(lastStartup)
	want :=
		RestartResponse{
			XMLName: xml.Name{"http://marklogic.com/manage", "restart"},
			LastStartup: LastStartupElement{
				XMLName: xml.Name{"http://marklogic.com/manage", "last-startup"},
				Value:   *lastStartupDateTimeNano,
				HostId:  "13544732455686476949",
			},
			Link: LinkElement{
				XMLName: xml.Name{"http://marklogic.com/manage", "link"},
				KindRef: "timestamp",
				UriRef:  "/admin/v1/timestamp",
			},
			Message: "Check for new timestamp to verify host restart.",
		}
	ih := InitHandle{}
	license := InitializeProperties{
		LicenseKey: "1234-5678-90AB",
		Licensee:   "Your Licensee",
	}
	ih.Serialize(license)

	// Using Basic Auth for test so initial call isn't actually made
	respHandle := RestartResponseHandle{Format: handle.XML}
	err := initialize(client, &ih, &respHandle)
	resp := respHandle.Get()
	if err != nil {
		t.Errorf("Error = %v", err)
	} else if resp == nil {
		t.Errorf("No response found")
	} else if !reflect.DeepEqual(want.LastStartup, resp.LastStartup) {
		t.Errorf("InstanceAdmin LastStartup = %+v, Want = %+v", spew.Sdump(resp.LastStartup), spew.Sdump(want.LastStartup))
	} else if !reflect.DeepEqual(resp.Link, want.Link) {
		t.Errorf("InstanceAdmin Link = %+v, Want = %+v", spew.Sdump(resp.Link), spew.Sdump(want.Link))
	} else if !reflect.DeepEqual(*resp, want) {
		t.Errorf("InstanceAdmin Response = %+v, Want = %+v", spew.Sdump(*resp), spew.Sdump(want))
	}
}

var timestampResponse = `2013-05-15T10:35:38.932514-07:00`

func TestTimestamp(t *testing.T) {
	client, server := test.AdminClient(timestampResponse)
	defer server.Close()
	want, _ := time.Parse(time.RFC3339Nano, timestampResponse)

	// Using Basic Auth for test so initial call isn't actually made
	respHandle := TimestampResponseHandle{Format: handle.TEXTPLAIN}
	err := timestamp(client, &respHandle)
	resp := respHandle.Get()
	if err != nil {
		t.Errorf("Error = %v", err)
	} else if resp == nil {
		t.Errorf("No response found")
	} else if !want.Equal(*resp) {
		t.Errorf("Timestamp Response = %+v, Want = %+v", spew.Sdump(*resp), spew.Sdump(want))
	}
}

var serverConfigResponse = `<host xmlns="http://marklogic.com/manage"><timestamp>2013-06-18T08:11:29.188561-07:00</timestamp><version>7.0</version><platform>linux</platform><edition>Essential Enterprise</edition><host-id>4808503609057420751</host-id><host-name>my-host.marklogic.com</host-name><bind-port>7999</bind-port><connect-port>7996</connect-port><foreign-bind-port>7998</foreign-bind-port><foreign-connect-port>7997</foreign-connect-port><ssl-certificate>...elided...</ssl-certificate></host>`

func TestServerConfig(t *testing.T) {
	client, server := test.AdminClient(serverConfigResponse)
	defer server.Close()
	timestamp, _ := time.Parse(time.RFC3339Nano, "2013-06-18T08:11:29.188561-07:00")
	timestampDateTimeNano, _ := util.NewDateTimeNano(timestamp)
	want :=
		ServerConfigResponse{
			XMLName:            xml.Name{"http://marklogic.com/manage", "host"},
			Timestamp:          *timestampDateTimeNano,
			Version:            "7.0",
			Platform:           "linux",
			Edition:            "Essential Enterprise",
			HostId:             "4808503609057420751",
			HostName:           "my-host.marklogic.com",
			BindPort:           7999,
			ConnectPort:        7996,
			ForeignBindPort:    7998,
			ForeignConnectPort: 7997,
			SslCertificate:     "...elided...",
		}

	// Using Basic Auth for test so initial call isn't actually made
	respHandle := ServerConfigResponseHandle{Format: handle.XML}
	err := serverConfig(client, &respHandle)
	resp := respHandle.Get()
	if err != nil {
		t.Errorf("Error = %v", err)
	} else if resp == nil {
		t.Errorf("No response found")
	} else if !reflect.DeepEqual(*resp, want) {
		t.Errorf("ServerConfig Response = %+v, Want = %+v", spew.Sdump(*resp), spew.Sdump(want))
	}
}
