// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/oregister/openregister-cli/internal/mocktest"
)

func TestCompanyGetContactV0(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-contact-v0",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetDetailsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-details-v1",
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
			t,
			"--api-key", "string",
			"company", "get-financials-v1",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetHistoricalOwnersV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-historical-owners-v1",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetHoldingsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-holdings-v1",
			"--company-id", "company_id",
		)
	})
}

func TestCompanyGetOwnersV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-owners-v1",
			"--company-id", "company_id",
			"--best-available=true",
			"--export=true",
			"--realtime=true",
		)
	})
}

func TestCompanyGetUbosV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"company", "get-ubos-v1",
			"--company-id", "company_id",
		)
	})
}
