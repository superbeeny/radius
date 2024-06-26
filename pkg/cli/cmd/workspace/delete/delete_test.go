/*
Copyright 2023 The Radius Authors.

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

package delete

import (
	"context"
	"fmt"
	"testing"

	"github.com/radius-project/radius/pkg/cli/framework"
	"github.com/radius-project/radius/pkg/cli/output"
	"github.com/radius-project/radius/pkg/cli/prompt"
	"github.com/radius-project/radius/pkg/cli/workspaces"
	"github.com/radius-project/radius/test/radcli"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_CommandValidation(t *testing.T) {
	radcli.SharedCommandValidation(t, NewCommand)
}

func Test_Validate(t *testing.T) {
	config := radcli.LoadConfigWithWorkspace(t)

	testcases := []radcli.ValidateInput{
		{
			Name:          "delete current workspace valid",
			Input:         []string{},
			ExpectedValid: true,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
		{
			Name:          "delete explicit workspace flag valid",
			Input:         []string{"-w", radcli.TestWorkspaceName},
			ExpectedValid: true,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
		{
			Name:          "delete explicit workspace positional valid",
			Input:         []string{radcli.TestWorkspaceName},
			ExpectedValid: true,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
		{
			Name:          "delete workspace with non-named workspace invalid",
			Input:         []string{},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: radcli.LoadEmptyConfig(t)},
		},
		{
			Name:          "delete workspace not-found invalid",
			Input:         []string{"other-workspace"},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
		{
			Name:          "delete workspace too-many-args invalid",
			Input:         []string{"other-workspace", "other-thing"},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
		{
			Name:          "delete workspace flag and positional invalid",
			Input:         []string{"other-workspace", "-w", "other-thing"},
			ExpectedValid: false,
			ConfigHolder:  framework.ConfigHolder{Config: config},
		},
	}
	radcli.SharedValidateValidation(t, NewCommand, testcases)
}

func Test_Run(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Delete workspace with confirmation", func(t *testing.T) {
		outputSink := &output.MockOutput{}

		configFile := framework.NewMockConfigFileInterface(ctrl)
		configFile.EXPECT().
			DeleteWorkspace(gomock.Any(), gomock.Any(), "test-workspace").
			Return(nil).
			Times(1)

		prompter := prompt.NewMockInterface(ctrl)
		prompter.EXPECT().
			GetListInput([]string{prompt.ConfirmNo, prompt.ConfirmYes}, fmt.Sprintf(deleteConfirmationFmt, "test-workspace")).
			Return(prompt.ConfirmYes, nil).
			Times(1)

		runner := &Runner{
			ConfigHolder:        &framework.ConfigHolder{},
			ConfigFileInterface: configFile,
			Output:              outputSink,
			InputPrompter:       prompter,
			Workspace: &workspaces.Workspace{
				Name: "test-workspace",
			},
		}

		err := runner.Run(context.Background())
		require.NoError(t, err)

		require.Empty(t, outputSink.Writes)
	})
	t.Run("Delete workspace bypass confirmation", func(t *testing.T) {
		outputSink := &output.MockOutput{}

		configFile := framework.NewMockConfigFileInterface(ctrl)
		configFile.EXPECT().
			DeleteWorkspace(gomock.Any(), gomock.Any(), "test-workspace").
			Return(nil).
			Times(1)

		prompter := prompt.NewMockInterface(ctrl)
		prompter.EXPECT().
			GetListInput(gomock.Any(), gomock.Any()).
			Return(prompt.ConfirmNo, nil).
			Times(0)

		runner := &Runner{
			ConfigHolder:        &framework.ConfigHolder{},
			ConfigFileInterface: configFile,
			Output:              outputSink,
			InputPrompter:       prompter,
			Workspace: &workspaces.Workspace{
				Name: "test-workspace",
			},

			Confirm: true,
		}

		err := runner.Run(context.Background())
		require.NoError(t, err)

		require.Empty(t, outputSink.Writes)
	})
	t.Run("Delete workspace not confirmed", func(t *testing.T) {
		outputSink := &output.MockOutput{}

		configFile := framework.NewMockConfigFileInterface(ctrl)
		configFile.EXPECT().
			DeleteWorkspace(gomock.Any(), gomock.Any(), "test-workspace").
			Return(nil).
			Times(0)

		prompter := prompt.NewMockInterface(ctrl)
		prompter.EXPECT().
			GetListInput([]string{prompt.ConfirmNo, prompt.ConfirmYes}, fmt.Sprintf(deleteConfirmationFmt, "test-workspace")).
			Return(prompt.ConfirmNo, nil).
			Times(1)

		runner := &Runner{
			ConfigHolder:        &framework.ConfigHolder{},
			ConfigFileInterface: configFile,
			Output:              outputSink,
			InputPrompter:       prompter,
			Workspace: &workspaces.Workspace{
				Name: "test-workspace",
			},
		}

		err := runner.Run(context.Background())
		require.NoError(t, err)

		require.Empty(t, outputSink.Writes)
	})

	t.Run("Exit Console with interrupt", func(t *testing.T) {
		outputSink := &output.MockOutput{}

		prompter := prompt.NewMockInterface(ctrl)
		prompter.EXPECT().
			GetListInput([]string{prompt.ConfirmNo, prompt.ConfirmYes}, fmt.Sprintf(deleteConfirmationFmt, "test-workspace")).
			Return("", &prompt.ErrExitConsole{}).
			Times(1)

		runner := &Runner{
			ConfigHolder:  &framework.ConfigHolder{},
			Output:        outputSink,
			InputPrompter: prompter,
			Workspace: &workspaces.Workspace{
				Name: "test-workspace",
			},
		}

		err := runner.Run(context.Background())
		require.Equal(t, &prompt.ErrExitConsole{}, err)
		require.Empty(t, outputSink.Writes)
	})
}
