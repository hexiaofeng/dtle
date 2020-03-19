package fingerprint

import (
	"testing"

	"github.com/actiontech/dtle/client/config"
	"github.com/actiontech/dtle/helper/testlog"
	"github.com/actiontech/dtle/nomad/structs"
)

func TestHostFingerprint(t *testing.T) {
	f := NewHostFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	request := &FingerprintRequest{Config: &config.Config{}, Node: node}
	var response FingerprintResponse
	err := f.Fingerprint(request, &response)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !response.Detected {
		t.Fatalf("expected response to be applicable")
	}

	if len(response.Attributes) == 0 {
		t.Fatalf("should generate a diff of node attributes")
	}

	// Host info
	for _, key := range []string{"os.name", "os.version", "unique.hostname", "kernel.name"} {
		assertNodeAttributeContains(t, response.Attributes, key)
	}
}