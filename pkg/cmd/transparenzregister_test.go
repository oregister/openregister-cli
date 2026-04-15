// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/oregister/openregister-cli/internal/mocktest"
)

func TestTransparenzregisterSetCredentialsV1(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"transparenzregister", "set-credentials-v1",
			"--password", "password",
			"--username", "username",
			"--credential-label", "credential_label",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"password: password\n" +
			"username: username\n" +
			"credential_label: credential_label\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"transparenzregister", "set-credentials-v1",
		)
	})
}
