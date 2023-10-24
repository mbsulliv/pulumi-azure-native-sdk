// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The list credential result response.
func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20230901:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	// Only apply to AAD clusters, specifies the format of returned kubeconfig. Format 'azure' will return azure auth-provider kubeconfig; format 'exec' will return exec format kubeconfig, which requires kubelogin binary in the path.
	Format *string `pulumi:"format"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn *string `pulumi:"serverFqdn"`
}

// The list credential result response.
type ListManagedClusterUserCredentialsResult struct {
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterUserCredentialsResult, error) {
			args := v.(ListManagedClusterUserCredentialsArgs)
			r, err := ListManagedClusterUserCredentials(ctx, &args, opts...)
			var s ListManagedClusterUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterUserCredentialsResultOutput)
}

type ListManagedClusterUserCredentialsOutputArgs struct {
	// Only apply to AAD clusters, specifies the format of returned kubeconfig. Format 'azure' will return azure auth-provider kubeconfig; format 'exec' will return exec format kubeconfig, which requires kubelogin binary in the path.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsArgs)(nil)).Elem()
}

// The list credential result response.
type ListManagedClusterUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutput() ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListManagedClusterUserCredentialsResult] {
	return pulumix.Output[ListManagedClusterUserCredentialsResult]{
		OutputState: o.OutputState,
	}
}

// Base64-encoded Kubernetes configuration file.
func (o ListManagedClusterUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterUserCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterUserCredentialsResultOutput{})
}
