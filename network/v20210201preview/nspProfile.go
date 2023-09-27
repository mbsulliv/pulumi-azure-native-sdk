// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The network security perimeter profile resource
type NspProfile struct {
	pulumi.CustomResourceState

	// Version number that increases with every update to access rules within the profile.
	AccessRulesVersion pulumi.StringOutput `pulumi:"accessRulesVersion"`
	// Version number that increases with every update to diagnostic settings within the profile.
	DiagnosticSettingsVersion pulumi.StringOutput `pulumi:"diagnosticSettingsVersion"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNspProfile registers a new resource with the given unique name, arguments, and options.
func NewNspProfile(ctx *pulumi.Context,
	name string, args *NspProfileArgs, opts ...pulumi.ResourceOption) (*NspProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NspProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NspProfile
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NspProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNspProfile gets an existing NspProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNspProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspProfileState, opts ...pulumi.ResourceOption) (*NspProfile, error) {
	var resource NspProfile
	err := ctx.ReadResource("azure-native:network/v20210201preview:NspProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NspProfile resources.
type nspProfileState struct {
}

type NspProfileState struct {
}

func (NspProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspProfileState)(nil)).Elem()
}

type nspProfileArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the profile resource that is unique within a perimeter. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// The name of the NSP profile.
	ProfileName *string `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NspProfile resource.
type NspProfileArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the profile resource that is unique within a perimeter. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput
	// The name of the NSP profile.
	ProfileName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NspProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspProfileArgs)(nil)).Elem()
}

type NspProfileInput interface {
	pulumi.Input

	ToNspProfileOutput() NspProfileOutput
	ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput
}

func (*NspProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NspProfile)(nil)).Elem()
}

func (i *NspProfile) ToNspProfileOutput() NspProfileOutput {
	return i.ToNspProfileOutputWithContext(context.Background())
}

func (i *NspProfile) ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspProfileOutput)
}

func (i *NspProfile) ToOutput(ctx context.Context) pulumix.Output[*NspProfile] {
	return pulumix.Output[*NspProfile]{
		OutputState: i.ToNspProfileOutputWithContext(ctx).OutputState,
	}
}

type NspProfileOutput struct{ *pulumi.OutputState }

func (NspProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspProfile)(nil)).Elem()
}

func (o NspProfileOutput) ToNspProfileOutput() NspProfileOutput {
	return o
}

func (o NspProfileOutput) ToNspProfileOutputWithContext(ctx context.Context) NspProfileOutput {
	return o
}

func (o NspProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*NspProfile] {
	return pulumix.Output[*NspProfile]{
		OutputState: o.OutputState,
	}
}

// Version number that increases with every update to access rules within the profile.
func (o NspProfileOutput) AccessRulesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.AccessRulesVersion }).(pulumi.StringOutput)
}

// Version number that increases with every update to diagnostic settings within the profile.
func (o NspProfileOutput) DiagnosticSettingsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.DiagnosticSettingsVersion }).(pulumi.StringOutput)
}

// Resource location.
func (o NspProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NspProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource tags.
func (o NspProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NspProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspProfileOutput{})
}
