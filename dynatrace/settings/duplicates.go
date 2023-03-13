package settings

import (
	"os"
	"strings"
)

const DYNATRACE_DUPLICATE_REJECT = "DYNATRACE_DUPLICATE_REJECT"
const DYNATRACE_DUPLICATE_HIJACK = "DYNATRACE_DUPLICATE_HIJACK"

func RejectDuplicate(resourceNames ...string) bool {
	return envVarContains(os.Getenv(DYNATRACE_DUPLICATE_REJECT), resourceNames...)
}

func HijackDuplicate(resourceNames ...string) bool {
	return envVarContains(os.Getenv(DYNATRACE_DUPLICATE_HIJACK), resourceNames...)
}

func envVarContains(envVar string, search ...string) bool {
	svalues := os.Getenv(envVar)
	if len(svalues) == 0 {
		return false
	}
	values := strings.Split(svalues, ",")
	for _, value := range values {
		value = strings.TrimSpace(value)
		for _, searchValue := range search {
			if value == searchValue {
				return true
			}
		}
	}
	return false
}
