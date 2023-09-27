// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbidedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents an instance of a Dedicated Capacity resource.
// Azure REST API version: 2021-01-01. Prior API version in Azure Native 1.x: 2021-01-01
type CapacityDetails struct {
	pulumi.CustomResourceState

	// A collection of Dedicated capacity administrators
	Administration DedicatedCapacityAdministratorsResponsePtrOutput `pulumi:"administration"`
	// Capacity name
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the PowerBI Dedicated resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the PowerBI Dedicated capacity resource.
	Sku CapacitySkuResponseOutput `pulumi:"sku"`
	// The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
	State pulumi.StringOutput `pulumi:"state"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponsePtrOutput `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Tenant ID for the capacity. Used for creating Pro Plus capacity.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the PowerBI Dedicated resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCapacityDetails registers a new resource with the given unique name, arguments, and options.
func NewCapacityDetails(ctx *pulumi.Context,
	name string, args *CapacityDetailsArgs, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
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
			Type: pulumi.String("azure-native:powerbidedicated/v20171001:CapacityDetails"),
		},
		{
			Type: pulumi.String("azure-native:powerbidedicated/v20210101:CapacityDetails"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CapacityDetails
	err := ctx.RegisterResource("azure-native:powerbidedicated:CapacityDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityDetails gets an existing CapacityDetails resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityDetailsState, opts ...pulumi.ResourceOption) (*CapacityDetails, error) {
	var resource CapacityDetails
	err := ctx.ReadResource("azure-native:powerbidedicated:CapacityDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityDetails resources.
type capacityDetailsState struct {
}

type CapacityDetailsState struct {
}

func (CapacityDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsState)(nil)).Elem()
}

type capacityDetailsArgs struct {
	// A collection of Dedicated capacity administrators
	Administration *DedicatedCapacityAdministrators `pulumi:"administration"`
	// The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	DedicatedCapacityName *string `pulumi:"dedicatedCapacityName"`
	// Location of the PowerBI Dedicated resource.
	Location *string `pulumi:"location"`
	// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
	Mode *string `pulumi:"mode"`
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the PowerBI Dedicated capacity resource.
	Sku CapacitySku `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CapacityDetails resource.
type CapacityDetailsArgs struct {
	// A collection of Dedicated capacity administrators
	Administration DedicatedCapacityAdministratorsPtrInput
	// The name of the Dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	DedicatedCapacityName pulumi.StringPtrInput
	// Location of the PowerBI Dedicated resource.
	Location pulumi.StringPtrInput
	// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
	Mode pulumi.StringPtrInput
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput
	// The SKU of the PowerBI Dedicated capacity resource.
	Sku CapacitySkuInput
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataPtrInput
	// Key-value pairs of additional resource provisioning properties.
	Tags pulumi.StringMapInput
}

func (CapacityDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityDetailsArgs)(nil)).Elem()
}

type CapacityDetailsInput interface {
	pulumi.Input

	ToCapacityDetailsOutput() CapacityDetailsOutput
	ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput
}

func (*CapacityDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityDetails)(nil)).Elem()
}

func (i *CapacityDetails) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return i.ToCapacityDetailsOutputWithContext(context.Background())
}

func (i *CapacityDetails) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityDetailsOutput)
}

func (i *CapacityDetails) ToOutput(ctx context.Context) pulumix.Output[*CapacityDetails] {
	return pulumix.Output[*CapacityDetails]{
		OutputState: i.ToCapacityDetailsOutputWithContext(ctx).OutputState,
	}
}

type CapacityDetailsOutput struct{ *pulumi.OutputState }

func (CapacityDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityDetails)(nil)).Elem()
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutput() CapacityDetailsOutput {
	return o
}

func (o CapacityDetailsOutput) ToCapacityDetailsOutputWithContext(ctx context.Context) CapacityDetailsOutput {
	return o
}

func (o CapacityDetailsOutput) ToOutput(ctx context.Context) pulumix.Output[*CapacityDetails] {
	return pulumix.Output[*CapacityDetails]{
		OutputState: o.OutputState,
	}
}

// A collection of Dedicated capacity administrators
func (o CapacityDetailsOutput) Administration() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v *CapacityDetails) DedicatedCapacityAdministratorsResponsePtrOutput { return v.Administration }).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

// Capacity name
func (o CapacityDetailsOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// Location of the PowerBI Dedicated resource.
func (o CapacityDetailsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
func (o CapacityDetailsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the PowerBI Dedicated resource.
func (o CapacityDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
func (o CapacityDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the PowerBI Dedicated capacity resource.
func (o CapacityDetailsOutput) Sku() CapacitySkuResponseOutput {
	return o.ApplyT(func(v *CapacityDetails) CapacitySkuResponseOutput { return v.Sku }).(CapacitySkuResponseOutput)
}

// The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
func (o CapacityDetailsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CapacityDetailsOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v *CapacityDetails) SystemDataResponsePtrOutput { return v.SystemData }).(SystemDataResponsePtrOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o CapacityDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Tenant ID for the capacity. Used for creating Pro Plus capacity.
func (o CapacityDetailsOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the PowerBI Dedicated resource.
func (o CapacityDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityDetailsOutput{})
}
