// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupRegistryEnvironmentContainer(ctx *pulumi.Context, args *LookupRegistryEnvironmentContainerArgs, opts ...pulumi.InvokeOption) (*LookupRegistryEnvironmentContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryEnvironmentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401:getRegistryEnvironmentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryEnvironmentContainerArgs struct {
	// Container name. This is case-sensitive.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryEnvironmentContainerResult struct {
	// [Required] Additional attributes of the entity.
	EnvironmentContainerProperties EnvironmentContainerResponse `pulumi:"environmentContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRegistryEnvironmentContainerResult
func (val *LookupRegistryEnvironmentContainerResult) Defaults() *LookupRegistryEnvironmentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentContainerProperties = *tmp.EnvironmentContainerProperties.Defaults()

	return &tmp
}

func LookupRegistryEnvironmentContainerOutput(ctx *pulumi.Context, args LookupRegistryEnvironmentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryEnvironmentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryEnvironmentContainerResult, error) {
			args := v.(LookupRegistryEnvironmentContainerArgs)
			r, err := LookupRegistryEnvironmentContainer(ctx, &args, opts...)
			var s LookupRegistryEnvironmentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryEnvironmentContainerResultOutput)
}

type LookupRegistryEnvironmentContainerOutputArgs struct {
	// Container name. This is case-sensitive.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryEnvironmentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryEnvironmentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryEnvironmentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentContainerResult)(nil)).Elem()
}

func (o LookupRegistryEnvironmentContainerResultOutput) ToLookupRegistryEnvironmentContainerResultOutput() LookupRegistryEnvironmentContainerResultOutput {
	return o
}

func (o LookupRegistryEnvironmentContainerResultOutput) ToLookupRegistryEnvironmentContainerResultOutputWithContext(ctx context.Context) LookupRegistryEnvironmentContainerResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryEnvironmentContainerResultOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) EnvironmentContainerResponse {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryEnvironmentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryEnvironmentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryEnvironmentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryEnvironmentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryEnvironmentContainerResultOutput{})
}
