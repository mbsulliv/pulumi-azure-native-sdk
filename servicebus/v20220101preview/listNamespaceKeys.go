// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the primary and secondary connection strings for the namespace.
func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:servicebus/v20220101preview:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/ServiceBus Connection String
type ListNamespaceKeysResult struct {
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
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNamespaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysArgs)(nil)).Elem()
}

// Namespace/ServiceBus Connection String
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

// Primary connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o ListNamespaceKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace authorization rule.
func (o ListNamespaceKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListNamespaceKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace authorization rule.
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
