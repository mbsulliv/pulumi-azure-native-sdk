// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The preview of Virtual Machine Cloud Management from the Azure supports deploying and managing VMs on your Azure Stack Edge device from Azure Portal.
// For more information, refer to: https://docs.microsoft.com/en-us/azure/databox-online/azure-stack-edge-gpu-virtual-machine-overview
// By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/ for additional details.
// Azure REST API version: 2022-03-01. Prior API version in Azure Native 1.x: 2020-12-01.
type CloudEdgeManagementRole struct {
	pulumi.CustomResourceState

	// Edge Profile of the resource
	EdgeProfile EdgeProfileResponseOutput `pulumi:"edgeProfile"`
	// Role type.
	// Expected value is 'CloudEdgeManagement'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Local Edge Management Status
	LocalManagementStatus pulumi.StringOutput `pulumi:"localManagementStatus"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Role status.
	RoleStatus pulumi.StringOutput `pulumi:"roleStatus"`
	// Metadata pertaining to creation and last modification of Role
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCloudEdgeManagementRole registers a new resource with the given unique name, arguments, and options.
func NewCloudEdgeManagementRole(ctx *pulumi.Context,
	name string, args *CloudEdgeManagementRoleArgs, opts ...pulumi.ResourceOption) (*CloudEdgeManagementRole, error) {
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
	args.Kind = pulumi.String("CloudEdgeManagement")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:CloudEdgeManagementRole"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:CloudEdgeManagementRole"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CloudEdgeManagementRole
	err := ctx.RegisterResource("azure-native:databoxedge:CloudEdgeManagementRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudEdgeManagementRole gets an existing CloudEdgeManagementRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudEdgeManagementRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudEdgeManagementRoleState, opts ...pulumi.ResourceOption) (*CloudEdgeManagementRole, error) {
	var resource CloudEdgeManagementRole
	err := ctx.ReadResource("azure-native:databoxedge:CloudEdgeManagementRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudEdgeManagementRole resources.
type cloudEdgeManagementRoleState struct {
}

type CloudEdgeManagementRoleState struct {
}

func (CloudEdgeManagementRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEdgeManagementRoleState)(nil)).Elem()
}

type cloudEdgeManagementRoleArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// Role type.
	// Expected value is 'CloudEdgeManagement'.
	Kind string `pulumi:"kind"`
	// The role name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Role status.
	RoleStatus string `pulumi:"roleStatus"`
}

// The set of arguments for constructing a CloudEdgeManagementRole resource.
type CloudEdgeManagementRoleArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// Role type.
	// Expected value is 'CloudEdgeManagement'.
	Kind pulumi.StringInput
	// The role name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Role status.
	RoleStatus pulumi.StringInput
}

func (CloudEdgeManagementRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudEdgeManagementRoleArgs)(nil)).Elem()
}

type CloudEdgeManagementRoleInput interface {
	pulumi.Input

	ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput
	ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput
}

func (*CloudEdgeManagementRole) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEdgeManagementRole)(nil)).Elem()
}

func (i *CloudEdgeManagementRole) ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput {
	return i.ToCloudEdgeManagementRoleOutputWithContext(context.Background())
}

func (i *CloudEdgeManagementRole) ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudEdgeManagementRoleOutput)
}

type CloudEdgeManagementRoleOutput struct{ *pulumi.OutputState }

func (CloudEdgeManagementRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudEdgeManagementRole)(nil)).Elem()
}

func (o CloudEdgeManagementRoleOutput) ToCloudEdgeManagementRoleOutput() CloudEdgeManagementRoleOutput {
	return o
}

func (o CloudEdgeManagementRoleOutput) ToCloudEdgeManagementRoleOutputWithContext(ctx context.Context) CloudEdgeManagementRoleOutput {
	return o
}

// Edge Profile of the resource
func (o CloudEdgeManagementRoleOutput) EdgeProfile() EdgeProfileResponseOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) EdgeProfileResponseOutput { return v.EdgeProfile }).(EdgeProfileResponseOutput)
}

// Role type.
// Expected value is 'CloudEdgeManagement'.
func (o CloudEdgeManagementRoleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Local Edge Management Status
func (o CloudEdgeManagementRoleOutput) LocalManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) pulumi.StringOutput { return v.LocalManagementStatus }).(pulumi.StringOutput)
}

// The object name.
func (o CloudEdgeManagementRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Role status.
func (o CloudEdgeManagementRoleOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) pulumi.StringOutput { return v.RoleStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Role
func (o CloudEdgeManagementRoleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o CloudEdgeManagementRoleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudEdgeManagementRole) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudEdgeManagementRoleOutput{})
}
