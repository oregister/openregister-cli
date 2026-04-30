// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/oregister/openregister-cli/internal/apiquery"
	"github.com/oregister/openregister-cli/internal/requestflag"
	"github.com/oregister/openregister-go/v2"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var transparenzregisterExtractCreateV1 = cli.Command{
	Name:    "create-v1",
	Usage:   "Submit a Transparenzregister extract request and return an extract resource with\nprocessing status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Usage:    "Unique company identifier.\nRequired unless `X-Credential-Name` is set to `sandbox`.\nIn sandbox mode this field should be omitted.\nExample: DE-HRB-F1103-267645\n",
			BodyPath: "company_id",
		},
		&requestflag.Flag[string]{
			Name:       "x-credential-name",
			Default:    "default",
			HeaderPath: "X-Credential-Name",
		},
	},
	Action:          handleTransparenzregisterExtractCreateV1,
	HideHelpCommand: true,
}

var transparenzregisterExtractGetV1 = cli.Command{
	Name:    "get-v1",
	Usage:   "Get the results of a Transparenzregister extract request. This endpoint handles\nall internal complexity including polling request status, selecting all\navailable documents, creating Transparenzregister baskets, and returning\ndownload URLs when ready. If the request is still processing, it will return a\npending status. Polling reuses the credential mode stored on the extract at\ncreate time. Sandbox extracts keep using the Transparenzregister test client\nautomatically; no credential header is accepted on this endpoint.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "extract-id",
			Required:  true,
			PathParam: "extract_id",
		},
	},
	Action:          handleTransparenzregisterExtractGetV1,
	HideHelpCommand: true,
}

func handleTransparenzregisterExtractCreateV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := openregister.TransparenzregisterExtractNewV1Params{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Transparenzregister.Extract.NewV1(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "transparenzregister:extract create-v1",
		Transform:      transform,
	})
}

func handleTransparenzregisterExtractGetV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("extract-id") && len(unusedArgs) > 0 {
		cmd.Set("extract-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Transparenzregister.Extract.GetV1(ctx, cmd.Value("extract-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "transparenzregister:extract get-v1",
		Transform:      transform,
	})
}
