// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the ExternalNetwork item.
type ExternalNetwork struct {
	pulumi.CustomResourceState

	// AdministrativeState of the externalNetwork. Example: Enabled | Disabled.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// List of resources the externalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
	DisabledOnResources pulumi.StringArrayOutput `pulumi:"disabledOnResources"`
	// ARM resource ID of exportRoutePolicy.
	ExportRoutePolicyId pulumi.StringPtrOutput `pulumi:"exportRoutePolicyId"`
	// ARM resource ID of importRoutePolicy.
	ImportRoutePolicyId pulumi.StringPtrOutput `pulumi:"importRoutePolicyId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the networkToNetworkInterconnectId of the resource.
	NetworkToNetworkInterconnectId pulumi.StringOutput `pulumi:"networkToNetworkInterconnectId"`
	// option A properties object
	OptionAProperties ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput `pulumi:"optionAProperties"`
	// option B properties object
	OptionBProperties OptionBPropertiesResponsePtrOutput `pulumi:"optionBProperties"`
	// Peering option list.
	PeeringOption pulumi.StringOutput `pulumi:"peeringOption"`
	// Gets the provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExternalNetwork registers a new resource with the given unique name, arguments, and options.
func NewExternalNetwork(ctx *pulumi.Context,
	name string, args *ExternalNetworkArgs, opts ...pulumi.ResourceOption) (*ExternalNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.L3IsolationDomainName == nil {
		return nil, errors.New("invalid value for required argument 'L3IsolationDomainName'")
	}
	if args.PeeringOption == nil {
		return nil, errors.New("invalid value for required argument 'PeeringOption'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.OptionAProperties != nil {
		args.OptionAProperties = args.OptionAProperties.ToExternalNetworkPropertiesOptionAPropertiesPtrOutput().ApplyT(func(v *ExternalNetworkPropertiesOptionAProperties) *ExternalNetworkPropertiesOptionAProperties {
			return v.Defaults()
		}).(ExternalNetworkPropertiesOptionAPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:ExternalNetwork"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:ExternalNetwork"),
		},
	})
	opts = append(opts, aliases)
	var resource ExternalNetwork
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230201preview:ExternalNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalNetwork gets an existing ExternalNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalNetworkState, opts ...pulumi.ResourceOption) (*ExternalNetwork, error) {
	var resource ExternalNetwork
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230201preview:ExternalNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalNetwork resources.
type externalNetworkState struct {
}

type ExternalNetworkState struct {
}

func (ExternalNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNetworkState)(nil)).Elem()
}

type externalNetworkArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// ARM resource ID of exportRoutePolicy.
	ExportRoutePolicyId *string `pulumi:"exportRoutePolicyId"`
	// Name of the ExternalNetwork
	ExternalNetworkName *string `pulumi:"externalNetworkName"`
	// ARM resource ID of importRoutePolicy.
	ImportRoutePolicyId *string `pulumi:"importRoutePolicyId"`
	// Name of the L3IsolationDomain
	L3IsolationDomainName string `pulumi:"l3IsolationDomainName"`
	// option A properties object
	OptionAProperties *ExternalNetworkPropertiesOptionAProperties `pulumi:"optionAProperties"`
	// option B properties object
	OptionBProperties *OptionBProperties `pulumi:"optionBProperties"`
	// Peering option list.
	PeeringOption string `pulumi:"peeringOption"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExternalNetwork resource.
type ExternalNetworkArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// ARM resource ID of exportRoutePolicy.
	ExportRoutePolicyId pulumi.StringPtrInput
	// Name of the ExternalNetwork
	ExternalNetworkName pulumi.StringPtrInput
	// ARM resource ID of importRoutePolicy.
	ImportRoutePolicyId pulumi.StringPtrInput
	// Name of the L3IsolationDomain
	L3IsolationDomainName pulumi.StringInput
	// option A properties object
	OptionAProperties ExternalNetworkPropertiesOptionAPropertiesPtrInput
	// option B properties object
	OptionBProperties OptionBPropertiesPtrInput
	// Peering option list.
	PeeringOption pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ExternalNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalNetworkArgs)(nil)).Elem()
}

type ExternalNetworkInput interface {
	pulumi.Input

	ToExternalNetworkOutput() ExternalNetworkOutput
	ToExternalNetworkOutputWithContext(ctx context.Context) ExternalNetworkOutput
}

func (*ExternalNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNetwork)(nil)).Elem()
}

func (i *ExternalNetwork) ToExternalNetworkOutput() ExternalNetworkOutput {
	return i.ToExternalNetworkOutputWithContext(context.Background())
}

func (i *ExternalNetwork) ToExternalNetworkOutputWithContext(ctx context.Context) ExternalNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalNetworkOutput)
}

type ExternalNetworkOutput struct{ *pulumi.OutputState }

func (ExternalNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalNetwork)(nil)).Elem()
}

func (o ExternalNetworkOutput) ToExternalNetworkOutput() ExternalNetworkOutput {
	return o
}

func (o ExternalNetworkOutput) ToExternalNetworkOutputWithContext(ctx context.Context) ExternalNetworkOutput {
	return o
}

// AdministrativeState of the externalNetwork. Example: Enabled | Disabled.
func (o ExternalNetworkOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o ExternalNetworkOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// List of resources the externalNetwork is disabled on. Can be either entire NetworkFabric or NetworkRack.
func (o ExternalNetworkOutput) DisabledOnResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringArrayOutput { return v.DisabledOnResources }).(pulumi.StringArrayOutput)
}

// ARM resource ID of exportRoutePolicy.
func (o ExternalNetworkOutput) ExportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringPtrOutput { return v.ExportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// ARM resource ID of importRoutePolicy.
func (o ExternalNetworkOutput) ImportRoutePolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringPtrOutput { return v.ImportRoutePolicyId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ExternalNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the networkToNetworkInterconnectId of the resource.
func (o ExternalNetworkOutput) NetworkToNetworkInterconnectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.NetworkToNetworkInterconnectId }).(pulumi.StringOutput)
}

// option A properties object
func (o ExternalNetworkOutput) OptionAProperties() ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput {
	return o.ApplyT(func(v *ExternalNetwork) ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput {
		return v.OptionAProperties
	}).(ExternalNetworkPropertiesResponseOptionAPropertiesPtrOutput)
}

// option B properties object
func (o ExternalNetworkOutput) OptionBProperties() OptionBPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ExternalNetwork) OptionBPropertiesResponsePtrOutput { return v.OptionBProperties }).(OptionBPropertiesResponsePtrOutput)
}

// Peering option list.
func (o ExternalNetworkOutput) PeeringOption() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.PeeringOption }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o ExternalNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ExternalNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ExternalNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ExternalNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalNetworkOutput{})
}
