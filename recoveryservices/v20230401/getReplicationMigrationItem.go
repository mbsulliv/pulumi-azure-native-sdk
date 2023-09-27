// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Migration item.
func LookupReplicationMigrationItem(ctx *pulumi.Context, args *LookupReplicationMigrationItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationMigrationItemResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationMigrationItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20230401:getReplicationMigrationItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationMigrationItemArgs struct {
	// Fabric unique name.
	FabricName string `pulumi:"fabricName"`
	// Migration item name.
	MigrationItemName string `pulumi:"migrationItemName"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Migration item.
type LookupReplicationMigrationItemResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The migration item properties.
	Properties MigrationItemPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationMigrationItemOutput(ctx *pulumi.Context, args LookupReplicationMigrationItemOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationMigrationItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationMigrationItemResult, error) {
			args := v.(LookupReplicationMigrationItemArgs)
			r, err := LookupReplicationMigrationItem(ctx, &args, opts...)
			var s LookupReplicationMigrationItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationMigrationItemResultOutput)
}

type LookupReplicationMigrationItemOutputArgs struct {
	// Fabric unique name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Migration item name.
	MigrationItemName pulumi.StringInput `pulumi:"migrationItemName"`
	// Protection container name.
	ProtectionContainerName pulumi.StringInput `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationMigrationItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationMigrationItemArgs)(nil)).Elem()
}

// Migration item.
type LookupReplicationMigrationItemResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationMigrationItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationMigrationItemResult)(nil)).Elem()
}

func (o LookupReplicationMigrationItemResultOutput) ToLookupReplicationMigrationItemResultOutput() LookupReplicationMigrationItemResultOutput {
	return o
}

func (o LookupReplicationMigrationItemResultOutput) ToLookupReplicationMigrationItemResultOutputWithContext(ctx context.Context) LookupReplicationMigrationItemResultOutput {
	return o
}

func (o LookupReplicationMigrationItemResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicationMigrationItemResult] {
	return pulumix.Output[LookupReplicationMigrationItemResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupReplicationMigrationItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationMigrationItemResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationMigrationItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// The migration item properties.
func (o LookupReplicationMigrationItemResultOutput) Properties() MigrationItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) MigrationItemPropertiesResponse { return v.Properties }).(MigrationItemPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationMigrationItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationMigrationItemResultOutput{})
}
