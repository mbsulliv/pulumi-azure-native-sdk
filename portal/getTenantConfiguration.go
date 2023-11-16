// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the tenant configuration.
// Azure REST API version: 2020-09-01-preview.
func LookupTenantConfiguration(ctx *pulumi.Context, args *LookupTenantConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupTenantConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTenantConfigurationResult
	err := ctx.Invoke("azure-native:portal:getTenantConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTenantConfigurationArgs struct {
	// The configuration name. Value must be 'default'
	ConfigurationName string `pulumi:"configurationName"`
}

// Tenant configuration.
type LookupTenantConfigurationResult struct {
	// When flag is set to true Markdown tile will require external storage configuration (URI). The inline content configuration will be prohibited.
	EnforcePrivateMarkdownStorage *bool `pulumi:"enforcePrivateMarkdownStorage"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTenantConfigurationOutput(ctx *pulumi.Context, args LookupTenantConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupTenantConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTenantConfigurationResult, error) {
			args := v.(LookupTenantConfigurationArgs)
			r, err := LookupTenantConfiguration(ctx, &args, opts...)
			var s LookupTenantConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTenantConfigurationResultOutput)
}

type LookupTenantConfigurationOutputArgs struct {
	// The configuration name. Value must be 'default'
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
}

func (LookupTenantConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantConfigurationArgs)(nil)).Elem()
}

// Tenant configuration.
type LookupTenantConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupTenantConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTenantConfigurationResult)(nil)).Elem()
}

func (o LookupTenantConfigurationResultOutput) ToLookupTenantConfigurationResultOutput() LookupTenantConfigurationResultOutput {
	return o
}

func (o LookupTenantConfigurationResultOutput) ToLookupTenantConfigurationResultOutputWithContext(ctx context.Context) LookupTenantConfigurationResultOutput {
	return o
}

// When flag is set to true Markdown tile will require external storage configuration (URI). The inline content configuration will be prohibited.
func (o LookupTenantConfigurationResultOutput) EnforcePrivateMarkdownStorage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) *bool { return v.EnforcePrivateMarkdownStorage }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTenantConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTenantConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTenantConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTenantConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTenantConfigurationResultOutput{})
}
