package common

import (
	"os"
	"strings"

	"knative.dev/func/pkg/openshift"
)

// Intended to provide setup configuration for E2E tests
const (
	DefaultRegistry = "localhost:50000/user"
)

var testRegistry = ""

func init() {
	// Setup test Registry.
	testRegistry = os.Getenv("E2E_REGISTRY_URL")
	if testRegistry == "" || testRegistry == "default" {
		if openshift.IsOpenShift() {
			testRegistry = openshift.GetDefaultRegistry()
		} else {
			testRegistry = DefaultRegistry
		}
	}
}

// GetRegistry returns registry
func GetRegistry() string {
	return testRegistry
}

// GetFuncBinaryPath should return the Path of 'func' binary under test
func GetFuncBinaryPath() string {
	return GetOsEnvOrDefault("E2E_FUNC_BIN_PATH", "")
}

// GetRuntime returns the runtime that should be tested.
func GetRuntime() string {
	return GetOsEnvOrDefault("E2E_RUNTIME", "node")
}

// IsUseKnFunc indicates that tests should be run against "kn func" instead of "func" binary
func IsUseKnFunc() bool {
	return strings.EqualFold(os.Getenv("E2E_USE_KN_FUNC"), "true")
}

func GetOsEnvOrDefault(env string, dflt string) string {
	e := os.Getenv(env)
	if e == "" {
		return dflt
	}
	return e
}
