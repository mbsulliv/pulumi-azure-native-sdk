// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsorchestrator

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Instance resource belonging to an Instance resource.
// Azure REST API version: 2023-10-04-preview.
type Instance struct {
	pulumi.CustomResourceState

	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Reconciliation Policy.
	ReconciliationPolicy ReconciliationPolicyResponsePtrOutput `pulumi:"reconciliationPolicy"`
	// Deployment scope (such as Kubernetes namespace).
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Name of the solution.
	Solution pulumi.StringPtrOutput `pulumi:"solution"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Defines the Target the Instance will deploy to.
	Target TargetSelectorPropertiesResponsePtrOutput `pulumi:"target"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the particular resource.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperationsorchestrator/v20231004preview:Instance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("azure-native:iotoperationsorchestrator:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("azure-native:iotoperationsorchestrator:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Edge location of the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of Instance.
	Name *string `pulumi:"name"`
	// Reconciliation Policy.
	ReconciliationPolicy *ReconciliationPolicy `pulumi:"reconciliationPolicy"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Deployment scope (such as Kubernetes namespace).
	Scope *string `pulumi:"scope"`
	// Name of the solution.
	Solution *string `pulumi:"solution"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Defines the Target the Instance will deploy to.
	Target *TargetSelectorProperties `pulumi:"target"`
	// Version of the particular resource.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of Instance.
	Name pulumi.StringPtrInput
	// Reconciliation Policy.
	ReconciliationPolicy ReconciliationPolicyPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Deployment scope (such as Kubernetes namespace).
	Scope pulumi.StringPtrInput
	// Name of the solution.
	Solution pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Defines the Target the Instance will deploy to.
	Target TargetSelectorPropertiesPtrInput
	// Version of the particular resource.
	Version pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Edge location of the resource.
func (o InstanceOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Instance) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o InstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Reconciliation Policy.
func (o InstanceOutput) ReconciliationPolicy() ReconciliationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Instance) ReconciliationPolicyResponsePtrOutput { return v.ReconciliationPolicy }).(ReconciliationPolicyResponsePtrOutput)
}

// Deployment scope (such as Kubernetes namespace).
func (o InstanceOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Name of the solution.
func (o InstanceOutput) Solution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Solution }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Instance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Defines the Target the Instance will deploy to.
func (o InstanceOutput) Target() TargetSelectorPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Instance) TargetSelectorPropertiesResponsePtrOutput { return v.Target }).(TargetSelectorPropertiesResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the particular resource.
func (o InstanceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
