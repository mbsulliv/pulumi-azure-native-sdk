// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a credential.
func LookupCredentialOperation(ctx *pulumi.Context, args *LookupCredentialOperationArgs, opts ...pulumi.InvokeOption) (*LookupCredentialOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCredentialOperationResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getCredentialOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialOperationArgs struct {
	// Credential name
	CredentialName string `pulumi:"credentialName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Credential resource type.
type LookupCredentialOperationResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Managed Identity Credential properties.
	Properties ManagedIdentityCredentialResponse `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupCredentialOperationOutput(ctx *pulumi.Context, args LookupCredentialOperationOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCredentialOperationResult, error) {
			args := v.(LookupCredentialOperationArgs)
			r, err := LookupCredentialOperation(ctx, &args, opts...)
			var s LookupCredentialOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCredentialOperationResultOutput)
}

type LookupCredentialOperationOutputArgs struct {
	// Credential name
	CredentialName pulumi.StringInput `pulumi:"credentialName"`
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCredentialOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationArgs)(nil)).Elem()
}

// Credential resource type.
type LookupCredentialOperationResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationResult)(nil)).Elem()
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutput() LookupCredentialOperationResultOutput {
	return o
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutputWithContext(ctx context.Context) LookupCredentialOperationResultOutput {
	return o
}

func (o LookupCredentialOperationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCredentialOperationResult] {
	return pulumix.Output[LookupCredentialOperationResult]{
		OutputState: o.OutputState,
	}
}

// Etag identifies change in the resource.
func (o LookupCredentialOperationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupCredentialOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupCredentialOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Managed Identity Credential properties.
func (o LookupCredentialOperationResultOutput) Properties() ManagedIdentityCredentialResponseOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) ManagedIdentityCredentialResponse { return v.Properties }).(ManagedIdentityCredentialResponseOutput)
}

// The resource type.
func (o LookupCredentialOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialOperationResultOutput{})
}
