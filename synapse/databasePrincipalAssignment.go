// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a database principal assignment.
// Azure REST API version: 2021-04-01-preview. Prior API version in Azure Native 1.x: 2021-04-01-preview.
type DatabasePrincipalAssignment struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The principal name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// Principal type.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Database principal role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id of the principal
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabasePrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, args *DatabasePrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:DatabasePrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:DatabasePrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabasePrincipalAssignment
	err := ctx.RegisterResource("azure-native:synapse:DatabasePrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasePrincipalAssignment gets an existing DatabasePrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasePrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePrincipalAssignmentState, opts ...pulumi.ResourceOption) (*DatabasePrincipalAssignment, error) {
	var resource DatabasePrincipalAssignment
	err := ctx.ReadResource("azure-native:synapse:DatabasePrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasePrincipalAssignment resources.
type databasePrincipalAssignmentState struct {
}

type DatabasePrincipalAssignmentState struct {
}

func (DatabasePrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalAssignmentState)(nil)).Elem()
}

type databasePrincipalAssignmentArgs struct {
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Database principal role.
	Role string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DatabasePrincipalAssignment resource.
type DatabasePrincipalAssignmentArgs struct {
	// The name of the database in the Kusto pool.
	DatabaseName pulumi.StringInput
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringPtrInput
	// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringInput
	// Principal type.
	PrincipalType pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Database principal role.
	Role pulumi.StringInput
	// The tenant id of the principal
	TenantId pulumi.StringPtrInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (DatabasePrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePrincipalAssignmentArgs)(nil)).Elem()
}

type DatabasePrincipalAssignmentInput interface {
	pulumi.Input

	ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput
	ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput
}

func (*DatabasePrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipalAssignment)(nil)).Elem()
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return i.ToDatabasePrincipalAssignmentOutputWithContext(context.Background())
}

func (i *DatabasePrincipalAssignment) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalAssignmentOutput)
}

type DatabasePrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipalAssignment)(nil)).Elem()
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutput() DatabasePrincipalAssignmentOutput {
	return o
}

func (o DatabasePrincipalAssignmentOutput) ToDatabasePrincipalAssignmentOutputWithContext(ctx context.Context) DatabasePrincipalAssignmentOutput {
	return o
}

// The name of the resource
func (o DatabasePrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
func (o DatabasePrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o DatabasePrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o DatabasePrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o DatabasePrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Database principal role.
func (o DatabasePrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DatabasePrincipalAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id of the principal
func (o DatabasePrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o DatabasePrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DatabasePrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabasePrincipalAssignmentOutput{})
}
