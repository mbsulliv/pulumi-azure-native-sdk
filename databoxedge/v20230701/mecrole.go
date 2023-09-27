// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// MEC role.
type MECRole struct {
	pulumi.CustomResourceState

	// Activation key of the MEC.
	ConnectionString AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"connectionString"`
	// Controller Endpoint.
	ControllerEndpoint pulumi.StringPtrOutput `pulumi:"controllerEndpoint"`
	// Role type.
	// Expected value is 'MEC'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique Id of the Resource.
	ResourceUniqueId pulumi.StringPtrOutput `pulumi:"resourceUniqueId"`
	// Role status.
	RoleStatus pulumi.StringOutput `pulumi:"roleStatus"`
	// Metadata pertaining to creation and last modification of Role
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMECRole registers a new resource with the given unique name, arguments, and options.
func NewMECRole(ctx *pulumi.Context,
	name string, args *MECRoleArgs, opts ...pulumi.ResourceOption) (*MECRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RoleStatus == nil {
		return nil, errors.New("invalid value for required argument 'RoleStatus'")
	}
	args.Kind = pulumi.String("MEC")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:MECRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:MECRole"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MECRole
	err := ctx.RegisterResource("azure-native:databoxedge/v20230701:MECRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMECRole gets an existing MECRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMECRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MECRoleState, opts ...pulumi.ResourceOption) (*MECRole, error) {
	var resource MECRole
	err := ctx.ReadResource("azure-native:databoxedge/v20230701:MECRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MECRole resources.
type mecroleState struct {
}

type MECRoleState struct {
}

func (MECRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*mecroleState)(nil)).Elem()
}

type mecroleArgs struct {
	// Activation key of the MEC.
	ConnectionString *AsymmetricEncryptedSecret `pulumi:"connectionString"`
	// Controller Endpoint.
	ControllerEndpoint *string `pulumi:"controllerEndpoint"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Role type.
	// Expected value is 'MEC'.
	Kind string `pulumi:"kind"`
	// The role name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Unique Id of the Resource.
	ResourceUniqueId *string `pulumi:"resourceUniqueId"`
	// Role status.
	RoleStatus string `pulumi:"roleStatus"`
}

// The set of arguments for constructing a MECRole resource.
type MECRoleArgs struct {
	// Activation key of the MEC.
	ConnectionString AsymmetricEncryptedSecretPtrInput
	// Controller Endpoint.
	ControllerEndpoint pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// Role type.
	// Expected value is 'MEC'.
	Kind pulumi.StringInput
	// The role name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Unique Id of the Resource.
	ResourceUniqueId pulumi.StringPtrInput
	// Role status.
	RoleStatus pulumi.StringInput
}

func (MECRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mecroleArgs)(nil)).Elem()
}

type MECRoleInput interface {
	pulumi.Input

	ToMECRoleOutput() MECRoleOutput
	ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput
}

func (*MECRole) ElementType() reflect.Type {
	return reflect.TypeOf((**MECRole)(nil)).Elem()
}

func (i *MECRole) ToMECRoleOutput() MECRoleOutput {
	return i.ToMECRoleOutputWithContext(context.Background())
}

func (i *MECRole) ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MECRoleOutput)
}

func (i *MECRole) ToOutput(ctx context.Context) pulumix.Output[*MECRole] {
	return pulumix.Output[*MECRole]{
		OutputState: i.ToMECRoleOutputWithContext(ctx).OutputState,
	}
}

type MECRoleOutput struct{ *pulumi.OutputState }

func (MECRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MECRole)(nil)).Elem()
}

func (o MECRoleOutput) ToMECRoleOutput() MECRoleOutput {
	return o
}

func (o MECRoleOutput) ToMECRoleOutputWithContext(ctx context.Context) MECRoleOutput {
	return o
}

func (o MECRoleOutput) ToOutput(ctx context.Context) pulumix.Output[*MECRole] {
	return pulumix.Output[*MECRole]{
		OutputState: o.OutputState,
	}
}

// Activation key of the MEC.
func (o MECRoleOutput) ConnectionString() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *MECRole) AsymmetricEncryptedSecretResponsePtrOutput { return v.ConnectionString }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// Controller Endpoint.
func (o MECRoleOutput) ControllerEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringPtrOutput { return v.ControllerEndpoint }).(pulumi.StringPtrOutput)
}

// Role type.
// Expected value is 'MEC'.
func (o MECRoleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o MECRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique Id of the Resource.
func (o MECRoleOutput) ResourceUniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringPtrOutput { return v.ResourceUniqueId }).(pulumi.StringPtrOutput)
}

// Role status.
func (o MECRoleOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringOutput { return v.RoleStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Role
func (o MECRoleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MECRole) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o MECRoleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MECRole) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MECRoleOutput{})
}
