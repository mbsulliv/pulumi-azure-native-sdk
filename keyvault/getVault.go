// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Azure key vault.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2018-02-14-preview, 2023-07-01.
func LookupVault(ctx *pulumi.Context, args *LookupVaultArgs, opts ...pulumi.InvokeOption) (*LookupVaultResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVaultResult
	err := ctx.Invoke("azure-native:keyvault:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVaultArgs struct {
	// The name of the Resource Group to which the vault belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the vault.
	VaultName string `pulumi:"vaultName"`
}

// Resource information with extended details.
type LookupVaultResult struct {
	// Fully qualified identifier of the key vault resource.
	Id string `pulumi:"id"`
	// Azure location of the key vault resource.
	Location *string `pulumi:"location"`
	// Name of the key vault resource.
	Name string `pulumi:"name"`
	// Properties of the vault
	Properties VaultPropertiesResponse `pulumi:"properties"`
	// System metadata for the key vault.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags assigned to the key vault resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupVaultResult
func (val *LookupVaultResult) Defaults() *LookupVaultResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupVaultOutput(ctx *pulumi.Context, args LookupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVaultResult, error) {
			args := v.(LookupVaultArgs)
			r, err := LookupVault(ctx, &args, opts...)
			var s LookupVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVaultResultOutput)
}

type LookupVaultOutputArgs struct {
	// The name of the Resource Group to which the vault belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultArgs)(nil)).Elem()
}

// Resource information with extended details.
type LookupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVaultResult)(nil)).Elem()
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutput() LookupVaultResultOutput {
	return o
}

func (o LookupVaultResultOutput) ToLookupVaultResultOutputWithContext(ctx context.Context) LookupVaultResultOutput {
	return o
}

// Fully qualified identifier of the key vault resource.
func (o LookupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location of the key vault resource.
func (o LookupVaultResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVaultResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the key vault resource.
func (o LookupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the vault
func (o LookupVaultResultOutput) Properties() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVaultResult) VaultPropertiesResponse { return v.Properties }).(VaultPropertiesResponseOutput)
}

// System metadata for the key vault.
func (o LookupVaultResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVaultResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags assigned to the key vault resource.
func (o LookupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type of the key vault resource.
func (o LookupVaultResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVaultResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVaultResultOutput{})
}
