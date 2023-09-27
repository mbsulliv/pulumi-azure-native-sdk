// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns Modeling resources for a given name.
func LookupModeling(ctx *pulumi.Context, args *LookupModelingArgs, opts ...pulumi.InvokeOption) (*LookupModelingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupModelingResult
	err := ctx.Invoke("azure-native:recommendationsservice/v20220301preview:getModeling", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelingArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName string `pulumi:"accountName"`
	// The name of the Modeling resource.
	ModelingName string `pulumi:"modelingName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Modeling resource details.
type LookupModelingResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Modeling resource properties.
	Properties ModelingResourceResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupModelingOutput(ctx *pulumi.Context, args LookupModelingOutputArgs, opts ...pulumi.InvokeOption) LookupModelingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelingResult, error) {
			args := v.(LookupModelingArgs)
			r, err := LookupModeling(ctx, &args, opts...)
			var s LookupModelingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelingResultOutput)
}

type LookupModelingOutputArgs struct {
	// The name of the RecommendationsService Account resource.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Modeling resource.
	ModelingName pulumi.StringInput `pulumi:"modelingName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupModelingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelingArgs)(nil)).Elem()
}

// Modeling resource details.
type LookupModelingResultOutput struct{ *pulumi.OutputState }

func (LookupModelingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelingResult)(nil)).Elem()
}

func (o LookupModelingResultOutput) ToLookupModelingResultOutput() LookupModelingResultOutput {
	return o
}

func (o LookupModelingResultOutput) ToLookupModelingResultOutputWithContext(ctx context.Context) LookupModelingResultOutput {
	return o
}

func (o LookupModelingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupModelingResult] {
	return pulumix.Output[LookupModelingResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupModelingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupModelingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupModelingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Modeling resource properties.
func (o LookupModelingResultOutput) Properties() ModelingResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupModelingResult) ModelingResourceResponseProperties { return v.Properties }).(ModelingResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupModelingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupModelingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupModelingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupModelingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupModelingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelingResultOutput{})
}
