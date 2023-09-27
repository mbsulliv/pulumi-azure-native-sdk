// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datareplication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the replication extension.
// Azure REST API version: 2021-02-16-preview.
func LookupReplicationExtension(ctx *pulumi.Context, args *LookupReplicationExtensionArgs, opts ...pulumi.InvokeOption) (*LookupReplicationExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationExtensionResult
	err := ctx.Invoke("azure-native:datareplication:getReplicationExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationExtensionArgs struct {
	// The replication extension name.
	ReplicationExtensionName string `pulumi:"replicationExtensionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName string `pulumi:"vaultName"`
}

// Replication extension model.
type LookupReplicationExtensionResult struct {
	// Gets or sets the Id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// Replication extension model properties.
	Properties ReplicationExtensionModelPropertiesResponse `pulumi:"properties"`
	SystemData ReplicationExtensionModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupReplicationExtensionOutput(ctx *pulumi.Context, args LookupReplicationExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationExtensionResult, error) {
			args := v.(LookupReplicationExtensionArgs)
			r, err := LookupReplicationExtension(ctx, &args, opts...)
			var s LookupReplicationExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationExtensionResultOutput)
}

type LookupReplicationExtensionOutputArgs struct {
	// The replication extension name.
	ReplicationExtensionName pulumi.StringInput `pulumi:"replicationExtensionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The vault name.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupReplicationExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationExtensionArgs)(nil)).Elem()
}

// Replication extension model.
type LookupReplicationExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationExtensionResult)(nil)).Elem()
}

func (o LookupReplicationExtensionResultOutput) ToLookupReplicationExtensionResultOutput() LookupReplicationExtensionResultOutput {
	return o
}

func (o LookupReplicationExtensionResultOutput) ToLookupReplicationExtensionResultOutputWithContext(ctx context.Context) LookupReplicationExtensionResultOutput {
	return o
}

func (o LookupReplicationExtensionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicationExtensionResult] {
	return pulumix.Output[LookupReplicationExtensionResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the Id of the resource.
func (o LookupReplicationExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o LookupReplicationExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Replication extension model properties.
func (o LookupReplicationExtensionResultOutput) Properties() ReplicationExtensionModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationExtensionResult) ReplicationExtensionModelPropertiesResponse {
		return v.Properties
	}).(ReplicationExtensionModelPropertiesResponseOutput)
}

func (o LookupReplicationExtensionResultOutput) SystemData() ReplicationExtensionModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupReplicationExtensionResult) ReplicationExtensionModelResponseSystemData {
		return v.SystemData
	}).(ReplicationExtensionModelResponseSystemDataOutput)
}

// Gets or sets the type of the resource.
func (o LookupReplicationExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationExtensionResultOutput{})
}
