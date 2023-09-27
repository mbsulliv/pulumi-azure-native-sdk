// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Administrative credentials for accessing vCenter and NSX-T
func ListPrivateCloudAdminCredentials(ctx *pulumi.Context, args *ListPrivateCloudAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListPrivateCloudAdminCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPrivateCloudAdminCredentialsResult
	err := ctx.Invoke("azure-native:avs/v20220501:listPrivateCloudAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateCloudAdminCredentialsArgs struct {
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Administrative credentials for accessing vCenter and NSX-T
type ListPrivateCloudAdminCredentialsResult struct {
	// NSX-T Manager password
	NsxtPassword string `pulumi:"nsxtPassword"`
	// NSX-T Manager username
	NsxtUsername string `pulumi:"nsxtUsername"`
	// vCenter admin password
	VcenterPassword string `pulumi:"vcenterPassword"`
	// vCenter admin username
	VcenterUsername string `pulumi:"vcenterUsername"`
}

func ListPrivateCloudAdminCredentialsOutput(ctx *pulumi.Context, args ListPrivateCloudAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListPrivateCloudAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateCloudAdminCredentialsResult, error) {
			args := v.(ListPrivateCloudAdminCredentialsArgs)
			r, err := ListPrivateCloudAdminCredentials(ctx, &args, opts...)
			var s ListPrivateCloudAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateCloudAdminCredentialsResultOutput)
}

type ListPrivateCloudAdminCredentialsOutputArgs struct {
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListPrivateCloudAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateCloudAdminCredentialsArgs)(nil)).Elem()
}

// Administrative credentials for accessing vCenter and NSX-T
type ListPrivateCloudAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListPrivateCloudAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateCloudAdminCredentialsResult)(nil)).Elem()
}

func (o ListPrivateCloudAdminCredentialsResultOutput) ToListPrivateCloudAdminCredentialsResultOutput() ListPrivateCloudAdminCredentialsResultOutput {
	return o
}

func (o ListPrivateCloudAdminCredentialsResultOutput) ToListPrivateCloudAdminCredentialsResultOutputWithContext(ctx context.Context) ListPrivateCloudAdminCredentialsResultOutput {
	return o
}

func (o ListPrivateCloudAdminCredentialsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListPrivateCloudAdminCredentialsResult] {
	return pulumix.Output[ListPrivateCloudAdminCredentialsResult]{
		OutputState: o.OutputState,
	}
}

// NSX-T Manager password
func (o ListPrivateCloudAdminCredentialsResultOutput) NsxtPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.NsxtPassword }).(pulumi.StringOutput)
}

// NSX-T Manager username
func (o ListPrivateCloudAdminCredentialsResultOutput) NsxtUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.NsxtUsername }).(pulumi.StringOutput)
}

// vCenter admin password
func (o ListPrivateCloudAdminCredentialsResultOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.VcenterPassword }).(pulumi.StringOutput)
}

// vCenter admin username
func (o ListPrivateCloudAdminCredentialsResultOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.VcenterUsername }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateCloudAdminCredentialsResultOutput{})
}
