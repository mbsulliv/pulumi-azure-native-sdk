// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Assessment project site resource.
// Azure REST API version: 2023-03-15.
type AssessmentProjectsOperation struct {
	pulumi.CustomResourceState

	// Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
	AssessmentSolutionId pulumi.StringPtrOutput `pulumi:"assessmentSolutionId"`
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// The ARM id of the storage account used for interactions when public access is
	// disabled.
	CustomerStorageAccountArmId pulumi.StringPtrOutput `pulumi:"customerStorageAccountArmId"`
	// The ARM id of service map workspace created by customer.
	CustomerWorkspaceId pulumi.StringPtrOutput `pulumi:"customerWorkspaceId"`
	// Location of service map workspace created by customer.
	CustomerWorkspaceLocation pulumi.StringPtrOutput `pulumi:"customerWorkspaceLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of private endpoint connections to the project.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Assessment project status.
	ProjectStatus pulumi.StringPtrOutput `pulumi:"projectStatus"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// This value can be set to 'enabled' to avoid breaking changes on existing
	// customer resources and templates. If set to 'disabled', traffic over public
	// interface is not allowed, and private endpoint connections would be the
	// exclusive access method.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Endpoint at which the collector agent can call agent REST API.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601
	// format.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
}

// NewAssessmentProjectsOperation registers a new resource with the given unique name, arguments, and options.
func NewAssessmentProjectsOperation(ctx *pulumi.Context,
	name string, args *AssessmentProjectsOperationArgs, opts ...pulumi.ResourceOption) (*AssessmentProjectsOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:AssessmentProjectsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:AssessmentProjectsOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AssessmentProjectsOperation
	err := ctx.RegisterResource("azure-native:migrate:AssessmentProjectsOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentProjectsOperation gets an existing AssessmentProjectsOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentProjectsOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentProjectsOperationState, opts ...pulumi.ResourceOption) (*AssessmentProjectsOperation, error) {
	var resource AssessmentProjectsOperation
	err := ctx.ReadResource("azure-native:migrate:AssessmentProjectsOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentProjectsOperation resources.
type assessmentProjectsOperationState struct {
}

type AssessmentProjectsOperationState struct {
}

func (AssessmentProjectsOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentProjectsOperationState)(nil)).Elem()
}

type assessmentProjectsOperationArgs struct {
	// Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
	AssessmentSolutionId *string `pulumi:"assessmentSolutionId"`
	// The ARM id of the storage account used for interactions when public access is
	// disabled.
	CustomerStorageAccountArmId *string `pulumi:"customerStorageAccountArmId"`
	// The ARM id of service map workspace created by customer.
	CustomerWorkspaceId *string `pulumi:"customerWorkspaceId"`
	// Location of service map workspace created by customer.
	CustomerWorkspaceLocation *string `pulumi:"customerWorkspaceLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Assessment Project Name
	ProjectName *string `pulumi:"projectName"`
	// Assessment project status.
	ProjectStatus *string `pulumi:"projectStatus"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// This value can be set to 'enabled' to avoid breaking changes on existing
	// customer resources and templates. If set to 'disabled', traffic over public
	// interface is not allowed, and private endpoint connections would be the
	// exclusive access method.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AssessmentProjectsOperation resource.
type AssessmentProjectsOperationArgs struct {
	// Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
	AssessmentSolutionId pulumi.StringPtrInput
	// The ARM id of the storage account used for interactions when public access is
	// disabled.
	CustomerStorageAccountArmId pulumi.StringPtrInput
	// The ARM id of service map workspace created by customer.
	CustomerWorkspaceId pulumi.StringPtrInput
	// Location of service map workspace created by customer.
	CustomerWorkspaceLocation pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Assessment Project Name
	ProjectName pulumi.StringPtrInput
	// Assessment project status.
	ProjectStatus pulumi.StringPtrInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// This value can be set to 'enabled' to avoid breaking changes on existing
	// customer resources and templates. If set to 'disabled', traffic over public
	// interface is not allowed, and private endpoint connections would be the
	// exclusive access method.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AssessmentProjectsOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentProjectsOperationArgs)(nil)).Elem()
}

type AssessmentProjectsOperationInput interface {
	pulumi.Input

	ToAssessmentProjectsOperationOutput() AssessmentProjectsOperationOutput
	ToAssessmentProjectsOperationOutputWithContext(ctx context.Context) AssessmentProjectsOperationOutput
}

func (*AssessmentProjectsOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentProjectsOperation)(nil)).Elem()
}

func (i *AssessmentProjectsOperation) ToAssessmentProjectsOperationOutput() AssessmentProjectsOperationOutput {
	return i.ToAssessmentProjectsOperationOutputWithContext(context.Background())
}

func (i *AssessmentProjectsOperation) ToAssessmentProjectsOperationOutputWithContext(ctx context.Context) AssessmentProjectsOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentProjectsOperationOutput)
}

type AssessmentProjectsOperationOutput struct{ *pulumi.OutputState }

func (AssessmentProjectsOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentProjectsOperation)(nil)).Elem()
}

func (o AssessmentProjectsOperationOutput) ToAssessmentProjectsOperationOutput() AssessmentProjectsOperationOutput {
	return o
}

func (o AssessmentProjectsOperationOutput) ToAssessmentProjectsOperationOutputWithContext(ctx context.Context) AssessmentProjectsOperationOutput {
	return o
}

// Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
func (o AssessmentProjectsOperationOutput) AssessmentSolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.AssessmentSolutionId }).(pulumi.StringPtrOutput)
}

// Time when this project was created. Date-Time represented in ISO-8601 format.
func (o AssessmentProjectsOperationOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// The ARM id of the storage account used for interactions when public access is
// disabled.
func (o AssessmentProjectsOperationOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.CustomerStorageAccountArmId }).(pulumi.StringPtrOutput)
}

// The ARM id of service map workspace created by customer.
func (o AssessmentProjectsOperationOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

// Location of service map workspace created by customer.
func (o AssessmentProjectsOperationOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.CustomerWorkspaceLocation }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o AssessmentProjectsOperationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AssessmentProjectsOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections to the project.
func (o AssessmentProjectsOperationOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Assessment project status.
func (o AssessmentProjectsOperationOutput) ProjectStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.ProjectStatus }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o AssessmentProjectsOperationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// This value can be set to 'enabled' to avoid breaking changes on existing
// customer resources and templates. If set to 'disabled', traffic over public
// interface is not allowed, and private endpoint connections would be the
// exclusive access method.
func (o AssessmentProjectsOperationOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Endpoint at which the collector agent can call agent REST API.
func (o AssessmentProjectsOperationOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AssessmentProjectsOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AssessmentProjectsOperationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AssessmentProjectsOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Time when this project was last updated. Date-Time represented in ISO-8601
// format.
func (o AssessmentProjectsOperationOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentProjectsOperation) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentProjectsOperationOutput{})
}
