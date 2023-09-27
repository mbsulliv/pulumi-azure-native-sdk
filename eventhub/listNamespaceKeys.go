// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the primary and secondary connection strings for the Namespace.
// Azure REST API version: 2022-10-01-preview.
func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:eventhub:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/EventHub Connection String
type ListNamespaceKeysResult struct {
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

func ListNamespaceKeysOutput(ctx *pulumi.Context, args ListNamespaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListNamespaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNamespaceKeysResult, error) {
			args := v.(ListNamespaceKeysArgs)
			r, err := ListNamespaceKeys(ctx, &args, opts...)
			var s ListNamespaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNamespaceKeysResultOutput)
}

type ListNamespaceKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNamespaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysArgs)(nil)).Elem()
}

// Namespace/EventHub Connection String
type ListNamespaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListNamespaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysResult)(nil)).Elem()
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutput() ListNamespaceKeysResultOutput {
	return o
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutputWithContext(ctx context.Context) ListNamespaceKeysResultOutput {
	return o
}

func (o ListNamespaceKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListNamespaceKeysResult] {
	return pulumix.Output[ListNamespaceKeysResult]{
		OutputState: o.OutputState,
	}
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the AuthorizationRule.
func (o ListNamespaceKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace AuthorizationRule.
func (o ListNamespaceKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListNamespaceKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace AuthorizationRule.
func (o ListNamespaceKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListNamespaceKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNamespaceKeysResultOutput{})
}
