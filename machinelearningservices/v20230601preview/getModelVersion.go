// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
func LookupModelVersion(ctx *pulumi.Context, args *LookupModelVersionArgs, opts ...pulumi.InvokeOption) (*LookupModelVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupModelVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230601preview:getModelVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupModelVersionArgs struct {
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
type LookupModelVersionResult struct {
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

// Defaults sets the appropriate defaults for LookupModelVersionResult
func (val *LookupModelVersionResult) Defaults() *LookupModelVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ModelVersionProperties = *tmp.ModelVersionProperties.Defaults()

	return &tmp
}

func LookupModelVersionOutput(ctx *pulumi.Context, args LookupModelVersionOutputArgs, opts ...pulumi.InvokeOption) LookupModelVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelVersionResult, error) {
			args := v.(LookupModelVersionArgs)
			r, err := LookupModelVersion(ctx, &args, opts...)
			var s LookupModelVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelVersionResultOutput)
}

type LookupModelVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupModelVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupModelVersionResultOutput struct{ *pulumi.OutputState }

func (LookupModelVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelVersionResult)(nil)).Elem()
}

func (o LookupModelVersionResultOutput) ToLookupModelVersionResultOutput() LookupModelVersionResultOutput {
	return o
}

func (o LookupModelVersionResultOutput) ToLookupModelVersionResultOutputWithContext(ctx context.Context) LookupModelVersionResultOutput {
	return o
}

func (o LookupModelVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupModelVersionResult] {
	return pulumix.Output[LookupModelVersionResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupModelVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupModelVersionResultOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v LookupModelVersionResult) ModelVersionResponse { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

// The name of the resource
func (o LookupModelVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupModelVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupModelVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupModelVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelVersionResultOutput{})
}
