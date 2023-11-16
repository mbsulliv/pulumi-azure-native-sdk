// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockchain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the available consortiums for a subscription.
// Azure REST API version: 2018-06-01-preview.
func ListLocationConsortiums(ctx *pulumi.Context, args *ListLocationConsortiumsArgs, opts ...pulumi.InvokeOption) (*ListLocationConsortiumsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListLocationConsortiumsResult
	err := ctx.Invoke("azure-native:blockchain:listLocationConsortiums", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocationConsortiumsArgs struct {
	// Location Name.
	LocationName string `pulumi:"locationName"`
}

// Collection of the consortium payload.
type ListLocationConsortiumsResult struct {
	// Gets or sets the collection of consortiums.
	Value []ConsortiumResponse `pulumi:"value"`
}

func ListLocationConsortiumsOutput(ctx *pulumi.Context, args ListLocationConsortiumsOutputArgs, opts ...pulumi.InvokeOption) ListLocationConsortiumsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLocationConsortiumsResult, error) {
			args := v.(ListLocationConsortiumsArgs)
			r, err := ListLocationConsortiums(ctx, &args, opts...)
			var s ListLocationConsortiumsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLocationConsortiumsResultOutput)
}

type ListLocationConsortiumsOutputArgs struct {
	// Location Name.
	LocationName pulumi.StringInput `pulumi:"locationName"`
}

func (ListLocationConsortiumsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocationConsortiumsArgs)(nil)).Elem()
}

// Collection of the consortium payload.
type ListLocationConsortiumsResultOutput struct{ *pulumi.OutputState }

func (ListLocationConsortiumsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLocationConsortiumsResult)(nil)).Elem()
}

func (o ListLocationConsortiumsResultOutput) ToListLocationConsortiumsResultOutput() ListLocationConsortiumsResultOutput {
	return o
}

func (o ListLocationConsortiumsResultOutput) ToListLocationConsortiumsResultOutputWithContext(ctx context.Context) ListLocationConsortiumsResultOutput {
	return o
}

// Gets or sets the collection of consortiums.
func (o ListLocationConsortiumsResultOutput) Value() ConsortiumResponseArrayOutput {
	return o.ApplyT(func(v ListLocationConsortiumsResult) []ConsortiumResponse { return v.Value }).(ConsortiumResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLocationConsortiumsResultOutput{})
}
