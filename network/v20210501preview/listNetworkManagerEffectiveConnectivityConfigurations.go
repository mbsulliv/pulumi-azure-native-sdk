// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all effective connectivity configurations applied on a virtual network.
func ListNetworkManagerEffectiveConnectivityConfigurations(ctx *pulumi.Context, args *ListNetworkManagerEffectiveConnectivityConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListNetworkManagerEffectiveConnectivityConfigurationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNetworkManagerEffectiveConnectivityConfigurationsResult
	err := ctx.Invoke("azure-native:network/v20210501preview:listNetworkManagerEffectiveConnectivityConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNetworkManagerEffectiveConnectivityConfigurationsArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// Result of the request to list networkManagerEffectiveConnectivityConfiguration. It contains a list of groups and a skiptoken to get the next set of results.
type ListNetworkManagerEffectiveConnectivityConfigurationsResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of NetworkManagerEffectiveConnectivityConfiguration
	Value []EffectiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListNetworkManagerEffectiveConnectivityConfigurationsOutput(ctx *pulumi.Context, args ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNetworkManagerEffectiveConnectivityConfigurationsResult, error) {
			args := v.(ListNetworkManagerEffectiveConnectivityConfigurationsArgs)
			r, err := ListNetworkManagerEffectiveConnectivityConfigurations(ctx, &args, opts...)
			var s ListNetworkManagerEffectiveConnectivityConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput)
}

type ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (ListNetworkManagerEffectiveConnectivityConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveConnectivityConfigurationsArgs)(nil)).Elem()
}

// Result of the request to list networkManagerEffectiveConnectivityConfiguration. It contains a list of groups and a skiptoken to get the next set of results.
type ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNetworkManagerEffectiveConnectivityConfigurationsResult)(nil)).Elem()
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ToListNetworkManagerEffectiveConnectivityConfigurationsResultOutput() ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) ToListNetworkManagerEffectiveConnectivityConfigurationsResultOutputWithContext(ctx context.Context) ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveConnectivityConfigurationsResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of NetworkManagerEffectiveConnectivityConfiguration
func (o ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput) Value() EffectiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListNetworkManagerEffectiveConnectivityConfigurationsResult) []EffectiveConnectivityConfigurationResponse {
		return v.Value
	}).(EffectiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNetworkManagerEffectiveConnectivityConfigurationsResultOutput{})
}
