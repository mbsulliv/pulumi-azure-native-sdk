// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a backup policy belonging to a backup vault
func LookupBackupPolicy(ctx *pulumi.Context, args *LookupBackupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupPolicyResult
	err := ctx.Invoke("azure-native:dataprotection/v20230601preview:getBackupPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupPolicyArgs struct {
	BackupPolicyName string `pulumi:"backupPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// BaseBackupPolicy resource
type LookupBackupPolicyResult struct {
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// BaseBackupPolicyResource properties
	Properties BackupPolicyResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupBackupPolicyOutput(ctx *pulumi.Context, args LookupBackupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupPolicyResult, error) {
			args := v.(LookupBackupPolicyArgs)
			r, err := LookupBackupPolicy(ctx, &args, opts...)
			var s LookupBackupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupPolicyResultOutput)
}

type LookupBackupPolicyOutputArgs struct {
	BackupPolicyName pulumi.StringInput `pulumi:"backupPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the backup vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupBackupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyArgs)(nil)).Elem()
}

// BaseBackupPolicy resource
type LookupBackupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupPolicyResult)(nil)).Elem()
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutput() LookupBackupPolicyResultOutput {
	return o
}

func (o LookupBackupPolicyResultOutput) ToLookupBackupPolicyResultOutputWithContext(ctx context.Context) LookupBackupPolicyResultOutput {
	return o
}

// Resource Id represents the complete path to the resource.
func (o LookupBackupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name associated with the resource.
func (o LookupBackupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// BaseBackupPolicyResource properties
func (o LookupBackupPolicyResultOutput) Properties() BackupPolicyResponseOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) BackupPolicyResponse { return v.Properties }).(BackupPolicyResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBackupPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupBackupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupPolicyResultOutput{})
}
