// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
func LookupRegistryModelVersion(ctx *pulumi.Context, args *LookupRegistryModelVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryModelVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryModelVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20231001:getRegistryModelVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryModelVersionArgs struct {
	// Container name. This is case-sensitive.
	ModelName string `pulumi:"modelName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version string `pulumi:"version"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryModelVersionResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionResponse `pulumi:"modelVersionProperties"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryModelVersionResult
func (val *LookupRegistryModelVersionResult) Defaults() *LookupRegistryModelVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelVersionProperties = *tmp.ModelVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryModelVersionOutput(ctx *pulumi.Context, args LookupRegistryModelVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryModelVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryModelVersionResult, error) {
			args := v.(LookupRegistryModelVersionArgs)
			r, err := LookupRegistryModelVersion(ctx, &args, opts...)
			var s LookupRegistryModelVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryModelVersionResultOutput)
}

type LookupRegistryModelVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	ModelName pulumi.StringInput `pulumi:"modelName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryModelVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryModelVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryModelVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelVersionResult)(nil)).Elem()
}

func (o LookupRegistryModelVersionResultOutput) ToLookupRegistryModelVersionResultOutput() LookupRegistryModelVersionResultOutput {
	return o
}

func (o LookupRegistryModelVersionResultOutput) ToLookupRegistryModelVersionResultOutputWithContext(ctx context.Context) LookupRegistryModelVersionResultOutput {
	return o
}

func (o LookupRegistryModelVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryModelVersionResult] {
	return pulumix.Output[LookupRegistryModelVersionResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryModelVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryModelVersionResultOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) ModelVersionResponse { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

// The name of the resource
func (o LookupRegistryModelVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryModelVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryModelVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryModelVersionResultOutput{})
}