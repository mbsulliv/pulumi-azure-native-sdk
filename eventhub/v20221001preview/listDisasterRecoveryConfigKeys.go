// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the primary and secondary connection strings for the Namespace.
func ListDisasterRecoveryConfigKeys(ctx *pulumi.Context, args *ListDisasterRecoveryConfigKeysArgs, opts ...pulumi.InvokeOption) (*ListDisasterRecoveryConfigKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDisasterRecoveryConfigKeysResult
	err := ctx.Invoke("azure-native:eventhub/v20221001preview:listDisasterRecoveryConfigKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDisasterRecoveryConfigKeysArgs struct {
	// The Disaster Recovery configuration name
	Alias string `pulumi:"alias"`
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/EventHub Connection String
type ListDisasterRecoveryConfigKeysResult struct {
	// Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString string `pulumi:"aliasPrimaryConnectionString"`
	// Secondary  connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	// A string that describes the AuthorizationRule.
	KeyName string `pulumi:"keyName"`
	// Primary connection string of the created namespace AuthorizationRule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey string `pulumi:"primaryKey"`
	// Secondary connection string of the created namespace AuthorizationRule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListDisasterRecoveryConfigKeysOutput(ctx *pulumi.Context, args ListDisasterRecoveryConfigKeysOutputArgs, opts ...pulumi.InvokeOption) ListDisasterRecoveryConfigKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDisasterRecoveryConfigKeysResult, error) {
			args := v.(ListDisasterRecoveryConfigKeysArgs)
			r, err := ListDisasterRecoveryConfigKeys(ctx, &args, opts...)
			var s ListDisasterRecoveryConfigKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDisasterRecoveryConfigKeysResultOutput)
}

type ListDisasterRecoveryConfigKeysOutputArgs struct {
	// The Disaster Recovery configuration name
	Alias pulumi.StringInput `pulumi:"alias"`
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDisasterRecoveryConfigKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisasterRecoveryConfigKeysArgs)(nil)).Elem()
}

// Namespace/EventHub Connection String
type ListDisasterRecoveryConfigKeysResultOutput struct{ *pulumi.OutputState }

func (ListDisasterRecoveryConfigKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisasterRecoveryConfigKeysResult)(nil)).Elem()
}

func (o ListDisasterRecoveryConfigKeysResultOutput) ToListDisasterRecoveryConfigKeysResultOutput() ListDisasterRecoveryConfigKeysResultOutput {
	return o
}

func (o ListDisasterRecoveryConfigKeysResultOutput) ToListDisasterRecoveryConfigKeysResultOutputWithContext(ctx context.Context) ListDisasterRecoveryConfigKeysResultOutput {
	return o
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListDisasterRecoveryConfigKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListDisasterRecoveryConfigKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the AuthorizationRule.
func (o ListDisasterRecoveryConfigKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace AuthorizationRule.
func (o ListDisasterRecoveryConfigKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListDisasterRecoveryConfigKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace AuthorizationRule.
func (o ListDisasterRecoveryConfigKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListDisasterRecoveryConfigKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDisasterRecoveryConfigKeysResultOutput{})
}
