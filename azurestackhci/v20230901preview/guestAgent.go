// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the GuestAgent.
type GuestAgent struct {
	pulumi.CustomResourceState

	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialResponsePtrOutput `pulumi:"credentials"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrOutput `pulumi:"provisioningAction"`
	// The provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The guest agent status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGuestAgent registers a new resource with the given unique name, arguments, and options.
func NewGuestAgent(ctx *pulumi.Context,
	name string, args *GuestAgentArgs, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230701preview:GuestAgent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GuestAgent
	err := ctx.RegisterResource("azure-native:azurestackhci/v20230901preview:GuestAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGuestAgent gets an existing GuestAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGuestAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestAgentState, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	var resource GuestAgent
	err := ctx.ReadResource("azure-native:azurestackhci/v20230901preview:GuestAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GuestAgent resources.
type guestAgentState struct {
}

type GuestAgentState struct {
}

func (GuestAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestAgentState)(nil)).Elem()
}

type guestAgentArgs struct {
	// Username / Password Credentials to provision guest agent.
	Credentials *GuestCredential `pulumi:"credentials"`
	// The guest agent provisioning action.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a GuestAgent resource.
type GuestAgentArgs struct {
	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialPtrInput
	// The guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrInput
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri pulumi.StringInput
}

func (GuestAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestAgentArgs)(nil)).Elem()
}

type GuestAgentInput interface {
	pulumi.Input

	ToGuestAgentOutput() GuestAgentOutput
	ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput
}

func (*GuestAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestAgent)(nil)).Elem()
}

func (i *GuestAgent) ToGuestAgentOutput() GuestAgentOutput {
	return i.ToGuestAgentOutputWithContext(context.Background())
}

func (i *GuestAgent) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestAgentOutput)
}

func (i *GuestAgent) ToOutput(ctx context.Context) pulumix.Output[*GuestAgent] {
	return pulumix.Output[*GuestAgent]{
		OutputState: i.ToGuestAgentOutputWithContext(ctx).OutputState,
	}
}

type GuestAgentOutput struct{ *pulumi.OutputState }

func (GuestAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestAgent)(nil)).Elem()
}

func (o GuestAgentOutput) ToGuestAgentOutput() GuestAgentOutput {
	return o
}

func (o GuestAgentOutput) ToGuestAgentOutputWithContext(ctx context.Context) GuestAgentOutput {
	return o
}

func (o GuestAgentOutput) ToOutput(ctx context.Context) pulumix.Output[*GuestAgent] {
	return pulumix.Output[*GuestAgent]{
		OutputState: o.OutputState,
	}
}

// Username / Password Credentials to provision guest agent.
func (o GuestAgentOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v *GuestAgent) GuestCredentialResponsePtrOutput { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

// The name of the resource
func (o GuestAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The guest agent provisioning action.
func (o GuestAgentOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringPtrOutput { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

// The provisioning state.
func (o GuestAgentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The guest agent status.
func (o GuestAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GuestAgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestAgent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GuestAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestAgentOutput{})
}