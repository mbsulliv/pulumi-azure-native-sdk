// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the ACS and SAS connection strings for the Event Hub.
func ListEventHubKeys(ctx *pulumi.Context, args *ListEventHubKeysArgs, opts ...pulumi.InvokeOption) (*ListEventHubKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListEventHubKeysResult
	err := ctx.Invoke("azure-native:eventhub/v20221001preview:listEventHubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEventHubKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/EventHub Connection String
type ListEventHubKeysResult struct {
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

func ListEventHubKeysOutput(ctx *pulumi.Context, args ListEventHubKeysOutputArgs, opts ...pulumi.InvokeOption) ListEventHubKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEventHubKeysResult, error) {
			args := v.(ListEventHubKeysArgs)
			r, err := ListEventHubKeys(ctx, &args, opts...)
			var s ListEventHubKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEventHubKeysResultOutput)
}

type ListEventHubKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName pulumi.StringInput `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListEventHubKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEventHubKeysArgs)(nil)).Elem()
}

// Namespace/EventHub Connection String
type ListEventHubKeysResultOutput struct{ *pulumi.OutputState }

func (ListEventHubKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEventHubKeysResult)(nil)).Elem()
}

func (o ListEventHubKeysResultOutput) ToListEventHubKeysResultOutput() ListEventHubKeysResultOutput {
	return o
}

func (o ListEventHubKeysResultOutput) ToListEventHubKeysResultOutputWithContext(ctx context.Context) ListEventHubKeysResultOutput {
	return o
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListEventHubKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListEventHubKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the AuthorizationRule.
func (o ListEventHubKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace AuthorizationRule.
func (o ListEventHubKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListEventHubKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace AuthorizationRule.
func (o ListEventHubKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListEventHubKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEventHubKeysResultOutput{})
}
