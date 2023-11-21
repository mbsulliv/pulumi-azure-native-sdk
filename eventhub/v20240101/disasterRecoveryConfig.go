// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in List or Get Alias(Disaster Recovery configuration) operation
type DisasterRecoveryConfig struct {
	pulumi.CustomResourceState

	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrOutput `pulumi:"alternateName"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace pulumi.StringPtrOutput `pulumi:"partnerNamespace"`
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount pulumi.Float64Output `pulumi:"pendingReplicationOperationsCount"`
	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role pulumi.StringOutput `pulumi:"role"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDisasterRecoveryConfig registers a new resource with the given unique name, arguments, and options.
func NewDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, args *DisasterRecoveryConfigArgs, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:DisasterRecoveryConfig"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:DisasterRecoveryConfig"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DisasterRecoveryConfig
	err := ctx.RegisterResource("azure-native:eventhub/v20240101:DisasterRecoveryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisasterRecoveryConfig gets an existing DisasterRecoveryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisasterRecoveryConfigState, opts ...pulumi.ResourceOption) (*DisasterRecoveryConfig, error) {
	var resource DisasterRecoveryConfig
	err := ctx.ReadResource("azure-native:eventhub/v20240101:DisasterRecoveryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DisasterRecoveryConfig resources.
type disasterRecoveryConfigState struct {
}

type DisasterRecoveryConfigState struct {
}

func (DisasterRecoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigState)(nil)).Elem()
}

type disasterRecoveryConfigArgs struct {
	// The Disaster Recovery configuration name
	Alias *string `pulumi:"alias"`
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `pulumi:"alternateName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string `pulumi:"partnerNamespace"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DisasterRecoveryConfig resource.
type DisasterRecoveryConfigArgs struct {
	// The Disaster Recovery configuration name
	Alias pulumi.StringPtrInput
	// Alternate name specified when alias and namespace names are same.
	AlternateName pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace pulumi.StringPtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (DisasterRecoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disasterRecoveryConfigArgs)(nil)).Elem()
}

type DisasterRecoveryConfigInput interface {
	pulumi.Input

	ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput
	ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput
}

func (*DisasterRecoveryConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DisasterRecoveryConfig)(nil)).Elem()
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return i.ToDisasterRecoveryConfigOutputWithContext(context.Background())
}

func (i *DisasterRecoveryConfig) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisasterRecoveryConfigOutput)
}

type DisasterRecoveryConfigOutput struct{ *pulumi.OutputState }

func (DisasterRecoveryConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisasterRecoveryConfig)(nil)).Elem()
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutput() DisasterRecoveryConfigOutput {
	return o
}

func (o DisasterRecoveryConfigOutput) ToDisasterRecoveryConfigOutputWithContext(ctx context.Context) DisasterRecoveryConfigOutput {
	return o
}

// Alternate name specified when alias and namespace names are same.
func (o DisasterRecoveryConfigOutput) AlternateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringPtrOutput { return v.AlternateName }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o DisasterRecoveryConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DisasterRecoveryConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
func (o DisasterRecoveryConfigOutput) PartnerNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringPtrOutput { return v.PartnerNamespace }).(pulumi.StringPtrOutput)
}

// Number of entities pending to be replicated.
func (o DisasterRecoveryConfigOutput) PendingReplicationOperationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.Float64Output { return v.PendingReplicationOperationsCount }).(pulumi.Float64Output)
}

// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
func (o DisasterRecoveryConfigOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
func (o DisasterRecoveryConfigOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o DisasterRecoveryConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o DisasterRecoveryConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DisasterRecoveryConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DisasterRecoveryConfigOutput{})
}
