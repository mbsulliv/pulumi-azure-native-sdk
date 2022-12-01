// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provider details.
type ReplicationRecoveryServicesProvider struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Provider properties.
	Properties RecoveryServicesProviderPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationRecoveryServicesProvider registers a new resource with the given unique name, arguments, and options.
func NewReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, args *ReplicationRecoveryServicesProviderArgs, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FabricName == nil {
		return nil, errors.New("invalid value for required argument 'FabricName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationRecoveryServicesProvider"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationRecoveryServicesProvider"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationRecoveryServicesProvider
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211101:ReplicationRecoveryServicesProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationRecoveryServicesProvider gets an existing ReplicationRecoveryServicesProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationRecoveryServicesProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationRecoveryServicesProviderState, opts ...pulumi.ResourceOption) (*ReplicationRecoveryServicesProvider, error) {
	var resource ReplicationRecoveryServicesProvider
	err := ctx.ReadResource("azure-native:recoveryservices/v20211101:ReplicationRecoveryServicesProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationRecoveryServicesProvider resources.
type replicationRecoveryServicesProviderState struct {
}

type ReplicationRecoveryServicesProviderState struct {
}

func (ReplicationRecoveryServicesProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderState)(nil)).Elem()
}

type replicationRecoveryServicesProviderArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// The properties of an add provider request.
	Properties AddRecoveryServicesProviderInputProperties `pulumi:"properties"`
	// Recovery services provider name.
	ProviderName *string `pulumi:"providerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationRecoveryServicesProvider resource.
type ReplicationRecoveryServicesProviderArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput
	// The properties of an add provider request.
	Properties AddRecoveryServicesProviderInputPropertiesInput
	// Recovery services provider name.
	ProviderName pulumi.StringPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationRecoveryServicesProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationRecoveryServicesProviderArgs)(nil)).Elem()
}

type ReplicationRecoveryServicesProviderInput interface {
	pulumi.Input

	ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput
	ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput
}

func (*ReplicationRecoveryServicesProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationRecoveryServicesProvider)(nil)).Elem()
}

func (i *ReplicationRecoveryServicesProvider) ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput {
	return i.ToReplicationRecoveryServicesProviderOutputWithContext(context.Background())
}

func (i *ReplicationRecoveryServicesProvider) ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationRecoveryServicesProviderOutput)
}

type ReplicationRecoveryServicesProviderOutput struct{ *pulumi.OutputState }

func (ReplicationRecoveryServicesProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationRecoveryServicesProvider)(nil)).Elem()
}

func (o ReplicationRecoveryServicesProviderOutput) ToReplicationRecoveryServicesProviderOutput() ReplicationRecoveryServicesProviderOutput {
	return o
}

func (o ReplicationRecoveryServicesProviderOutput) ToReplicationRecoveryServicesProviderOutputWithContext(ctx context.Context) ReplicationRecoveryServicesProviderOutput {
	return o
}

// Resource Location
func (o ReplicationRecoveryServicesProviderOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationRecoveryServicesProvider) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationRecoveryServicesProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationRecoveryServicesProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provider properties.
func (o ReplicationRecoveryServicesProviderOutput) Properties() RecoveryServicesProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationRecoveryServicesProvider) RecoveryServicesProviderPropertiesResponseOutput {
		return v.Properties
	}).(RecoveryServicesProviderPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationRecoveryServicesProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationRecoveryServicesProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationRecoveryServicesProviderOutput{})
}
