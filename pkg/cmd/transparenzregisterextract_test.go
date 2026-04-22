// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/oregister/openregister-cli/internal/mocktest"
)

func TestTransparenzregisterExtractCreateV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transparenzregister:extract", "create-v1",
			"--company-id", "company_id",
			"--x-credential-name", "X-Credential-Name",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("company_id: company_id")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transparenzregister:extract", "create-v1",
			"--x-credential-name", "X-Credential-Name",
		)
	})
}

func TestTransparenzregisterExtractGetV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transparenzregister:extract", "get-v1",
			"--extract-id", "extract_id",
		)
	})
}
