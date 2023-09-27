// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220904

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The operation returns the admin kubeconfig.
func ListOpenShiftClusterAdminCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterAdminCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListOpenShiftClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20220904:listOpenShiftClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterAdminCredentialsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// OpenShiftClusterAdminKubeconfig represents an OpenShift cluster's admin kubeconfig.
type ListOpenShiftClusterAdminCredentialsResult struct {
	// The base64-encoded kubeconfig file.
	Kubeconfig *string `pulumi:"kubeconfig"`
}

func ListOpenShiftClusterAdminCredentialsOutput(ctx *pulumi.Context, args ListOpenShiftClusterAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListOpenShiftClusterAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOpenShiftClusterAdminCredentialsResult, error) {
			args := v.(ListOpenShiftClusterAdminCredentialsArgs)
			r, err := ListOpenShiftClusterAdminCredentials(ctx, &args, opts...)
			var s ListOpenShiftClusterAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOpenShiftClusterAdminCredentialsResultOutput)
}

type ListOpenShiftClusterAdminCredentialsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListOpenShiftClusterAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterAdminCredentialsArgs)(nil)).Elem()
}

// OpenShiftClusterAdminKubeconfig represents an OpenShift cluster's admin kubeconfig.
type ListOpenShiftClusterAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListOpenShiftClusterAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterAdminCredentialsResult)(nil)).Elem()
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) ToListOpenShiftClusterAdminCredentialsResultOutput() ListOpenShiftClusterAdminCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) ToListOpenShiftClusterAdminCredentialsResultOutputWithContext(ctx context.Context) ListOpenShiftClusterAdminCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListOpenShiftClusterAdminCredentialsResult] {
	return pulumix.Output[ListOpenShiftClusterAdminCredentialsResult]{
		OutputState: o.OutputState,
	}
}

// The base64-encoded kubeconfig file.
func (o ListOpenShiftClusterAdminCredentialsResultOutput) Kubeconfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterAdminCredentialsResult) *string { return v.Kubeconfig }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOpenShiftClusterAdminCredentialsResultOutput{})
}
