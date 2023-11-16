// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for iSCSI Target requests.
type IscsiTarget struct {
	pulumi.CustomResourceState

	// Mode for Target connectivity.
	AclMode pulumi.StringOutput `pulumi:"aclMode"`
	// List of private IPv4 addresses to connect to the iSCSI Target.
	Endpoints pulumi.StringArrayOutput `pulumi:"endpoints"`
	// List of LUNs to be exposed through iSCSI Target.
	Luns IscsiLunResponseArrayOutput `pulumi:"luns"`
	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// List of Azure resource ids that manage this resource.
	ManagedByExtended pulumi.StringArrayOutput `pulumi:"managedByExtended"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The port used by iSCSI Target portal group.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// State of the operation on the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of identifiers for active sessions on the iSCSI target
	Sessions pulumi.StringArrayOutput `pulumi:"sessions"`
	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticAcls AclResponseArrayOutput `pulumi:"staticAcls"`
	// Operational status of the iSCSI Target.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource metadata required by ARM RPC
	SystemData SystemMetadataResponseOutput `pulumi:"systemData"`
	// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn pulumi.StringOutput `pulumi:"targetIqn"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIscsiTarget registers a new resource with the given unique name, arguments, and options.
func NewIscsiTarget(ctx *pulumi.Context,
	name string, args *IscsiTargetArgs, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclMode == nil {
		return nil, errors.New("invalid value for required argument 'AclMode'")
	}
	if args.DiskPoolName == nil {
		return nil, errors.New("invalid value for required argument 'DiskPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagepool:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20200315preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210401preview:IscsiTarget"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IscsiTarget
	err := ctx.RegisterResource("azure-native:storagepool/v20210801:IscsiTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIscsiTarget gets an existing IscsiTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIscsiTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiTargetState, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	var resource IscsiTarget
	err := ctx.ReadResource("azure-native:storagepool/v20210801:IscsiTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IscsiTarget resources.
type iscsiTargetState struct {
}

type IscsiTargetState struct {
}

func (IscsiTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetState)(nil)).Elem()
}

type iscsiTargetArgs struct {
	// Mode for Target connectivity.
	AclMode string `pulumi:"aclMode"`
	// The name of the Disk Pool.
	DiskPoolName string `pulumi:"diskPoolName"`
	// The name of the iSCSI Target.
	IscsiTargetName *string `pulumi:"iscsiTargetName"`
	// List of LUNs to be exposed through iSCSI Target.
	Luns []IscsiLun `pulumi:"luns"`
	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string `pulumi:"managedBy"`
	// List of Azure resource ids that manage this resource.
	ManagedByExtended []string `pulumi:"managedByExtended"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticAcls []Acl `pulumi:"staticAcls"`
	// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string `pulumi:"targetIqn"`
}

// The set of arguments for constructing a IscsiTarget resource.
type IscsiTargetArgs struct {
	// Mode for Target connectivity.
	AclMode pulumi.StringInput
	// The name of the Disk Pool.
	DiskPoolName pulumi.StringInput
	// The name of the iSCSI Target.
	IscsiTargetName pulumi.StringPtrInput
	// List of LUNs to be exposed through iSCSI Target.
	Luns IscsiLunArrayInput
	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy pulumi.StringPtrInput
	// List of Azure resource ids that manage this resource.
	ManagedByExtended pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticAcls AclArrayInput
	// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn pulumi.StringPtrInput
}

func (IscsiTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetArgs)(nil)).Elem()
}

type IscsiTargetInput interface {
	pulumi.Input

	ToIscsiTargetOutput() IscsiTargetOutput
	ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput
}

func (*IscsiTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (i *IscsiTarget) ToIscsiTargetOutput() IscsiTargetOutput {
	return i.ToIscsiTargetOutputWithContext(context.Background())
}

func (i *IscsiTarget) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiTargetOutput)
}

type IscsiTargetOutput struct{ *pulumi.OutputState }

func (IscsiTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (o IscsiTargetOutput) ToIscsiTargetOutput() IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return o
}

// Mode for Target connectivity.
func (o IscsiTargetOutput) AclMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.AclMode }).(pulumi.StringOutput)
}

// List of private IPv4 addresses to connect to the iSCSI Target.
func (o IscsiTargetOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringArrayOutput { return v.Endpoints }).(pulumi.StringArrayOutput)
}

// List of LUNs to be exposed through iSCSI Target.
func (o IscsiTargetOutput) Luns() IscsiLunResponseArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) IscsiLunResponseArrayOutput { return v.Luns }).(IscsiLunResponseArrayOutput)
}

// Azure resource id. Indicates if this resource is managed by another Azure resource.
func (o IscsiTargetOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// List of Azure resource ids that manage this resource.
func (o IscsiTargetOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringArrayOutput { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o IscsiTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port used by iSCSI Target portal group.
func (o IscsiTargetOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// State of the operation on the resource.
func (o IscsiTargetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of identifiers for active sessions on the iSCSI target
func (o IscsiTargetOutput) Sessions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringArrayOutput { return v.Sessions }).(pulumi.StringArrayOutput)
}

// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
func (o IscsiTargetOutput) StaticAcls() AclResponseArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) AclResponseArrayOutput { return v.StaticAcls }).(AclResponseArrayOutput)
}

// Operational status of the iSCSI Target.
func (o IscsiTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource metadata required by ARM RPC
func (o IscsiTargetOutput) SystemData() SystemMetadataResponseOutput {
	return o.ApplyT(func(v *IscsiTarget) SystemMetadataResponseOutput { return v.SystemData }).(SystemMetadataResponseOutput)
}

// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
func (o IscsiTargetOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.TargetIqn }).(pulumi.StringOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o IscsiTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IscsiTargetOutput{})
}
