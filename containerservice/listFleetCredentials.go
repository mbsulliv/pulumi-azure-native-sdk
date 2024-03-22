// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the user credentials of a Fleet.
// Azure REST API version: 2023-03-15-preview.
//
// Other available API versions: 2022-07-02-preview, 2023-06-15-preview, 2023-08-15-preview, 2023-10-15, 2024-02-02-preview.
func ListFleetCredentials(ctx *pulumi.Context, args *ListFleetCredentialsArgs, opts ...pulumi.InvokeOption) (*ListFleetCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListFleetCredentialsResult
	err := ctx.Invoke("azure-native:containerservice:listFleetCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListFleetCredentialsArgs struct {
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Credential results response.
type ListFleetCredentialsResult struct {
	// Array of base64-encoded Kubernetes configuration files.
	Kubeconfigs []FleetCredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListFleetCredentialsOutput(ctx *pulumi.Context, args ListFleetCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListFleetCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListFleetCredentialsResult, error) {
			args := v.(ListFleetCredentialsArgs)
			r, err := ListFleetCredentials(ctx, &args, opts...)
			var s ListFleetCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListFleetCredentialsResultOutput)
}

type ListFleetCredentialsOutputArgs struct {
	// The name of the Fleet resource.
	FleetName pulumi.StringInput `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListFleetCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFleetCredentialsArgs)(nil)).Elem()
}

// The Credential results response.
type ListFleetCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListFleetCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListFleetCredentialsResult)(nil)).Elem()
}

func (o ListFleetCredentialsResultOutput) ToListFleetCredentialsResultOutput() ListFleetCredentialsResultOutput {
	return o
}

func (o ListFleetCredentialsResultOutput) ToListFleetCredentialsResultOutputWithContext(ctx context.Context) ListFleetCredentialsResultOutput {
	return o
}

// Array of base64-encoded Kubernetes configuration files.
func (o ListFleetCredentialsResultOutput) Kubeconfigs() FleetCredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListFleetCredentialsResult) []FleetCredentialResultResponse { return v.Kubeconfigs }).(FleetCredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListFleetCredentialsResultOutput{})
}
