// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An export resource.
type Export struct {
	pulumi.CustomResourceState

	// Has the definition for the export.
	Definition ExportDefinitionResponseOutput `pulumi:"definition"`
	// Has delivery information for the export.
	DeliveryInfo ExportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// The format of the export being delivered. Currently only 'Csv' is supported.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The managed identity associated with Export
	Identity SystemAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the Export's managed identity. Only required when utilizing managed identity.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// If the export has an active schedule, provides an estimate of the next run time.
	NextRunTimeEstimate pulumi.StringOutput `pulumi:"nextRunTimeEstimate"`
	// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
	PartitionData pulumi.BoolPtrOutput `pulumi:"partitionData"`
	// If requested, has the most recent run history for the export.
	RunHistory ExportExecutionListResultResponsePtrOutput `pulumi:"runHistory"`
	// Has schedule information for the export.
	Schedule ExportScheduleResponsePtrOutput `pulumi:"schedule"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExport registers a new resource with the given unique name, arguments, and options.
func NewExport(ctx *pulumi.Context,
	name string, args *ExportArgs, opts ...pulumi.ResourceOption) (*Export, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190101:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20190901:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191001:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20191101:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20200601:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20201201preview:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20210101:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20211001:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230301:Export"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230401preview:Export"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Export
	err := ctx.RegisterResource("azure-native:costmanagement/v20230801:Export", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExport gets an existing Export resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportState, opts ...pulumi.ResourceOption) (*Export, error) {
	var resource Export
	err := ctx.ReadResource("azure-native:costmanagement/v20230801:Export", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Export resources.
type exportState struct {
}

type ExportState struct {
}

func (ExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportState)(nil)).Elem()
}

type exportArgs struct {
	// Has the definition for the export.
	Definition ExportDefinition `pulumi:"definition"`
	// Has delivery information for the export.
	DeliveryInfo ExportDeliveryInfo `pulumi:"deliveryInfo"`
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag *string `pulumi:"eTag"`
	// Export Name.
	ExportName *string `pulumi:"exportName"`
	// The format of the export being delivered. Currently only 'Csv' is supported.
	Format *string `pulumi:"format"`
	// The managed identity associated with Export
	Identity *SystemAssignedServiceIdentity `pulumi:"identity"`
	// The location of the Export's managed identity. Only required when utilizing managed identity.
	Location *string `pulumi:"location"`
	// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
	PartitionData *bool `pulumi:"partitionData"`
	// Has schedule information for the export.
	Schedule *ExportSchedule `pulumi:"schedule"`
	// The scope associated with export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a Export resource.
type ExportArgs struct {
	// Has the definition for the export.
	Definition ExportDefinitionInput
	// Has delivery information for the export.
	DeliveryInfo ExportDeliveryInfoInput
	// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
	ETag pulumi.StringPtrInput
	// Export Name.
	ExportName pulumi.StringPtrInput
	// The format of the export being delivered. Currently only 'Csv' is supported.
	Format pulumi.StringPtrInput
	// The managed identity associated with Export
	Identity SystemAssignedServiceIdentityPtrInput
	// The location of the Export's managed identity. Only required when utilizing managed identity.
	Location pulumi.StringPtrInput
	// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
	PartitionData pulumi.BoolPtrInput
	// Has schedule information for the export.
	Schedule ExportSchedulePtrInput
	// The scope associated with export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
	Scope pulumi.StringInput
}

func (ExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportArgs)(nil)).Elem()
}

type ExportInput interface {
	pulumi.Input

	ToExportOutput() ExportOutput
	ToExportOutputWithContext(ctx context.Context) ExportOutput
}

func (*Export) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (i *Export) ToExportOutput() ExportOutput {
	return i.ToExportOutputWithContext(context.Background())
}

func (i *Export) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportOutput)
}

type ExportOutput struct{ *pulumi.OutputState }

func (ExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (o ExportOutput) ToExportOutput() ExportOutput {
	return o
}

func (o ExportOutput) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return o
}

// Has the definition for the export.
func (o ExportOutput) Definition() ExportDefinitionResponseOutput {
	return o.ApplyT(func(v *Export) ExportDefinitionResponseOutput { return v.Definition }).(ExportDefinitionResponseOutput)
}

// Has delivery information for the export.
func (o ExportOutput) DeliveryInfo() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *Export) ExportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ExportDeliveryInfoResponseOutput)
}

// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
func (o ExportOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// The format of the export being delivered. Currently only 'Csv' is supported.
func (o ExportOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The managed identity associated with Export
func (o ExportOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Export) SystemAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

// The location of the Export's managed identity. Only required when utilizing managed identity.
func (o ExportOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ExportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Export) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// If the export has an active schedule, provides an estimate of the next run time.
func (o ExportOutput) NextRunTimeEstimate() pulumi.StringOutput {
	return o.ApplyT(func(v *Export) pulumi.StringOutput { return v.NextRunTimeEstimate }).(pulumi.StringOutput)
}

// If set to true, exported data will be partitioned by size and placed in a blob directory together with a manifest file. Note: this option is currently available only for Microsoft Customer Agreement commerce scopes.
func (o ExportOutput) PartitionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.BoolPtrOutput { return v.PartitionData }).(pulumi.BoolPtrOutput)
}

// If requested, has the most recent run history for the export.
func (o ExportOutput) RunHistory() ExportExecutionListResultResponsePtrOutput {
	return o.ApplyT(func(v *Export) ExportExecutionListResultResponsePtrOutput { return v.RunHistory }).(ExportExecutionListResultResponsePtrOutput)
}

// Has schedule information for the export.
func (o ExportOutput) Schedule() ExportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *Export) ExportScheduleResponsePtrOutput { return v.Schedule }).(ExportScheduleResponsePtrOutput)
}

// Resource type.
func (o ExportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Export) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportOutput{})
}
