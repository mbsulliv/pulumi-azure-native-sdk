// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get domain ownership identifier for web app.
func LookupWebAppDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A domain specific resource identifier.
type LookupWebAppDomainOwnershipIdentifierResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
	// String representation of the identity.
	Value *string `pulumi:"value"`
}

func LookupWebAppDomainOwnershipIdentifierOutput(ctx *pulumi.Context, args LookupWebAppDomainOwnershipIdentifierOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDomainOwnershipIdentifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDomainOwnershipIdentifierResult, error) {
			args := v.(LookupWebAppDomainOwnershipIdentifierArgs)
			r, err := LookupWebAppDomainOwnershipIdentifier(ctx, &args, opts...)
			var s LookupWebAppDomainOwnershipIdentifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDomainOwnershipIdentifierResultOutput)
}

type LookupWebAppDomainOwnershipIdentifierOutputArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName pulumi.StringInput `pulumi:"domainOwnershipIdentifierName"`
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDomainOwnershipIdentifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierArgs)(nil)).Elem()
}

// A domain specific resource identifier.
type LookupWebAppDomainOwnershipIdentifierResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDomainOwnershipIdentifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierResult)(nil)).Elem()
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) ToLookupWebAppDomainOwnershipIdentifierResultOutput() LookupWebAppDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) ToLookupWebAppDomainOwnershipIdentifierResultOutputWithContext(ctx context.Context) LookupWebAppDomainOwnershipIdentifierResultOutput {
	return o
}

// Resource Id.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Name }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Type }).(pulumi.StringOutput)
}

// String representation of the identity.
func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDomainOwnershipIdentifierResultOutput{})
}
