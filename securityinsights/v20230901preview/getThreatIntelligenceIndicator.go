// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// View a threat intelligence indicator by name.
func LookupThreatIntelligenceIndicator(ctx *pulumi.Context, args *LookupThreatIntelligenceIndicatorArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelligenceIndicatorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupThreatIntelligenceIndicatorResult
	err := ctx.Invoke("azure-native:securityinsights/v20230901preview:getThreatIntelligenceIndicator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelligenceIndicatorArgs struct {
	// Threat intelligence indicator name field.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Threat intelligence information object.
type LookupThreatIntelligenceIndicatorResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The kind of the entity.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupThreatIntelligenceIndicatorOutput(ctx *pulumi.Context, args LookupThreatIntelligenceIndicatorOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelligenceIndicatorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelligenceIndicatorResult, error) {
			args := v.(LookupThreatIntelligenceIndicatorArgs)
			r, err := LookupThreatIntelligenceIndicator(ctx, &args, opts...)
			var s LookupThreatIntelligenceIndicatorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThreatIntelligenceIndicatorResultOutput)
}

type LookupThreatIntelligenceIndicatorOutputArgs struct {
	// Threat intelligence indicator name field.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupThreatIntelligenceIndicatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceIndicatorArgs)(nil)).Elem()
}

// Threat intelligence information object.
type LookupThreatIntelligenceIndicatorResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelligenceIndicatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceIndicatorResult)(nil)).Elem()
}

func (o LookupThreatIntelligenceIndicatorResultOutput) ToLookupThreatIntelligenceIndicatorResultOutput() LookupThreatIntelligenceIndicatorResultOutput {
	return o
}

func (o LookupThreatIntelligenceIndicatorResultOutput) ToLookupThreatIntelligenceIndicatorResultOutputWithContext(ctx context.Context) LookupThreatIntelligenceIndicatorResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupThreatIntelligenceIndicatorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupThreatIntelligenceIndicatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The kind of the entity.
func (o LookupThreatIntelligenceIndicatorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupThreatIntelligenceIndicatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupThreatIntelligenceIndicatorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupThreatIntelligenceIndicatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceIndicatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelligenceIndicatorResultOutput{})
}
