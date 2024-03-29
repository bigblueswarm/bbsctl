package root

import (
	"testing"

	"github.com/bigblueswarm/bbsctl/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	type test struct {
		name      string
		args      []string
		validator func(t *testing.T, err error)
	}

	tests := []test{
		{
			name: "Root command should initialize configuration",
			args: []string{"--config", "../../../test/config.yml"},
			validator: func(t *testing.T, err error) {
				assert.Equal(t, "bbs_admin_api_key", config.APIKey)
				assert.Equal(t, "http://localhost:8090", config.BBS)
			},
		},
	}

	for _, test := range tests {
		cmd := NewCmd()
		cmd.SetArgs(test.args)
		cmd.Execute()
	}
}
