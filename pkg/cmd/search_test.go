// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/openregister-cli/internal/mocktest"
	"github.com/stainless-sdks/openregister-cli/internal/requestflag"
)

func TestSearchAutocompleteCompaniesV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "autocomplete-companies-v1",
			"--api-key", "string",
			"--query", "query",
		)
	})
}

func TestSearchFindCompaniesV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "find-companies-v1",
			"--api-key", "string",
			"--filter", "{keywords: [string], max: max, min: min, value: value, values: [string], field: status}",
			"--location", "{latitude: 0, longitude: 0, radius: 0}",
			"--pagination", "{page: 0, per_page: 0}",
			"--query", "{value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(searchFindCompaniesV1)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "search", "find-companies-v1",
			"--api-key", "string",
			"--filter", "{keywords: [string], max: max, min: min, value: value, values: [string], field: status}",
			"--location.latitude", "0",
			"--location.longitude", "0",
			"--location.radius", "0",
			"--pagination.page", "0",
			"--pagination.per-page", "0",
			"--query.value", "value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"filters:\n" +
			"  - keywords:\n" +
			"      - string\n" +
			"    max: max\n" +
			"    min: min\n" +
			"    value: value\n" +
			"    values:\n" +
			"      - string\n" +
			"    field: status\n" +
			"location:\n" +
			"  latitude: 0\n" +
			"  longitude: 0\n" +
			"  radius: 0\n" +
			"pagination:\n" +
			"  page: 0\n" +
			"  per_page: 0\n" +
			"query:\n" +
			"  value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "search", "find-companies-v1",
			"--api-key", "string",
		)
	})
}

func TestSearchFindPersonV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "find-person-v1",
			"--api-key", "string",
			"--filter", "{keywords: [string], max: max, min: min, value: value, values: [string], field: date_of_birth}",
			"--pagination", "{page: 0, per_page: 0}",
			"--query", "{value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(searchFindPersonV1)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "search", "find-person-v1",
			"--api-key", "string",
			"--filter", "{keywords: [string], max: max, min: min, value: value, values: [string], field: date_of_birth}",
			"--pagination.page", "0",
			"--pagination.per-page", "0",
			"--query.value", "value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"filters:\n" +
			"  - keywords:\n" +
			"      - string\n" +
			"    max: max\n" +
			"    min: min\n" +
			"    value: value\n" +
			"    values:\n" +
			"      - string\n" +
			"    field: date_of_birth\n" +
			"pagination:\n" +
			"  page: 0\n" +
			"  per_page: 0\n" +
			"query:\n" +
			"  value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "search", "find-person-v1",
			"--api-key", "string",
		)
	})
}

func TestSearchLookupCompanyByURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "search", "lookup-company-by-url",
			"--api-key", "string",
			"--url", "https://example.com",
		)
	})
}
