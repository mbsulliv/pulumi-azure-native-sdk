// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automanage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a configuration profile
// Azure REST API version: 2022-05-04.
func LookupConfigurationProfile(ctx *pulumi.Context, args *LookupConfigurationProfileArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationProfileResult
	err := ctx.Invoke("azure-native:automanage:getConfigurationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationProfileArgs struct {
	// The configuration profile name.
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the configuration profile.
type LookupConfigurationProfileResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of the configuration profile.
	Properties ConfigurationProfilePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupConfigurationProfileOutput(ctx *pulumi.Context, args LookupConfigurationProfileOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationProfileResult, error) {
			args := v.(LookupConfigurationProfileArgs)
			r, err := LookupConfigurationProfile(ctx, &args, opts...)
			var s LookupConfigurationProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationProfileResultOutput)
}

type LookupConfigurationProfileOutputArgs struct {
	// The configuration profile name.
	ConfigurationProfileName pulumi.StringInput `pulumi:"configurationProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConfigurationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileArgs)(nil)).Elem()
}

// Definition of the configuration profile.
type LookupConfigurationProfileResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationProfileResult)(nil)).Elem()
}

func (o LookupConfigurationProfileResultOutput) ToLookupConfigurationProfileResultOutput() LookupConfigurationProfileResultOutput {
	return o
}

func (o LookupConfigurationProfileResultOutput) ToLookupConfigurationProfileResultOutputWithContext(ctx context.Context) LookupConfigurationProfileResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConfigurationProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupConfigurationProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConfigurationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile.
func (o LookupConfigurationProfileResultOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) ConfigurationProfilePropertiesResponse { return v.Properties }).(ConfigurationProfilePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConfigurationProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupConfigurationProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConfigurationProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationProfileResultOutput{})
}