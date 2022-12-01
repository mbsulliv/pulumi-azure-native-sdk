// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fabric definition.
type ReplicationFabric struct {
	pulumi.CustomResourceState

	// Resource Location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Fabric related data.
	Properties FabricPropertiesResponseOutput `pulumi:"properties"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationFabric registers a new resource with the given unique name, arguments, and options.
func NewReplicationFabric(ctx *pulumi.Context,
	name string, args *ReplicationFabricArgs, opts ...pulumi.ResourceOption) (*ReplicationFabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20160810:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180110:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20180710:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210210:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211101:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220501:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220801:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220910:ReplicationFabric"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ReplicationFabric"),
		},
	})
	opts = append(opts, aliases)
	var resource ReplicationFabric
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210601:ReplicationFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationFabric gets an existing ReplicationFabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationFabricState, opts ...pulumi.ResourceOption) (*ReplicationFabric, error) {
	var resource ReplicationFabric
	err := ctx.ReadResource("azure-native:recoveryservices/v20210601:ReplicationFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationFabric resources.
type replicationFabricState struct {
}

type ReplicationFabricState struct {
}

func (ReplicationFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationFabricState)(nil)).Elem()
}

type replicationFabricArgs struct {
	// Name of the ASR fabric.
	FabricName *string `pulumi:"fabricName"`
	// Fabric creation input.
	Properties *FabricCreationInputProperties `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a ReplicationFabric resource.
type ReplicationFabricArgs struct {
	// Name of the ASR fabric.
	FabricName pulumi.StringPtrInput
	// Fabric creation input.
	Properties FabricCreationInputPropertiesPtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput
}

func (ReplicationFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationFabricArgs)(nil)).Elem()
}

type ReplicationFabricInput interface {
	pulumi.Input

	ToReplicationFabricOutput() ReplicationFabricOutput
	ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput
}

func (*ReplicationFabric) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationFabric)(nil)).Elem()
}

func (i *ReplicationFabric) ToReplicationFabricOutput() ReplicationFabricOutput {
	return i.ToReplicationFabricOutputWithContext(context.Background())
}

func (i *ReplicationFabric) ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationFabricOutput)
}

type ReplicationFabricOutput struct{ *pulumi.OutputState }

func (ReplicationFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationFabric)(nil)).Elem()
}

func (o ReplicationFabricOutput) ToReplicationFabricOutput() ReplicationFabricOutput {
	return o
}

func (o ReplicationFabricOutput) ToReplicationFabricOutputWithContext(ctx context.Context) ReplicationFabricOutput {
	return o
}

// Resource Location
func (o ReplicationFabricOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o ReplicationFabricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Fabric related data.
func (o ReplicationFabricOutput) Properties() FabricPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationFabric) FabricPropertiesResponseOutput { return v.Properties }).(FabricPropertiesResponseOutput)
}

// Resource Type
func (o ReplicationFabricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationFabric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationFabricOutput{})
}
