// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A profile is a logical grouping of endpoints that share the same settings.
type Profile struct {
	pulumi.CustomResourceState

	// Key-Value pair representing additional properties for profiles.
	ExtendedProperties pulumi.StringMapOutput `pulumi:"extendedProperties"`
	// The Id of the frontdoor.
	FrontDoorId pulumi.StringOutput `pulumi:"frontDoorId"`
	// Managed service identity (system assigned and/or user assigned identities).
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
	OriginResponseTimeoutSeconds pulumi.IntPtrOutput `pulumi:"originResponseTimeoutSeconds"`
	// Provisioning status of the profile.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the profile.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:Profile"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20240201:Profile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Profile
	err := ctx.RegisterResource("azure-native:cdn/v20230501:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure-native:cdn/v20230501:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
}

type ProfileState struct {
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// Managed service identity (system assigned and/or user assigned identities).
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
	OriginResponseTimeoutSeconds *int `pulumi:"originResponseTimeoutSeconds"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the resource group.
	ProfileName *string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
	Sku Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// Managed service identity (system assigned and/or user assigned identities).
	Identity ManagedServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
	OriginResponseTimeoutSeconds pulumi.IntPtrInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium or CDN profile which is unique within the resource group.
	ProfileName pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
	Sku SkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}

type ProfileInput interface {
	pulumi.Input

	ToProfileOutput() ProfileOutput
	ToProfileOutputWithContext(ctx context.Context) ProfileOutput
}

func (*Profile) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (i *Profile) ToProfileOutput() ProfileOutput {
	return i.ToProfileOutputWithContext(context.Background())
}

func (i *Profile) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileOutput)
}

type ProfileOutput struct{ *pulumi.OutputState }

func (ProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Profile)(nil)).Elem()
}

func (o ProfileOutput) ToProfileOutput() ProfileOutput {
	return o
}

func (o ProfileOutput) ToProfileOutputWithContext(ctx context.Context) ProfileOutput {
	return o
}

// Key-Value pair representing additional properties for profiles.
func (o ProfileOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

// The Id of the frontdoor.
func (o ProfileOutput) FrontDoorId() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.FrontDoorId }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities).
func (o ProfileOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Profile) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Kind of the profile. Used by portal to differentiate traditional CDN profile and new AFD profile.
func (o ProfileOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o ProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o ProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Send and receive timeout on forwarding request to the origin. When timeout is reached, the request fails and returns.
func (o ProfileOutput) OriginResponseTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Profile) pulumi.IntPtrOutput { return v.OriginResponseTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// Provisioning status of the profile.
func (o ProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the profile.
func (o ProfileOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The pricing tier (defines Azure Front Door Standard or Premium or a CDN provider, feature list and rate) of the profile.
func (o ProfileOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Profile) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Read only system data
func (o ProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Profile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Profile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileOutput{})
}
