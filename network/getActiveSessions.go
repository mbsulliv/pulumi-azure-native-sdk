// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the list of currently active sessions on the Bastion.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01.
func GetActiveSessions(ctx *pulumi.Context, args *GetActiveSessionsArgs, opts ...pulumi.InvokeOption) (*GetActiveSessionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetActiveSessionsResult
	err := ctx.Invoke("azure-native:network:getActiveSessions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetActiveSessionsArgs struct {
	// The name of the Bastion Host.
	BastionHostName string `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for GetActiveSessions.
type GetActiveSessionsResult struct {
	// The URL to get the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// List of active sessions on the bastion.
	Value []BastionActiveSessionResponse `pulumi:"value"`
}

func GetActiveSessionsOutput(ctx *pulumi.Context, args GetActiveSessionsOutputArgs, opts ...pulumi.InvokeOption) GetActiveSessionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActiveSessionsResult, error) {
			args := v.(GetActiveSessionsArgs)
			r, err := GetActiveSessions(ctx, &args, opts...)
			var s GetActiveSessionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActiveSessionsResultOutput)
}

type GetActiveSessionsOutputArgs struct {
	// The name of the Bastion Host.
	BastionHostName pulumi.StringInput `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetActiveSessionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsArgs)(nil)).Elem()
}

// Response for GetActiveSessions.
type GetActiveSessionsResultOutput struct{ *pulumi.OutputState }

func (GetActiveSessionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveSessionsResult)(nil)).Elem()
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutput() GetActiveSessionsResultOutput {
	return o
}

func (o GetActiveSessionsResultOutput) ToGetActiveSessionsResultOutputWithContext(ctx context.Context) GetActiveSessionsResultOutput {
	return o
}

func (o GetActiveSessionsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetActiveSessionsResult] {
	return pulumix.Output[GetActiveSessionsResult]{
		OutputState: o.OutputState,
	}
}

// The URL to get the next set of results.
func (o GetActiveSessionsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of active sessions on the bastion.
func (o GetActiveSessionsResultOutput) Value() BastionActiveSessionResponseArrayOutput {
	return o.ApplyT(func(v GetActiveSessionsResult) []BastionActiveSessionResponse { return v.Value }).(BastionActiveSessionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActiveSessionsResultOutput{})
}
