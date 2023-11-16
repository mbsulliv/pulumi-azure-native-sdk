// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231002preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The list credential result response.
func ListManagedClusterAdminCredentials(ctx *pulumi.Context, args *ListManagedClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAdminCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagedClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20231002preview:listManagedClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAdminCredentialsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn *string `pulumi:"serverFqdn"`
}

// The list credential result response.
type ListManagedClusterAdminCredentialsResult struct {
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterAdminCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterAdminCredentialsResult, error) {
			args := v.(ListManagedClusterAdminCredentialsArgs)
			r, err := ListManagedClusterAdminCredentials(ctx, &args, opts...)
			var s ListManagedClusterAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterAdminCredentialsResultOutput)
}

type ListManagedClusterAdminCredentialsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAdminCredentialsArgs)(nil)).Elem()
}

// The list credential result response.
type ListManagedClusterAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAdminCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterAdminCredentialsResultOutput) ToListManagedClusterAdminCredentialsResultOutput() ListManagedClusterAdminCredentialsResultOutput {
	return o
}

func (o ListManagedClusterAdminCredentialsResultOutput) ToListManagedClusterAdminCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterAdminCredentialsResultOutput {
	return o
}

// Base64-encoded Kubernetes configuration file.
func (o ListManagedClusterAdminCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterAdminCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterAdminCredentialsResultOutput{})
}
