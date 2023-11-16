// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a definition for a Developer Machine.
type DevBoxDefinition struct {
	pulumi.CustomResourceState

	// Image reference information for the currently active image (only populated during updates).
	ActiveImageReference ImageReferenceResponseOutput `pulumi:"activeImageReference"`
	// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
	HibernateSupport pulumi.StringPtrOutput `pulumi:"hibernateSupport"`
	// Image reference information.
	ImageReference ImageReferenceResponseOutput `pulumi:"imageReference"`
	// Details for image validator error. Populated when the image validation is not successful.
	ImageValidationErrorDetails ImageValidationErrorDetailsResponseOutput `pulumi:"imageValidationErrorDetails"`
	// Validation status of the configured image.
	ImageValidationStatus pulumi.StringOutput `pulumi:"imageValidationStatus"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage type used for the Operating System disk of Dev Boxes created using this definition.
	OsStorageType pulumi.StringPtrOutput `pulumi:"osStorageType"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU for Dev Boxes created using this definition.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Validation status for the Dev Box Definition.
	ValidationStatus pulumi.StringOutput `pulumi:"validationStatus"`
}

// NewDevBoxDefinition registers a new resource with the given unique name, arguments, and options.
func NewDevBoxDefinition(ctx *pulumi.Context,
	name string, args *DevBoxDefinitionArgs, opts ...pulumi.ResourceOption) (*DevBoxDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ImageReference == nil {
		return nil, errors.New("invalid value for required argument 'ImageReference'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:DevBoxDefinition"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:DevBoxDefinition"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DevBoxDefinition
	err := ctx.RegisterResource("azure-native:devcenter/v20230801preview:DevBoxDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevBoxDefinition gets an existing DevBoxDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevBoxDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevBoxDefinitionState, opts ...pulumi.ResourceOption) (*DevBoxDefinition, error) {
	var resource DevBoxDefinition
	err := ctx.ReadResource("azure-native:devcenter/v20230801preview:DevBoxDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevBoxDefinition resources.
type devBoxDefinitionState struct {
}

type DevBoxDefinitionState struct {
}

func (DevBoxDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*devBoxDefinitionState)(nil)).Elem()
}

type devBoxDefinitionArgs struct {
	// The name of the Dev Box definition.
	DevBoxDefinitionName *string `pulumi:"devBoxDefinitionName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
	HibernateSupport *string `pulumi:"hibernateSupport"`
	// Image reference information.
	ImageReference ImageReference `pulumi:"imageReference"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The storage type used for the Operating System disk of Dev Boxes created using this definition.
	OsStorageType *string `pulumi:"osStorageType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU for Dev Boxes created using this definition.
	Sku Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DevBoxDefinition resource.
type DevBoxDefinitionArgs struct {
	// The name of the Dev Box definition.
	DevBoxDefinitionName pulumi.StringPtrInput
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
	HibernateSupport pulumi.StringPtrInput
	// Image reference information.
	ImageReference ImageReferenceInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The storage type used for the Operating System disk of Dev Boxes created using this definition.
	OsStorageType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU for Dev Boxes created using this definition.
	Sku SkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DevBoxDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devBoxDefinitionArgs)(nil)).Elem()
}

type DevBoxDefinitionInput interface {
	pulumi.Input

	ToDevBoxDefinitionOutput() DevBoxDefinitionOutput
	ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput
}

func (*DevBoxDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**DevBoxDefinition)(nil)).Elem()
}

func (i *DevBoxDefinition) ToDevBoxDefinitionOutput() DevBoxDefinitionOutput {
	return i.ToDevBoxDefinitionOutputWithContext(context.Background())
}

func (i *DevBoxDefinition) ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevBoxDefinitionOutput)
}

type DevBoxDefinitionOutput struct{ *pulumi.OutputState }

func (DevBoxDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevBoxDefinition)(nil)).Elem()
}

func (o DevBoxDefinitionOutput) ToDevBoxDefinitionOutput() DevBoxDefinitionOutput {
	return o
}

func (o DevBoxDefinitionOutput) ToDevBoxDefinitionOutputWithContext(ctx context.Context) DevBoxDefinitionOutput {
	return o
}

// Image reference information for the currently active image (only populated during updates).
func (o DevBoxDefinitionOutput) ActiveImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageReferenceResponseOutput { return v.ActiveImageReference }).(ImageReferenceResponseOutput)
}

// Indicates whether Dev Boxes created with this definition are capable of hibernation. Not all images are capable of supporting hibernation. To find out more see https://aka.ms/devbox/hibernate
func (o DevBoxDefinitionOutput) HibernateSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringPtrOutput { return v.HibernateSupport }).(pulumi.StringPtrOutput)
}

// Image reference information.
func (o DevBoxDefinitionOutput) ImageReference() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageReferenceResponseOutput { return v.ImageReference }).(ImageReferenceResponseOutput)
}

// Details for image validator error. Populated when the image validation is not successful.
func (o DevBoxDefinitionOutput) ImageValidationErrorDetails() ImageValidationErrorDetailsResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) ImageValidationErrorDetailsResponseOutput {
		return v.ImageValidationErrorDetails
	}).(ImageValidationErrorDetailsResponseOutput)
}

// Validation status of the configured image.
func (o DevBoxDefinitionOutput) ImageValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.ImageValidationStatus }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DevBoxDefinitionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DevBoxDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The storage type used for the Operating System disk of Dev Boxes created using this definition.
func (o DevBoxDefinitionOutput) OsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringPtrOutput { return v.OsStorageType }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o DevBoxDefinitionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU for Dev Boxes created using this definition.
func (o DevBoxDefinitionOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DevBoxDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevBoxDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DevBoxDefinitionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DevBoxDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Validation status for the Dev Box Definition.
func (o DevBoxDefinitionOutput) ValidationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DevBoxDefinition) pulumi.StringOutput { return v.ValidationStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevBoxDefinitionOutput{})
}
