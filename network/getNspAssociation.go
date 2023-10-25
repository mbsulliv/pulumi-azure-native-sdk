// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified NSP association by name.
// Azure REST API version: 2021-02-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-08-01-preview.
func LookupNspAssociation(ctx *pulumi.Context, args *LookupNspAssociationArgs, opts ...pulumi.InvokeOption) (*LookupNspAssociationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNspAssociationResult
	err := ctx.Invoke("azure-native:network:getNspAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspAssociationArgs struct {
	// The name of the NSP association.
	AssociationName string `pulumi:"associationName"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The NSP resource association resource
type LookupNspAssociationResult struct {
	// Access mode on the association.
	AccessMode *string `pulumi:"accessMode"`
	// Specifies if there are provisioning issues
	HasProvisioningIssues string `pulumi:"hasProvisioningIssues"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The PaaS resource to be associated.
	PrivateLinkResource *SubResourceResponse `pulumi:"privateLinkResource"`
	// Profile id to which the PaaS resource is associated.
	Profile *SubResourceResponse `pulumi:"profile"`
	// The provisioning state of the resource  association resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNspAssociationOutput(ctx *pulumi.Context, args LookupNspAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupNspAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspAssociationResult, error) {
			args := v.(LookupNspAssociationArgs)
			r, err := LookupNspAssociation(ctx, &args, opts...)
			var s LookupNspAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspAssociationResultOutput)
}

type LookupNspAssociationOutputArgs struct {
	// The name of the NSP association.
	AssociationName pulumi.StringInput `pulumi:"associationName"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationArgs)(nil)).Elem()
}

// The NSP resource association resource
type LookupNspAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupNspAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAssociationResult)(nil)).Elem()
}

func (o LookupNspAssociationResultOutput) ToLookupNspAssociationResultOutput() LookupNspAssociationResultOutput {
	return o
}

func (o LookupNspAssociationResultOutput) ToLookupNspAssociationResultOutputWithContext(ctx context.Context) LookupNspAssociationResultOutput {
	return o
}

func (o LookupNspAssociationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNspAssociationResult] {
	return pulumix.Output[LookupNspAssociationResult]{
		OutputState: o.OutputState,
	}
}

// Access mode on the association.
func (o LookupNspAssociationResultOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *string { return v.AccessMode }).(pulumi.StringPtrOutput)
}

// Specifies if there are provisioning issues
func (o LookupNspAssociationResultOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNspAssociationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupNspAssociationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNspAssociationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The PaaS resource to be associated.
func (o LookupNspAssociationResultOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *SubResourceResponse { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

// Profile id to which the PaaS resource is associated.
func (o LookupNspAssociationResultOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) *SubResourceResponse { return v.Profile }).(SubResourceResponsePtrOutput)
}

// The provisioning state of the resource  association resource.
func (o LookupNspAssociationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupNspAssociationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNspAssociationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAssociationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspAssociationResultOutput{})
}
