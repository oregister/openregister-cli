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

var transparenzregisterRequestCreateV1 = cli.Command{
	Name:    "create-v1",
	Usage:   "Submit a Transparenzregister request for a company using its company ID. This\nendpoint will initiate the Transparenzregister request process and return a\nrequest ID for tracking.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "company-id",
			Usage:     "Unique company identifier. Required unless using X-Credential-Label=test.\nExample: DE-HRB-F1103-267645\n",
			QueryPath: "company_id",
		},
		&requestflag.Flag[string]{
			Name:       "x-credential-label",
			Default:    "default",
			HeaderPath: "X-Credential-Label",
		},
	},
	Action:          handleTransparenzregisterRequestCreateV1,
	HideHelpCommand: true,
}

var transparenzregisterRequestGetV1 = cli.Command{
	Name:    "get-v1",
	Usage:   "Get the results of a Transparenzregister request. This endpoint handles all\ninternal complexity including polling request status, selecting all available\ndocuments, creating Transparenzregister baskets, and returning download URLs\nwhen ready. If the request is still processing, it will return a pending status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "request-id",
			Required: true,
		},
	},
	Action:          handleTransparenzregisterRequestGetV1,
	HideHelpCommand: true,
}

func handleTransparenzregisterRequestCreateV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.TransparenzregisterRequestNewV1Params{}

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
	_, err = client.Transparenzregister.Request.NewV1(ctx, params, options...)
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
		Title:          "transparenzregister:request create-v1",
		Transform:      transform,
	})
}

func handleTransparenzregisterRequestGetV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("request-id") && len(unusedArgs) > 0 {
		cmd.Set("request-id", unusedArgs[0])
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
	_, err = client.Transparenzregister.Request.GetV1(ctx, cmd.Value("request-id").(string), options...)
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
		Title:          "transparenzregister:request get-v1",
		Transform:      transform,
	})
}
