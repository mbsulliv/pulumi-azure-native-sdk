// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
// Azure REST API version: 2023-04-01.
func LookupRegistryCodeVersion(ctx *pulumi.Context, args *LookupRegistryCodeVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryCodeVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryCodeVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices:getRegistryCodeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryCodeVersionArgs struct {
	// Container name.
	CodeName string `pulumi:"codeName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version string `pulumi:"version"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryCodeVersionResult struct {
	// [Required] Additional attributes of the entity.
	CodeVersionProperties CodeVersionResponse `pulumi:"codeVersionProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryCodeVersionResult
func (val *LookupRegistryCodeVersionResult) Defaults() *LookupRegistryCodeVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeVersionProperties = *tmp.CodeVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryCodeVersionOutput(ctx *pulumi.Context, args LookupRegistryCodeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryCodeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryCodeVersionResult, error) {
			args := v.(LookupRegistryCodeVersionArgs)
			r, err := LookupRegistryCodeVersion(ctx, &args, opts...)
			var s LookupRegistryCodeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryCodeVersionResultOutput)
}

type LookupRegistryCodeVersionOutputArgs struct {
	// Container name.
	CodeName pulumi.StringInput `pulumi:"codeName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier.
	Version pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryCodeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryCodeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryCodeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeVersionResult)(nil)).Elem()
}

func (o LookupRegistryCodeVersionResultOutput) ToLookupRegistryCodeVersionResultOutput() LookupRegistryCodeVersionResultOutput {
	return o
}

func (o LookupRegistryCodeVersionResultOutput) ToLookupRegistryCodeVersionResultOutputWithContext(ctx context.Context) LookupRegistryCodeVersionResultOutput {
	return o
}

func (o LookupRegistryCodeVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryCodeVersionResult] {
	return pulumix.Output[LookupRegistryCodeVersionResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryCodeVersionResultOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) CodeVersionResponse { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryCodeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryCodeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryCodeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryCodeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryCodeVersionResultOutput{})
}
