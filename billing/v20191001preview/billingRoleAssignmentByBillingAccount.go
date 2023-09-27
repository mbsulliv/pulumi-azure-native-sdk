// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The role assignment
type BillingRoleAssignmentByBillingAccount struct {
	pulumi.CustomResourceState

	// The principal Id of the user who created the role assignment.
	CreatedByPrincipalId pulumi.StringOutput `pulumi:"createdByPrincipalId"`
	// The tenant Id of the user who created the role assignment.
	CreatedByPrincipalTenantId pulumi.StringOutput `pulumi:"createdByPrincipalTenantId"`
	// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
	CreatedByUserEmailAddress pulumi.StringOutput `pulumi:"createdByUserEmailAddress"`
	// The date the role assignment was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrOutput `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The scope at which the role was assigned.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrOutput `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrOutput `pulumi:"userEmailAddress"`
}

// NewBillingRoleAssignmentByBillingAccount registers a new resource with the given unique name, arguments, and options.
func NewBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, args *BillingRoleAssignmentByBillingAccountArgs, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:billing:BillingRoleAssignmentByBillingAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.RegisterResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingRoleAssignmentByBillingAccount gets an existing BillingRoleAssignmentByBillingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingRoleAssignmentByBillingAccountState, opts ...pulumi.ResourceOption) (*BillingRoleAssignmentByBillingAccount, error) {
	var resource BillingRoleAssignmentByBillingAccount
	err := ctx.ReadResource("azure-native:billing/v20191001preview:BillingRoleAssignmentByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingRoleAssignmentByBillingAccount resources.
type billingRoleAssignmentByBillingAccountState struct {
}

type BillingRoleAssignmentByBillingAccountState struct {
}

func (BillingRoleAssignmentByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountState)(nil)).Elem()
}

type billingRoleAssignmentByBillingAccountArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName *string `pulumi:"billingRoleAssignmentName"`
	// The principal id of the user to whom the role was assigned.
	PrincipalId *string `pulumi:"principalId"`
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId *string `pulumi:"principalTenantId"`
	// The ID of the role definition.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType *string `pulumi:"userAuthenticationType"`
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress *string `pulumi:"userEmailAddress"`
}

// The set of arguments for constructing a BillingRoleAssignmentByBillingAccount resource.
type BillingRoleAssignmentByBillingAccountArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput
	// The ID that uniquely identifies a role assignment.
	BillingRoleAssignmentName pulumi.StringPtrInput
	// The principal id of the user to whom the role was assigned.
	PrincipalId pulumi.StringPtrInput
	// The principal tenant id of the user to whom the role was assigned.
	PrincipalTenantId pulumi.StringPtrInput
	// The ID of the role definition.
	RoleDefinitionId pulumi.StringPtrInput
	// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserAuthenticationType pulumi.StringPtrInput
	// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
	UserEmailAddress pulumi.StringPtrInput
}

func (BillingRoleAssignmentByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingRoleAssignmentByBillingAccountArgs)(nil)).Elem()
}

type BillingRoleAssignmentByBillingAccountInput interface {
	pulumi.Input

	ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput
	ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput
}

func (*BillingRoleAssignmentByBillingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByBillingAccount)(nil)).Elem()
}

func (i *BillingRoleAssignmentByBillingAccount) ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput {
	return i.ToBillingRoleAssignmentByBillingAccountOutputWithContext(context.Background())
}

func (i *BillingRoleAssignmentByBillingAccount) ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingRoleAssignmentByBillingAccountOutput)
}

func (i *BillingRoleAssignmentByBillingAccount) ToOutput(ctx context.Context) pulumix.Output[*BillingRoleAssignmentByBillingAccount] {
	return pulumix.Output[*BillingRoleAssignmentByBillingAccount]{
		OutputState: i.ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx).OutputState,
	}
}

type BillingRoleAssignmentByBillingAccountOutput struct{ *pulumi.OutputState }

func (BillingRoleAssignmentByBillingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingRoleAssignmentByBillingAccount)(nil)).Elem()
}

func (o BillingRoleAssignmentByBillingAccountOutput) ToBillingRoleAssignmentByBillingAccountOutput() BillingRoleAssignmentByBillingAccountOutput {
	return o
}

func (o BillingRoleAssignmentByBillingAccountOutput) ToBillingRoleAssignmentByBillingAccountOutputWithContext(ctx context.Context) BillingRoleAssignmentByBillingAccountOutput {
	return o
}

func (o BillingRoleAssignmentByBillingAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*BillingRoleAssignmentByBillingAccount] {
	return pulumix.Output[*BillingRoleAssignmentByBillingAccount]{
		OutputState: o.OutputState,
	}
}

// The principal Id of the user who created the role assignment.
func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

// The tenant Id of the user who created the role assignment.
func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput {
		return v.CreatedByPrincipalTenantId
	}).(pulumi.StringOutput)
}

// The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o BillingRoleAssignmentByBillingAccountOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

// The date the role assignment was created.
func (o BillingRoleAssignmentByBillingAccountOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// Resource name.
func (o BillingRoleAssignmentByBillingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principal id of the user to whom the role was assigned.
func (o BillingRoleAssignmentByBillingAccountOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The principal tenant id of the user to whom the role was assigned.
func (o BillingRoleAssignmentByBillingAccountOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

// The ID of the role definition.
func (o BillingRoleAssignmentByBillingAccountOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The scope at which the role was assigned.
func (o BillingRoleAssignmentByBillingAccountOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Resource type.
func (o BillingRoleAssignmentByBillingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o BillingRoleAssignmentByBillingAccountOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

// The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
func (o BillingRoleAssignmentByBillingAccountOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingRoleAssignmentByBillingAccount) pulumi.StringPtrOutput { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BillingRoleAssignmentByBillingAccountOutput{})
}
