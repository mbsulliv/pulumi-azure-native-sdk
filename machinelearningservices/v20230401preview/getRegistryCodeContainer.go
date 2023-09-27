// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
func LookupRegistryCodeContainer(ctx *pulumi.Context, args *LookupRegistryCodeContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryCodeContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryCodeContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getRegistryCodeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryCodeContainerArgs struct {
	// Container name.
	CodeName string `pulumi:"codeName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryCodeContainerResult struct {
	// [Required] Additional attributes of the entity.
	CodeContainerProperties CodeContainerResponse `pulumi:"codeContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryCodeContainerResult
func (val *LookupRegistryCodeContainerResult) Defaults() *LookupRegistryCodeContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeContainerProperties = *tmp.CodeContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryCodeContainerOutput(ctx *pulumi.Context, args LookupRegistryCodeContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryCodeContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryCodeContainerResult, error) {
			args := v.(LookupRegistryCodeContainerArgs)
			r, err := LookupRegistryCodeContainer(ctx, &args, opts...)
			var s LookupRegistryCodeContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryCodeContainerResultOutput)
}

type LookupRegistryCodeContainerOutputArgs struct {
	// Container name.
	CodeName pulumi.StringInput `pulumi:"codeName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryCodeContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryCodeContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryCodeContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeContainerResult)(nil)).Elem()
}

func (o LookupRegistryCodeContainerResultOutput) ToLookupRegistryCodeContainerResultOutput() LookupRegistryCodeContainerResultOutput {
	return o
}

func (o LookupRegistryCodeContainerResultOutput) ToLookupRegistryCodeContainerResultOutputWithContext(ctx context.Context) LookupRegistryCodeContainerResultOutput {
	return o
}

func (o LookupRegistryCodeContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryCodeContainerResult] {
	return pulumix.Output[LookupRegistryCodeContainerResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryCodeContainerResultOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) CodeContainerResponse { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryCodeContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryCodeContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryCodeContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryCodeContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryCodeContainerResultOutput{})
}
