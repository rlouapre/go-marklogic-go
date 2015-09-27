package admin

import (
	"bytes"
	"encoding/xml"
	"net/http"

	clients "github.com/ryanjdew/go-marklogic-go/clients"
	handle "github.com/ryanjdew/go-marklogic-go/handle"
	util "github.com/ryanjdew/go-marklogic-go/util"
)

// ServerConfigResponse is a handle that places the results into
// a Response struct
type ServerConfigResponseHandle struct {
	*bytes.Buffer
	Format   int
	response ServerConfigResponse
}

type ServerConfigResponse struct {
	XMLName            xml.Name          `xml:"http://marklogic.com/manage host"`
	Timestamp          util.DateTimeNano `xml:"timestamp"`
	Version            string            `xml:"version"`
	Platform           string            `xml:"platform"`
	Edition            string            `xml:"edition"`
	HostId             string            `xml:"host-id"`
	HostName           string            `xml:"host-name"`
	BindPort           int64             `xml:"bind-port"`
	ConnectPort        int64             `xml:"connect-port"`
	ForeignBindPort    int64             `xml:"foreign-bind-port"`
	ForeignConnectPort int64             `xml:"foreign-connect-port"`
	SslCertificate     string            `xml:"ssl-certificate"`
}

// GetFormat returns int that represents XML
func (rh *ServerConfigResponseHandle) GetFormat() int {
	return rh.Format
}

func (rh *ServerConfigResponseHandle) resetBuffer() {
	if rh.Buffer == nil {
		rh.Buffer = new(bytes.Buffer)
	}
	rh.Reset()
}

// Deserialize returns Response struct that represents XML
func (rh *ServerConfigResponseHandle) Deserialize(bytes []byte) {
	rh.resetBuffer()
	rh.Write(bytes)
	rh.response = ServerConfigResponse{}
	if rh.GetFormat() == handle.XML {
		xml.Unmarshal(bytes, &rh.response)
	}
}

// AcceptResponse handles an *http.Response
func (rh *ServerConfigResponseHandle) AcceptResponse(resp *http.Response) error {
	return handle.CommonHandleAcceptResponse(rh, resp)
}

// Serialize returns []byte of XML that represents the Response struct
func (rh *ServerConfigResponseHandle) Serialize(response interface{}) {
	rh.response = response.(ServerConfigResponse)
	rh.resetBuffer()
	if rh.GetFormat() == handle.XML {
		enc := xml.NewEncoder(rh.Buffer)
		enc.Encode(&rh.response)
	}
}

// Get returns string of XML
func (rh *ServerConfigResponseHandle) Get() *ServerConfigResponse {
	return &rh.response
}

// Serialized returns string of XML
func (rh *ServerConfigResponseHandle) Serialized() string {
	rh.Serialize(rh.response)
	return rh.String()
}

// Retrieve MarkLogic Server configuration information, suitable for use in joining a cluster.
// https://docs.marklogic.com/REST/GET/admin/v1/server-config
func serverConfig(ac *clients.AdminClient, response handle.ResponseHandle) error {
	req, err := util.BuildRequestFromHandle(ac, "GET", "/server-config", nil)
	if err != nil {
		return err
	}
	return util.Execute(ac, req, response)
}
