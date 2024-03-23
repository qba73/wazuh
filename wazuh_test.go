package wazuh_test

import (
	"testing"

	"github.com/qba73/wazuh"
)

func TestCreateWazuhClient(t *testing.T) {
	t.Parallel()

	_, err := wazuh.NewClient("apiKey")
	if err != nil {
		t.Fatal(err)
	}
}
