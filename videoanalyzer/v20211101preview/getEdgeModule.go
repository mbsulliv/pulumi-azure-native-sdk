// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an existing edge module resource with the given name.
func LookupEdgeModule(ctx *pulumi.Context, args *LookupEdgeModuleArgs, opts ...pulumi.InvokeOption) (*LookupEdgeModuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEdgeModuleResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20211101preview:getEdgeModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEdgeModuleArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// The Edge Module name.
	EdgeModuleName string `pulumi:"edgeModuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The representation of an edge module.
type LookupEdgeModuleResult struct {
	// Internal ID generated for the instance of the Video Analyzer edge module.
	EdgeModuleId string `pulumi:"edgeModuleId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEdgeModuleOutput(ctx *pulumi.Context, args LookupEdgeModuleOutputArgs, opts ...pulumi.InvokeOption) LookupEdgeModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEdgeModuleResult, error) {
			args := v.(LookupEdgeModuleArgs)
			r, err := LookupEdgeModule(ctx, &args, opts...)
			var s LookupEdgeModuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEdgeModuleResultOutput)
}

type LookupEdgeModuleOutputArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Edge Module name.
	EdgeModuleName pulumi.StringInput `pulumi:"edgeModuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEdgeModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeModuleArgs)(nil)).Elem()
}

// The representation of an edge module.
type LookupEdgeModuleResultOutput struct{ *pulumi.OutputState }

func (LookupEdgeModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeModuleResult)(nil)).Elem()
}

func (o LookupEdgeModuleResultOutput) ToLookupEdgeModuleResultOutput() LookupEdgeModuleResultOutput {
	return o
}

func (o LookupEdgeModuleResultOutput) ToLookupEdgeModuleResultOutputWithContext(ctx context.Context) LookupEdgeModuleResultOutput {
	return o
}

// Internal ID generated for the instance of the Video Analyzer edge module.
func (o LookupEdgeModuleResultOutput) EdgeModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.EdgeModuleId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEdgeModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEdgeModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEdgeModuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEdgeModuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeModuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEdgeModuleResultOutput{})
}
