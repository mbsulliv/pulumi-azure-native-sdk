// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The operation returns the credentials.
func ListOpenShiftClusterCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListOpenShiftClusterCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20230701preview:listOpenShiftClusterCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterCredentialsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
type ListOpenShiftClusterCredentialsResult struct {
	// The password for the kubeadmin user.
	KubeadminPassword *string `pulumi:"kubeadminPassword"`
	// The username for the kubeadmin user.
	KubeadminUsername *string `pulumi:"kubeadminUsername"`
}

func ListOpenShiftClusterCredentialsOutput(ctx *pulumi.Context, args ListOpenShiftClusterCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListOpenShiftClusterCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOpenShiftClusterCredentialsResult, error) {
			args := v.(ListOpenShiftClusterCredentialsArgs)
			r, err := ListOpenShiftClusterCredentials(ctx, &args, opts...)
			var s ListOpenShiftClusterCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOpenShiftClusterCredentialsResultOutput)
}

type ListOpenShiftClusterCredentialsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListOpenShiftClusterCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterCredentialsArgs)(nil)).Elem()
}

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
type ListOpenShiftClusterCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListOpenShiftClusterCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterCredentialsResult)(nil)).Elem()
}

func (o ListOpenShiftClusterCredentialsResultOutput) ToListOpenShiftClusterCredentialsResultOutput() ListOpenShiftClusterCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterCredentialsResultOutput) ToListOpenShiftClusterCredentialsResultOutputWithContext(ctx context.Context) ListOpenShiftClusterCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterCredentialsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListOpenShiftClusterCredentialsResult] {
	return pulumix.Output[ListOpenShiftClusterCredentialsResult]{
		OutputState: o.OutputState,
	}
}

// The password for the kubeadmin user.
func (o ListOpenShiftClusterCredentialsResultOutput) KubeadminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterCredentialsResult) *string { return v.KubeadminPassword }).(pulumi.StringPtrOutput)
}

// The username for the kubeadmin user.
func (o ListOpenShiftClusterCredentialsResultOutput) KubeadminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterCredentialsResult) *string { return v.KubeadminUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOpenShiftClusterCredentialsResultOutput{})
}
