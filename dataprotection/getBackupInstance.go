// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a backup instance with name in a backup vault
// Azure REST API version: 2023-01-01.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01, 2023-06-01-preview.
func LookupBackupInstance(ctx *pulumi.Context, args *LookupBackupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupBackupInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupInstanceResult
	err := ctx.Invoke("azure-native:dataprotection:getBackupInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupInstanceArgs struct {
	// The name of the backup instance.
	BackupInstanceName string `pulumi:"backupInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// BackupInstance Resource
type LookupBackupInstanceResult struct {
	// Proxy Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Proxy Resource name associated with the resource.
	Name string `pulumi:"name"`
	// BackupInstanceResource properties
	Properties BackupInstanceResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Proxy Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Proxy Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupBackupInstanceOutput(ctx *pulumi.Context, args LookupBackupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupBackupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupInstanceResult, error) {
			args := v.(LookupBackupInstanceArgs)
			r, err := LookupBackupInstance(ctx, &args, opts...)
			var s LookupBackupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupInstanceResultOutput)
}

type LookupBackupInstanceOutputArgs struct {
	// The name of the backup instance.
	BackupInstanceName pulumi.StringInput `pulumi:"backupInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the backup vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupBackupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupInstanceArgs)(nil)).Elem()
}

// BackupInstance Resource
type LookupBackupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupBackupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupInstanceResult)(nil)).Elem()
}

func (o LookupBackupInstanceResultOutput) ToLookupBackupInstanceResultOutput() LookupBackupInstanceResultOutput {
	return o
}

func (o LookupBackupInstanceResultOutput) ToLookupBackupInstanceResultOutputWithContext(ctx context.Context) LookupBackupInstanceResultOutput {
	return o
}

func (o LookupBackupInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBackupInstanceResult] {
	return pulumix.Output[LookupBackupInstanceResult]{
		OutputState: o.OutputState,
	}
}

// Proxy Resource Id represents the complete path to the resource.
func (o LookupBackupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Proxy Resource name associated with the resource.
func (o LookupBackupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// BackupInstanceResource properties
func (o LookupBackupInstanceResultOutput) Properties() BackupInstanceResponseOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) BackupInstanceResponse { return v.Properties }).(BackupInstanceResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBackupInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Proxy Resource tags.
func (o LookupBackupInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Proxy Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupBackupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupInstanceResultOutput{})
}
