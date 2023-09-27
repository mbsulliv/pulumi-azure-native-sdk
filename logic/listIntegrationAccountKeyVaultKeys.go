// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the integration account's Key Vault keys.
// Azure REST API version: 2019-05-01.
func ListIntegrationAccountKeyVaultKeys(ctx *pulumi.Context, args *ListIntegrationAccountKeyVaultKeysArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountKeyVaultKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListIntegrationAccountKeyVaultKeysResult
	err := ctx.Invoke("azure-native:logic:listIntegrationAccountKeyVaultKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountKeyVaultKeysArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The key vault reference.
	KeyVault KeyVaultReference `pulumi:"keyVault"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The skip token.
	SkipToken *string `pulumi:"skipToken"`
}

// Collection of key vault keys.
type ListIntegrationAccountKeyVaultKeysResult struct {
	// The skip token.
	SkipToken *string `pulumi:"skipToken"`
	// The key vault keys.
	Value []KeyVaultKeyResponse `pulumi:"value"`
}

func ListIntegrationAccountKeyVaultKeysOutput(ctx *pulumi.Context, args ListIntegrationAccountKeyVaultKeysOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountKeyVaultKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountKeyVaultKeysResult, error) {
			args := v.(ListIntegrationAccountKeyVaultKeysArgs)
			r, err := ListIntegrationAccountKeyVaultKeys(ctx, &args, opts...)
			var s ListIntegrationAccountKeyVaultKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountKeyVaultKeysResultOutput)
}

type ListIntegrationAccountKeyVaultKeysOutputArgs struct {
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The key vault reference.
	KeyVault KeyVaultReferenceInput `pulumi:"keyVault"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The skip token.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListIntegrationAccountKeyVaultKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountKeyVaultKeysArgs)(nil)).Elem()
}

// Collection of key vault keys.
type ListIntegrationAccountKeyVaultKeysResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountKeyVaultKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountKeyVaultKeysResult)(nil)).Elem()
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) ToListIntegrationAccountKeyVaultKeysResultOutput() ListIntegrationAccountKeyVaultKeysResultOutput {
	return o
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) ToListIntegrationAccountKeyVaultKeysResultOutputWithContext(ctx context.Context) ListIntegrationAccountKeyVaultKeysResultOutput {
	return o
}

func (o ListIntegrationAccountKeyVaultKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListIntegrationAccountKeyVaultKeysResult] {
	return pulumix.Output[ListIntegrationAccountKeyVaultKeysResult]{
		OutputState: o.OutputState,
	}
}

// The skip token.
func (o ListIntegrationAccountKeyVaultKeysResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountKeyVaultKeysResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// The key vault keys.
func (o ListIntegrationAccountKeyVaultKeysResultOutput) Value() KeyVaultKeyResponseArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountKeyVaultKeysResult) []KeyVaultKeyResponse { return v.Value }).(KeyVaultKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountKeyVaultKeysResultOutput{})
}
