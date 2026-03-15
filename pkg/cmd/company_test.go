// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/openregister-cli/internal/mocktest"
)

func TestCompanyGetContactV0(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-contact-v0",
			"--api-key", "string",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetDetailsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-details-v1",
			"--api-key", "string",
			"--company-id", "company_id",
			"--export=true",
			"--realtime=true",
		)
	})
}

func TestCompanyGetFinancialsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-financials-v1",
			"--api-key", "string",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetHistoricalOwnersV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-historical-owners-v1",
			"--api-key", "string",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetHoldingsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-holdings-v1",
			"--api-key", "string",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetOwnersV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-owners-v1",
			"--api-key", "string",
			"--company-id", "company_id",
			"--export=true",
			"--realtime=true",
		)
	})
}

func TestCompanyGetUbosV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "company", "get-ubos-v1",
			"--api-key", "string",
			"--company-id", "company_id",
		)
	})
}
