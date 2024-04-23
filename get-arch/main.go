package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

func getCustomFlagName(arch string) string {
	return fmt.Sprintf("custom-%s", arch)
}

func getArchesWithCustomFlag() map[string][]string {
	arches := getArches()

	for arch := range arches {
		arches[arch] = append(arches[arch], getCustomFlagName(arch))
	}

	return arches
}

func getArches() map[string][]string {
	return map[string][]string{
		"amd64": []string{"amd64", "x86_64", "x86-64"},
		"arm64": []string{"aarch64", "arm64"},
	}
}

func getFlagsFromArches() []cli.Flag {
	out := []cli.Flag{}

	for arch, subArches := range getArches() {
		archCategory := fmt.Sprintf("Options for %s:", arch)

		for _, subArch := range subArches {
			archFlag := &cli.BoolFlag{
				Name:     subArch,
				Category: archCategory,
				Value:    false,
				Usage:    fmt.Sprintf("use to select %s", subArch),
			}

			out = append(out, archFlag)
		}

		customArchFlag := &cli.StringFlag{
			Name:     getCustomFlagName(arch),
			Category: archCategory,
			Usage:    fmt.Sprintf("use to provide a custom architecture for %s", arch),
		}

		out = append(out, customArchFlag)
	}

	return out
}

func isMoreThanOneFlagForArchSet(cliCtx *cli.Context, arch string) bool {
	foundCount := 0

	for _, subArch := range getArchesWithCustomFlag()[arch] {
		if cliCtx.IsSet(subArch) {
			foundCount += 1
		}
	}

	if foundCount == 0 {
		return false
	}

	return foundCount == 1
}

func isFlagForArchSet(cliCtx *cli.Context, arch string) bool {
	for _, subArch := range getArchesWithCustomFlag()[arch] {
		if cliCtx.IsSet(subArch) {
			return true
		}
	}

	return false
}

func validateFlags(cliCtx *cli.Context) error {
	for arch := range getArches() {
		if !isFlagForArchSet(cliCtx, arch) {
			return fmt.Errorf("no flag set for architecture %q", arch)
		}
	}

	for arch := range getArches() {
		if !isMoreThanOneFlagForArchSet(cliCtx, arch) {
			return fmt.Errorf("only one flag for architecture %q can be set", arch)
		}
	}

	return nil
}

func lookupArchForFlag(cliCtx *cli.Context, systemArch string) (string, error) {
	arches := getArches()
	if _, ok := arches[systemArch]; !ok {
		return "", fmt.Errorf("unknown architecture %s", systemArch)
	}

	custom := getCustomFlagName(systemArch)
	if cliCtx.IsSet(custom) {
		return cliCtx.String(custom), nil
	}

	for _, arch := range arches[systemArch] {
		if cliCtx.IsSet(arch) {
			return arch, nil
		}
	}

	return "", fmt.Errorf("unexpected error")
}

func main() {
	app := &cli.App{
		Name:  "get-arch",
		Usage: "Select the correct arch for your download.",
		Flags: getFlagsFromArches(),
		Action: func(cliCtx *cli.Context) error {
			if err := validateFlags(cliCtx); err != nil {
				return err
			}

			arch, err := lookupArchForFlag(cliCtx, runtime.GOARCH)
			if err != nil {
				return err
			}

			fmt.Println(arch)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
