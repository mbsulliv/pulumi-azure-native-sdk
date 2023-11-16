// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a devcenter resource.
type DevCenter struct {
	pulumi.CustomResourceState

	// The URI of the Dev Center.
	DevCenterUri pulumi.StringOutput `pulumi:"devCenterUri"`
	// Managed identity properties
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDevCenter registers a new resource with the given unique name, arguments, and options.
func NewDevCenter(ctx *pulumi.Context,
	name string, args *DevCenterArgs, opts ...pulumi.ResourceOption) (*DevCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:DevCenter"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:DevCenter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DevCenter
	err := ctx.RegisterResource("azure-native:devcenter/v20230401:DevCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevCenter gets an existing DevCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DevCenterState, opts ...pulumi.ResourceOption) (*DevCenter, error) {
	var resource DevCenter
	err := ctx.ReadResource("azure-native:devcenter/v20230401:DevCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DevCenter resources.
type devCenterState struct {
}

type DevCenterState struct {
}

func (DevCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*devCenterState)(nil)).Elem()
}

type devCenterArgs struct {
	// The name of the devcenter.
	DevCenterName *string `pulumi:"devCenterName"`
	// Managed identity properties
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DevCenter resource.
type DevCenterArgs struct {
	// The name of the devcenter.
	DevCenterName pulumi.StringPtrInput
	// Managed identity properties
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DevCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*devCenterArgs)(nil)).Elem()
}

type DevCenterInput interface {
	pulumi.Input

	ToDevCenterOutput() DevCenterOutput
	ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput
}

func (*DevCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**DevCenter)(nil)).Elem()
}

func (i *DevCenter) ToDevCenterOutput() DevCenterOutput {
	return i.ToDevCenterOutputWithContext(context.Background())
}

func (i *DevCenter) ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevCenterOutput)
}

type DevCenterOutput struct{ *pulumi.OutputState }

func (DevCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DevCenter)(nil)).Elem()
}

func (o DevCenterOutput) ToDevCenterOutput() DevCenterOutput {
	return o
}

func (o DevCenterOutput) ToDevCenterOutputWithContext(ctx context.Context) DevCenterOutput {
	return o
}

// The URI of the Dev Center.
func (o DevCenterOutput) DevCenterUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.DevCenterUri }).(pulumi.StringOutput)
}

// Managed identity properties
func (o DevCenterOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DevCenter) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o DevCenterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DevCenterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o DevCenterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DevCenterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DevCenter) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DevCenterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DevCenterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DevCenter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DevCenterOutput{})
}
