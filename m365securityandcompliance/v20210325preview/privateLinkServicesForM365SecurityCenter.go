// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210325preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The description of the service.
type PrivateLinkServicesForM365SecurityCenter struct {
	pulumi.CustomResourceState

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	// The kind of the service.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponseOutput `pulumi:"properties"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkServicesForM365SecurityCenter registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForM365SecurityCenterArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-native:m365securityandcompliance:PrivateLinkServicesForM365SecurityCenter"),
		},
		{
			Type: pulumi.String("azure-native:m365securityandcompliance:privateLinkServicesForM365SecurityCenter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.RegisterResource("azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForM365SecurityCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkServicesForM365SecurityCenter gets an existing PrivateLinkServicesForM365SecurityCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkServicesForM365SecurityCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForM365SecurityCenterState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForM365SecurityCenter, error) {
	var resource PrivateLinkServicesForM365SecurityCenter
	err := ctx.ReadResource("azure-native:m365securityandcompliance/v20210325preview:PrivateLinkServicesForM365SecurityCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkServicesForM365SecurityCenter resources.
type privateLinkServicesForM365SecurityCenterState struct {
}

type PrivateLinkServicesForM365SecurityCenterState struct {
}

func (PrivateLinkServicesForM365SecurityCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterState)(nil)).Elem()
}

type privateLinkServicesForM365SecurityCenterArgs struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind Kind `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The common properties of a service.
	Properties *ServicesProperties `pulumi:"properties"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName *string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkServicesForM365SecurityCenter resource.
type PrivateLinkServicesForM365SecurityCenterArgs struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ServicesResourceIdentityPtrInput
	// The kind of the service.
	Kind KindInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The common properties of a service.
	Properties ServicesPropertiesPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateLinkServicesForM365SecurityCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForM365SecurityCenterArgs)(nil)).Elem()
}

type PrivateLinkServicesForM365SecurityCenterInput interface {
	pulumi.Input

	ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput
	ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput
}

func (*PrivateLinkServicesForM365SecurityCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365SecurityCenter)(nil)).Elem()
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return i.ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForM365SecurityCenterOutput)
}

func (i *PrivateLinkServicesForM365SecurityCenter) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkServicesForM365SecurityCenter] {
	return pulumix.Output[*PrivateLinkServicesForM365SecurityCenter]{
		OutputState: i.ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx).OutputState,
	}
}

type PrivateLinkServicesForM365SecurityCenterOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForM365SecurityCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForM365SecurityCenter)(nil)).Elem()
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutput() PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToPrivateLinkServicesForM365SecurityCenterOutputWithContext(ctx context.Context) PrivateLinkServicesForM365SecurityCenterOutput {
	return o
}

func (o PrivateLinkServicesForM365SecurityCenterOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateLinkServicesForM365SecurityCenter] {
	return pulumix.Output[*PrivateLinkServicesForM365SecurityCenter]{
		OutputState: o.OutputState,
	}
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

// The kind of the service.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) ServicesPropertiesResponseOutput {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

// Required property for system data
func (o PrivateLinkServicesForM365SecurityCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o PrivateLinkServicesForM365SecurityCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForM365SecurityCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForM365SecurityCenterOutput{})
}
