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

var monitorCreate = cli.Command{
	Name:    "create",
	Usage:   "Create webhook monitor item",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Usage:    "For `company` this is the register ID (e.g. `DE-HRB-F1103-267645`).\nFor `person` this is the person UUID.\n",
			Required: true,
			BodyPath: "entity_id",
		},
		&requestflag.Flag[string]{
			Name:     "entity-type",
			Usage:    "Type of the entity to monitor.\n",
			Required: true,
			BodyPath: "entity_type",
		},
		&requestflag.Flag[[]string]{
			Name:     "preference",
			Usage:    "Preferences for the entity to monitor.\nUse `WebhookMonitorCompanyPreference` values when `entity_type` is `company`,\nand `WebhookMonitorPersonPreference` values when `entity_type` is `person`.\n",
			Required: true,
			BodyPath: "preferences",
		},
	},
	Action:          handleMonitorCreate,
	HideHelpCommand: true,
}

var monitorList = cli.Command{
	Name:            "list",
	Usage:           "List webhook monitor items",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleMonitorList,
	HideHelpCommand: true,
}

var monitorDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete webhook monitor item",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "entity-id",
			Required: true,
		},
	},
	Action:          handleMonitorDelete,
	HideHelpCommand: true,
}

func handleMonitorCreate(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := openregister.MonitorNewParams{}

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
	_, err = client.Monitor.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "monitor create", obj, format, explicitFormat, transform)
}

func handleMonitorList(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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
	_, err = client.Monitor.List(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, os.Stderr, "monitor list", obj, format, explicitFormat, transform)
}

func handleMonitorDelete(ctx context.Context, cmd *cli.Command) error {
	client := openregister.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("entity-id") && len(unusedArgs) > 0 {
		cmd.Set("entity-id", unusedArgs[0])
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

	return client.Monitor.Delete(ctx, cmd.Value("entity-id").(string), options...)
}
