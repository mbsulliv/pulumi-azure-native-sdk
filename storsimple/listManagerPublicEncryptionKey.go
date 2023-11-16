// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the symmetric encrypted public encryption key of the manager.
// Azure REST API version: 2017-06-01.
func ListManagerPublicEncryptionKey(ctx *pulumi.Context, args *ListManagerPublicEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*ListManagerPublicEncryptionKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagerPublicEncryptionKeyResult
	err := ctx.Invoke("azure-native:storsimple:listManagerPublicEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagerPublicEncryptionKeyArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents the secrets encrypted using Symmetric Encryption Key.
type ListManagerPublicEncryptionKeyResult struct {
	// The algorithm used to encrypt the "Value".
	EncryptionAlgorithm string `pulumi:"encryptionAlgorithm"`
	// The value of the secret itself. If the secret is in plaintext or null then EncryptionAlgorithm will be none.
	Value string `pulumi:"value"`
	// The thumbprint of the cert that was used to encrypt "Value".
	ValueCertificateThumbprint *string `pulumi:"valueCertificateThumbprint"`
}

func ListManagerPublicEncryptionKeyOutput(ctx *pulumi.Context, args ListManagerPublicEncryptionKeyOutputArgs, opts ...pulumi.InvokeOption) ListManagerPublicEncryptionKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagerPublicEncryptionKeyResult, error) {
			args := v.(ListManagerPublicEncryptionKeyArgs)
			r, err := ListManagerPublicEncryptionKey(ctx, &args, opts...)
			var s ListManagerPublicEncryptionKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagerPublicEncryptionKeyResultOutput)
}

type ListManagerPublicEncryptionKeyOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListManagerPublicEncryptionKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerPublicEncryptionKeyArgs)(nil)).Elem()
}

// Represents the secrets encrypted using Symmetric Encryption Key.
type ListManagerPublicEncryptionKeyResultOutput struct{ *pulumi.OutputState }

func (ListManagerPublicEncryptionKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagerPublicEncryptionKeyResult)(nil)).Elem()
}

func (o ListManagerPublicEncryptionKeyResultOutput) ToListManagerPublicEncryptionKeyResultOutput() ListManagerPublicEncryptionKeyResultOutput {
	return o
}

func (o ListManagerPublicEncryptionKeyResultOutput) ToListManagerPublicEncryptionKeyResultOutputWithContext(ctx context.Context) ListManagerPublicEncryptionKeyResultOutput {
	return o
}

// The algorithm used to encrypt the "Value".
func (o ListManagerPublicEncryptionKeyResultOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

// The value of the secret itself. If the secret is in plaintext or null then EncryptionAlgorithm will be none.
func (o ListManagerPublicEncryptionKeyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) string { return v.Value }).(pulumi.StringOutput)
}

// The thumbprint of the cert that was used to encrypt "Value".
func (o ListManagerPublicEncryptionKeyResultOutput) ValueCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListManagerPublicEncryptionKeyResult) *string { return v.ValueCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagerPublicEncryptionKeyResultOutput{})
}
