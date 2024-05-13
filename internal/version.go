/*
Copyright 2024 Richard Kosegi

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

package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"runtime"
)

var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
	GoOS      = runtime.GOOS
	GoArch    = runtime.GOARCH
)

func newVersionCommand(out io.Writer) *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Display tool version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(out, "Version: %s / %s / %s\n", Version, Branch, Revision)
			fmt.Fprintf(out, "Build by: %s\n", BuildUser)
			fmt.Fprintf(out, "Build at: %s\n", BuildDate)
			fmt.Fprintf(out, "Go : %s / %s / %s\n", GoVersion, GoOS, GoArch)
			return nil
		},
	}
	return c
}
