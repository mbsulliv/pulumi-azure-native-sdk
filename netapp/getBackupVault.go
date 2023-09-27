// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the Backup Vault
// Azure REST API version: 2022-11-01-preview.
func LookupBackupVault(ctx *pulumi.Context, args *LookupBackupVaultArgs, opts ...pulumi.InvokeOption) (*LookupBackupVaultResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupVaultResult
	err := ctx.Invoke("azure-native:netapp:getBackupVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupVaultArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the Backup Vault
	BackupVaultName string `pulumi:"backupVaultName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Backup Vault information
type LookupBackupVaultResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBackupVaultOutput(ctx *pulumi.Context, args LookupBackupVaultOutputArgs, opts ...pulumi.InvokeOption) LookupBackupVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupVaultResult, error) {
			args := v.(LookupBackupVaultArgs)
			r, err := LookupBackupVault(ctx, &args, opts...)
			var s LookupBackupVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupVaultResultOutput)
}

type LookupBackupVaultOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Backup Vault
	BackupVaultName pulumi.StringInput `pulumi:"backupVaultName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBackupVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultArgs)(nil)).Elem()
}

// Backup Vault information
type LookupBackupVaultResultOutput struct{ *pulumi.OutputState }

func (LookupBackupVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupVaultResult)(nil)).Elem()
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutput() LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ToLookupBackupVaultResultOutputWithContext(ctx context.Context) LookupBackupVaultResultOutput {
	return o
}

func (o LookupBackupVaultResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBackupVaultResult] {
	return pulumix.Output[LookupBackupVaultResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBackupVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupBackupVaultResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBackupVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupBackupVaultResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBackupVaultResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBackupVaultResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBackupVaultResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupVaultResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupVaultResultOutput{})
}
