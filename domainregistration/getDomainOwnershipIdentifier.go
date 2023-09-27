// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domainregistration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Get ownership identifier for domain
// Azure REST API version: 2022-09-01.
func LookupDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupDomainOwnershipIdentifierResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:domainregistration:getDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainOwnershipIdentifierArgs struct {
	// Name of domain.
	DomainName string `pulumi:"domainName"`
	// Name of identifier.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Domain ownership Identifier.
type LookupDomainOwnershipIdentifierResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Ownership Id.
	OwnershipId *string `pulumi:"ownershipId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupDomainOwnershipIdentifierOutput(ctx *pulumi.Context, args LookupDomainOwnershipIdentifierOutputArgs, opts ...pulumi.InvokeOption) LookupDomainOwnershipIdentifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainOwnershipIdentifierResult, error) {
			args := v.(LookupDomainOwnershipIdentifierArgs)
			r, err := LookupDomainOwnershipIdentifier(ctx, &args, opts...)
			var s LookupDomainOwnershipIdentifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainOwnershipIdentifierResultOutput)
}

type LookupDomainOwnershipIdentifierOutputArgs struct {
	// Name of domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of identifier.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOwnershipIdentifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainOwnershipIdentifierArgs)(nil)).Elem()
}

// Domain ownership Identifier.
type LookupDomainOwnershipIdentifierResultOutput struct{ *pulumi.OutputState }

func (LookupDomainOwnershipIdentifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainOwnershipIdentifierResult)(nil)).Elem()
}

func (o LookupDomainOwnershipIdentifierResultOutput) ToLookupDomainOwnershipIdentifierResultOutput() LookupDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupDomainOwnershipIdentifierResultOutput) ToLookupDomainOwnershipIdentifierResultOutputWithContext(ctx context.Context) LookupDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupDomainOwnershipIdentifierResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDomainOwnershipIdentifierResult] {
	return pulumix.Output[LookupDomainOwnershipIdentifierResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o LookupDomainOwnershipIdentifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupDomainOwnershipIdentifierResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupDomainOwnershipIdentifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Name }).(pulumi.StringOutput)
}

// Ownership Id.
func (o LookupDomainOwnershipIdentifierResultOutput) OwnershipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) *string { return v.OwnershipId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupDomainOwnershipIdentifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainOwnershipIdentifierResultOutput{})
}
