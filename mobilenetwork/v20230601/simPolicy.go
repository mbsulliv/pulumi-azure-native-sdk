// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SIM policy resource.
type SimPolicy struct {
	pulumi.CustomResourceState

	// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
	DefaultSlice SliceResourceIdResponseOutput `pulumi:"defaultSlice"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the SIM policy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
	RegistrationTimer pulumi.IntPtrOutput `pulumi:"registrationTimer"`
	// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
	RfspIndex pulumi.IntPtrOutput `pulumi:"rfspIndex"`
	// A dictionary of sites to the provisioning state of this SIM policy on that site.
	SiteProvisioningState pulumi.StringMapOutput `pulumi:"siteProvisioningState"`
	// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
	SliceConfigurations SliceConfigurationResponseArrayOutput `pulumi:"sliceConfigurations"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
	UeAmbr AmbrResponseOutput `pulumi:"ueAmbr"`
}

// NewSimPolicy registers a new resource with the given unique name, arguments, and options.
func NewSimPolicy(ctx *pulumi.Context,
	name string, args *SimPolicyArgs, opts ...pulumi.ResourceOption) (*SimPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultSlice == nil {
		return nil, errors.New("invalid value for required argument 'DefaultSlice'")
	}
	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SliceConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'SliceConfigurations'")
	}
	if args.UeAmbr == nil {
		return nil, errors.New("invalid value for required argument 'UeAmbr'")
	}
	if args.RegistrationTimer == nil {
		args.RegistrationTimer = pulumi.IntPtr(3240)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:SimPolicy"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220301preview:SimPolicy"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:SimPolicy"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:SimPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SimPolicy
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20230601:SimPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSimPolicy gets an existing SimPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSimPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SimPolicyState, opts ...pulumi.ResourceOption) (*SimPolicy, error) {
	var resource SimPolicy
	err := ctx.ReadResource("azure-native:mobilenetwork/v20230601:SimPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SimPolicy resources.
type simPolicyState struct {
}

type SimPolicyState struct {
}

func (SimPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*simPolicyState)(nil)).Elem()
}

type simPolicyArgs struct {
	// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
	DefaultSlice SliceResourceId `pulumi:"defaultSlice"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
	RegistrationTimer *int `pulumi:"registrationTimer"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
	RfspIndex *int `pulumi:"rfspIndex"`
	// The name of the SIM policy.
	SimPolicyName *string `pulumi:"simPolicyName"`
	// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
	SliceConfigurations []SliceConfiguration `pulumi:"sliceConfigurations"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
	UeAmbr Ambr `pulumi:"ueAmbr"`
}

// The set of arguments for constructing a SimPolicy resource.
type SimPolicyArgs struct {
	// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
	DefaultSlice SliceResourceIdInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput
	// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
	RegistrationTimer pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
	RfspIndex pulumi.IntPtrInput
	// The name of the SIM policy.
	SimPolicyName pulumi.StringPtrInput
	// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
	SliceConfigurations SliceConfigurationArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
	UeAmbr AmbrInput
}

func (SimPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*simPolicyArgs)(nil)).Elem()
}

type SimPolicyInput interface {
	pulumi.Input

	ToSimPolicyOutput() SimPolicyOutput
	ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput
}

func (*SimPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicy)(nil)).Elem()
}

func (i *SimPolicy) ToSimPolicyOutput() SimPolicyOutput {
	return i.ToSimPolicyOutputWithContext(context.Background())
}

func (i *SimPolicy) ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimPolicyOutput)
}

type SimPolicyOutput struct{ *pulumi.OutputState }

func (SimPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SimPolicy)(nil)).Elem()
}

func (o SimPolicyOutput) ToSimPolicyOutput() SimPolicyOutput {
	return o
}

func (o SimPolicyOutput) ToSimPolicyOutputWithContext(ctx context.Context) SimPolicyOutput {
	return o
}

// The default slice to use if the UE does not explicitly specify it. This slice must exist in the `sliceConfigurations` map. The slice must be in the same location as the SIM policy.
func (o SimPolicyOutput) DefaultSlice() SliceResourceIdResponseOutput {
	return o.ApplyT(func(v *SimPolicy) SliceResourceIdResponseOutput { return v.DefaultSlice }).(SliceResourceIdResponseOutput)
}

// The geo-location where the resource lives
func (o SimPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SimPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the SIM policy resource.
func (o SimPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// UE periodic registration update timer (5G) or UE periodic tracking area update timer (4G), in seconds.
func (o SimPolicyOutput) RegistrationTimer() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.IntPtrOutput { return v.RegistrationTimer }).(pulumi.IntPtrOutput)
}

// RAT/Frequency Selection Priority Index, defined in 3GPP TS 36.413. This is an optional setting and by default is unspecified.
func (o SimPolicyOutput) RfspIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.IntPtrOutput { return v.RfspIndex }).(pulumi.IntPtrOutput)
}

// A dictionary of sites to the provisioning state of this SIM policy on that site.
func (o SimPolicyOutput) SiteProvisioningState() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringMapOutput { return v.SiteProvisioningState }).(pulumi.StringMapOutput)
}

// The allowed slices and the settings to use for them. The list must not contain duplicate items and must contain at least one item.
func (o SimPolicyOutput) SliceConfigurations() SliceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SimPolicy) SliceConfigurationResponseArrayOutput { return v.SliceConfigurations }).(SliceConfigurationResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SimPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SimPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SimPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SimPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SimPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Aggregate maximum bit rate across all non-GBR QoS flows of all PDU sessions of a given UE. See 3GPP TS23.501 section 5.7.2.6 for a full description of the UE-AMBR.
func (o SimPolicyOutput) UeAmbr() AmbrResponseOutput {
	return o.ApplyT(func(v *SimPolicy) AmbrResponseOutput { return v.UeAmbr }).(AmbrResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SimPolicyOutput{})
}