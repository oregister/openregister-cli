// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/oregister/openregister-cli/internal/mocktest"
)

func TestDocumentGetCachedV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "document", "get-cached-v1",
			"--api-key", "string",
			"--document-id", "document_id",
		)
	})
}

func TestDocumentGetRealtimeV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "document", "get-realtime-v1",
			"--api-key", "string",
			"--company-id", "company_id",
			"--document-category", "current_printout",
		)
	})
}
