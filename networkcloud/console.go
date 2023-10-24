// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2023-05-01-preview. Prior API version in Azure Native 1.x: 2022-12-12-preview.
//
// Other available API versions: 2023-07-01.
type Console struct {
	pulumi.CustomResourceState

	// The more detailed status of the console.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The indicator of whether the console access is enabled.
	Enabled pulumi.StringOutput `pulumi:"enabled"`
	// The date and time after which the key will be disallowed access.
	Expiration pulumi.StringPtrOutput `pulumi:"expiration"`
	// The extended location of the cluster manager associated with the cluster this virtual machine is created on.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the private link service that is used to provide virtual machine console access.
	PrivateLinkServiceId pulumi.StringOutput `pulumi:"privateLinkServiceId"`
	// The provisioning state of the virtual machine console.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
	SshPublicKey SshPublicKeyResponseOutput `pulumi:"sshPublicKey"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique identifier for the virtual machine that is used to access the console.
	VirtualMachineAccessId pulumi.StringOutput `pulumi:"virtualMachineAccessId"`
}

// NewConsole registers a new resource with the given unique name, arguments, and options.
func NewConsole(ctx *pulumi.Context,
	name string, args *ConsoleArgs, opts ...pulumi.ResourceOption) (*Console, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SshPublicKey == nil {
		return nil, errors.New("invalid value for required argument 'SshPublicKey'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:Console"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:Console"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:Console"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Console
	err := ctx.RegisterResource("azure-native:networkcloud:Console", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsole gets an existing Console resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsoleState, opts ...pulumi.ResourceOption) (*Console, error) {
	var resource Console
	err := ctx.ReadResource("azure-native:networkcloud:Console", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Console resources.
type consoleState struct {
}

type ConsoleState struct {
}

func (ConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleState)(nil)).Elem()
}

type consoleArgs struct {
	// The name of the virtual machine console.
	ConsoleName *string `pulumi:"consoleName"`
	// The indicator of whether the console access is enabled.
	Enabled string `pulumi:"enabled"`
	// The date and time after which the key will be disallowed access.
	Expiration *string `pulumi:"expiration"`
	// The extended location of the cluster manager associated with the cluster this virtual machine is created on.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
	SshPublicKey SshPublicKey `pulumi:"sshPublicKey"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the virtual machine.
	VirtualMachineName string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a Console resource.
type ConsoleArgs struct {
	// The name of the virtual machine console.
	ConsoleName pulumi.StringPtrInput
	// The indicator of whether the console access is enabled.
	Enabled pulumi.StringInput
	// The date and time after which the key will be disallowed access.
	Expiration pulumi.StringPtrInput
	// The extended location of the cluster manager associated with the cluster this virtual machine is created on.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
	SshPublicKey SshPublicKeyInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the virtual machine.
	VirtualMachineName pulumi.StringInput
}

func (ConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleArgs)(nil)).Elem()
}

type ConsoleInput interface {
	pulumi.Input

	ToConsoleOutput() ConsoleOutput
	ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput
}

func (*Console) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (i *Console) ToConsoleOutput() ConsoleOutput {
	return i.ToConsoleOutputWithContext(context.Background())
}

func (i *Console) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleOutput)
}

func (i *Console) ToOutput(ctx context.Context) pulumix.Output[*Console] {
	return pulumix.Output[*Console]{
		OutputState: i.ToConsoleOutputWithContext(ctx).OutputState,
	}
}

type ConsoleOutput struct{ *pulumi.OutputState }

func (ConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (o ConsoleOutput) ToConsoleOutput() ConsoleOutput {
	return o
}

func (o ConsoleOutput) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return o
}

func (o ConsoleOutput) ToOutput(ctx context.Context) pulumix.Output[*Console] {
	return pulumix.Output[*Console]{
		OutputState: o.OutputState,
	}
}

// The more detailed status of the console.
func (o ConsoleOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o ConsoleOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The indicator of whether the console access is enabled.
func (o ConsoleOutput) Enabled() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Enabled }).(pulumi.StringOutput)
}

// The date and time after which the key will be disallowed access.
func (o ConsoleOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Console) pulumi.StringPtrOutput { return v.Expiration }).(pulumi.StringPtrOutput)
}

// The extended location of the cluster manager associated with the cluster this virtual machine is created on.
func (o ConsoleOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Console) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o ConsoleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConsoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the private link service that is used to provide virtual machine console access.
func (o ConsoleOutput) PrivateLinkServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.PrivateLinkServiceId }).(pulumi.StringOutput)
}

// The provisioning state of the virtual machine console.
func (o ConsoleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SSH public key that will be provisioned for user access. The user is expected to have the corresponding SSH private key for logging in.
func (o ConsoleOutput) SshPublicKey() SshPublicKeyResponseOutput {
	return o.ApplyT(func(v *Console) SshPublicKeyResponseOutput { return v.SshPublicKey }).(SshPublicKeyResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConsoleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Console) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConsoleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Console) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConsoleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique identifier for the virtual machine that is used to access the console.
func (o ConsoleOutput) VirtualMachineAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.VirtualMachineAccessId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConsoleOutput{})
}
