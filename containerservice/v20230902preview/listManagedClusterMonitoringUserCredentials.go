// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230902preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The list credential result response.
func ListManagedClusterMonitoringUserCredentials(ctx *pulumi.Context, args *ListManagedClusterMonitoringUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterMonitoringUserCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagedClusterMonitoringUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20230902preview:listManagedClusterMonitoringUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterMonitoringUserCredentialsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn *string `pulumi:"serverFqdn"`
}

// The list credential result response.
type ListManagedClusterMonitoringUserCredentialsResult struct {
	// Base64-encoded Kubernetes configuration file.
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterMonitoringUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterMonitoringUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterMonitoringUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterMonitoringUserCredentialsResult, error) {
			args := v.(ListManagedClusterMonitoringUserCredentialsArgs)
			r, err := ListManagedClusterMonitoringUserCredentials(ctx, &args, opts...)
			var s ListManagedClusterMonitoringUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterMonitoringUserCredentialsResultOutput)
}

type ListManagedClusterMonitoringUserCredentialsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// server fqdn type for credentials to be returned
	ServerFqdn pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterMonitoringUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterMonitoringUserCredentialsArgs)(nil)).Elem()
}

// The list credential result response.
type ListManagedClusterMonitoringUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterMonitoringUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterMonitoringUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterMonitoringUserCredentialsResultOutput) ToListManagedClusterMonitoringUserCredentialsResultOutput() ListManagedClusterMonitoringUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterMonitoringUserCredentialsResultOutput) ToListManagedClusterMonitoringUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterMonitoringUserCredentialsResultOutput {
	return o
}

// Base64-encoded Kubernetes configuration file.
func (o ListManagedClusterMonitoringUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterMonitoringUserCredentialsResult) []CredentialResultResponse {
		return v.Kubeconfigs
	}).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterMonitoringUserCredentialsResultOutput{})
}
