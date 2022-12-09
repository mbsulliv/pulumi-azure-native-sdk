// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Association Subnet.
type AssociationSubnet struct {
	// Association ID.
	Id string `pulumi:"id"`
}

// AssociationSubnetInput is an input type that accepts AssociationSubnetArgs and AssociationSubnetOutput values.
// You can construct a concrete instance of `AssociationSubnetInput` via:
//
//	AssociationSubnetArgs{...}
type AssociationSubnetInput interface {
	pulumi.Input

	ToAssociationSubnetOutput() AssociationSubnetOutput
	ToAssociationSubnetOutputWithContext(context.Context) AssociationSubnetOutput
}

// Association Subnet.
type AssociationSubnetArgs struct {
	// Association ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (AssociationSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnet)(nil)).Elem()
}

func (i AssociationSubnetArgs) ToAssociationSubnetOutput() AssociationSubnetOutput {
	return i.ToAssociationSubnetOutputWithContext(context.Background())
}

func (i AssociationSubnetArgs) ToAssociationSubnetOutputWithContext(ctx context.Context) AssociationSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetOutput)
}

func (i AssociationSubnetArgs) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return i.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (i AssociationSubnetArgs) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetOutput).ToAssociationSubnetPtrOutputWithContext(ctx)
}

// AssociationSubnetPtrInput is an input type that accepts AssociationSubnetArgs, AssociationSubnetPtr and AssociationSubnetPtrOutput values.
// You can construct a concrete instance of `AssociationSubnetPtrInput` via:
//
//	        AssociationSubnetArgs{...}
//
//	or:
//
//	        nil
type AssociationSubnetPtrInput interface {
	pulumi.Input

	ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput
	ToAssociationSubnetPtrOutputWithContext(context.Context) AssociationSubnetPtrOutput
}

type associationSubnetPtrType AssociationSubnetArgs

func AssociationSubnetPtr(v *AssociationSubnetArgs) AssociationSubnetPtrInput {
	return (*associationSubnetPtrType)(v)
}

func (*associationSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnet)(nil)).Elem()
}

func (i *associationSubnetPtrType) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return i.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (i *associationSubnetPtrType) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationSubnetPtrOutput)
}

// Association Subnet.
type AssociationSubnetOutput struct{ *pulumi.OutputState }

func (AssociationSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnet)(nil)).Elem()
}

func (o AssociationSubnetOutput) ToAssociationSubnetOutput() AssociationSubnetOutput {
	return o
}

func (o AssociationSubnetOutput) ToAssociationSubnetOutputWithContext(ctx context.Context) AssociationSubnetOutput {
	return o
}

func (o AssociationSubnetOutput) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return o.ToAssociationSubnetPtrOutputWithContext(context.Background())
}

func (o AssociationSubnetOutput) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssociationSubnet) *AssociationSubnet {
		return &v
	}).(AssociationSubnetPtrOutput)
}

// Association ID.
func (o AssociationSubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AssociationSubnet) string { return v.Id }).(pulumi.StringOutput)
}

type AssociationSubnetPtrOutput struct{ *pulumi.OutputState }

func (AssociationSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnet)(nil)).Elem()
}

func (o AssociationSubnetPtrOutput) ToAssociationSubnetPtrOutput() AssociationSubnetPtrOutput {
	return o
}

func (o AssociationSubnetPtrOutput) ToAssociationSubnetPtrOutputWithContext(ctx context.Context) AssociationSubnetPtrOutput {
	return o
}

func (o AssociationSubnetPtrOutput) Elem() AssociationSubnetOutput {
	return o.ApplyT(func(v *AssociationSubnet) AssociationSubnet {
		if v != nil {
			return *v
		}
		var ret AssociationSubnet
		return ret
	}).(AssociationSubnetOutput)
}

// Association ID.
func (o AssociationSubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssociationSubnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// Association Subnet.
type AssociationSubnetResponse struct {
	// Association ID.
	Id string `pulumi:"id"`
}

// Association Subnet.
type AssociationSubnetResponseOutput struct{ *pulumi.OutputState }

func (AssociationSubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationSubnetResponse)(nil)).Elem()
}

func (o AssociationSubnetResponseOutput) ToAssociationSubnetResponseOutput() AssociationSubnetResponseOutput {
	return o
}

func (o AssociationSubnetResponseOutput) ToAssociationSubnetResponseOutputWithContext(ctx context.Context) AssociationSubnetResponseOutput {
	return o
}

// Association ID.
func (o AssociationSubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AssociationSubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type AssociationSubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (AssociationSubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssociationSubnetResponse)(nil)).Elem()
}

func (o AssociationSubnetResponsePtrOutput) ToAssociationSubnetResponsePtrOutput() AssociationSubnetResponsePtrOutput {
	return o
}

func (o AssociationSubnetResponsePtrOutput) ToAssociationSubnetResponsePtrOutputWithContext(ctx context.Context) AssociationSubnetResponsePtrOutput {
	return o
}

func (o AssociationSubnetResponsePtrOutput) Elem() AssociationSubnetResponseOutput {
	return o.ApplyT(func(v *AssociationSubnetResponse) AssociationSubnetResponse {
		if v != nil {
			return *v
		}
		var ret AssociationSubnetResponse
		return ret
	}).(AssociationSubnetResponseOutput)
}

// Association ID.
func (o AssociationSubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssociationSubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// Frontend IP Address.
type FrontendPropertiesIPAddress struct {
	// IP Address.
	Id string `pulumi:"id"`
}

// FrontendPropertiesIPAddressInput is an input type that accepts FrontendPropertiesIPAddressArgs and FrontendPropertiesIPAddressOutput values.
// You can construct a concrete instance of `FrontendPropertiesIPAddressInput` via:
//
//	FrontendPropertiesIPAddressArgs{...}
type FrontendPropertiesIPAddressInput interface {
	pulumi.Input

	ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput
	ToFrontendPropertiesIPAddressOutputWithContext(context.Context) FrontendPropertiesIPAddressOutput
}

// Frontend IP Address.
type FrontendPropertiesIPAddressArgs struct {
	// IP Address.
	Id pulumi.StringInput `pulumi:"id"`
}

func (FrontendPropertiesIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddress)(nil)).Elem()
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput {
	return i.ToFrontendPropertiesIPAddressOutputWithContext(context.Background())
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressOutput)
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return i.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (i FrontendPropertiesIPAddressArgs) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressOutput).ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx)
}

// FrontendPropertiesIPAddressPtrInput is an input type that accepts FrontendPropertiesIPAddressArgs, FrontendPropertiesIPAddressPtr and FrontendPropertiesIPAddressPtrOutput values.
// You can construct a concrete instance of `FrontendPropertiesIPAddressPtrInput` via:
//
//	        FrontendPropertiesIPAddressArgs{...}
//
//	or:
//
//	        nil
type FrontendPropertiesIPAddressPtrInput interface {
	pulumi.Input

	ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput
	ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Context) FrontendPropertiesIPAddressPtrOutput
}

type frontendPropertiesIPAddressPtrType FrontendPropertiesIPAddressArgs

func FrontendPropertiesIPAddressPtr(v *FrontendPropertiesIPAddressArgs) FrontendPropertiesIPAddressPtrInput {
	return (*frontendPropertiesIPAddressPtrType)(v)
}

func (*frontendPropertiesIPAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddress)(nil)).Elem()
}

func (i *frontendPropertiesIPAddressPtrType) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return i.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (i *frontendPropertiesIPAddressPtrType) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendPropertiesIPAddressPtrOutput)
}

// Frontend IP Address.
type FrontendPropertiesIPAddressOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddress)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressOutput() FrontendPropertiesIPAddressOutput {
	return o
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressOutput {
	return o
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return o.ToFrontendPropertiesIPAddressPtrOutputWithContext(context.Background())
}

func (o FrontendPropertiesIPAddressOutput) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendPropertiesIPAddress) *FrontendPropertiesIPAddress {
		return &v
	}).(FrontendPropertiesIPAddressPtrOutput)
}

// IP Address.
func (o FrontendPropertiesIPAddressOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendPropertiesIPAddress) string { return v.Id }).(pulumi.StringOutput)
}

type FrontendPropertiesIPAddressPtrOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddress)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressPtrOutput) ToFrontendPropertiesIPAddressPtrOutput() FrontendPropertiesIPAddressPtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressPtrOutput) ToFrontendPropertiesIPAddressPtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressPtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressPtrOutput) Elem() FrontendPropertiesIPAddressOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddress) FrontendPropertiesIPAddress {
		if v != nil {
			return *v
		}
		var ret FrontendPropertiesIPAddress
		return ret
	}).(FrontendPropertiesIPAddressOutput)
}

// IP Address.
func (o FrontendPropertiesIPAddressPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddress) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// Frontend IP Address.
type FrontendPropertiesIPAddressResponse struct {
	// IP Address.
	Id string `pulumi:"id"`
}

// Frontend IP Address.
type FrontendPropertiesIPAddressResponseOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendPropertiesIPAddressResponse)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressResponseOutput) ToFrontendPropertiesIPAddressResponseOutput() FrontendPropertiesIPAddressResponseOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponseOutput) ToFrontendPropertiesIPAddressResponseOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressResponseOutput {
	return o
}

// IP Address.
func (o FrontendPropertiesIPAddressResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendPropertiesIPAddressResponse) string { return v.Id }).(pulumi.StringOutput)
}

type FrontendPropertiesIPAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontendPropertiesIPAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendPropertiesIPAddressResponse)(nil)).Elem()
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) ToFrontendPropertiesIPAddressResponsePtrOutput() FrontendPropertiesIPAddressResponsePtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) ToFrontendPropertiesIPAddressResponsePtrOutputWithContext(ctx context.Context) FrontendPropertiesIPAddressResponsePtrOutput {
	return o
}

func (o FrontendPropertiesIPAddressResponsePtrOutput) Elem() FrontendPropertiesIPAddressResponseOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddressResponse) FrontendPropertiesIPAddressResponse {
		if v != nil {
			return *v
		}
		var ret FrontendPropertiesIPAddressResponse
		return ret
	}).(FrontendPropertiesIPAddressResponseOutput)
}

// IP Address.
func (o FrontendPropertiesIPAddressResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendPropertiesIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// Resource ID definition used by parent to reference child resources.
type ResourceIDResponse struct {
	// Resource ID of child resource.
	Id string `pulumi:"id"`
}

// Resource ID definition used by parent to reference child resources.
type ResourceIDResponseOutput struct{ *pulumi.OutputState }

func (ResourceIDResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIDResponse)(nil)).Elem()
}

func (o ResourceIDResponseOutput) ToResourceIDResponseOutput() ResourceIDResponseOutput {
	return o
}

func (o ResourceIDResponseOutput) ToResourceIDResponseOutputWithContext(ctx context.Context) ResourceIDResponseOutput {
	return o
}

// Resource ID of child resource.
func (o ResourceIDResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIDResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIDResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceIDResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceIDResponse)(nil)).Elem()
}

func (o ResourceIDResponseArrayOutput) ToResourceIDResponseArrayOutput() ResourceIDResponseArrayOutput {
	return o
}

func (o ResourceIDResponseArrayOutput) ToResourceIDResponseArrayOutputWithContext(ctx context.Context) ResourceIDResponseArrayOutput {
	return o
}

func (o ResourceIDResponseArrayOutput) Index(i pulumi.IntInput) ResourceIDResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceIDResponse {
		return vs[0].([]ResourceIDResponse)[vs[1].(int)]
	}).(ResourceIDResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssociationSubnetOutput{})
	pulumi.RegisterOutputType(AssociationSubnetPtrOutput{})
	pulumi.RegisterOutputType(AssociationSubnetResponseOutput{})
	pulumi.RegisterOutputType(AssociationSubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressPtrOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressResponseOutput{})
	pulumi.RegisterOutputType(FrontendPropertiesIPAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIDResponseOutput{})
	pulumi.RegisterOutputType(ResourceIDResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
