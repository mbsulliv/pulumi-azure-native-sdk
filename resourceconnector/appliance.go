// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Appliances definition.
// Azure REST API version: 2022-10-27. Prior API version in Azure Native 1.x: 2021-10-31-preview
type Appliance struct {
	pulumi.CustomResourceState

	// Represents a supported Fabric/Infra. (AKSEdge etc...).
	Distro pulumi.StringPtrOutput `pulumi:"distro"`
	// Identity for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Contains infrastructure information about the Appliance
	InfrastructureConfig AppliancePropertiesResponseInfrastructureConfigPtrOutput `pulumi:"infrastructureConfig"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Certificates pair used to download MSI certificate from HIS. Can only be set once.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Appliance’s health and state of connection to on-prem
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the Appliance
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewAppliance registers a new resource with the given unique name, arguments, and options.
func NewAppliance(ctx *pulumi.Context,
	name string, args *ApplianceArgs, opts ...pulumi.ResourceOption) (*Appliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Distro == nil {
		args.Distro = pulumi.StringPtr("AKSEdge")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resourceconnector/v20211031preview:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:resourceconnector/v20220415preview:Appliance"),
		},
		{
			Type: pulumi.String("azure-native:resourceconnector/v20221027:Appliance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Appliance
	err := ctx.RegisterResource("azure-native:resourceconnector:Appliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppliance gets an existing Appliance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplianceState, opts ...pulumi.ResourceOption) (*Appliance, error) {
	var resource Appliance
	err := ctx.ReadResource("azure-native:resourceconnector:Appliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Appliance resources.
type applianceState struct {
}

type ApplianceState struct {
}

func (ApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceState)(nil)).Elem()
}

type applianceArgs struct {
	// Represents a supported Fabric/Infra. (AKSEdge etc...).
	Distro *string `pulumi:"distro"`
	// Identity for the resource.
	Identity *Identity `pulumi:"identity"`
	// Contains infrastructure information about the Appliance
	InfrastructureConfig *AppliancePropertiesInfrastructureConfig `pulumi:"infrastructureConfig"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Certificates pair used to download MSI certificate from HIS. Can only be set once.
	PublicKey *string `pulumi:"publicKey"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Appliances name.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Version of the Appliance
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Appliance resource.
type ApplianceArgs struct {
	// Represents a supported Fabric/Infra. (AKSEdge etc...).
	Distro pulumi.StringPtrInput
	// Identity for the resource.
	Identity IdentityPtrInput
	// Contains infrastructure information about the Appliance
	InfrastructureConfig AppliancePropertiesInfrastructureConfigPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Certificates pair used to download MSI certificate from HIS. Can only be set once.
	PublicKey pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Appliances name.
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Version of the Appliance
	Version pulumi.StringPtrInput
}

func (ApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applianceArgs)(nil)).Elem()
}

type ApplianceInput interface {
	pulumi.Input

	ToApplianceOutput() ApplianceOutput
	ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput
}

func (*Appliance) ElementType() reflect.Type {
	return reflect.TypeOf((**Appliance)(nil)).Elem()
}

func (i *Appliance) ToApplianceOutput() ApplianceOutput {
	return i.ToApplianceOutputWithContext(context.Background())
}

func (i *Appliance) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplianceOutput)
}

func (i *Appliance) ToOutput(ctx context.Context) pulumix.Output[*Appliance] {
	return pulumix.Output[*Appliance]{
		OutputState: i.ToApplianceOutputWithContext(ctx).OutputState,
	}
}

type ApplianceOutput struct{ *pulumi.OutputState }

func (ApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Appliance)(nil)).Elem()
}

func (o ApplianceOutput) ToApplianceOutput() ApplianceOutput {
	return o
}

func (o ApplianceOutput) ToApplianceOutputWithContext(ctx context.Context) ApplianceOutput {
	return o
}

func (o ApplianceOutput) ToOutput(ctx context.Context) pulumix.Output[*Appliance] {
	return pulumix.Output[*Appliance]{
		OutputState: o.OutputState,
	}
}

// Represents a supported Fabric/Infra. (AKSEdge etc...).
func (o ApplianceOutput) Distro() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.Distro }).(pulumi.StringPtrOutput)
}

// Identity for the resource.
func (o ApplianceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Appliance) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// Contains infrastructure information about the Appliance
func (o ApplianceOutput) InfrastructureConfig() AppliancePropertiesResponseInfrastructureConfigPtrOutput {
	return o.ApplyT(func(v *Appliance) AppliancePropertiesResponseInfrastructureConfigPtrOutput {
		return v.InfrastructureConfig
	}).(AppliancePropertiesResponseInfrastructureConfigPtrOutput)
}

// The geo-location where the resource lives
func (o ApplianceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApplianceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o ApplianceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Certificates pair used to download MSI certificate from HIS. Can only be set once.
func (o ApplianceOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Appliance’s health and state of connection to on-prem
func (o ApplianceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApplianceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Appliance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ApplianceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApplianceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the Appliance
func (o ApplianceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Appliance) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplianceOutput{})
}
