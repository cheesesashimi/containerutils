package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"
	"sigs.k8s.io/yaml"
)

// TODO: Use a golang library to get this natively without shelling out.
func getGitRev() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	outBuf := bytes.NewBuffer([]byte{})
	cmd.Stdout = outBuf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("could not get git rev: %w", err)
	}

	return outBuf.String(), nil
}

func getCommonLabels() ([]Label, error) {
	gitRev, err := getGitRev()
	if err != nil {
		return []Label{}, err
	}

	labels := []Label{
		{
			Name: "org.opencontainers.image.url",
			// TODO: Make this a spec value.
			Value: "https://github.com/cheesesashimi/containerfiles",
		},
		{
			Name:  "org.opencontainers.image.revision",
			Value: gitRev,
		},
	}

	return labels, nil
}

func dumpContainerfilesSpec(path string) error {
	contents, err := os.ReadFile("containerfiles-spec.yaml")
	if err != nil {
		return err
	}

	out := &ContainerfileSpec{}

	if err := yaml.Unmarshal(contents, out); err != nil {
		return err
	}

	commonLabels, err := getCommonLabels()
	if err != nil {
		return err
	}

	for i, containerfile := range out.Containerfiles {
		out.Containerfiles[i].Labels = append(containerfile.Labels, commonLabels...)
	}

	spew.Dump(out)
	return nil
}

func main() {
	if err := dumpContainerfilesSpec("containerfiles-spec.yaml"); err != nil {
		panic(err)
	}
}
