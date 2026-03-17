// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/oregister/openregister-cli/internal/apiquery"
	"github.com/oregister/openregister-cli/internal/requestflag"
	"github.com/oregister/openregister-go/v2"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var documentGetCachedV1 = cli.Command{
	Name:    "get-cached-v1",
	Usage:   "Get document information",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "document-id",
			Required: true,
		},
	},
	Action:          handleDocumentGetCachedV1,
	HideHelpCommand: true,
}

var documentGetRealtimeV1 = cli.Command{
	Name:    "get-realtime-v1",
	Usage:   "Fetch a document in realtime.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "company-id",
			Required:  true,
			QueryPath: "company_id",
		},
		&requestflag.Flag[string]{
			Name:      "document-category",
			Required:  true,
			QueryPath: "document_category",
		},
	},
	Action:          handleDocumentGetRealtimeV1,
	HideHelpCommand: true,
}

func handleDocumentGetCachedV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("document-id") && len(unusedArgs) > 0 {
		cmd.Set("document-id", unusedArgs[0])
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
	_, err = client.Document.GetCachedV1(ctx, cmd.Value("document-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "document get-cached-v1", obj, format, transform)
}

func handleDocumentGetRealtimeV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.DocumentGetRealtimeV1Params{}

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
	_, err = client.Document.GetRealtimeV1(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "document get-realtime-v1", obj, format, transform)
}
