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

// The replica resource.
type Replica struct {
	pulumi.CustomResourceState

	// The URI of the replica where the replica API will be available.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The location of the replica.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the replica.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the replica.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource system metadata.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplica registers a new resource with the given unique name, arguments, and options.
func NewReplica(ctx *pulumi.Context,
	name string, args *ReplicaArgs, opts ...pulumi.ResourceOption) (*Replica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigStoreName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigStoreName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appconfiguration:Replica"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220301preview:Replica"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20230301:Replica"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20230901preview:Replica"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Replica
	err := ctx.RegisterResource("azure-native:appconfiguration/v20230801preview:Replica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplica gets an existing Replica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicaState, opts ...pulumi.ResourceOption) (*Replica, error) {
	var resource Replica
	err := ctx.ReadResource("azure-native:appconfiguration/v20230801preview:Replica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Replica resources.
type replicaState struct {
}

type ReplicaState struct {
}

func (ReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaState)(nil)).Elem()
}

type replicaArgs struct {
	// The name of the configuration store.
	ConfigStoreName string `pulumi:"configStoreName"`
	// The location of the replica.
	Location *string `pulumi:"location"`
	// The name of the replica.
	ReplicaName *string `pulumi:"replicaName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Replica resource.
type ReplicaArgs struct {
	// The name of the configuration store.
	ConfigStoreName pulumi.StringInput
	// The location of the replica.
	Location pulumi.StringPtrInput
	// The name of the replica.
	ReplicaName pulumi.StringPtrInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
}

func (ReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicaArgs)(nil)).Elem()
}

type ReplicaInput interface {
	pulumi.Input

	ToReplicaOutput() ReplicaOutput
	ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput
}

func (*Replica) ElementType() reflect.Type {
	return reflect.TypeOf((**Replica)(nil)).Elem()
}

func (i *Replica) ToReplicaOutput() ReplicaOutput {
	return i.ToReplicaOutputWithContext(context.Background())
}

func (i *Replica) ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicaOutput)
}

type ReplicaOutput struct{ *pulumi.OutputState }

func (ReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Replica)(nil)).Elem()
}

func (o ReplicaOutput) ToReplicaOutput() ReplicaOutput {
	return o
}

func (o ReplicaOutput) ToReplicaOutputWithContext(ctx context.Context) ReplicaOutput {
	return o
}

// The URI of the replica where the replica API will be available.
func (o ReplicaOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The location of the replica.
func (o ReplicaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the replica.
func (o ReplicaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the replica.
func (o ReplicaOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o ReplicaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Replica) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ReplicaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Replica) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicaOutput{})
}
