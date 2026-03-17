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

var searchAutocompleteCompaniesV1 = cli.Command{
	Name:    "autocomplete-companies-v1",
	Usage:   "Autocomplete company search",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "query",
			Usage:     "Text search query to find companies by name.\nExample: \"Descartes Technologies UG\"\n",
			Required:  true,
			QueryPath: "query",
		},
	},
	Action:          handleSearchAutocompleteCompaniesV1,
	HideHelpCommand: true,
}

var searchFindCompaniesV1 = requestflag.WithInnerFlags(cli.Command{
	Name:    "find-companies-v1",
	Usage:   "Search for companies",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]any]{
			Name:     "filter",
			Usage:    "Filters to filter companies.\n",
			BodyPath: "filters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "location",
			Usage:    "Location to filter companies.\n",
			BodyPath: "location",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pagination",
			BodyPath: "pagination",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "query",
			Usage:    "Search query to filter companies.\n",
			BodyPath: "query",
		},
	},
	Action:          handleSearchFindCompaniesV1,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"location": {
		&requestflag.InnerFlag[float64]{
			Name:       "location.latitude",
			Usage:      "Latitude to filter on.\n",
			InnerField: "latitude",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "location.longitude",
			Usage:      "Longitude to filter on.\n",
			InnerField: "longitude",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "location.radius",
			Usage:      "Radius in kilometers to filter on.\nExample: 10\n",
			InnerField: "radius",
		},
	},
	"pagination": {
		&requestflag.InnerFlag[int64]{
			Name:       "pagination.page",
			Usage:      "Page number to return.\n",
			InnerField: "page",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pagination.per-page",
			Usage:      "Number of results per page.\n",
			InnerField: "per_page",
		},
	},
	"query": {
		&requestflag.InnerFlag[string]{
			Name:       "query.value",
			Usage:      "Search query to filter companies.\n",
			InnerField: "value",
		},
	},
})

var searchFindPersonV1 = requestflag.WithInnerFlags(cli.Command{
	Name:    "find-person-v1",
	Usage:   "Search for people",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]any]{
			Name:     "filter",
			Usage:    "Filters to filter people.\n",
			BodyPath: "filters",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "pagination",
			BodyPath: "pagination",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "query",
			Usage:    "Search query to filter people.\n",
			BodyPath: "query",
		},
	},
	Action:          handleSearchFindPersonV1,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pagination": {
		&requestflag.InnerFlag[int64]{
			Name:       "pagination.page",
			Usage:      "Page number to return.\n",
			InnerField: "page",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pagination.per-page",
			Usage:      "Number of results per page.\n",
			InnerField: "per_page",
		},
	},
	"query": {
		&requestflag.InnerFlag[string]{
			Name:       "query.value",
			Usage:      "Search query to filter people.\n",
			InnerField: "value",
		},
	},
})

var searchLookupCompanyByURL = cli.Command{
	Name:    "lookup-company-by-url",
	Usage:   "Find company by website URL",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Website URL to search for.\nExample: \"https://openregister.de\"\n",
			Required:  true,
			QueryPath: "url",
		},
	},
	Action:          handleSearchLookupCompanyByURL,
	HideHelpCommand: true,
}

func handleSearchAutocompleteCompaniesV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.SearchAutocompleteCompaniesV1Params{}

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
	_, err = client.Search.AutocompleteCompaniesV1(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search autocomplete-companies-v1", obj, format, transform)
}

func handleSearchFindCompaniesV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.SearchFindCompaniesV1Params{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.FindCompaniesV1(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search find-companies-v1", obj, format, transform)
}

func handleSearchFindPersonV1(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.SearchFindPersonV1Params{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search.FindPersonV1(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search find-person-v1", obj, format, transform)
}

func handleSearchLookupCompanyByURL(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.SearchLookupCompanyByURLParams{}

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
	_, err = client.Search.LookupCompanyByURL(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search lookup-company-by-url", obj, format, transform)
}
