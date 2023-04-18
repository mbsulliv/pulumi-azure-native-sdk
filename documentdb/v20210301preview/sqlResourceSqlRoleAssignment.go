// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Role Assignment
type SqlResourceSqlRoleAssignment struct {
	pulumi.CustomResourceState

	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique identifier for the associated AAD principal in the AAD graph to which access is being granted through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The unique identifier for the associated Role Definition.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The data plane resource path for which access is being granted through this Role Assignment.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlResourceSqlRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewSqlResourceSqlRoleAssignment(ctx *pulumi.Context,
	name string, args *SqlResourceSqlRoleAssignmentArgs, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20221115:SqlResourceSqlRoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315:SqlResourceSqlRoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlResourceSqlRoleAssignment
	err := ctx.RegisterResource("azure-native:documentdb/v20210301preview:SqlResourceSqlRoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlResourceSqlRoleAssignment gets an existing SqlResourceSqlRoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlResourceSqlRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlResourceSqlRoleAssignmentState, opts ...pulumi.ResourceOption) (*SqlResourceSqlRoleAssignment, error) {
	var resource SqlResourceSqlRoleAssignment
	err := ctx.ReadResource("azure-native:documentdb/v20210301preview:SqlResourceSqlRoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlResourceSqlRoleAssignment resources.
type sqlResourceSqlRoleAssignmentState struct {
}

type SqlResourceSqlRoleAssignmentState struct {
}

func (SqlResourceSqlRoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleAssignmentState)(nil)).Elem()
}

type sqlResourceSqlRoleAssignmentArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The unique identifier for the associated AAD principal in the AAD graph to which access is being granted through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId *string `pulumi:"principalId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The GUID for the Role Assignment.
	RoleAssignmentId *string `pulumi:"roleAssignmentId"`
	// The unique identifier for the associated Role Definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The data plane resource path for which access is being granted through this Role Assignment.
	Scope *string `pulumi:"scope"`
}

// The set of arguments for constructing a SqlResourceSqlRoleAssignment resource.
type SqlResourceSqlRoleAssignmentArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// The unique identifier for the associated AAD principal in the AAD graph to which access is being granted through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The GUID for the Role Assignment.
	RoleAssignmentId pulumi.StringPtrInput
	// The unique identifier for the associated Role Definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The data plane resource path for which access is being granted through this Role Assignment.
	Scope pulumi.StringPtrInput
}

func (SqlResourceSqlRoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlResourceSqlRoleAssignmentArgs)(nil)).Elem()
}

type SqlResourceSqlRoleAssignmentInput interface {
	pulumi.Input

	ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput
	ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput
}

func (*SqlResourceSqlRoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlRoleAssignment)(nil)).Elem()
}

func (i *SqlResourceSqlRoleAssignment) ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput {
	return i.ToSqlResourceSqlRoleAssignmentOutputWithContext(context.Background())
}

func (i *SqlResourceSqlRoleAssignment) ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlResourceSqlRoleAssignmentOutput)
}

type SqlResourceSqlRoleAssignmentOutput struct{ *pulumi.OutputState }

func (SqlResourceSqlRoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlResourceSqlRoleAssignment)(nil)).Elem()
}

func (o SqlResourceSqlRoleAssignmentOutput) ToSqlResourceSqlRoleAssignmentOutput() SqlResourceSqlRoleAssignmentOutput {
	return o
}

func (o SqlResourceSqlRoleAssignmentOutput) ToSqlResourceSqlRoleAssignmentOutputWithContext(ctx context.Context) SqlResourceSqlRoleAssignmentOutput {
	return o
}

// The name of the database account.
func (o SqlResourceSqlRoleAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlRoleAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique identifier for the associated AAD principal in the AAD graph to which access is being granted through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
func (o SqlResourceSqlRoleAssignmentOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlRoleAssignment) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The unique identifier for the associated Role Definition.
func (o SqlResourceSqlRoleAssignmentOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlRoleAssignment) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The data plane resource path for which access is being granted through this Role Assignment.
func (o SqlResourceSqlRoleAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlResourceSqlRoleAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The type of Azure resource.
func (o SqlResourceSqlRoleAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlResourceSqlRoleAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlResourceSqlRoleAssignmentOutput{})
}
