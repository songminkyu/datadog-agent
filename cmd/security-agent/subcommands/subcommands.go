// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package subcommands implement security agent subcommands
package subcommands

import (
	"github.com/DataDog/datadog-agent/cmd/security-agent/command"
	cmdcheck "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/check"
	cmdcompliance "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/compliance"
	cmdconfig "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/config"
	cmdcoverage "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/coverage"
	cmdflare "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/flare"
	cmdstart "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/start"
	cmdstatus "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/status"
	cmdversion "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/version"
	cmdworkloadlist "github.com/DataDog/datadog-agent/cmd/security-agent/subcommands/workloadlist"
)

// SecurityAgentSubcommands returns SubcommandFactories for the subcommands supported
// with the current build flags. The build tags in use right now are
// !windows && kubeapiserver (check, and any parent command that uses check),
// kubeapiserver (config),
// and linux (runtime).
func SecurityAgentSubcommands() []command.SubcommandFactory {
	return []command.SubcommandFactory{
		cmdcheck.SecurityAgentCommands,
		cmdcompliance.Commands,
		cmdconfig.Commands,
		cmdflare.Commands,
		cmdstart.Commands,
		cmdstatus.Commands,
		cmdversion.Commands,
		cmdworkloadlist.Commands,
		cmdcoverage.Commands,
	}
}
