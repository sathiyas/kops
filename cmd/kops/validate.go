/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/cobra"
)

// ValidateCmd represents the get command
type ValidateCmd struct {
	output string

	cobraCommand *cobra.Command
}

var validateCmd = ValidateCmd{
	cobraCommand: &cobra.Command{
		Use:        "validate",
		SuggestFor: []string{"list"},
		Short:      "validate a kubernetes cluster",
		Long:       `validate a kubernetes cluster`,
	},
}

func init() {
	cmd := validateCmd.cobraCommand

	rootCommand.AddCommand(cmd)

	cmd.PersistentFlags().StringVarP(&validateCmd.output, "output", "o", OutputTable, "output format.  One of: table, yaml")
}
