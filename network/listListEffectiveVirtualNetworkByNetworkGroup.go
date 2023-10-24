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

// Lists all effective virtual networks by specified network group.
// Azure REST API version: 2022-04-01-preview.
//
// Other available API versions: 2021-05-01-preview.
func ListListEffectiveVirtualNetworkByNetworkGroup(ctx *pulumi.Context, args *ListListEffectiveVirtualNetworkByNetworkGroupArgs, opts ...pulumi.InvokeOption) (*ListListEffectiveVirtualNetworkByNetworkGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListEffectiveVirtualNetworkByNetworkGroupResult
	err := ctx.Invoke("azure-native:network:listListEffectiveVirtualNetworkByNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListEffectiveVirtualNetworkByNetworkGroupArgs struct {
	// The name of the network group.
	NetworkGroupName string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
}

// Result of the request to list Effective Virtual Network. It contains a list of groups and a URL link to get the next set of results.
type ListListEffectiveVirtualNetworkByNetworkGroupResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of EffectiveVirtualNetwork
	Value []EffectiveVirtualNetworkResponse `pulumi:"value"`
}

func ListListEffectiveVirtualNetworkByNetworkGroupOutput(ctx *pulumi.Context, args ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListEffectiveVirtualNetworkByNetworkGroupResult, error) {
			args := v.(ListListEffectiveVirtualNetworkByNetworkGroupArgs)
			r, err := ListListEffectiveVirtualNetworkByNetworkGroup(ctx, &args, opts...)
			var s ListListEffectiveVirtualNetworkByNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput)
}

type ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs struct {
	// The name of the network group.
	NetworkGroupName pulumi.StringInput `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupArgs)(nil)).Elem()
}

// Result of the request to list Effective Virtual Network. It contains a list of groups and a URL link to get the next set of results.
type ListListEffectiveVirtualNetworkByNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupResult)(nil)).Elem()
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutput() ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutputWithContext(ctx context.Context) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListListEffectiveVirtualNetworkByNetworkGroupResult] {
	return pulumix.Output[ListListEffectiveVirtualNetworkByNetworkGroupResult]{
		OutputState: o.OutputState,
	}
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of EffectiveVirtualNetwork
func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) Value() EffectiveVirtualNetworkResponseArrayOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) []EffectiveVirtualNetworkResponse {
		return v.Value
	}).(EffectiveVirtualNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput{})
}
