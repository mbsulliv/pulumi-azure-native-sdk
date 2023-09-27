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
func LookupCodeContainer(ctx *pulumi.Context, args *LookupCodeContainerArgs, opts ...pulumi.InvokeOption) (*LookupCodeContainerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCodeContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getCodeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCodeContainerArgs struct {
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupCodeContainerResult struct {
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

// Defaults sets the appropriate defaults for LookupCodeContainerResult
func (val *LookupCodeContainerResult) Defaults() *LookupCodeContainerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeContainerProperties = *tmp.CodeContainerProperties.Defaults()

	return &tmp
}

func LookupCodeContainerOutput(ctx *pulumi.Context, args LookupCodeContainerOutputArgs, opts ...pulumi.InvokeOption) LookupCodeContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeContainerResult, error) {
			args := v.(LookupCodeContainerArgs)
			r, err := LookupCodeContainer(ctx, &args, opts...)
			var s LookupCodeContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeContainerResultOutput)
}

type LookupCodeContainerOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodeContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeContainerArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupCodeContainerResultOutput struct{ *pulumi.OutputState }

func (LookupCodeContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeContainerResult)(nil)).Elem()
}

func (o LookupCodeContainerResultOutput) ToLookupCodeContainerResultOutput() LookupCodeContainerResultOutput {
	return o
}

func (o LookupCodeContainerResultOutput) ToLookupCodeContainerResultOutputWithContext(ctx context.Context) LookupCodeContainerResultOutput {
	return o
}

func (o LookupCodeContainerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCodeContainerResult] {
	return pulumix.Output[LookupCodeContainerResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupCodeContainerResultOutput) CodeContainerProperties() CodeContainerResponseOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) CodeContainerResponse { return v.CodeContainerProperties }).(CodeContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCodeContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCodeContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCodeContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCodeContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeContainerResultOutput{})
}
