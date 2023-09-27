// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Primary and secondary connection strings to the queue.
func ListQueueKeys(ctx *pulumi.Context, args *ListQueueKeysArgs, opts ...pulumi.InvokeOption) (*ListQueueKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListQueueKeysResult
	err := ctx.Invoke("azure-native:servicebus/v20220101preview:listQueueKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListQueueKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The queue name.
	QueueName string `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/ServiceBus Connection String
type ListQueueKeysResult struct {
	// Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString string `pulumi:"aliasPrimaryConnectionString"`
	// Secondary  connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	// A string that describes the authorization rule.
	KeyName string `pulumi:"keyName"`
	// Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey string `pulumi:"primaryKey"`
	// Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListQueueKeysOutput(ctx *pulumi.Context, args ListQueueKeysOutputArgs, opts ...pulumi.InvokeOption) ListQueueKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListQueueKeysResult, error) {
			args := v.(ListQueueKeysArgs)
			r, err := ListQueueKeys(ctx, &args, opts...)
			var s ListQueueKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListQueueKeysResultOutput)
}

type ListQueueKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The queue name.
	QueueName pulumi.StringInput `pulumi:"queueName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListQueueKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueueKeysArgs)(nil)).Elem()
}

// Namespace/ServiceBus Connection String
type ListQueueKeysResultOutput struct{ *pulumi.OutputState }

func (ListQueueKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueueKeysResult)(nil)).Elem()
}

func (o ListQueueKeysResultOutput) ToListQueueKeysResultOutput() ListQueueKeysResultOutput {
	return o
}

func (o ListQueueKeysResultOutput) ToListQueueKeysResultOutputWithContext(ctx context.Context) ListQueueKeysResultOutput {
	return o
}

func (o ListQueueKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListQueueKeysResult] {
	return pulumix.Output[ListQueueKeysResult]{
		OutputState: o.OutputState,
	}
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListQueueKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListQueueKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o ListQueueKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace authorization rule.
func (o ListQueueKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListQueueKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace authorization rule.
func (o ListQueueKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListQueueKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueueKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListQueueKeysResultOutput{})
}
