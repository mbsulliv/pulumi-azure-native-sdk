// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupComponentVersion(ctx *pulumi.Context, args *LookupComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupComponentVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupComponentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupComponentVersionArgs struct {
	// Container name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupComponentVersionResult struct {
	// [Required] Additional attributes of the entity.
	ComponentVersionProperties ComponentVersionResponse `pulumi:"componentVersionProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupComponentVersionResult
func (val *LookupComponentVersionResult) Defaults() *LookupComponentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentVersionProperties = *tmp.ComponentVersionProperties.Defaults()

	return &tmp
}

func LookupComponentVersionOutput(ctx *pulumi.Context, args LookupComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentVersionResult, error) {
			args := v.(LookupComponentVersionArgs)
			r, err := LookupComponentVersion(ctx, &args, opts...)
			var s LookupComponentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentVersionResultOutput)
}

type LookupComponentVersionOutputArgs struct {
	// Container name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionResult)(nil)).Elem()
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutput() LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutputWithContext(ctx context.Context) LookupComponentVersionResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupComponentVersionResultOutput) ComponentVersionProperties() ComponentVersionResponseOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) ComponentVersionResponse { return v.ComponentVersionProperties }).(ComponentVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupComponentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupComponentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupComponentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupComponentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentVersionResultOutput{})
}
