// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the sku details for the given resource type and sku name.
// Azure REST API version: 2021-09-01-preview.
func LookupSkus(ctx *pulumi.Context, args *LookupSkusArgs, opts ...pulumi.InvokeOption) (*LookupSkusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSkusResult
	err := ctx.Invoke("azure-native:providerhub:getSkus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku string `pulumi:"sku"`
}

type LookupSkusResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSkusOutput(ctx *pulumi.Context, args LookupSkusOutputArgs, opts ...pulumi.InvokeOption) LookupSkusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusResult, error) {
			args := v.(LookupSkusArgs)
			r, err := LookupSkus(ctx, &args, opts...)
			var s LookupSkusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusResultOutput)
}

type LookupSkusOutputArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// The SKU.
	Sku pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusArgs)(nil)).Elem()
}

type LookupSkusResultOutput struct{ *pulumi.OutputState }

func (LookupSkusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusResult)(nil)).Elem()
}

func (o LookupSkusResultOutput) ToLookupSkusResultOutput() LookupSkusResultOutput {
	return o
}

func (o LookupSkusResultOutput) ToLookupSkusResultOutputWithContext(ctx context.Context) LookupSkusResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSkusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSkusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSkusResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSkusResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusResultOutput{})
}
