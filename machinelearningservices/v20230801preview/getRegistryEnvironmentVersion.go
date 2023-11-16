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
func LookupRegistryEnvironmentVersion(ctx *pulumi.Context, args *LookupRegistryEnvironmentVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryEnvironmentVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryEnvironmentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getRegistryEnvironmentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryEnvironmentVersionArgs struct {
	// Container name. This is case-sensitive.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version string `pulumi:"version"`
}

// Azure Resource Manager resource envelope.
type LookupRegistryEnvironmentVersionResult struct {
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

// Defaults sets the appropriate defaults for LookupRegistryEnvironmentVersionResult
func (val *LookupRegistryEnvironmentVersionResult) Defaults() *LookupRegistryEnvironmentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentVersionProperties = *tmp.EnvironmentVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryEnvironmentVersionOutput(ctx *pulumi.Context, args LookupRegistryEnvironmentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryEnvironmentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryEnvironmentVersionResult, error) {
			args := v.(LookupRegistryEnvironmentVersionArgs)
			r, err := LookupRegistryEnvironmentVersion(ctx, &args, opts...)
			var s LookupRegistryEnvironmentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryEnvironmentVersionResultOutput)
}

type LookupRegistryEnvironmentVersionOutputArgs struct {
	// Container name. This is case-sensitive.
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryEnvironmentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupRegistryEnvironmentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryEnvironmentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryEnvironmentVersionResult)(nil)).Elem()
}

func (o LookupRegistryEnvironmentVersionResultOutput) ToLookupRegistryEnvironmentVersionResultOutput() LookupRegistryEnvironmentVersionResultOutput {
	return o
}

func (o LookupRegistryEnvironmentVersionResultOutput) ToLookupRegistryEnvironmentVersionResultOutputWithContext(ctx context.Context) LookupRegistryEnvironmentVersionResultOutput {
	return o
}

// [Required] Additional attributes of the entity.
func (o LookupRegistryEnvironmentVersionResultOutput) EnvironmentVersionProperties() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) EnvironmentVersionResponse {
		return v.EnvironmentVersionProperties
	}).(EnvironmentVersionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupRegistryEnvironmentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRegistryEnvironmentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRegistryEnvironmentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRegistryEnvironmentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryEnvironmentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryEnvironmentVersionResultOutput{})
}
