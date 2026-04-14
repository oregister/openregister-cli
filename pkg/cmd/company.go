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

var companyGetContactV0 = cli.Command{
	Name:    "get-contact-v0",
	Usage:   "Get company contact information",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
	},
	Action:          handleCompanyGetContactV0,
	HideHelpCommand: true,
}

var companyGetDetailsV1 = cli.Command{
	Name:    "get-details-v1",
	Usage:   "Get detailed company information",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "export",
			Usage:     "Setting this to true will return the company without sources.\n",
			QueryPath: "export",
		},
		&requestflag.Flag[bool]{
			Name:      "realtime",
			Usage:     "Get the most up-to-date company information directly from the Handelsregister.\nWhen set to true, we fetch the latest data in real-time from the official German commercial register, ensuring you receive the most current company details.\nNote: Real-time requests take longer but guarantee the freshest data available.\n",
			QueryPath: "realtime",
		},
	},
	Action:          handleCompanyGetDetailsV1,
	HideHelpCommand: true,
}

var companyGetFinancialsV1 = cli.Command{
	Name:    "get-financials-v1",
	Usage:   "Get financial reports",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
	},
	Action:          handleCompanyGetFinancialsV1,
	HideHelpCommand: true,
}

var companyGetHistoricalOwnersV1 = cli.Command{
	Name:    "get-historical-owners-v1",
	Usage:   "Get historical owner changes",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
	},
	Action:          handleCompanyGetHistoricalOwnersV1,
	HideHelpCommand: true,
}

var companyGetHoldingsV1 = cli.Command{
	Name:    "get-holdings-v1",
	Usage:   "Get company holdings",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
	},
	Action:          handleCompanyGetHoldingsV1,
	HideHelpCommand: true,
}

var companyGetOwnersV1 = cli.Command{
	Name:    "get-owners-v1",
	Usage:   "Get company owners",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:      "best-available",
			Usage:     "When set to true, returns the best available owner data for AG and SE companies.\nThis data is extracted from Handelsregister documents and may not reflect the most current ownership state, as these\ndocument types are not filed on every ownership change.\nRequests for AG/SE companies without this flag return 404.\nNote: realtime and best_available cannot be used together at the moment.\n",
			QueryPath: "best_available",
		},
		&requestflag.Flag[bool]{
			Name:      "export",
			Usage:     "Setting this to true will return the owners of the company if they exist\nbut will skip processing the documents in case they weren't processed yet.\n",
			QueryPath: "export",
		},
		&requestflag.Flag[bool]{
			Name:      "realtime",
			Usage:     "Get the most up-to-date company information directly from the Handelsregister.\nWhen set to true, we fetch the latest data in real-time from the official German commercial register, ensuring you receive the most current company details.\nNote: Real-time requests take longer but guarantee the freshest data available.\n",
			QueryPath: "realtime",
		},
	},
	Action:          handleCompanyGetOwnersV1,
	HideHelpCommand: true,
}

var companyGetUbosV1 = cli.Command{
	Name:    "get-ubos-v1",
	Usage:   "Get company end owners",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "company-id",
			Required: true,
		},
	},
	Action:          handleCompanyGetUbosV1,
	HideHelpCommand: true,
}

func handleCompanyGetContactV0(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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
	_, err = client.Company.GetContactV0(ctx, cmd.Value("company-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-contact-v0", obj, format, transform)
}

func handleCompanyGetDetailsV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.CompanyGetDetailsV1Params{}

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
	_, err = client.Company.GetDetailsV1(
		ctx,
		cmd.Value("company-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-details-v1", obj, format, transform)
}

func handleCompanyGetFinancialsV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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
	_, err = client.Company.GetFinancialsV1(ctx, cmd.Value("company-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-financials-v1", obj, format, transform)
}

func handleCompanyGetHistoricalOwnersV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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
	_, err = client.Company.GetHistoricalOwnersV1(ctx, cmd.Value("company-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-historical-owners-v1", obj, format, transform)
}

func handleCompanyGetHoldingsV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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
	_, err = client.Company.GetHoldingsV1(ctx, cmd.Value("company-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-holdings-v1", obj, format, transform)
}

func handleCompanyGetOwnersV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.CompanyGetOwnersV1Params{}

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
	_, err = client.Company.GetOwnersV1(
		ctx,
		cmd.Value("company-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-owners-v1", obj, format, transform)
}

func handleCompanyGetUbosV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("company-id") && len(unusedArgs) > 0 {
		cmd.Set("company-id", unusedArgs[0])
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
	_, err = client.Company.GetUbosV1(ctx, cmd.Value("company-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "company get-ubos-v1", obj, format, transform)
}
