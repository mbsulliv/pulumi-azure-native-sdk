// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Primary and secondary connection strings to the hybrid connection.
func ListHybridConnectionKeys(ctx *pulumi.Context, args *ListHybridConnectionKeysArgs, opts ...pulumi.InvokeOption) (*ListHybridConnectionKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListHybridConnectionKeysResult
	err := ctx.Invoke("azure-native:relay/v20211101:listHybridConnectionKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListHybridConnectionKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The hybrid connection name.
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/Relay Connection String
type ListHybridConnectionKeysResult struct {
	// A string that describes the authorization rule.
	KeyName *string `pulumi:"keyName"`
	// Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// A base64-encoded 256-bit secondary key for signing and validating the SAS token.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListHybridConnectionKeysOutput(ctx *pulumi.Context, args ListHybridConnectionKeysOutputArgs, opts ...pulumi.InvokeOption) ListHybridConnectionKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListHybridConnectionKeysResult, error) {
			args := v.(ListHybridConnectionKeysArgs)
			r, err := ListHybridConnectionKeys(ctx, &args, opts...)
			var s ListHybridConnectionKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListHybridConnectionKeysResultOutput)
}

type ListHybridConnectionKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The hybrid connection name.
	HybridConnectionName pulumi.StringInput `pulumi:"hybridConnectionName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListHybridConnectionKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHybridConnectionKeysArgs)(nil)).Elem()
}

// Namespace/Relay Connection String
type ListHybridConnectionKeysResultOutput struct{ *pulumi.OutputState }

func (ListHybridConnectionKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListHybridConnectionKeysResult)(nil)).Elem()
}

func (o ListHybridConnectionKeysResultOutput) ToListHybridConnectionKeysResultOutput() ListHybridConnectionKeysResultOutput {
	return o
}

func (o ListHybridConnectionKeysResultOutput) ToListHybridConnectionKeysResultOutputWithContext(ctx context.Context) ListHybridConnectionKeysResultOutput {
	return o
}

func (o ListHybridConnectionKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListHybridConnectionKeysResult] {
	return pulumix.Output[ListHybridConnectionKeysResult]{
		OutputState: o.OutputState,
	}
}

// A string that describes the authorization rule.
func (o ListHybridConnectionKeysResultOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

// Primary connection string of the created namespace authorization rule.
func (o ListHybridConnectionKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListHybridConnectionKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// Secondary connection string of the created namespace authorization rule.
func (o ListHybridConnectionKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

// A base64-encoded 256-bit secondary key for signing and validating the SAS token.
func (o ListHybridConnectionKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListHybridConnectionKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListHybridConnectionKeysResultOutput{})
}
