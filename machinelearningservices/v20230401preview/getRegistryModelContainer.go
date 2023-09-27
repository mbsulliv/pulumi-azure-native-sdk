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
func LookupRegistryModelContainer(ctx *pulumi.Context, args *LookupRegistryModelContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryModelContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryModelContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getRegistryModelContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryModelContainerArgs struct {
	// Container name. This is case-sensitive.
	ModelName string `pulumi:"modelName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryModelContainerResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// [Required] Additional attributes of the entity.
	ModelContainerProperties ModelContainerResponse `pulumi:"modelContainerProperties"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryModelContainerResult
func (val *LookupRegistryModelContainerResult) Defaults() *LookupRegistryModelContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelContainerProperties = *tmp.ModelContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryModelContainerOutput(ctx *pulumi.Context, args LookupRegistryModelContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryModelContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryModelContainerResult, error) {
			args := v.(LookupRegistryModelContainerArgs)
			r, err := LookupRegistryModelContainer(ctx, &args, opts...)
			var s LookupRegistryModelContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryModelContainerResultOutput)
}

type LookupRegistryModelContainerOutputArgs struct {
	// Container name. This is case-sensitive.
	ModelName pulumi.StringInput `pulumi:"modelName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryModelContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryModelContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryModelContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryModelContainerResult)(nil)).Elem()
}

func (o LookupRegistryModelContainerResultOutput) ToLookupRegistryModelContainerResultOutput() LookupRegistryModelContainerResultOutput {
	return o
}

func (o LookupRegistryModelContainerResultOutput) ToLookupRegistryModelContainerResultOutputWithContext(ctx context.Context) LookupRegistryModelContainerResultOutput {
	return o
}

func (o LookupRegistryModelContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryModelContainerResult] {
	return pulumix.Output[LookupRegistryModelContainerResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryModelContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryModelContainerResultOutput) ModelContainerProperties() ModelContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) ModelContainerResponse { return v.ModelContainerProperties }).(ModelContainerResponseOutput)
}

// The name of the resource
func (o LookupRegistryModelContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryModelContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryModelContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryModelContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryModelContainerResultOutput{})
}
