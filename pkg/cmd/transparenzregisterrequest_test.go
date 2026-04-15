// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/oregister/openregister-cli/internal/mocktest"
)

func TestTransparenzregisterRequestCreateV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transparenzregister:request", "create-v1",
			"--company-id", "company_id",
			"--x-credential-label", "X-Credential-Label",
		)
	})
}

func TestTransparenzregisterRequestGetV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transparenzregister:request", "get-v1",
			"--request-id", "request_id",
		)
	})
}
