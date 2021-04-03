package crypto_service

import (
	"strings"
	"testing"
)

// test function
func TestSampleSearch(t *testing.T) {
	testResult, err := SearchByName("bitcoin")
	if err != nil {
		t.Errorf("Search currencies returned an error")
	}
	if len(testResult) > 0 {
		target := testResult[0]
		if strings.ToUpper(target.Name) != "BITCOIN" {
			t.Errorf("Search for target failed.")
		}
	}
}
