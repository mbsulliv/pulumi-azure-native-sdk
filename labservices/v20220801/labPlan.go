// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lab Plans act as a permission container for creating labs via labs.azure.com. Additionally, they can provide a set of default configurations that will apply at the time of creating a lab, but these defaults can still be overwritten.
type LabPlan struct {
	pulumi.CustomResourceState

	// The allowed regions for the lab creator to use when creating labs using this lab plan.
	AllowedRegions pulumi.StringArrayOutput `pulumi:"allowedRegions"`
	// The default lab shutdown profile. This can be changed on a lab resource and only provides a default profile.
	DefaultAutoShutdownProfile AutoShutdownProfileResponsePtrOutput `pulumi:"defaultAutoShutdownProfile"`
	// The default lab connection profile. This can be changed on a lab resource and only provides a default profile.
	DefaultConnectionProfile ConnectionProfileResponsePtrOutput `pulumi:"defaultConnectionProfile"`
	// The lab plan network profile. To enforce lab network policies they must be defined here and cannot be changed when there are existing labs associated with this lab plan.
	DefaultNetworkProfile LabPlanNetworkProfileResponsePtrOutput `pulumi:"defaultNetworkProfile"`
	// Managed Identity Information
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Base Url of the lms instance this lab plan can link lab rosters against.
	LinkedLmsInstance pulumi.StringPtrOutput `pulumi:"linkedLmsInstance"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Current provisioning state of the lab plan.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource ID of the Shared Image Gallery attached to this lab plan. When saving a lab template virtual machine image it will be persisted in this gallery. Shared images from the gallery can be made available to use when creating new labs.
	SharedGalleryId pulumi.StringPtrOutput `pulumi:"sharedGalleryId"`
	// Support contact information and instructions for users of the lab plan. This information is displayed to lab owners and virtual machine users for all labs in the lab plan.
	SupportInfo SupportInfoResponsePtrOutput `pulumi:"supportInfo"`
	// Metadata pertaining to creation and last modification of the lab plan.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLabPlan registers a new resource with the given unique name, arguments, and options.
func NewLabPlan(ctx *pulumi.Context,
	name string, args *LabPlanArgs, opts ...pulumi.ResourceOption) (*LabPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DefaultAutoShutdownProfile != nil {
		args.DefaultAutoShutdownProfile = args.DefaultAutoShutdownProfile.ToAutoShutdownProfilePtrOutput().ApplyT(func(v *AutoShutdownProfile) *AutoShutdownProfile { return v.Defaults() }).(AutoShutdownProfilePtrOutput)
	}
	if args.DefaultConnectionProfile != nil {
		args.DefaultConnectionProfile = args.DefaultConnectionProfile.ToConnectionProfilePtrOutput().ApplyT(func(v *ConnectionProfile) *ConnectionProfile { return v.Defaults() }).(ConnectionProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices:LabPlan"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211001preview:LabPlan"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:LabPlan"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20230607:LabPlan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LabPlan
	err := ctx.RegisterResource("azure-native:labservices/v20220801:LabPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabPlan gets an existing LabPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabPlanState, opts ...pulumi.ResourceOption) (*LabPlan, error) {
	var resource LabPlan
	err := ctx.ReadResource("azure-native:labservices/v20220801:LabPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LabPlan resources.
type labPlanState struct {
}

type LabPlanState struct {
}

func (LabPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*labPlanState)(nil)).Elem()
}

type labPlanArgs struct {
	// The allowed regions for the lab creator to use when creating labs using this lab plan.
	AllowedRegions []string `pulumi:"allowedRegions"`
	// The default lab shutdown profile. This can be changed on a lab resource and only provides a default profile.
	DefaultAutoShutdownProfile *AutoShutdownProfile `pulumi:"defaultAutoShutdownProfile"`
	// The default lab connection profile. This can be changed on a lab resource and only provides a default profile.
	DefaultConnectionProfile *ConnectionProfile `pulumi:"defaultConnectionProfile"`
	// The lab plan network profile. To enforce lab network policies they must be defined here and cannot be changed when there are existing labs associated with this lab plan.
	DefaultNetworkProfile *LabPlanNetworkProfile `pulumi:"defaultNetworkProfile"`
	// Managed Identity Information
	Identity *Identity `pulumi:"identity"`
	// The name of the lab plan that uniquely identifies it within containing resource group. Used in resource URIs and in UI.
	LabPlanName *string `pulumi:"labPlanName"`
	// Base Url of the lms instance this lab plan can link lab rosters against.
	LinkedLmsInstance *string `pulumi:"linkedLmsInstance"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource ID of the Shared Image Gallery attached to this lab plan. When saving a lab template virtual machine image it will be persisted in this gallery. Shared images from the gallery can be made available to use when creating new labs.
	SharedGalleryId *string `pulumi:"sharedGalleryId"`
	// Support contact information and instructions for users of the lab plan. This information is displayed to lab owners and virtual machine users for all labs in the lab plan.
	SupportInfo *SupportInfo `pulumi:"supportInfo"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LabPlan resource.
type LabPlanArgs struct {
	// The allowed regions for the lab creator to use when creating labs using this lab plan.
	AllowedRegions pulumi.StringArrayInput
	// The default lab shutdown profile. This can be changed on a lab resource and only provides a default profile.
	DefaultAutoShutdownProfile AutoShutdownProfilePtrInput
	// The default lab connection profile. This can be changed on a lab resource and only provides a default profile.
	DefaultConnectionProfile ConnectionProfilePtrInput
	// The lab plan network profile. To enforce lab network policies they must be defined here and cannot be changed when there are existing labs associated with this lab plan.
	DefaultNetworkProfile LabPlanNetworkProfilePtrInput
	// Managed Identity Information
	Identity IdentityPtrInput
	// The name of the lab plan that uniquely identifies it within containing resource group. Used in resource URIs and in UI.
	LabPlanName pulumi.StringPtrInput
	// Base Url of the lms instance this lab plan can link lab rosters against.
	LinkedLmsInstance pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource ID of the Shared Image Gallery attached to this lab plan. When saving a lab template virtual machine image it will be persisted in this gallery. Shared images from the gallery can be made available to use when creating new labs.
	SharedGalleryId pulumi.StringPtrInput
	// Support contact information and instructions for users of the lab plan. This information is displayed to lab owners and virtual machine users for all labs in the lab plan.
	SupportInfo SupportInfoPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LabPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labPlanArgs)(nil)).Elem()
}

type LabPlanInput interface {
	pulumi.Input

	ToLabPlanOutput() LabPlanOutput
	ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput
}

func (*LabPlan) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlan)(nil)).Elem()
}

func (i *LabPlan) ToLabPlanOutput() LabPlanOutput {
	return i.ToLabPlanOutputWithContext(context.Background())
}

func (i *LabPlan) ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabPlanOutput)
}

type LabPlanOutput struct{ *pulumi.OutputState }

func (LabPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabPlan)(nil)).Elem()
}

func (o LabPlanOutput) ToLabPlanOutput() LabPlanOutput {
	return o
}

func (o LabPlanOutput) ToLabPlanOutputWithContext(ctx context.Context) LabPlanOutput {
	return o
}

// The allowed regions for the lab creator to use when creating labs using this lab plan.
func (o LabPlanOutput) AllowedRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringArrayOutput { return v.AllowedRegions }).(pulumi.StringArrayOutput)
}

// The default lab shutdown profile. This can be changed on a lab resource and only provides a default profile.
func (o LabPlanOutput) DefaultAutoShutdownProfile() AutoShutdownProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) AutoShutdownProfileResponsePtrOutput { return v.DefaultAutoShutdownProfile }).(AutoShutdownProfileResponsePtrOutput)
}

// The default lab connection profile. This can be changed on a lab resource and only provides a default profile.
func (o LabPlanOutput) DefaultConnectionProfile() ConnectionProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) ConnectionProfileResponsePtrOutput { return v.DefaultConnectionProfile }).(ConnectionProfileResponsePtrOutput)
}

// The lab plan network profile. To enforce lab network policies they must be defined here and cannot be changed when there are existing labs associated with this lab plan.
func (o LabPlanOutput) DefaultNetworkProfile() LabPlanNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) LabPlanNetworkProfileResponsePtrOutput { return v.DefaultNetworkProfile }).(LabPlanNetworkProfileResponsePtrOutput)
}

// Managed Identity Information
func (o LabPlanOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Base Url of the lms instance this lab plan can link lab rosters against.
func (o LabPlanOutput) LinkedLmsInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringPtrOutput { return v.LinkedLmsInstance }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LabPlanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LabPlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Current provisioning state of the lab plan.
func (o LabPlanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource ID of the Shared Image Gallery attached to this lab plan. When saving a lab template virtual machine image it will be persisted in this gallery. Shared images from the gallery can be made available to use when creating new labs.
func (o LabPlanOutput) SharedGalleryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringPtrOutput { return v.SharedGalleryId }).(pulumi.StringPtrOutput)
}

// Support contact information and instructions for users of the lab plan. This information is displayed to lab owners and virtual machine users for all labs in the lab plan.
func (o LabPlanOutput) SupportInfo() SupportInfoResponsePtrOutput {
	return o.ApplyT(func(v *LabPlan) SupportInfoResponsePtrOutput { return v.SupportInfo }).(SupportInfoResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the lab plan.
func (o LabPlanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabPlan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LabPlanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LabPlanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabPlan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabPlanOutput{})
}
