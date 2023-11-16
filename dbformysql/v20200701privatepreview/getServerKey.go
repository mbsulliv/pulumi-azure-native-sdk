// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a server key.
func LookupServerKey(ctx *pulumi.Context, args *LookupServerKeyArgs, opts ...pulumi.InvokeOption) (*LookupServerKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerKeyResult
	err := ctx.Invoke("azure-native:dbformysql/v20200701privatepreview:getServerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerKeyArgs struct {
	// The name of the server key.
	KeyName string `pulumi:"keyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A MySQL Server key.
type LookupServerKeyResult struct {
	// The key creation date.
	CreationDate string `pulumi:"creationDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Kind of encryption protector used to protect the key.
	Kind string `pulumi:"kind"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The key type like 'AzureKeyVault'.
	ServerKeyType string `pulumi:"serverKeyType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The URI of the key.
	Uri *string `pulumi:"uri"`
}

func LookupServerKeyOutput(ctx *pulumi.Context, args LookupServerKeyOutputArgs, opts ...pulumi.InvokeOption) LookupServerKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerKeyResult, error) {
			args := v.(LookupServerKeyArgs)
			r, err := LookupServerKey(ctx, &args, opts...)
			var s LookupServerKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerKeyResultOutput)
}

type LookupServerKeyOutputArgs struct {
	// The name of the server key.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerKeyArgs)(nil)).Elem()
}

// A MySQL Server key.
type LookupServerKeyResultOutput struct{ *pulumi.OutputState }

func (LookupServerKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerKeyResult)(nil)).Elem()
}

func (o LookupServerKeyResultOutput) ToLookupServerKeyResultOutput() LookupServerKeyResultOutput {
	return o
}

func (o LookupServerKeyResultOutput) ToLookupServerKeyResultOutputWithContext(ctx context.Context) LookupServerKeyResultOutput {
	return o
}

// The key creation date.
func (o LookupServerKeyResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of encryption protector used to protect the key.
func (o LookupServerKeyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServerKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The key type like 'AzureKeyVault'.
func (o LookupServerKeyResultOutput) ServerKeyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.ServerKeyType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

// The URI of the key.
func (o LookupServerKeyResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerKeyResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerKeyResultOutput{})
}
