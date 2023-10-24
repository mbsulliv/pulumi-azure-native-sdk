// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2020-01-01. Prior API version in Azure Native 1.x: 2020-01-01.
type JitNetworkAccessPolicy struct {
	pulumi.CustomResourceState

	// Kind of the resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the Just-in-Time policy.
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	Requests          JitNetworkAccessRequestResponseArrayOutput `pulumi:"requests"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewJitNetworkAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, args *JitNetworkAccessPolicyArgs, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscLocation == nil {
		return nil, errors.New("invalid value for required argument 'AscLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachines == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachines'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20150601preview:JitNetworkAccessPolicy"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:JitNetworkAccessPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JitNetworkAccessPolicy
	err := ctx.RegisterResource("azure-native:security:JitNetworkAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJitNetworkAccessPolicy gets an existing JitNetworkAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJitNetworkAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JitNetworkAccessPolicyState, opts ...pulumi.ResourceOption) (*JitNetworkAccessPolicy, error) {
	var resource JitNetworkAccessPolicy
	err := ctx.ReadResource("azure-native:security:JitNetworkAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JitNetworkAccessPolicy resources.
type jitNetworkAccessPolicyState struct {
}

type JitNetworkAccessPolicyState struct {
}

func (JitNetworkAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyState)(nil)).Elem()
}

type jitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName *string `pulumi:"jitNetworkAccessPolicyName"`
	// Kind of the resource
	Kind     *string                   `pulumi:"kind"`
	Requests []JitNetworkAccessRequest `pulumi:"requests"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines []JitNetworkAccessPolicyVirtualMachine `pulumi:"virtualMachines"`
}

// The set of arguments for constructing a JitNetworkAccessPolicy resource.
type JitNetworkAccessPolicyArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// Name of a Just-in-Time access configuration policy.
	JitNetworkAccessPolicyName pulumi.StringPtrInput
	// Kind of the resource
	Kind     pulumi.StringPtrInput
	Requests JitNetworkAccessRequestArrayInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Configurations for Microsoft.Compute/virtualMachines resource type.
	VirtualMachines JitNetworkAccessPolicyVirtualMachineArrayInput
}

func (JitNetworkAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jitNetworkAccessPolicyArgs)(nil)).Elem()
}

type JitNetworkAccessPolicyInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput
	ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput
}

func (*JitNetworkAccessPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**JitNetworkAccessPolicy)(nil)).Elem()
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return i.ToJitNetworkAccessPolicyOutputWithContext(context.Background())
}

func (i *JitNetworkAccessPolicy) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyOutput)
}

func (i *JitNetworkAccessPolicy) ToOutput(ctx context.Context) pulumix.Output[*JitNetworkAccessPolicy] {
	return pulumix.Output[*JitNetworkAccessPolicy]{
		OutputState: i.ToJitNetworkAccessPolicyOutputWithContext(ctx).OutputState,
	}
}

type JitNetworkAccessPolicyOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JitNetworkAccessPolicy)(nil)).Elem()
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutput() JitNetworkAccessPolicyOutput {
	return o
}

func (o JitNetworkAccessPolicyOutput) ToJitNetworkAccessPolicyOutputWithContext(ctx context.Context) JitNetworkAccessPolicyOutput {
	return o
}

func (o JitNetworkAccessPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*JitNetworkAccessPolicy] {
	return pulumix.Output[*JitNetworkAccessPolicy]{
		OutputState: o.OutputState,
	}
}

// Kind of the resource
func (o JitNetworkAccessPolicyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o JitNetworkAccessPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o JitNetworkAccessPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the Just-in-Time policy.
func (o JitNetworkAccessPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyOutput) Requests() JitNetworkAccessRequestResponseArrayOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) JitNetworkAccessRequestResponseArrayOutput { return v.Requests }).(JitNetworkAccessRequestResponseArrayOutput)
}

// Resource type
func (o JitNetworkAccessPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Configurations for Microsoft.Compute/virtualMachines resource type.
func (o JitNetworkAccessPolicyOutput) VirtualMachines() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v *JitNetworkAccessPolicy) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
		return v.VirtualMachines
	}).(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(JitNetworkAccessPolicyOutput{})
}
