package admin

import (
	"testing"
	"time"

	handle "github.com/ryanjdew/go-marklogic-go/handle"
)

func TestTimestampResponseHandleDeserialize(t *testing.T) {
	tm := `2013-05-15T10:34:38.932514-07:00`
	want, err := time.Parse(time.RFC3339Nano, tm)
	result := TimestampResponseHandle{Format: handle.TEXTPLAIN}
	result.Deserialize([]byte(tm))
	if err != nil {
		t.Errorf("Error = %v", err)
	} else if !result.timestamp.Equal(want) {
		t.Errorf("Not equal - TimestampResponseHandle timestamp = %+v, Want = %+v", result.timestamp, want)
	}
}
