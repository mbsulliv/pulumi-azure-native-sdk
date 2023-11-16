// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of an Asset Filter associated with the specified Asset.
func LookupAssetFilter(ctx *pulumi.Context, args *LookupAssetFilterArgs, opts ...pulumi.InvokeOption) (*LookupAssetFilterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAssetFilterResult
	err := ctx.Invoke("azure-native:media/v20230101:getAssetFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The Asset Filter name
	FilterName string `pulumi:"filterName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Asset Filter.
type LookupAssetFilterResult struct {
	// The first quality.
	FirstQuality *FirstQualityResponse `pulumi:"firstQuality"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The presentation time range.
	PresentationTimeRange *PresentationTimeRangeResponse `pulumi:"presentationTimeRange"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tracks selection conditions.
	Tracks []FilterTrackSelectionResponse `pulumi:"tracks"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAssetFilterOutput(ctx *pulumi.Context, args LookupAssetFilterOutputArgs, opts ...pulumi.InvokeOption) LookupAssetFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssetFilterResult, error) {
			args := v.(LookupAssetFilterArgs)
			r, err := LookupAssetFilter(ctx, &args, opts...)
			var s LookupAssetFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssetFilterResultOutput)
}

type LookupAssetFilterOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Asset name.
	AssetName pulumi.StringInput `pulumi:"assetName"`
	// The Asset Filter name
	FilterName pulumi.StringInput `pulumi:"filterName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAssetFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetFilterArgs)(nil)).Elem()
}

// An Asset Filter.
type LookupAssetFilterResultOutput struct{ *pulumi.OutputState }

func (LookupAssetFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetFilterResult)(nil)).Elem()
}

func (o LookupAssetFilterResultOutput) ToLookupAssetFilterResultOutput() LookupAssetFilterResultOutput {
	return o
}

func (o LookupAssetFilterResultOutput) ToLookupAssetFilterResultOutputWithContext(ctx context.Context) LookupAssetFilterResultOutput {
	return o
}

// The first quality.
func (o LookupAssetFilterResultOutput) FirstQuality() FirstQualityResponsePtrOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) *FirstQualityResponse { return v.FirstQuality }).(FirstQualityResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAssetFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAssetFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The presentation time range.
func (o LookupAssetFilterResultOutput) PresentationTimeRange() PresentationTimeRangeResponsePtrOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) *PresentationTimeRangeResponse { return v.PresentationTimeRange }).(PresentationTimeRangeResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o LookupAssetFilterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tracks selection conditions.
func (o LookupAssetFilterResultOutput) Tracks() FilterTrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) []FilterTrackSelectionResponse { return v.Tracks }).(FilterTrackSelectionResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAssetFilterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssetFilterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetFilterResultOutput{})
}
