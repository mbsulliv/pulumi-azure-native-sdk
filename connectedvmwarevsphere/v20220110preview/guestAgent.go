// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the GuestAgent.
type GuestAgent struct {
	pulumi.CustomResourceState

	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialResponsePtrOutput `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig HttpProxyConfigurationResponsePtrOutput `pulumi:"httpProxyConfig"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrOutput `pulumi:"provisioningAction"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the guest agent status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The resource status information.
	Statuses ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewGuestAgent registers a new resource with the given unique name, arguments, and options.
func NewGuestAgent(ctx *pulumi.Context,
	name string, args *GuestAgentArgs, opts ...pulumi.ResourceOption) (*GuestAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:GuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:GuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:GuestAgent"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestAgent
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20220110preview:GuestAgent", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20220110preview:GuestAgent", name, id, state, &resource, opts...)
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
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig *HttpProxyConfiguration `pulumi:"httpProxyConfig"`
	// Name of the guestAgents.
	Name *string `pulumi:"name"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vm.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a GuestAgent resource.
type GuestAgentArgs struct {
	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialPtrInput
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig HttpProxyConfigurationPtrInput
	// Name of the guestAgents.
	Name pulumi.StringPtrInput
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Name of the vm.
	VirtualMachineName pulumi.StringInput
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

// Username / Password Credentials to provision guest agent.
func (o GuestAgentOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v *GuestAgent) GuestCredentialResponsePtrOutput { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o GuestAgentOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// HTTP Proxy configuration for the VM.
func (o GuestAgentOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GuestAgent) HttpProxyConfigurationResponsePtrOutput { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

// The name of the resource
func (o GuestAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the guest agent provisioning action.
func (o GuestAgentOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringPtrOutput { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

// Gets or sets the provisioning state.
func (o GuestAgentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the guest agent status.
func (o GuestAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The resource status information.
func (o GuestAgentOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *GuestAgent) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o GuestAgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestAgent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GuestAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o GuestAgentOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestAgent) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestAgentOutput{})
}
