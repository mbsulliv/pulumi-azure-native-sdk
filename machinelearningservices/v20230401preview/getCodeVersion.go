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
func LookupCodeVersion(ctx *pulumi.Context, args *LookupCodeVersionArgs, opts ...pulumi.InvokeOption) (*LookupCodeVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCodeVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getCodeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCodeVersionArgs struct {
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupCodeVersionResult struct {
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

// Defaults sets the appropriate defaults for LookupCodeVersionResult
func (val *LookupCodeVersionResult) Defaults() *LookupCodeVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeVersionProperties = *tmp.CodeVersionProperties.Defaults()

	return &tmp
}

func LookupCodeVersionOutput(ctx *pulumi.Context, args LookupCodeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupCodeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeVersionResult, error) {
			args := v.(LookupCodeVersionArgs)
			r, err := LookupCodeVersion(ctx, &args, opts...)
			var s LookupCodeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeVersionResultOutput)
}

type LookupCodeVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupCodeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupCodeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeVersionResult)(nil)).Elem()
}

func (o LookupCodeVersionResultOutput) ToLookupCodeVersionResultOutput() LookupCodeVersionResultOutput {
	return o
}

func (o LookupCodeVersionResultOutput) ToLookupCodeVersionResultOutputWithContext(ctx context.Context) LookupCodeVersionResultOutput {
	return o
}

func (o LookupCodeVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCodeVersionResult] {
	return pulumix.Output[LookupCodeVersionResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupCodeVersionResultOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) CodeVersionResponse { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCodeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCodeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCodeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCodeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeVersionResultOutput{})
}
