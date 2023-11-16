// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the available Service Providers for creating Connection Settings
func ListBotConnectionServiceProviders(ctx *pulumi.Context, args *ListBotConnectionServiceProvidersArgs, opts ...pulumi.InvokeOption) (*ListBotConnectionServiceProvidersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListBotConnectionServiceProvidersResult
	err := ctx.Invoke("azure-native:botservice/v20220615preview:listBotConnectionServiceProviders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBotConnectionServiceProvidersArgs struct {
}

// The list of bot service providers response.
type ListBotConnectionServiceProvidersResult struct {
	// The link used to get the next page of bot service providers.
	NextLink *string `pulumi:"nextLink"`
	// Gets the list of bot service providers and their properties.
	Value []ServiceProviderResponse `pulumi:"value"`
}

func ListBotConnectionServiceProvidersOutput(ctx *pulumi.Context, args ListBotConnectionServiceProvidersOutputArgs, opts ...pulumi.InvokeOption) ListBotConnectionServiceProvidersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBotConnectionServiceProvidersResult, error) {
			args := v.(ListBotConnectionServiceProvidersArgs)
			r, err := ListBotConnectionServiceProviders(ctx, &args, opts...)
			var s ListBotConnectionServiceProvidersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBotConnectionServiceProvidersResultOutput)
}

type ListBotConnectionServiceProvidersOutputArgs struct {
}

func (ListBotConnectionServiceProvidersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotConnectionServiceProvidersArgs)(nil)).Elem()
}

// The list of bot service providers response.
type ListBotConnectionServiceProvidersResultOutput struct{ *pulumi.OutputState }

func (ListBotConnectionServiceProvidersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBotConnectionServiceProvidersResult)(nil)).Elem()
}

func (o ListBotConnectionServiceProvidersResultOutput) ToListBotConnectionServiceProvidersResultOutput() ListBotConnectionServiceProvidersResultOutput {
	return o
}

func (o ListBotConnectionServiceProvidersResultOutput) ToListBotConnectionServiceProvidersResultOutputWithContext(ctx context.Context) ListBotConnectionServiceProvidersResultOutput {
	return o
}

// The link used to get the next page of bot service providers.
func (o ListBotConnectionServiceProvidersResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBotConnectionServiceProvidersResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Gets the list of bot service providers and their properties.
func (o ListBotConnectionServiceProvidersResultOutput) Value() ServiceProviderResponseArrayOutput {
	return o.ApplyT(func(v ListBotConnectionServiceProvidersResult) []ServiceProviderResponse { return v.Value }).(ServiceProviderResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBotConnectionServiceProvidersResultOutput{})
}
