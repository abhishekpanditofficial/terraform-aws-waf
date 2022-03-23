package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformWafExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/waf",
		Vars: map[string]interface{}{
			"acl_name": "test-acl",
			"acl_description": "test-description",
			"metric_name": "test-name",
			"project": "test-project",
		},
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "acl_name")
	assert.Equal(t, "test-acl", output)
}