// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/oregister/openregister-cli/internal/autocomplete"
	"github.com/oregister/openregister-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "openregister",
		Usage:     "CLI for the openregister API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.BoolFlag{
				Name:    "raw-output",
				Aliases: []string{"r"},
				Usage:   "If the result is a string, print it without JSON quotes. This can be useful for making output transforms talk to non-JSON-based systems.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Usage:   "API Key Authentication\nProvide your API key as a Bearer token in the Authorization header.\n",
				Sources: cli.EnvVars("OPENREGISTER_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchAutocompleteCompaniesV1,
					&searchFindCompaniesV1,
					&searchFindPersonV1,
					&searchLookupCompanyByURL,
				},
			},
			{
				Name:     "company",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&companyGetContactV0,
					&companyGetDetailsV1,
					&companyGetFinancialsV1,
					&companyGetHistoricalOwnersV1,
					&companyGetHoldingsV1,
					&companyGetOwnersV1,
					&companyGetUbosV1,
				},
			},
			{
				Name:     "document",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&documentGetCachedV1,
					&documentGetRealtimeV1,
				},
			},
			{
				Name:     "person",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&personGetDetailsV1,
					&personGetHoldingsV1,
				},
			},
			{
				Name:     "monitor",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&monitorCreate,
					&monitorList,
					&monitorDelete,
				},
			},
			{
				Name:     "transparenzregister",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transparenzregisterSetCredentialsV1,
				},
			},
			{
				Name:     "transparenzregister:request",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&transparenzregisterRequestCreateV1,
					&transparenzregisterRequestGetV1,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "openregister @manpages [-o openregister.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "openregister.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "openregister.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
