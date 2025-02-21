package test

import (
	"testing"
	"time"
	"fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureLinuxVMCreation(t *testing.T) {
	// Define Terraform options
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // Assuming Terraform files are in the root folder
		VarFiles:     []string{"terraform.tfvars"},
	}

	// Ensure resources are destroyed after test completion
	defer terraform.Destroy(t, terraformOptions)

	// Run Terraform Init and Apply
	terraform.InitAndApply(t, terraformOptions)

	// Get output variables
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	publicIP := terraform.Output(t, terraformOptions, "vm_public_ip")

	// Print output values for debugging
	fmt.Println("VM Name:", vmName)
	fmt.Println("Public IP:", publicIP)

	// Assert outputs are not empty
	assert.NotEmpty(t, vmName, "VM name should not be empty")
	assert.NotEmpty(t, publicIP, "Public IP should not be empty")

	// Wait for the instance to be fully available (optional)
	time.Sleep(10 * time.Second)
}
