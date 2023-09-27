// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A class represent a replica resource.
// Azure REST API version: 2023-03-01-preview.
type SignalRReplica struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The billing information of the resource.
	Sku ResourceSkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSignalRReplica registers a new resource with the given unique name, arguments, and options.
func NewSignalRReplica(ctx *pulumi.Context,
	name string, args *SignalRReplicaArgs, opts ...pulumi.ResourceOption) (*SignalRReplica, error) {
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
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalRReplica"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalRReplica"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230801preview:SignalRReplica"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalRReplica
	err := ctx.RegisterResource("azure-native:signalrservice:SignalRReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRReplica gets an existing SignalRReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRReplicaState, opts ...pulumi.ResourceOption) (*SignalRReplica, error) {
	var resource SignalRReplica
	err := ctx.ReadResource("azure-native:signalrservice:SignalRReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRReplica resources.
type signalRReplicaState struct {
}

type SignalRReplicaState struct {
}

func (SignalRReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRReplicaState)(nil)).Elem()
}

type signalRReplicaArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the replica.
	ReplicaName *string `pulumi:"replicaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The billing information of the resource.
	Sku *ResourceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SignalRReplica resource.
type SignalRReplicaArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the replica.
	ReplicaName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
	// The billing information of the resource.
	Sku ResourceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SignalRReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRReplicaArgs)(nil)).Elem()
}

type SignalRReplicaInput interface {
	pulumi.Input

	ToSignalRReplicaOutput() SignalRReplicaOutput
	ToSignalRReplicaOutputWithContext(ctx context.Context) SignalRReplicaOutput
}

func (*SignalRReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRReplica)(nil)).Elem()
}

func (i *SignalRReplica) ToSignalRReplicaOutput() SignalRReplicaOutput {
	return i.ToSignalRReplicaOutputWithContext(context.Background())
}

func (i *SignalRReplica) ToSignalRReplicaOutputWithContext(ctx context.Context) SignalRReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRReplicaOutput)
}

func (i *SignalRReplica) ToOutput(ctx context.Context) pulumix.Output[*SignalRReplica] {
	return pulumix.Output[*SignalRReplica]{
		OutputState: i.ToSignalRReplicaOutputWithContext(ctx).OutputState,
	}
}

type SignalRReplicaOutput struct{ *pulumi.OutputState }

func (SignalRReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRReplica)(nil)).Elem()
}

func (o SignalRReplicaOutput) ToSignalRReplicaOutput() SignalRReplicaOutput {
	return o
}

func (o SignalRReplicaOutput) ToSignalRReplicaOutputWithContext(ctx context.Context) SignalRReplicaOutput {
	return o
}

func (o SignalRReplicaOutput) ToOutput(ctx context.Context) pulumix.Output[*SignalRReplica] {
	return pulumix.Output[*SignalRReplica]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o SignalRReplicaOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRReplica) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SignalRReplicaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRReplica) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o SignalRReplicaOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRReplica) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The billing information of the resource.
func (o SignalRReplicaOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v *SignalRReplica) ResourceSkuResponsePtrOutput { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SignalRReplicaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRReplica) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SignalRReplicaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SignalRReplica) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SignalRReplicaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRReplica) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRReplicaOutput{})
}
