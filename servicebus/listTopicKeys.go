// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the primary and secondary connection strings for the topic.
// Azure REST API version: 2022-01-01-preview.
func ListTopicKeys(ctx *pulumi.Context, args *ListTopicKeysArgs, opts ...pulumi.InvokeOption) (*ListTopicKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListTopicKeysResult
	err := ctx.Invoke("azure-native:servicebus:listTopicKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTopicKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName string `pulumi:"topicName"`
}

// Namespace/ServiceBus Connection String
type ListTopicKeysResult struct {
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

func ListTopicKeysOutput(ctx *pulumi.Context, args ListTopicKeysOutputArgs, opts ...pulumi.InvokeOption) ListTopicKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTopicKeysResult, error) {
			args := v.(ListTopicKeysArgs)
			r, err := ListTopicKeys(ctx, &args, opts...)
			var s ListTopicKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTopicKeysResultOutput)
}

type ListTopicKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The topic name.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (ListTopicKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicKeysArgs)(nil)).Elem()
}

// Namespace/ServiceBus Connection String
type ListTopicKeysResultOutput struct{ *pulumi.OutputState }

func (ListTopicKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTopicKeysResult)(nil)).Elem()
}

func (o ListTopicKeysResultOutput) ToListTopicKeysResultOutput() ListTopicKeysResultOutput {
	return o
}

func (o ListTopicKeysResultOutput) ToListTopicKeysResultOutputWithContext(ctx context.Context) ListTopicKeysResultOutput {
	return o
}

func (o ListTopicKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListTopicKeysResult] {
	return pulumix.Output[ListTopicKeysResult]{
		OutputState: o.OutputState,
	}
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListTopicKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListTopicKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o ListTopicKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace authorization rule.
func (o ListTopicKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListTopicKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace authorization rule.
func (o ListTopicKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListTopicKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListTopicKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTopicKeysResultOutput{})
}
