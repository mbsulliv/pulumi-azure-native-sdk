// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220415preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the cluster customer user credentials for the dedicated appliance.
func ListApplianceClusterCustomerUserCredential(ctx *pulumi.Context, args *ListApplianceClusterCustomerUserCredentialArgs, opts ...pulumi.InvokeOption) (*ListApplianceClusterCustomerUserCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListApplianceClusterCustomerUserCredentialResult
	err := ctx.Invoke("azure-native:resourceconnector/v20220415preview:listApplianceClusterCustomerUserCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplianceClusterCustomerUserCredentialArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Appliances name.
	ResourceName string `pulumi:"resourceName"`
}

// The List Cluster Customer User Credential Results appliance.
type ListApplianceClusterCustomerUserCredentialResult struct {
	// The list of appliance kubeconfigs.
	Kubeconfigs []ApplianceCredentialKubeconfigResponse `pulumi:"kubeconfigs"`
	// Map of Customer User Public and Private SSH Keys
	SshKeys map[string]SSHKeyResponse `pulumi:"sshKeys"`
}

func ListApplianceClusterCustomerUserCredentialOutput(ctx *pulumi.Context, args ListApplianceClusterCustomerUserCredentialOutputArgs, opts ...pulumi.InvokeOption) ListApplianceClusterCustomerUserCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplianceClusterCustomerUserCredentialResult, error) {
			args := v.(ListApplianceClusterCustomerUserCredentialArgs)
			r, err := ListApplianceClusterCustomerUserCredential(ctx, &args, opts...)
			var s ListApplianceClusterCustomerUserCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplianceClusterCustomerUserCredentialResultOutput)
}

type ListApplianceClusterCustomerUserCredentialOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Appliances name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListApplianceClusterCustomerUserCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterCustomerUserCredentialArgs)(nil)).Elem()
}

// The List Cluster Customer User Credential Results appliance.
type ListApplianceClusterCustomerUserCredentialResultOutput struct{ *pulumi.OutputState }

func (ListApplianceClusterCustomerUserCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplianceClusterCustomerUserCredentialResult)(nil)).Elem()
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) ToListApplianceClusterCustomerUserCredentialResultOutput() ListApplianceClusterCustomerUserCredentialResultOutput {
	return o
}

func (o ListApplianceClusterCustomerUserCredentialResultOutput) ToListApplianceClusterCustomerUserCredentialResultOutputWithContext(ctx context.Context) ListApplianceClusterCustomerUserCredentialResultOutput {
	return o
}

// The list of appliance kubeconfigs.
func (o ListApplianceClusterCustomerUserCredentialResultOutput) Kubeconfigs() ApplianceCredentialKubeconfigResponseArrayOutput {
	return o.ApplyT(func(v ListApplianceClusterCustomerUserCredentialResult) []ApplianceCredentialKubeconfigResponse {
		return v.Kubeconfigs
	}).(ApplianceCredentialKubeconfigResponseArrayOutput)
}

// Map of Customer User Public and Private SSH Keys
func (o ListApplianceClusterCustomerUserCredentialResultOutput) SshKeys() SSHKeyResponseMapOutput {
	return o.ApplyT(func(v ListApplianceClusterCustomerUserCredentialResult) map[string]SSHKeyResponse { return v.SshKeys }).(SSHKeyResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplianceClusterCustomerUserCredentialResultOutput{})
}
