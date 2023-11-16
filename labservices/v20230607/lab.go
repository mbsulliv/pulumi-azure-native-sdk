// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230607

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The lab resource.
type Lab struct {
	pulumi.CustomResourceState

	// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
	AutoShutdownProfile AutoShutdownProfileResponseOutput `pulumi:"autoShutdownProfile"`
	// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
	ConnectionProfile ConnectionProfileResponseOutput `pulumi:"connectionProfile"`
	// The description of the lab.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
	LabPlanId pulumi.StringPtrOutput `pulumi:"labPlanId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
	NetworkProfile LabNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Current provisioning state of the lab.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Error details of last operation done on lab.
	ResourceOperationError ResourceOperationErrorResponseOutput `pulumi:"resourceOperationError"`
	// The lab user list management profile.
	RosterProfile RosterProfileResponsePtrOutput `pulumi:"rosterProfile"`
	// The lab security profile.
	SecurityProfile SecurityProfileResponseOutput `pulumi:"securityProfile"`
	// The lab state.
	State pulumi.StringOutput `pulumi:"state"`
	// Metadata pertaining to creation and last modification of the lab.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The title of the lab.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The profile used for creating lab virtual machines.
	VirtualMachineProfile VirtualMachineProfileResponseOutput `pulumi:"virtualMachineProfile"`
}

// NewLab registers a new resource with the given unique name, arguments, and options.
func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoShutdownProfile == nil {
		return nil, errors.New("invalid value for required argument 'AutoShutdownProfile'")
	}
	if args.ConnectionProfile == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecurityProfile == nil {
		return nil, errors.New("invalid value for required argument 'SecurityProfile'")
	}
	if args.VirtualMachineProfile == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineProfile'")
	}
	args.AutoShutdownProfile = args.AutoShutdownProfile.ToAutoShutdownProfileOutput().ApplyT(func(v AutoShutdownProfile) AutoShutdownProfile { return *v.Defaults() }).(AutoShutdownProfileOutput)
	args.ConnectionProfile = args.ConnectionProfile.ToConnectionProfileOutput().ApplyT(func(v ConnectionProfile) ConnectionProfile { return *v.Defaults() }).(ConnectionProfileOutput)
	args.VirtualMachineProfile = args.VirtualMachineProfile.ToVirtualMachineProfileOutput().ApplyT(func(v VirtualMachineProfile) VirtualMachineProfile { return *v.Defaults() }).(VirtualMachineProfileOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices:Lab"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211001preview:Lab"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:Lab"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20220801:Lab"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Lab
	err := ctx.RegisterResource("azure-native:labservices/v20230607:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLab gets an existing Lab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:labservices/v20230607:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Lab resources.
type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
	AutoShutdownProfile AutoShutdownProfile `pulumi:"autoShutdownProfile"`
	// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
	ConnectionProfile ConnectionProfile `pulumi:"connectionProfile"`
	// The description of the lab.
	Description *string `pulumi:"description"`
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName *string `pulumi:"labName"`
	// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
	LabPlanId *string `pulumi:"labPlanId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
	NetworkProfile *LabNetworkProfile `pulumi:"networkProfile"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The lab user list management profile.
	RosterProfile *RosterProfile `pulumi:"rosterProfile"`
	// The lab security profile.
	SecurityProfile SecurityProfile `pulumi:"securityProfile"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The title of the lab.
	Title *string `pulumi:"title"`
	// The profile used for creating lab virtual machines.
	VirtualMachineProfile VirtualMachineProfile `pulumi:"virtualMachineProfile"`
}

// The set of arguments for constructing a Lab resource.
type LabArgs struct {
	// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
	AutoShutdownProfile AutoShutdownProfileInput
	// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
	ConnectionProfile ConnectionProfileInput
	// The description of the lab.
	Description pulumi.StringPtrInput
	// The name of the lab that uniquely identifies it within containing lab plan. Used in resource URIs.
	LabName pulumi.StringPtrInput
	// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
	LabPlanId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
	NetworkProfile LabNetworkProfilePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The lab user list management profile.
	RosterProfile RosterProfilePtrInput
	// The lab security profile.
	SecurityProfile SecurityProfileInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The title of the lab.
	Title pulumi.StringPtrInput
	// The profile used for creating lab virtual machines.
	VirtualMachineProfile VirtualMachineProfileInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

// The resource auto shutdown configuration for the lab. This controls whether actions are taken on resources that are sitting idle.
func (o LabOutput) AutoShutdownProfile() AutoShutdownProfileResponseOutput {
	return o.ApplyT(func(v *Lab) AutoShutdownProfileResponseOutput { return v.AutoShutdownProfile }).(AutoShutdownProfileResponseOutput)
}

// The connection profile for the lab. This controls settings such as web access to lab resources or whether RDP or SSH ports are open.
func (o LabOutput) ConnectionProfile() ConnectionProfileResponseOutput {
	return o.ApplyT(func(v *Lab) ConnectionProfileResponseOutput { return v.ConnectionProfile }).(ConnectionProfileResponseOutput)
}

// The description of the lab.
func (o LabOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the lab plan. Used during resource creation to provide defaults and acts as a permission container when creating a lab via labs.azure.com. Setting a labPlanId on an existing lab provides organization..
func (o LabOutput) LabPlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.LabPlanId }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LabOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network profile for the lab, typically applied via a lab plan. This profile cannot be modified once a lab has been created.
func (o LabOutput) NetworkProfile() LabNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *Lab) LabNetworkProfileResponsePtrOutput { return v.NetworkProfile }).(LabNetworkProfileResponsePtrOutput)
}

// Current provisioning state of the lab.
func (o LabOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Error details of last operation done on lab.
func (o LabOutput) ResourceOperationError() ResourceOperationErrorResponseOutput {
	return o.ApplyT(func(v *Lab) ResourceOperationErrorResponseOutput { return v.ResourceOperationError }).(ResourceOperationErrorResponseOutput)
}

// The lab user list management profile.
func (o LabOutput) RosterProfile() RosterProfileResponsePtrOutput {
	return o.ApplyT(func(v *Lab) RosterProfileResponsePtrOutput { return v.RosterProfile }).(RosterProfileResponsePtrOutput)
}

// The lab security profile.
func (o LabOutput) SecurityProfile() SecurityProfileResponseOutput {
	return o.ApplyT(func(v *Lab) SecurityProfileResponseOutput { return v.SecurityProfile }).(SecurityProfileResponseOutput)
}

// The lab state.
func (o LabOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the lab.
func (o LabOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Lab) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LabOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The title of the lab.
func (o LabOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The profile used for creating lab virtual machines.
func (o LabOutput) VirtualMachineProfile() VirtualMachineProfileResponseOutput {
	return o.ApplyT(func(v *Lab) VirtualMachineProfileResponseOutput { return v.VirtualMachineProfile }).(VirtualMachineProfileResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
