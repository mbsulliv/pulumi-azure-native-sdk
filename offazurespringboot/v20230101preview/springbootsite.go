// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The springbootsites envelope resource definition.
type Springbootsite struct {
	pulumi.CustomResourceState

	// The extended location definition.
	ExtendedLocation SpringbootsitesModelResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The springbootsites resource definition.
	Properties SpringbootsitesPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSpringbootsite registers a new resource with the given unique name, arguments, and options.
func NewSpringbootsite(ctx *pulumi.Context,
	name string, args *SpringbootsiteArgs, opts ...pulumi.ResourceOption) (*Springbootsite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazurespringboot/v20230101preview:springbootsite"),
		},
		{
			Type: pulumi.String("azure-native:offazurespringboot:Springbootsite"),
		},
		{
			Type: pulumi.String("azure-native:offazurespringboot:springbootsite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Springbootsite
	err := ctx.RegisterResource("azure-native:offazurespringboot/v20230101preview:Springbootsite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpringbootsite gets an existing Springbootsite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpringbootsite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpringbootsiteState, opts ...pulumi.ResourceOption) (*Springbootsite, error) {
	var resource Springbootsite
	err := ctx.ReadResource("azure-native:offazurespringboot/v20230101preview:Springbootsite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Springbootsite resources.
type springbootsiteState struct {
}

type SpringbootsiteState struct {
}

func (SpringbootsiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*springbootsiteState)(nil)).Elem()
}

type springbootsiteArgs struct {
	// The extended location definition.
	ExtendedLocation *SpringbootsitesModelExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The springbootsites resource definition.
	Properties *SpringbootsitesProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The springbootsites name.
	SpringbootsitesName *string `pulumi:"springbootsitesName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Springbootsite resource.
type SpringbootsiteArgs struct {
	// The extended location definition.
	ExtendedLocation SpringbootsitesModelExtendedLocationPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The springbootsites resource definition.
	Properties SpringbootsitesPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The springbootsites name.
	SpringbootsitesName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SpringbootsiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*springbootsiteArgs)(nil)).Elem()
}

type SpringbootsiteInput interface {
	pulumi.Input

	ToSpringbootsiteOutput() SpringbootsiteOutput
	ToSpringbootsiteOutputWithContext(ctx context.Context) SpringbootsiteOutput
}

func (*Springbootsite) ElementType() reflect.Type {
	return reflect.TypeOf((**Springbootsite)(nil)).Elem()
}

func (i *Springbootsite) ToSpringbootsiteOutput() SpringbootsiteOutput {
	return i.ToSpringbootsiteOutputWithContext(context.Background())
}

func (i *Springbootsite) ToSpringbootsiteOutputWithContext(ctx context.Context) SpringbootsiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpringbootsiteOutput)
}

type SpringbootsiteOutput struct{ *pulumi.OutputState }

func (SpringbootsiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Springbootsite)(nil)).Elem()
}

func (o SpringbootsiteOutput) ToSpringbootsiteOutput() SpringbootsiteOutput {
	return o
}

func (o SpringbootsiteOutput) ToSpringbootsiteOutputWithContext(ctx context.Context) SpringbootsiteOutput {
	return o
}

// The extended location definition.
func (o SpringbootsiteOutput) ExtendedLocation() SpringbootsitesModelResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *Springbootsite) SpringbootsitesModelResponseExtendedLocationPtrOutput {
		return v.ExtendedLocation
	}).(SpringbootsitesModelResponseExtendedLocationPtrOutput)
}

// The geo-location where the resource lives
func (o SpringbootsiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Springbootsite) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SpringbootsiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Springbootsite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The springbootsites resource definition.
func (o SpringbootsiteOutput) Properties() SpringbootsitesPropertiesResponseOutput {
	return o.ApplyT(func(v *Springbootsite) SpringbootsitesPropertiesResponseOutput { return v.Properties }).(SpringbootsitesPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SpringbootsiteOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Springbootsite) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SpringbootsiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Springbootsite) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SpringbootsiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Springbootsite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SpringbootsiteOutput{})
}