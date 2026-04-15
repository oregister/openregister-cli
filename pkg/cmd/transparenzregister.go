// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/oregister/openregister-cli/internal/apiquery"
	"github.com/oregister/openregister-cli/internal/requestflag"
	"github.com/oregister/openregister-go/v2"
	"github.com/urfave/cli/v3"
)

var transparenzregisterSetCredentialsV1 = cli.Command{
	Name:    "set-credentials-v1",
	Usage:   "Store username and password credentials for accessing the Transparenzregister\nAPI. These credentials will be used for subsequent requests to retrieve company\ndocuments.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "password",
			Usage:    "Password for Transparenzregister API access.\n",
			Required: true,
			BodyPath: "password",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			Usage:    "Username for Transparenzregister API access.\nExample: \"testnutzer-eis@transparenzregister.de\"\n",
			Required: true,
			BodyPath: "username",
		},
		&requestflag.Flag[string]{
			Name:     "credential-label",
			Usage:    "Label to identify this set of credentials. Allows storing multiple\nTransparenzregister credentials per user (e.g., for different accounts\nor clients). Defaults to 'default' if not provided.\nExample: \"client_a\"\n",
			Default:  "default",
			BodyPath: "credential_label",
		},
	},
	Action:          handleTransparenzregisterSetCredentialsV1,
	HideHelpCommand: true,
}

func handleTransparenzregisterSetCredentialsV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.TransparenzregisterSetCredentialsV1Params{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Transparenzregister.SetCredentialsV1(ctx, params, options...)
}
