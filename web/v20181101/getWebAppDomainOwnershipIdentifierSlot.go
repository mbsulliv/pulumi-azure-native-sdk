// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get domain ownership identifier for web app.
func LookupWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppDomainOwnershipIdentifierSlotResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppDomainOwnershipIdentifierSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierSlotArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot string `pulumi:"slot"`
}

// A domain specific resource identifier.
type LookupWebAppDomainOwnershipIdentifierSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppDomainOwnershipIdentifierSlotOutput(ctx *pulumi.Context, args LookupWebAppDomainOwnershipIdentifierSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDomainOwnershipIdentifierSlotResult, error) {
			args := v.(LookupWebAppDomainOwnershipIdentifierSlotArgs)
			r, err := LookupWebAppDomainOwnershipIdentifierSlot(ctx, &args, opts...)
			var s LookupWebAppDomainOwnershipIdentifierSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDomainOwnershipIdentifierSlotResultOutput)
}

type LookupWebAppDomainOwnershipIdentifierSlotOutputArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName pulumi.StringInput `pulumi:"domainOwnershipIdentifierName"`
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppDomainOwnershipIdentifierSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierSlotArgs)(nil)).Elem()
}

// A domain specific resource identifier.
type LookupWebAppDomainOwnershipIdentifierSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierSlotResult)(nil)).Elem()
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ToLookupWebAppDomainOwnershipIdentifierSlotResultOutput() LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ToLookupWebAppDomainOwnershipIdentifierSlotResultOutputWithContext(ctx context.Context) LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppDomainOwnershipIdentifierSlotResult] {
	return pulumix.Output[LookupWebAppDomainOwnershipIdentifierSlotResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDomainOwnershipIdentifierSlotResultOutput{})
}
