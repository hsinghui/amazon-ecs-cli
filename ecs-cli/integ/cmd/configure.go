// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package cmd contains ECS CLI command tests.
package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aws/amazon-ecs-cli/ecs-cli/integ"

	"github.com/aws/amazon-ecs-cli/ecs-cli/integ/stdout"
)

// A CLIConfig holds the basic lookup information used by the ECS CLI for a cluster.
type CLIConfig struct {
	ClusterName string
	ConfigName  string
}

// TestFargateTutorialConfig runs `ecs-cli configure` with a FARGATE launch type.
func TestFargateTutorialConfig(t *testing.T) *CLIConfig {
	conf := CLIConfig{
		ClusterName: integ.SuggestedResourceName("fargate-tutorial", "cluster"),
		ConfigName:  integ.SuggestedResourceName("fargate-tutorial", "config"),
	}
	testConfig(t, conf.ClusterName, "FARGATE", conf.ConfigName)
	t.Logf("Created config %s", conf.ConfigName)
	return &conf
}

// TestEC2TutorialConfig runs `ecs-cli configure` with a EC2 launch type.
func TestEC2TutorialConfig(t *testing.T) *CLIConfig {
	conf := CLIConfig{
		ClusterName: integ.SuggestedResourceName("ec2-tutorial", "cluster"),
		ConfigName:  integ.SuggestedResourceName("ec2-tutorial", "config"),
	}
	testConfig(t, conf.ClusterName, "EC2", conf.ConfigName)
	t.Logf("Created config %s", conf.ConfigName)
	return &conf
}

func testConfig(t *testing.T, clusterName, launchType, configName string) {
	args := []string{
		"configure",
		"--cluster",
		clusterName,
		"--region",
		os.Getenv("AWS_REGION"),
		"--default-launch-type",
		launchType,
		"--config-name",
		configName,
	}
	cmd := integ.GetCommand(args)

	// When
	out, err := cmd.Output()
	require.NoErrorf(t, err, "Failed to configure CLI", "error %v, running %v, out: %s", err, args, string(out))

	// Then
	stdout.Stdout(out).TestHasAllSubstrings(t, []string{
		fmt.Sprintf("Saved ECS CLI cluster configuration %s", configName),
	})
}
