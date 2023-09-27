// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ResourceGuardProxyBaseResource object, used for response and request bodies for ResourceGuardProxy APIs
func LookupDppResourceGuardProxy(ctx *pulumi.Context, args *LookupDppResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupDppResourceGuardProxyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDppResourceGuardProxyResult
	err := ctx.Invoke("azure-native:dataprotection/v20230601preview:getDppResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDppResourceGuardProxyArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// name of the resource guard proxy
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// ResourceGuardProxyBaseResource object, used for response and request bodies for ResourceGuardProxy APIs
type LookupDppResourceGuardProxyResult struct {
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupDppResourceGuardProxyOutput(ctx *pulumi.Context, args LookupDppResourceGuardProxyOutputArgs, opts ...pulumi.InvokeOption) LookupDppResourceGuardProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDppResourceGuardProxyResult, error) {
			args := v.(LookupDppResourceGuardProxyArgs)
			r, err := LookupDppResourceGuardProxy(ctx, &args, opts...)
			var s LookupDppResourceGuardProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDppResourceGuardProxyResultOutput)
}

type LookupDppResourceGuardProxyOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// name of the resource guard proxy
	ResourceGuardProxyName pulumi.StringInput `pulumi:"resourceGuardProxyName"`
	// The name of the backup vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupDppResourceGuardProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDppResourceGuardProxyArgs)(nil)).Elem()
}

// ResourceGuardProxyBaseResource object, used for response and request bodies for ResourceGuardProxy APIs
type LookupDppResourceGuardProxyResultOutput struct{ *pulumi.OutputState }

func (LookupDppResourceGuardProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDppResourceGuardProxyResult)(nil)).Elem()
}

func (o LookupDppResourceGuardProxyResultOutput) ToLookupDppResourceGuardProxyResultOutput() LookupDppResourceGuardProxyResultOutput {
	return o
}

func (o LookupDppResourceGuardProxyResultOutput) ToLookupDppResourceGuardProxyResultOutputWithContext(ctx context.Context) LookupDppResourceGuardProxyResultOutput {
	return o
}

func (o LookupDppResourceGuardProxyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDppResourceGuardProxyResult] {
	return pulumix.Output[LookupDppResourceGuardProxyResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id represents the complete path to the resource.
func (o LookupDppResourceGuardProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name associated with the resource.
func (o LookupDppResourceGuardProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardProxyBaseResource properties
func (o LookupDppResourceGuardProxyResultOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) ResourceGuardProxyBaseResponse { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDppResourceGuardProxyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupDppResourceGuardProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDppResourceGuardProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDppResourceGuardProxyResultOutput{})
}
