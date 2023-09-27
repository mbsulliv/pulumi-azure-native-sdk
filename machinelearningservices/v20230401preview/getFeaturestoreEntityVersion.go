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
func LookupFeaturestoreEntityVersion(ctx *pulumi.Context, args *LookupFeaturestoreEntityVersionArgs, opts ...pulumi.InvokeOption) (*LookupFeaturestoreEntityVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFeaturestoreEntityVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getFeaturestoreEntityVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFeaturestoreEntityVersionArgs struct {
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
type LookupFeaturestoreEntityVersionResult struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityVersionProperties FeaturestoreEntityVersionResponse `pulumi:"featurestoreEntityVersionProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupFeaturestoreEntityVersionResult
func (val *LookupFeaturestoreEntityVersionResult) Defaults() *LookupFeaturestoreEntityVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FeaturestoreEntityVersionProperties = *tmp.FeaturestoreEntityVersionProperties.Defaults()

	return &tmp
}

func LookupFeaturestoreEntityVersionOutput(ctx *pulumi.Context, args LookupFeaturestoreEntityVersionOutputArgs, opts ...pulumi.InvokeOption) LookupFeaturestoreEntityVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeaturestoreEntityVersionResult, error) {
			args := v.(LookupFeaturestoreEntityVersionArgs)
			r, err := LookupFeaturestoreEntityVersion(ctx, &args, opts...)
			var s LookupFeaturestoreEntityVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeaturestoreEntityVersionResultOutput)
}

type LookupFeaturestoreEntityVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFeaturestoreEntityVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeaturestoreEntityVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupFeaturestoreEntityVersionResultOutput struct{ *pulumi.OutputState }

func (LookupFeaturestoreEntityVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeaturestoreEntityVersionResult)(nil)).Elem()
}

func (o LookupFeaturestoreEntityVersionResultOutput) ToLookupFeaturestoreEntityVersionResultOutput() LookupFeaturestoreEntityVersionResultOutput {
	return o
}

func (o LookupFeaturestoreEntityVersionResultOutput) ToLookupFeaturestoreEntityVersionResultOutputWithContext(ctx context.Context) LookupFeaturestoreEntityVersionResultOutput {
	return o
}

func (o LookupFeaturestoreEntityVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeaturestoreEntityVersionResult] {
	return pulumix.Output[LookupFeaturestoreEntityVersionResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupFeaturestoreEntityVersionResultOutput) FeaturestoreEntityVersionProperties() FeaturestoreEntityVersionResponseOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityVersionResult) FeaturestoreEntityVersionResponse {
		return v.FeaturestoreEntityVersionProperties
	}).(FeaturestoreEntityVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFeaturestoreEntityVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFeaturestoreEntityVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFeaturestoreEntityVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFeaturestoreEntityVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeaturestoreEntityVersionResultOutput{})
}
