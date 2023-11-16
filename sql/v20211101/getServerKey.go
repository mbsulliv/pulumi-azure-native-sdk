// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

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
	err := ctx.Invoke("azure-native:sql/v20211101:getServerKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerKeyArgs struct {
	// The name of the server key to be retrieved.
	KeyName string `pulumi:"keyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A server key.
type LookupServerKeyResult struct {
	// Key auto rotation opt-in flag. Either true or false.
	AutoRotationEnabled bool `pulumi:"autoRotationEnabled"`
	// The server key creation date.
	CreationDate string `pulumi:"creationDate"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Kind of encryption protector. This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Subregion of the server key.
	Subregion string `pulumi:"subregion"`
	// Thumbprint of the server key.
	Thumbprint string `pulumi:"thumbprint"`
	// Resource type.
	Type string `pulumi:"type"`
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
	// The name of the server key to be retrieved.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerKeyArgs)(nil)).Elem()
}

// A server key.
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

// Key auto rotation opt-in flag. Either true or false.
func (o LookupServerKeyResultOutput) AutoRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerKeyResult) bool { return v.AutoRotationEnabled }).(pulumi.BoolOutput)
}

// The server key creation date.
func (o LookupServerKeyResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServerKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of encryption protector. This is metadata used for the Azure portal experience.
func (o LookupServerKeyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupServerKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServerKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Subregion of the server key.
func (o LookupServerKeyResultOutput) Subregion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Subregion }).(pulumi.StringOutput)
}

// Thumbprint of the server key.
func (o LookupServerKeyResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupServerKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerKeyResultOutput{})
}
