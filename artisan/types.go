package main

import "fmt"

type BuildArg struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (b *BuildArg) String() string {
	return fmt.Sprintf("%s=%s", b.Name, b.Value)
}

type Label struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (l *Label) String() string {
	return fmt.Sprintf("%s=%s", l.Name, l.Value)
}

type ContainerfileName string

type Containerfile struct {
	ContextPath string              `json:"contextPath,omitempty"`
	Path        string              `json:"path,omitempty"`
	Tags        []string            `json:"tags,omitempty"`
	Parents     []ContainerfileName `json:"parents,omitempty"`
	BuildArgs   []BuildArg          `json:"buildArgs,omitempty"`
	Labels      []Label             `json:"labels,omitempty"`
	Transient   bool                `json:"transient,omitempty"`
}

type ContainerfileSpec struct {
	Containerfiles []Containerfile `json:"containerfiles"`
}
