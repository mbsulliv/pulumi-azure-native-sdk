// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a SQL pool's transparent data encryption configuration.
// Azure REST API version: 2021-06-01.
//
// Other available API versions: 2021-06-01-preview.
func LookupSqlPoolTransparentDataEncryption(ctx *pulumi.Context, args *LookupSqlPoolTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolTransparentDataEncryptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlPoolTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:synapse:getSqlPoolTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolTransparentDataEncryptionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the transparent data encryption configuration.
	TransparentDataEncryptionName string `pulumi:"transparentDataEncryptionName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Sql pool transparent data encryption configuration.
type LookupSqlPoolTransparentDataEncryptionResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the database transparent data encryption.
	Status *string `pulumi:"status"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlPoolTransparentDataEncryptionOutput(ctx *pulumi.Context, args LookupSqlPoolTransparentDataEncryptionOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolTransparentDataEncryptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolTransparentDataEncryptionResult, error) {
			args := v.(LookupSqlPoolTransparentDataEncryptionArgs)
			r, err := LookupSqlPoolTransparentDataEncryption(ctx, &args, opts...)
			var s LookupSqlPoolTransparentDataEncryptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolTransparentDataEncryptionResultOutput)
}

type LookupSqlPoolTransparentDataEncryptionOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName pulumi.StringInput `pulumi:"sqlPoolName"`
	// The name of the transparent data encryption configuration.
	TransparentDataEncryptionName pulumi.StringInput `pulumi:"transparentDataEncryptionName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolTransparentDataEncryptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolTransparentDataEncryptionArgs)(nil)).Elem()
}

// Represents a Sql pool transparent data encryption configuration.
type LookupSqlPoolTransparentDataEncryptionResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolTransparentDataEncryptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolTransparentDataEncryptionResult)(nil)).Elem()
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) ToLookupSqlPoolTransparentDataEncryptionResultOutput() LookupSqlPoolTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) ToLookupSqlPoolTransparentDataEncryptionResultOutputWithContext(ctx context.Context) LookupSqlPoolTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupSqlPoolTransparentDataEncryptionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlPoolTransparentDataEncryptionResult] {
	return pulumix.Output[LookupSqlPoolTransparentDataEncryptionResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the database transparent data encryption.
func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlPoolTransparentDataEncryptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolTransparentDataEncryptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolTransparentDataEncryptionResultOutput{})
}
