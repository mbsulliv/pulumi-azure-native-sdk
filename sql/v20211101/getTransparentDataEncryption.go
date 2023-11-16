// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a logical database's transparent data encryption.
func LookupTransparentDataEncryption(ctx *pulumi.Context, args *LookupTransparentDataEncryptionArgs, opts ...pulumi.InvokeOption) (*LookupTransparentDataEncryptionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTransparentDataEncryptionResult
	err := ctx.Invoke("azure-native:sql/v20211101:getTransparentDataEncryption", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTransparentDataEncryptionArgs struct {
	// The name of the logical database for which the transparent data encryption is defined.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the transparent data encryption configuration.
	TdeName string `pulumi:"tdeName"`
}

// A logical database transparent data encryption state.
type LookupTransparentDataEncryptionResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the state of the transparent data encryption.
	State string `pulumi:"state"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupTransparentDataEncryptionOutput(ctx *pulumi.Context, args LookupTransparentDataEncryptionOutputArgs, opts ...pulumi.InvokeOption) LookupTransparentDataEncryptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTransparentDataEncryptionResult, error) {
			args := v.(LookupTransparentDataEncryptionArgs)
			r, err := LookupTransparentDataEncryption(ctx, &args, opts...)
			var s LookupTransparentDataEncryptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTransparentDataEncryptionResultOutput)
}

type LookupTransparentDataEncryptionOutputArgs struct {
	// The name of the logical database for which the transparent data encryption is defined.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the transparent data encryption configuration.
	TdeName pulumi.StringInput `pulumi:"tdeName"`
}

func (LookupTransparentDataEncryptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransparentDataEncryptionArgs)(nil)).Elem()
}

// A logical database transparent data encryption state.
type LookupTransparentDataEncryptionResultOutput struct{ *pulumi.OutputState }

func (LookupTransparentDataEncryptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTransparentDataEncryptionResult)(nil)).Elem()
}

func (o LookupTransparentDataEncryptionResultOutput) ToLookupTransparentDataEncryptionResultOutput() LookupTransparentDataEncryptionResultOutput {
	return o
}

func (o LookupTransparentDataEncryptionResultOutput) ToLookupTransparentDataEncryptionResultOutputWithContext(ctx context.Context) LookupTransparentDataEncryptionResultOutput {
	return o
}

// Resource ID.
func (o LookupTransparentDataEncryptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupTransparentDataEncryptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the state of the transparent data encryption.
func (o LookupTransparentDataEncryptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupTransparentDataEncryptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTransparentDataEncryptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTransparentDataEncryptionResultOutput{})
}
