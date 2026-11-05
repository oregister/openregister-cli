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

var jobCreateV1 = cli.Command{
	Name:    "create-v1",
	Usage:   "Start a realtime request without waiting for it. The request returns at once\nwith a pending job. Use this instead of `realtime=true` when your client cannot\nkeep a request open for the minutes the Handelsregister may take.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Usage:    "Unique company identifier.\nExample: DE-HRB-F1103-267645\n",
			Required: true,
			BodyPath: "company_id",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "company", "owners", "document".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "document-category",
			Usage:    `Allowed values: "current_printout", "chronological_printout", "historical_printout", "structured_information", "shareholder_list", "articles_of_association".`,
			BodyPath: "document_category",
		},
	},
	Action:          handleJobCreateV1,
	HideHelpCommand: true,
}

var jobGetV1 = cli.Command{
	Name:    "get-v1",
	Usage:   "Get a job's status and, once it completed, its result in `results`. Polling is\nfree.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "job-id",
			Required:  true,
			PathParam: "job_id",
		},
	},
	Action:          handleJobGetV1,
	HideHelpCommand: true,
}

func handleJobCreateV1(ctx context.Context, cmd *cli.Command) error {
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

	params := openregister.JobNewV1Params{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Job.NewV1(ctx, params, options...)
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
		Title:          "job create-v1",
		Transform:      transform,
	})
}

func handleJobGetV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("job-id") && len(unusedArgs) > 0 {
		cmd.Set("job-id", unusedArgs[0])
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
	_, err = client.Job.GetV1(ctx, cmd.Value("job-id").(string), options...)
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
		Title:          "job get-v1",
		Transform:      transform,
	})
}
