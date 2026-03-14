// Copyright (C) 2019-2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	"github.com/rink-labs/rinkgo/tests/fixture/stacktrace"
	"github.com/rink-labs/rinkgo/tests/fixture/tmpnet"
)

const (
	processRuntime   = "process"
	processDocPrefix = "[process runtime] "

	rinkgoPathFlag = "rinkgo-path"
)

var errRinkGoRequired = fmt.Errorf("--%s or %s are required", rinkgoPathFlag, tmpnet.RinkGoPathEnvName)

type processRuntimeVars struct {
	config tmpnet.ProcessRuntimeConfig
}

func (v *processRuntimeVars) registerWithFlag() {
	v.register(flag.StringVar, flag.BoolVar)
}

func (v *processRuntimeVars) registerWithFlagSet(flagSet *pflag.FlagSet) {
	v.register(flagSet.StringVar, flagSet.BoolVar)
}

func (v *processRuntimeVars) register(stringVar varFunc[string], boolVar varFunc[bool]) {
	stringVar(
		&v.config.RinkGoPath,
		rinkgoPathFlag,
		os.Getenv(tmpnet.RinkGoPathEnvName),
		processDocPrefix+fmt.Sprintf(
			"The rinkgo executable path. Also possible to configure via the %s env variable.",
			tmpnet.RinkGoPathEnvName,
		),
	)
	stringVar(
		&v.config.PluginDir,
		"plugin-dir",
		tmpnet.GetEnvWithDefault(tmpnet.RinkGoPluginDirEnvName, os.ExpandEnv("$HOME/.rinkgo/plugins")),
		processDocPrefix+fmt.Sprintf(
			"The dir containing VM plugins. Also possible to configure via the %s env variable.",
			tmpnet.RinkGoPluginDirEnvName,
		),
	)
	boolVar(
		&v.config.ReuseDynamicPorts,
		"reuse-dynamic-ports",
		false,
		processDocPrefix+"Whether to attempt to reuse dynamically allocated ports across node restarts.",
	)
}

func (v *processRuntimeVars) getProcessRuntimeConfig() (*tmpnet.ProcessRuntimeConfig, error) {
	if err := v.validate(); err != nil {
		return nil, stacktrace.Wrap(err)
	}
	return &v.config, nil
}

func (v *processRuntimeVars) validate() error {
	path := v.config.RinkGoPath

	if len(path) == 0 {
		return stacktrace.Wrap(errRinkGoRequired)
	}

	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err != nil {
			return stacktrace.Errorf("--%s (%s) not found: %w", rinkgoPathFlag, path, err)
		}
		return nil
	}

	// A relative path must be resolvable to an absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return stacktrace.Errorf(
			"--%s (%s) is a relative path but its absolute path cannot be determined: %w",
			rinkgoPathFlag,
			path,
			err,
		)
	}

	// The absolute path must exist
	if _, err := os.Stat(absPath); err != nil {
		return stacktrace.Errorf(
			"--%s (%s) is a relative path but its absolute path (%s) is not found: %w",
			rinkgoPathFlag,
			path,
			absPath,
			err,
		)
	}
	return nil
}
