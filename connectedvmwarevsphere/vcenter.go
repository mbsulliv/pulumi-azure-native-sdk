// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the vCenter.
// Azure REST API version: 2022-07-15-preview. Prior API version in Azure Native 1.x: 2020-10-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-10-01.
type VCenter struct {
	pulumi.CustomResourceState

	// Gets or sets the connection status to the vCenter.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// Username / Password Credentials to connect to vcenter.
	Credentials VICredentialResponsePtrOutput `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Gets or sets the FQDN/IPAddress of the vCenter.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Gets or sets the instance UUID of the vCenter.
	InstanceUuid pulumi.StringOutput `pulumi:"instanceUuid"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets or sets the location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Gets or sets the name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the port of the vCenter.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource status information.
	Statuses ResourceStatusResponseArrayOutput `pulumi:"statuses"`
	// The system data.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Gets or sets the version of the vCenter.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewVCenter registers a new resource with the given unique name, arguments, and options.
func NewVCenter(ctx *pulumi.Context,
	name string, args *VCenterArgs, opts ...pulumi.ResourceOption) (*VCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fqdn == nil {
		return nil, errors.New("invalid value for required argument 'Fqdn'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20230301preview:VCenter"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20231001:VCenter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VCenter
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:VCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVCenter gets an existing VCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VCenterState, opts ...pulumi.ResourceOption) (*VCenter, error) {
	var resource VCenter
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:VCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VCenter resources.
type vcenterState struct {
}

type VCenterState struct {
}

func (VCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*vcenterState)(nil)).Elem()
}

type vcenterArgs struct {
	// Username / Password Credentials to connect to vcenter.
	Credentials *VICredential `pulumi:"credentials"`
	// Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Gets or sets the FQDN/IPAddress of the vCenter.
	Fqdn string `pulumi:"fqdn"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the location.
	Location *string `pulumi:"location"`
	// Gets or sets the port of the vCenter.
	Port *int `pulumi:"port"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Name of the vCenter.
	VcenterName *string `pulumi:"vcenterName"`
}

// The set of arguments for constructing a VCenter resource.
type VCenterArgs struct {
	// Username / Password Credentials to connect to vcenter.
	Credentials VICredentialPtrInput
	// Gets or sets the extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// Gets or sets the FQDN/IPAddress of the vCenter.
	Fqdn pulumi.StringInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// Gets or sets the location.
	Location pulumi.StringPtrInput
	// Gets or sets the port of the vCenter.
	Port pulumi.IntPtrInput
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the Resource tags.
	Tags pulumi.StringMapInput
	// Name of the vCenter.
	VcenterName pulumi.StringPtrInput
}

func (VCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vcenterArgs)(nil)).Elem()
}

type VCenterInput interface {
	pulumi.Input

	ToVCenterOutput() VCenterOutput
	ToVCenterOutputWithContext(ctx context.Context) VCenterOutput
}

func (*VCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenter)(nil)).Elem()
}

func (i *VCenter) ToVCenterOutput() VCenterOutput {
	return i.ToVCenterOutputWithContext(context.Background())
}

func (i *VCenter) ToVCenterOutputWithContext(ctx context.Context) VCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCenterOutput)
}

type VCenterOutput struct{ *pulumi.OutputState }

func (VCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenter)(nil)).Elem()
}

func (o VCenterOutput) ToVCenterOutput() VCenterOutput {
	return o
}

func (o VCenterOutput) ToVCenterOutputWithContext(ctx context.Context) VCenterOutput {
	return o
}

// Gets or sets the connection status to the vCenter.
func (o VCenterOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// Username / Password Credentials to connect to vcenter.
func (o VCenterOutput) Credentials() VICredentialResponsePtrOutput {
	return o.ApplyT(func(v *VCenter) VICredentialResponsePtrOutput { return v.Credentials }).(VICredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o VCenterOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// Gets or sets the extended location.
func (o VCenterOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VCenter) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Gets or sets the FQDN/IPAddress of the vCenter.
func (o VCenterOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Gets or sets the instance UUID of the vCenter.
func (o VCenterOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.InstanceUuid }).(pulumi.StringOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o VCenterOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o VCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the name.
func (o VCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the port of the vCenter.
func (o VCenterOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VCenter) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// Gets or sets the provisioning state.
func (o VCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource status information.
func (o VCenterOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VCenter) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

// The system data.
func (o VCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets the Resource tags.
func (o VCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o VCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o VCenterOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Gets or sets the version of the vCenter.
func (o VCenterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *VCenter) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VCenterOutput{})
}
