// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupEnvironmentVersion(ctx *pulumi.Context, args *LookupEnvironmentVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230601preview:getEnvironmentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEnvironmentVersionArgs struct {
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
type LookupEnvironmentVersionResult struct {
	// [Required] Additional attributes of the entity.
	EnvironmentVersionProperties EnvironmentVersionResponse `pulumi:"environmentVersionProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEnvironmentVersionResult
func (val *LookupEnvironmentVersionResult) Defaults() *LookupEnvironmentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentVersionProperties = *tmp.EnvironmentVersionProperties.Defaults()

	return &tmp
}

func LookupEnvironmentVersionOutput(ctx *pulumi.Context, args LookupEnvironmentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentVersionResult, error) {
			args := v.(LookupEnvironmentVersionArgs)
			r, err := LookupEnvironmentVersion(ctx, &args, opts...)
			var s LookupEnvironmentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentVersionResultOutput)
}

type LookupEnvironmentVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionResult)(nil)).Elem()
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutput() LookupEnvironmentVersionResultOutput {
	return o
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutputWithContext(ctx context.Context) LookupEnvironmentVersionResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupEnvironmentVersionResultOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) EnvironmentVersionResponse {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnvironmentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnvironmentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnvironmentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnvironmentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentVersionResultOutput{})
}
