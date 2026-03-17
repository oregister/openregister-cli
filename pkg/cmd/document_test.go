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
			t,
			"--api-key", "string",
			"document", "get-cached-v1",
			"--document-id", "document_id",
		)
	})
}

func TestDocumentGetRealtimeV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"document", "get-realtime-v1",
			"--company-id", "company_id",
			"--document-category", "current_printout",
		)
	})
}
