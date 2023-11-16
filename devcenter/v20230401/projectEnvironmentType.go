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

// Represents an environment type.
type ProjectEnvironmentType struct {
	pulumi.CustomResourceState

	// The role definition assigned to the environment creator on backing resources.
	CreatorRoleAssignment ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput `pulumi:"creatorRoleAssignment"`
	// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
	DeploymentTargetId pulumi.StringPtrOutput `pulumi:"deploymentTargetId"`
	// Managed identity properties
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location for the environment type
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines whether this Environment Type can be used in this Project.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
	UserRoleAssignments UserRoleAssignmentResponseMapOutput `pulumi:"userRoleAssignments"`
}

// NewProjectEnvironmentType registers a new resource with the given unique name, arguments, and options.
func NewProjectEnvironmentType(ctx *pulumi.Context,
	name string, args *ProjectEnvironmentTypeArgs, opts ...pulumi.ResourceOption) (*ProjectEnvironmentType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230801preview:ProjectEnvironmentType"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:ProjectEnvironmentType"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProjectEnvironmentType
	err := ctx.RegisterResource("azure-native:devcenter/v20230401:ProjectEnvironmentType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectEnvironmentType gets an existing ProjectEnvironmentType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectEnvironmentType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectEnvironmentTypeState, opts ...pulumi.ResourceOption) (*ProjectEnvironmentType, error) {
	var resource ProjectEnvironmentType
	err := ctx.ReadResource("azure-native:devcenter/v20230401:ProjectEnvironmentType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectEnvironmentType resources.
type projectEnvironmentTypeState struct {
}

type ProjectEnvironmentTypeState struct {
}

func (ProjectEnvironmentTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentTypeState)(nil)).Elem()
}

type projectEnvironmentTypeArgs struct {
	// The role definition assigned to the environment creator on backing resources.
	CreatorRoleAssignment *ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment `pulumi:"creatorRoleAssignment"`
	// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
	DeploymentTargetId *string `pulumi:"deploymentTargetId"`
	// The name of the environment type.
	EnvironmentTypeName *string `pulumi:"environmentTypeName"`
	// Managed identity properties
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location for the environment type
	Location *string `pulumi:"location"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines whether this Environment Type can be used in this Project.
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
	UserRoleAssignments map[string]UserRoleAssignment `pulumi:"userRoleAssignments"`
}

// The set of arguments for constructing a ProjectEnvironmentType resource.
type ProjectEnvironmentTypeArgs struct {
	// The role definition assigned to the environment creator on backing resources.
	CreatorRoleAssignment ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignmentPtrInput
	// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
	DeploymentTargetId pulumi.StringPtrInput
	// The name of the environment type.
	EnvironmentTypeName pulumi.StringPtrInput
	// Managed identity properties
	Identity ManagedServiceIdentityPtrInput
	// The geo-location for the environment type
	Location pulumi.StringPtrInput
	// The name of the project.
	ProjectName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Defines whether this Environment Type can be used in this Project.
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
	UserRoleAssignments UserRoleAssignmentMapInput
}

func (ProjectEnvironmentTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectEnvironmentTypeArgs)(nil)).Elem()
}

type ProjectEnvironmentTypeInput interface {
	pulumi.Input

	ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput
	ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput
}

func (*ProjectEnvironmentType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentType)(nil)).Elem()
}

func (i *ProjectEnvironmentType) ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput {
	return i.ToProjectEnvironmentTypeOutputWithContext(context.Background())
}

func (i *ProjectEnvironmentType) ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentTypeOutput)
}

type ProjectEnvironmentTypeOutput struct{ *pulumi.OutputState }

func (ProjectEnvironmentTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectEnvironmentType)(nil)).Elem()
}

func (o ProjectEnvironmentTypeOutput) ToProjectEnvironmentTypeOutput() ProjectEnvironmentTypeOutput {
	return o
}

func (o ProjectEnvironmentTypeOutput) ToProjectEnvironmentTypeOutputWithContext(ctx context.Context) ProjectEnvironmentTypeOutput {
	return o
}

// The role definition assigned to the environment creator on backing resources.
func (o ProjectEnvironmentTypeOutput) CreatorRoleAssignment() ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput {
		return v.CreatorRoleAssignment
	}).(ProjectEnvironmentTypeUpdatePropertiesResponseCreatorRoleAssignmentPtrOutput)
}

// Id of a subscription that the environment type will be mapped to. The environment's resources will be deployed into this subscription.
func (o ProjectEnvironmentTypeOutput) DeploymentTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.DeploymentTargetId }).(pulumi.StringPtrOutput)
}

// Managed identity properties
func (o ProjectEnvironmentTypeOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location for the environment type
func (o ProjectEnvironmentTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ProjectEnvironmentTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o ProjectEnvironmentTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines whether this Environment Type can be used in this Project.
func (o ProjectEnvironmentTypeOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ProjectEnvironmentTypeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ProjectEnvironmentTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProjectEnvironmentTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Role Assignments created on environment backing resources. This is a mapping from a user object ID to an object of role definition IDs.
func (o ProjectEnvironmentTypeOutput) UserRoleAssignments() UserRoleAssignmentResponseMapOutput {
	return o.ApplyT(func(v *ProjectEnvironmentType) UserRoleAssignmentResponseMapOutput { return v.UserRoleAssignments }).(UserRoleAssignmentResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectEnvironmentTypeOutput{})
}
