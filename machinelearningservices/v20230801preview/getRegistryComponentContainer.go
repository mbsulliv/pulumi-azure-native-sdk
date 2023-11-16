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
func LookupRegistryComponentContainer(ctx *pulumi.Context, args *LookupRegistryComponentContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryComponentContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryComponentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getRegistryComponentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryComponentContainerArgs struct {
	// Container name.
	ComponentName string `pulumi:"componentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryComponentContainerResult struct {
	// [Required] Additional attributes of the entity.
	ComponentContainerProperties ComponentContainerResponse `pulumi:"componentContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryComponentContainerResult
func (val *LookupRegistryComponentContainerResult) Defaults() *LookupRegistryComponentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentContainerProperties = *tmp.ComponentContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryComponentContainerOutput(ctx *pulumi.Context, args LookupRegistryComponentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryComponentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryComponentContainerResult, error) {
			args := v.(LookupRegistryComponentContainerArgs)
			r, err := LookupRegistryComponentContainer(ctx, &args, opts...)
			var s LookupRegistryComponentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryComponentContainerResultOutput)
}

type LookupRegistryComponentContainerOutputArgs struct {
	// Container name.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryComponentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryComponentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryComponentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryComponentContainerResult)(nil)).Elem()
}

func (o LookupRegistryComponentContainerResultOutput) ToLookupRegistryComponentContainerResultOutput() LookupRegistryComponentContainerResultOutput {
	return o
}

func (o LookupRegistryComponentContainerResultOutput) ToLookupRegistryComponentContainerResultOutputWithContext(ctx context.Context) LookupRegistryComponentContainerResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryComponentContainerResultOutput) ComponentContainerProperties() ComponentContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) ComponentContainerResponse {
		return v.ComponentContainerProperties
	}).(ComponentContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryComponentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryComponentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryComponentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryComponentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryComponentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryComponentContainerResultOutput{})
}
