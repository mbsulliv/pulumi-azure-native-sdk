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
func LookupEnvironmentContainer(ctx *pulumi.Context, args *LookupEnvironmentContainerArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20231001:getEnvironmentContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEnvironmentContainerArgs struct {
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentContainerResult struct {
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

// Defaults sets the appropriate defaults for LookupEnvironmentContainerResult
func (val *LookupEnvironmentContainerResult) Defaults() *LookupEnvironmentContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentContainerProperties = *tmp.EnvironmentContainerProperties.Defaults()

	return &tmp
}

func LookupEnvironmentContainerOutput(ctx *pulumi.Context, args LookupEnvironmentContainerOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentContainerResult, error) {
			args := v.(LookupEnvironmentContainerArgs)
			r, err := LookupEnvironmentContainer(ctx, &args, opts...)
			var s LookupEnvironmentContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentContainerResultOutput)
}

type LookupEnvironmentContainerOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentContainerResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentContainerResult)(nil)).Elem()
}

func (o LookupEnvironmentContainerResultOutput) ToLookupEnvironmentContainerResultOutput() LookupEnvironmentContainerResultOutput {
	return o
}

func (o LookupEnvironmentContainerResultOutput) ToLookupEnvironmentContainerResultOutputWithContext(ctx context.Context) LookupEnvironmentContainerResultOutput {
	return o
}

func (o LookupEnvironmentContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnvironmentContainerResult] {
	return pulumix.Output[LookupEnvironmentContainerResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupEnvironmentContainerResultOutput) EnvironmentContainerProperties() EnvironmentContainerResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) EnvironmentContainerResponse {
		return v.EnvironmentContainerProperties
	}).(EnvironmentContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnvironmentContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnvironmentContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnvironmentContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnvironmentContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentContainerResultOutput{})
}