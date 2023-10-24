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
func LookupRegistryDataContainer(ctx *pulumi.Context, args *LookupRegistryDataContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryDataContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryDataContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20231001:getRegistryDataContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryDataContainerArgs struct {
	// Container name.
	Name string `pulumi:"name"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryDataContainerResult struct {
	// [Required] Additional attributes of the entity.
	DataContainerProperties DataContainerResponse `pulumi:"dataContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryDataContainerResult
func (val *LookupRegistryDataContainerResult) Defaults() *LookupRegistryDataContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DataContainerProperties = *tmp.DataContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryDataContainerOutput(ctx *pulumi.Context, args LookupRegistryDataContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryDataContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryDataContainerResult, error) {
			args := v.(LookupRegistryDataContainerArgs)
			r, err := LookupRegistryDataContainer(ctx, &args, opts...)
			var s LookupRegistryDataContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryDataContainerResultOutput)
}

type LookupRegistryDataContainerOutputArgs struct {
	// Container name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryDataContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryDataContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryDataContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryDataContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryDataContainerResult)(nil)).Elem()
}

func (o LookupRegistryDataContainerResultOutput) ToLookupRegistryDataContainerResultOutput() LookupRegistryDataContainerResultOutput {
	return o
}

func (o LookupRegistryDataContainerResultOutput) ToLookupRegistryDataContainerResultOutputWithContext(ctx context.Context) LookupRegistryDataContainerResultOutput {
	return o
}

func (o LookupRegistryDataContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryDataContainerResult] {
	return pulumix.Output[LookupRegistryDataContainerResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryDataContainerResultOutput) DataContainerProperties() DataContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryDataContainerResult) DataContainerResponse { return v.DataContainerProperties }).(DataContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryDataContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryDataContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryDataContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryDataContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryDataContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryDataContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryDataContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryDataContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryDataContainerResultOutput{})
}
