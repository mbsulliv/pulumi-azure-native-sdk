// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The VmmServers resource definition.
type VmmServer struct {
	pulumi.CustomResourceState

	// Gets or sets the connection status to the vmmServer.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// Credentials to connect to VMMServer.
	Credentials VMMServerPropertiesResponseCredentialsPtrOutput `pulumi:"credentials"`
	// Gets or sets any error message if connection to vmmServer is having any issue.
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Fqdn is the hostname/ip of the vmmServer.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Port is the port on which the vmmServer is listening.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique ID of vmmServer.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Version is the version of the vmmSever.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewVmmServer registers a new resource with the given unique name, arguments, and options.
func NewVmmServer(ctx *pulumi.Context,
	name string, args *VmmServerArgs, opts ...pulumi.ResourceOption) (*VmmServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.Fqdn == nil {
		return nil, errors.New("invalid value for required argument 'Fqdn'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VmmServer"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:VmmServer"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20220521preview:VmmServer"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20231007:VmmServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VmmServer
	err := ctx.RegisterResource("azure-native:scvmm/v20230401preview:VmmServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmmServer gets an existing VmmServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmmServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmmServerState, opts ...pulumi.ResourceOption) (*VmmServer, error) {
	var resource VmmServer
	err := ctx.ReadResource("azure-native:scvmm/v20230401preview:VmmServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmmServer resources.
type vmmServerState struct {
}

type VmmServerState struct {
}

func (VmmServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmmServerState)(nil)).Elem()
}

type vmmServerArgs struct {
	// Credentials to connect to VMMServer.
	Credentials *VMMServerPropertiesCredentials `pulumi:"credentials"`
	// The extended location.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Fqdn is the hostname/ip of the vmmServer.
	Fqdn string `pulumi:"fqdn"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Port is the port on which the vmmServer is listening.
	Port *int `pulumi:"port"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Name of the VMMServer.
	VmmServerName *string `pulumi:"vmmServerName"`
}

// The set of arguments for constructing a VmmServer resource.
type VmmServerArgs struct {
	// Credentials to connect to VMMServer.
	Credentials VMMServerPropertiesCredentialsPtrInput
	// The extended location.
	ExtendedLocation ExtendedLocationInput
	// Fqdn is the hostname/ip of the vmmServer.
	Fqdn pulumi.StringInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Port is the port on which the vmmServer is listening.
	Port pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Name of the VMMServer.
	VmmServerName pulumi.StringPtrInput
}

func (VmmServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmmServerArgs)(nil)).Elem()
}

type VmmServerInput interface {
	pulumi.Input

	ToVmmServerOutput() VmmServerOutput
	ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput
}

func (*VmmServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VmmServer)(nil)).Elem()
}

func (i *VmmServer) ToVmmServerOutput() VmmServerOutput {
	return i.ToVmmServerOutputWithContext(context.Background())
}

func (i *VmmServer) ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmServerOutput)
}

func (i *VmmServer) ToOutput(ctx context.Context) pulumix.Output[*VmmServer] {
	return pulumix.Output[*VmmServer]{
		OutputState: i.ToVmmServerOutputWithContext(ctx).OutputState,
	}
}

type VmmServerOutput struct{ *pulumi.OutputState }

func (VmmServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmmServer)(nil)).Elem()
}

func (o VmmServerOutput) ToVmmServerOutput() VmmServerOutput {
	return o
}

func (o VmmServerOutput) ToVmmServerOutputWithContext(ctx context.Context) VmmServerOutput {
	return o
}

func (o VmmServerOutput) ToOutput(ctx context.Context) pulumix.Output[*VmmServer] {
	return pulumix.Output[*VmmServer]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the connection status to the vmmServer.
func (o VmmServerOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Credentials to connect to VMMServer.
func (o VmmServerOutput) Credentials() VMMServerPropertiesResponseCredentialsPtrOutput {
	return o.ApplyT(func(v *VmmServer) VMMServerPropertiesResponseCredentialsPtrOutput { return v.Credentials }).(VMMServerPropertiesResponseCredentialsPtrOutput)
}

// Gets or sets any error message if connection to vmmServer is having any issue.
func (o VmmServerOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The extended location.
func (o VmmServerOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VmmServer) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fqdn is the hostname/ip of the vmmServer.
func (o VmmServerOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Gets or sets the location.
func (o VmmServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o VmmServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port is the port on which the vmmServer is listening.
func (o VmmServerOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// Gets or sets the provisioning state.
func (o VmmServerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system data.
func (o VmmServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VmmServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o VmmServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o VmmServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of vmmServer.
func (o VmmServerOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Version is the version of the vmmSever.
func (o VmmServerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *VmmServer) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VmmServerOutput{})
}
