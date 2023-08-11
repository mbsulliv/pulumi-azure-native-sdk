// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Maps Creator resource.
func LookupCreator(ctx *pulumi.Context, args *LookupCreatorArgs, opts ...pulumi.InvokeOption) (*LookupCreatorResult, error) {
	var rv LookupCreatorResult
	err := ctx.Invoke("azure-native:maps/v20230601:getCreator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCreatorArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the Maps Creator instance.
	CreatorName string `pulumi:"creatorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type LookupCreatorResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Creator resource properties.
	Properties CreatorPropertiesResponse `pulumi:"properties"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCreatorOutput(ctx *pulumi.Context, args LookupCreatorOutputArgs, opts ...pulumi.InvokeOption) LookupCreatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCreatorResult, error) {
			args := v.(LookupCreatorArgs)
			r, err := LookupCreator(ctx, &args, opts...)
			var s LookupCreatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCreatorResultOutput)
}

type LookupCreatorOutputArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Maps Creator instance.
	CreatorName pulumi.StringInput `pulumi:"creatorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCreatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCreatorArgs)(nil)).Elem()
}

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type LookupCreatorResultOutput struct{ *pulumi.OutputState }

func (LookupCreatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCreatorResult)(nil)).Elem()
}

func (o LookupCreatorResultOutput) ToLookupCreatorResultOutput() LookupCreatorResultOutput {
	return o
}

func (o LookupCreatorResultOutput) ToLookupCreatorResultOutputWithContext(ctx context.Context) LookupCreatorResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCreatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupCreatorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCreatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Creator resource properties.
func (o LookupCreatorResultOutput) Properties() CreatorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCreatorResult) CreatorPropertiesResponse { return v.Properties }).(CreatorPropertiesResponseOutput)
}

// The system meta data relating to this resource.
func (o LookupCreatorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCreatorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCreatorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCreatorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCreatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCreatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCreatorResultOutput{})
}
