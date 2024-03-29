package apply

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/bigblueswarm/bbsctl/internal/mock"
	"github.com/bigblueswarm/bbsctl/internal/test"
	"github.com/bigblueswarm/bigblueswarm/v2/pkg/admin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNewCmd(t *testing.T) {
	mock.InitAdminMock()
	tests := []test.CmdTest{
		{
			Name: "an error returned by file loader should be returned",
			Args: []string{"-f", "/etc/dummy_folder/instances.yml"},
			Mock: func() {},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			Name: "an error return by admin apply method should be returned",
			Args: []string{"-f", "/tmp/instances.yml"},
			Mock: func() {
				instances := &admin.InstanceList{
					Kind:      "InstanceList",
					Instances: map[string]string{},
				}

				out, err := yaml.Marshal(instances)
				if err != nil {
					t.Fatal(err)
				}

				if err := os.WriteFile("/tmp/instances.yml", out, os.FileMode(0644)); err != nil {
					t.Fatal(err)
				}

				mock.ApplyFunc = func(kind string, resource *interface{}) error {
					return errors.New("admin error")
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, "unable to apply file. admin error", err.Error())
			},
		},
		{
			Name: "applying an InstanceList file should print a valid `InstanceList resource created` message",
			Args: []string{"-f", "/tmp/instances.yml"},
			Mock: func() {
				mock.ApplyFunc = func(kind string, resource *interface{}) error {
					return nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "InstanceList resource created", strings.TrimSpace(string(output.Bytes())))
			},
		},
		{
			Name: "applying a Tenant file should print `Tenant resource created`",
			Args: []string{"-f", "/tmp/localhost.tenant.yml"},
			Mock: func() {
				instances := &admin.Tenant{
					Kind:      "Tenant",
					Spec:      &admin.TenantSpec{},
					Instances: []string{},
				}

				out, err := yaml.Marshal(instances)
				if err != nil {
					t.Fatal(err)
				}

				if err := os.WriteFile("/tmp/localhost.tenant.yml", out, os.FileMode(0644)); err != nil {
					t.Fatal(err)
				}

				mock.ApplyFunc = func(kind string, resource *interface{}) error {
					return nil
				}
			},
			Validator: func(t *testing.T, output *bytes.Buffer, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "Tenant resource created", strings.TrimSpace(string(output.Bytes())))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewCmd()
			cmd.SetArgs(test.Args)
			cmd.SetOut(b)
			test.Mock()
			err := cmd.Execute()
			test.Validator(t, b, err)
		})
	}
}
