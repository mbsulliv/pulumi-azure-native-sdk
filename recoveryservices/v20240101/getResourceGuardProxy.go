// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns ResourceGuardProxy under vault and with the name referenced in request
func LookupResourceGuardProxy(ctx *pulumi.Context, args *LookupResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardProxyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceGuardProxyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20240101:getResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardProxyArgs struct {
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

type LookupResourceGuardProxyResult struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupResourceGuardProxyOutput(ctx *pulumi.Context, args LookupResourceGuardProxyOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGuardProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGuardProxyResult, error) {
			args := v.(LookupResourceGuardProxyArgs)
			r, err := LookupResourceGuardProxy(ctx, &args, opts...)
			var s LookupResourceGuardProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceGuardProxyResultOutput)
}

type LookupResourceGuardProxyOutputArgs struct {
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceGuardProxyName pulumi.StringInput `pulumi:"resourceGuardProxyName"`
	// The name of the recovery services vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupResourceGuardProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyArgs)(nil)).Elem()
}

type LookupResourceGuardProxyResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGuardProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyResult)(nil)).Elem()
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutput() LookupResourceGuardProxyResultOutput {
	return o
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutputWithContext(ctx context.Context) LookupResourceGuardProxyResultOutput {
	return o
}

// Optional ETag.
func (o LookupResourceGuardProxyResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id represents the complete path to the resource.
func (o LookupResourceGuardProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupResourceGuardProxyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o LookupResourceGuardProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardProxyBaseResource properties
func (o LookupResourceGuardProxyResultOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) ResourceGuardProxyBaseResponse { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

// Resource tags.
func (o LookupResourceGuardProxyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupResourceGuardProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGuardProxyResultOutput{})
}
