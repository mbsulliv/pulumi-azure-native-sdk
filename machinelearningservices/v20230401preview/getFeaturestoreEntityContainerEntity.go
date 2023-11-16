// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupFeaturestoreEntityContainerEntity(ctx *pulumi.Context, args *LookupFeaturestoreEntityContainerEntityArgs, opts ...pulumi.InvokeOption) (*LookupFeaturestoreEntityContainerEntityResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFeaturestoreEntityContainerEntityResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getFeaturestoreEntityContainerEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFeaturestoreEntityContainerEntityArgs struct {
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupFeaturestoreEntityContainerEntityResult struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityContainerProperties FeaturestoreEntityContainerResponse `pulumi:"featurestoreEntityContainerProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupFeaturestoreEntityContainerEntityResult
func (val *LookupFeaturestoreEntityContainerEntityResult) Defaults() *LookupFeaturestoreEntityContainerEntityResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FeaturestoreEntityContainerProperties = *tmp.FeaturestoreEntityContainerProperties.Defaults()

	return &tmp
}

func LookupFeaturestoreEntityContainerEntityOutput(ctx *pulumi.Context, args LookupFeaturestoreEntityContainerEntityOutputArgs, opts ...pulumi.InvokeOption) LookupFeaturestoreEntityContainerEntityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeaturestoreEntityContainerEntityResult, error) {
			args := v.(LookupFeaturestoreEntityContainerEntityArgs)
			r, err := LookupFeaturestoreEntityContainerEntity(ctx, &args, opts...)
			var s LookupFeaturestoreEntityContainerEntityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeaturestoreEntityContainerEntityResultOutput)
}

type LookupFeaturestoreEntityContainerEntityOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFeaturestoreEntityContainerEntityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeaturestoreEntityContainerEntityArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupFeaturestoreEntityContainerEntityResultOutput struct{ *pulumi.OutputState }

func (LookupFeaturestoreEntityContainerEntityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeaturestoreEntityContainerEntityResult)(nil)).Elem()
}

func (o LookupFeaturestoreEntityContainerEntityResultOutput) ToLookupFeaturestoreEntityContainerEntityResultOutput() LookupFeaturestoreEntityContainerEntityResultOutput {
	return o
}

func (o LookupFeaturestoreEntityContainerEntityResultOutput) ToLookupFeaturestoreEntityContainerEntityResultOutputWithContext(ctx context.Context) LookupFeaturestoreEntityContainerEntityResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupFeaturestoreEntityContainerEntityResultOutput) FeaturestoreEntityContainerProperties() FeaturestoreEntityContainerResponseOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityContainerEntityResult) FeaturestoreEntityContainerResponse {
		return v.FeaturestoreEntityContainerProperties
	}).(FeaturestoreEntityContainerResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFeaturestoreEntityContainerEntityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityContainerEntityResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFeaturestoreEntityContainerEntityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityContainerEntityResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFeaturestoreEntityContainerEntityResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityContainerEntityResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFeaturestoreEntityContainerEntityResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeaturestoreEntityContainerEntityResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeaturestoreEntityContainerEntityResultOutput{})
}
