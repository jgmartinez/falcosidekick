// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

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

package outputs

import (
	"encoding/json"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"

	"github.com/falcosecurity/falcosidekick/types"
)

func TestNewRocketchatPayload(t *testing.T) {
	expectedOutput := slackPayload{
		Text:     "Rule: Test rule Priority: Debug",
		Username: "Falcosidekick",
		IconURL:  DefaultIconURL,
		Attachments: []slackAttachment{
			{
				Fallback: "This is a test from falcosidekick",
				Color:    PaleCyan,
				Text:     "This is a test from falcosidekick",
				Footer:   "",
				Fields: []slackAttachmentField{
					{
						Title: "rule",
						Value: "Test rule",
						Short: true,
					},
					{
						Title: "priority",
						Value: "Debug",
						Short: true,
					},
					{
						Title: "source",
						Value: "syscalls",
						Short: true,
					},
					{
						Title: "tags",
						Value: "test, example",
						Short: true,
					},
					{
						Title: "proc.name",
						Value: "falcosidekick",
						Short: true,
					},
					{
						Title: "time",
						Value: "2001-01-01 01:10:00 +0000 UTC",
						Short: false,
					},
					{
						Title: "hostname",
						Value: "test-host",
						Short: true,
					},
				},
			},
		},
	}

	var f types.FalcoPayload
	require.Nil(t, json.Unmarshal([]byte(falcoTestInput), &f))
	config := &types.Configuration{
		Rocketchat: types.RocketchatOutputConfig{
			Username: "Falcosidekick",
			Icon:     DefaultIconURL,
		},
	}

	var err error
	config.Rocketchat.MessageFormatTemplate, err = template.New("").Parse("Rule: {{ .Rule }} Priority: {{ .Priority }}")
	require.Nil(t, err)

	output := newRocketchatPayload(f, config)
	require.Equal(t, output, expectedOutput)
}
