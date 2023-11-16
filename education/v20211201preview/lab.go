// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lab details.
type Lab struct {
	pulumi.CustomResourceState

	// Default monetary cap for each student in this lab
	BudgetPerStudent AmountResponseOutput `pulumi:"budgetPerStudent"`
	// The type of currency being used for the value.
	Currency pulumi.StringPtrOutput `pulumi:"currency"`
	// Detail description of this lab
	Description pulumi.StringOutput `pulumi:"description"`
	// Lab Display Name
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Lab creation date
	EffectiveDate pulumi.StringOutput `pulumi:"effectiveDate"`
	// Default expiration date for each student in this lab
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// invitation code for redeemable lab
	InvitationCode pulumi.StringOutput `pulumi:"invitationCode"`
	// the total number of students that can be accepted to the lab.
	MaxStudentCount pulumi.Float64Output `pulumi:"maxStudentCount"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of this lab
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Amount value.
	Value pulumi.Float64PtrOutput `pulumi:"value"`
}

// NewLab registers a new resource with the given unique name, arguments, and options.
func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.BillingProfileName == nil {
		return nil, errors.New("invalid value for required argument 'BillingProfileName'")
	}
	if args.BudgetPerStudent == nil {
		return nil, errors.New("invalid value for required argument 'BudgetPerStudent'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ExpirationDate == nil {
		return nil, errors.New("invalid value for required argument 'ExpirationDate'")
	}
	if args.InvoiceSectionName == nil {
		return nil, errors.New("invalid value for required argument 'InvoiceSectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:education:Lab"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Lab
	err := ctx.RegisterResource("azure-native:education/v20211201preview:Lab", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:education/v20211201preview:Lab", name, id, state, &resource, opts...)
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
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a billing profile.
	BillingProfileName string `pulumi:"billingProfileName"`
	// Default monetary cap for each student in this lab
	BudgetPerStudent Amount `pulumi:"budgetPerStudent"`
	// The type of currency being used for the value.
	Currency *string `pulumi:"currency"`
	// Detail description of this lab
	Description string `pulumi:"description"`
	// Lab Display Name
	DisplayName string `pulumi:"displayName"`
	// Default expiration date for each student in this lab
	ExpirationDate string `pulumi:"expirationDate"`
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName string `pulumi:"invoiceSectionName"`
	// Amount value.
	Value *float64 `pulumi:"value"`
}

// The set of arguments for constructing a Lab resource.
type LabArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput
	// The ID that uniquely identifies a billing profile.
	BillingProfileName pulumi.StringInput
	// Default monetary cap for each student in this lab
	BudgetPerStudent AmountInput
	// The type of currency being used for the value.
	Currency pulumi.StringPtrInput
	// Detail description of this lab
	Description pulumi.StringInput
	// Lab Display Name
	DisplayName pulumi.StringInput
	// Default expiration date for each student in this lab
	ExpirationDate pulumi.StringInput
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName pulumi.StringInput
	// Amount value.
	Value pulumi.Float64PtrInput
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

// Default monetary cap for each student in this lab
func (o LabOutput) BudgetPerStudent() AmountResponseOutput {
	return o.ApplyT(func(v *Lab) AmountResponseOutput { return v.BudgetPerStudent }).(AmountResponseOutput)
}

// The type of currency being used for the value.
func (o LabOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Currency }).(pulumi.StringPtrOutput)
}

// Detail description of this lab
func (o LabOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Lab Display Name
func (o LabOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Lab creation date
func (o LabOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.EffectiveDate }).(pulumi.StringOutput)
}

// Default expiration date for each student in this lab
func (o LabOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// invitation code for redeemable lab
func (o LabOutput) InvitationCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.InvitationCode }).(pulumi.StringOutput)
}

// the total number of students that can be accepted to the lab.
func (o LabOutput) MaxStudentCount() pulumi.Float64Output {
	return o.ApplyT(func(v *Lab) pulumi.Float64Output { return v.MaxStudentCount }).(pulumi.Float64Output)
}

// The name of the resource
func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of this lab
func (o LabOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LabOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Lab) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Amount value.
func (o LabOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.Float64PtrOutput { return v.Value }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
