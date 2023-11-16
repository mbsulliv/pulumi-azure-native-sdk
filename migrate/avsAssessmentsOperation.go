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

// AVS assessment resource.
// Azure REST API version: 2023-03-15.
type AvsAssessmentsOperation struct {
	pulumi.CustomResourceState

	// Gets the assessment error summary.
	//             This is the number of machines
	// affected by each type of error in this assessment.
	AssessmentErrorSummary pulumi.IntMapOutput `pulumi:"assessmentErrorSummary"`
	// Assessment type of the assessment.
	AssessmentType pulumi.StringOutput `pulumi:"assessmentType"`
	// Azure Location or Azure region where to which the machines will be migrated.
	AzureLocation pulumi.StringPtrOutput `pulumi:"azureLocation"`
	// Azure Offer code according to which cost estimation is done.
	AzureOfferCode pulumi.StringPtrOutput `pulumi:"azureOfferCode"`
	// Confidence Rating in Percentage.
	ConfidenceRatingInPercentage pulumi.Float64Output `pulumi:"confidenceRatingInPercentage"`
	// Predicted CPU utilization.
	CpuUtilization pulumi.Float64Output `pulumi:"cpuUtilization"`
	// Date and Time when assessment was created.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// Currency in which prices should be reported.
	Currency pulumi.StringPtrOutput `pulumi:"currency"`
	// De-duplication compression.
	DedupeCompression pulumi.Float64PtrOutput `pulumi:"dedupeCompression"`
	// Custom discount percentage.
	DiscountPercentage pulumi.Float64PtrOutput `pulumi:"discountPercentage"`
	// Failures to tolerate and RAID level in a common property.
	FailuresToTolerateAndRaidLevel pulumi.StringPtrOutput `pulumi:"failuresToTolerateAndRaidLevel"`
	// Gets the group type for the assessment.
	GroupType pulumi.StringOutput `pulumi:"groupType"`
	// Is Stretch Cluster Enabled.
	IsStretchClusterEnabled pulumi.BoolPtrOutput `pulumi:"isStretchClusterEnabled"`
	// Limiting factor.
	LimitingFactor pulumi.StringOutput `pulumi:"limitingFactor"`
	// Memory overcommit.
	MemOvercommit pulumi.Float64PtrOutput `pulumi:"memOvercommit"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// AVS node type.
	NodeType pulumi.StringPtrOutput `pulumi:"nodeType"`
	// Number of machines part of the assessment.
	NumberOfMachines pulumi.IntOutput `pulumi:"numberOfMachines"`
	// Recommended number of nodes.
	NumberOfNodes pulumi.IntOutput `pulumi:"numberOfNodes"`
	// Percentile of the utilization data values to be considered while assessing
	// machines.
	Percentile pulumi.StringPtrOutput `pulumi:"percentile"`
	// Gets or sets the end time to consider performance data for assessment.
	PerfDataEndTime pulumi.StringPtrOutput `pulumi:"perfDataEndTime"`
	// Gets or sets the start time to consider performance data for assessment.
	PerfDataStartTime pulumi.StringPtrOutput `pulumi:"perfDataStartTime"`
	// Time when the Azure Prices were queried. Date-Time represented in ISO-8601
	// format.
	PricesTimestamp pulumi.StringOutput `pulumi:"pricesTimestamp"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Predicted RAM utilization.
	RamUtilization pulumi.Float64Output `pulumi:"ramUtilization"`
	// Reserved instance.
	ReservedInstance pulumi.StringPtrOutput `pulumi:"reservedInstance"`
	// Percentage of buffer that user wants on performance metrics when recommending
	// Azure sizes.
	ScalingFactor pulumi.Float64PtrOutput `pulumi:"scalingFactor"`
	// Schema version.
	SchemaVersion pulumi.StringOutput `pulumi:"schemaVersion"`
	// Assessment sizing criterion.
	SizingCriterion pulumi.StringPtrOutput `pulumi:"sizingCriterion"`
	// User configurable setting to display the Stage of Assessment.
	Stage pulumi.StringOutput `pulumi:"stage"`
	// Whether assessment is in valid state and all machines have been assessed.
	Status pulumi.StringOutput `pulumi:"status"`
	// Predicted storage utilization.
	StorageUtilization pulumi.Float64Output `pulumi:"storageUtilization"`
	// Gets or sets the Assessment cloud suitability.
	Suitability pulumi.StringOutput `pulumi:"suitability"`
	// Gets or sets the Assessment suitability explanation.
	SuitabilityExplanation pulumi.StringOutput `pulumi:"suitabilityExplanation"`
	// Cloud suitability summary for all the machines in the assessment.
	SuitabilitySummary pulumi.IntMapOutput `pulumi:"suitabilitySummary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Time Range for which the historic utilization data should be considered for
	// assessment.
	TimeRange pulumi.StringPtrOutput `pulumi:"timeRange"`
	// Predicted total CPU cores used.
	TotalCpuCores pulumi.Float64Output `pulumi:"totalCpuCores"`
	// Total monthly cost.
	TotalMonthlyCost pulumi.Float64Output `pulumi:"totalMonthlyCost"`
	// Predicted total RAM used in GB.
	TotalRamInGB pulumi.Float64Output `pulumi:"totalRamInGB"`
	// Predicted total Storage used in GB.
	TotalStorageInGB pulumi.Float64Output `pulumi:"totalStorageInGB"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Date and Time when assessment was last updated.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
	// VCPU over subscription.
	VcpuOversubscription pulumi.Float64PtrOutput `pulumi:"vcpuOversubscription"`
}

// NewAvsAssessmentsOperation registers a new resource with the given unique name, arguments, and options.
func NewAvsAssessmentsOperation(ctx *pulumi.Context,
	name string, args *AvsAssessmentsOperationArgs, opts ...pulumi.ResourceOption) (*AvsAssessmentsOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20230315:AvsAssessmentsOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AvsAssessmentsOperation
	err := ctx.RegisterResource("azure-native:migrate:AvsAssessmentsOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvsAssessmentsOperation gets an existing AvsAssessmentsOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvsAssessmentsOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvsAssessmentsOperationState, opts ...pulumi.ResourceOption) (*AvsAssessmentsOperation, error) {
	var resource AvsAssessmentsOperation
	err := ctx.ReadResource("azure-native:migrate:AvsAssessmentsOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvsAssessmentsOperation resources.
type avsAssessmentsOperationState struct {
}

type AvsAssessmentsOperationState struct {
}

func (AvsAssessmentsOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*avsAssessmentsOperationState)(nil)).Elem()
}

type avsAssessmentsOperationArgs struct {
	// AVS Assessment ARM name
	AssessmentName *string `pulumi:"assessmentName"`
	// Azure Location or Azure region where to which the machines will be migrated.
	AzureLocation *string `pulumi:"azureLocation"`
	// Azure Offer code according to which cost estimation is done.
	AzureOfferCode *string `pulumi:"azureOfferCode"`
	// Currency in which prices should be reported.
	Currency *string `pulumi:"currency"`
	// De-duplication compression.
	DedupeCompression *float64 `pulumi:"dedupeCompression"`
	// Custom discount percentage.
	DiscountPercentage *float64 `pulumi:"discountPercentage"`
	// Failures to tolerate and RAID level in a common property.
	FailuresToTolerateAndRaidLevel *string `pulumi:"failuresToTolerateAndRaidLevel"`
	// Group ARM name
	GroupName string `pulumi:"groupName"`
	// Is Stretch Cluster Enabled.
	IsStretchClusterEnabled *bool `pulumi:"isStretchClusterEnabled"`
	// Memory overcommit.
	MemOvercommit *float64 `pulumi:"memOvercommit"`
	// AVS node type.
	NodeType *string `pulumi:"nodeType"`
	// Percentile of the utilization data values to be considered while assessing
	// machines.
	Percentile *string `pulumi:"percentile"`
	// Gets or sets the end time to consider performance data for assessment.
	PerfDataEndTime *string `pulumi:"perfDataEndTime"`
	// Gets or sets the start time to consider performance data for assessment.
	PerfDataStartTime *string `pulumi:"perfDataStartTime"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Reserved instance.
	ReservedInstance *string `pulumi:"reservedInstance"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Percentage of buffer that user wants on performance metrics when recommending
	// Azure sizes.
	ScalingFactor *float64 `pulumi:"scalingFactor"`
	// Assessment sizing criterion.
	SizingCriterion *string `pulumi:"sizingCriterion"`
	// Time Range for which the historic utilization data should be considered for
	// assessment.
	TimeRange *string `pulumi:"timeRange"`
	// VCPU over subscription.
	VcpuOversubscription *float64 `pulumi:"vcpuOversubscription"`
}

// The set of arguments for constructing a AvsAssessmentsOperation resource.
type AvsAssessmentsOperationArgs struct {
	// AVS Assessment ARM name
	AssessmentName pulumi.StringPtrInput
	// Azure Location or Azure region where to which the machines will be migrated.
	AzureLocation pulumi.StringPtrInput
	// Azure Offer code according to which cost estimation is done.
	AzureOfferCode pulumi.StringPtrInput
	// Currency in which prices should be reported.
	Currency pulumi.StringPtrInput
	// De-duplication compression.
	DedupeCompression pulumi.Float64PtrInput
	// Custom discount percentage.
	DiscountPercentage pulumi.Float64PtrInput
	// Failures to tolerate and RAID level in a common property.
	FailuresToTolerateAndRaidLevel pulumi.StringPtrInput
	// Group ARM name
	GroupName pulumi.StringInput
	// Is Stretch Cluster Enabled.
	IsStretchClusterEnabled pulumi.BoolPtrInput
	// Memory overcommit.
	MemOvercommit pulumi.Float64PtrInput
	// AVS node type.
	NodeType pulumi.StringPtrInput
	// Percentile of the utilization data values to be considered while assessing
	// machines.
	Percentile pulumi.StringPtrInput
	// Gets or sets the end time to consider performance data for assessment.
	PerfDataEndTime pulumi.StringPtrInput
	// Gets or sets the start time to consider performance data for assessment.
	PerfDataStartTime pulumi.StringPtrInput
	// Assessment Project Name
	ProjectName pulumi.StringInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// Reserved instance.
	ReservedInstance pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Percentage of buffer that user wants on performance metrics when recommending
	// Azure sizes.
	ScalingFactor pulumi.Float64PtrInput
	// Assessment sizing criterion.
	SizingCriterion pulumi.StringPtrInput
	// Time Range for which the historic utilization data should be considered for
	// assessment.
	TimeRange pulumi.StringPtrInput
	// VCPU over subscription.
	VcpuOversubscription pulumi.Float64PtrInput
}

func (AvsAssessmentsOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*avsAssessmentsOperationArgs)(nil)).Elem()
}

type AvsAssessmentsOperationInput interface {
	pulumi.Input

	ToAvsAssessmentsOperationOutput() AvsAssessmentsOperationOutput
	ToAvsAssessmentsOperationOutputWithContext(ctx context.Context) AvsAssessmentsOperationOutput
}

func (*AvsAssessmentsOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**AvsAssessmentsOperation)(nil)).Elem()
}

func (i *AvsAssessmentsOperation) ToAvsAssessmentsOperationOutput() AvsAssessmentsOperationOutput {
	return i.ToAvsAssessmentsOperationOutputWithContext(context.Background())
}

func (i *AvsAssessmentsOperation) ToAvsAssessmentsOperationOutputWithContext(ctx context.Context) AvsAssessmentsOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvsAssessmentsOperationOutput)
}

type AvsAssessmentsOperationOutput struct{ *pulumi.OutputState }

func (AvsAssessmentsOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvsAssessmentsOperation)(nil)).Elem()
}

func (o AvsAssessmentsOperationOutput) ToAvsAssessmentsOperationOutput() AvsAssessmentsOperationOutput {
	return o
}

func (o AvsAssessmentsOperationOutput) ToAvsAssessmentsOperationOutputWithContext(ctx context.Context) AvsAssessmentsOperationOutput {
	return o
}

// Gets the assessment error summary.
//
//	This is the number of machines
//
// affected by each type of error in this assessment.
func (o AvsAssessmentsOperationOutput) AssessmentErrorSummary() pulumi.IntMapOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.IntMapOutput { return v.AssessmentErrorSummary }).(pulumi.IntMapOutput)
}

// Assessment type of the assessment.
func (o AvsAssessmentsOperationOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.AssessmentType }).(pulumi.StringOutput)
}

// Azure Location or Azure region where to which the machines will be migrated.
func (o AvsAssessmentsOperationOutput) AzureLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.AzureLocation }).(pulumi.StringPtrOutput)
}

// Azure Offer code according to which cost estimation is done.
func (o AvsAssessmentsOperationOutput) AzureOfferCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.AzureOfferCode }).(pulumi.StringPtrOutput)
}

// Confidence Rating in Percentage.
func (o AvsAssessmentsOperationOutput) ConfidenceRatingInPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.ConfidenceRatingInPercentage }).(pulumi.Float64Output)
}

// Predicted CPU utilization.
func (o AvsAssessmentsOperationOutput) CpuUtilization() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.CpuUtilization }).(pulumi.Float64Output)
}

// Date and Time when assessment was created.
func (o AvsAssessmentsOperationOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Currency in which prices should be reported.
func (o AvsAssessmentsOperationOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.Currency }).(pulumi.StringPtrOutput)
}

// De-duplication compression.
func (o AvsAssessmentsOperationOutput) DedupeCompression() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64PtrOutput { return v.DedupeCompression }).(pulumi.Float64PtrOutput)
}

// Custom discount percentage.
func (o AvsAssessmentsOperationOutput) DiscountPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64PtrOutput { return v.DiscountPercentage }).(pulumi.Float64PtrOutput)
}

// Failures to tolerate and RAID level in a common property.
func (o AvsAssessmentsOperationOutput) FailuresToTolerateAndRaidLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.FailuresToTolerateAndRaidLevel }).(pulumi.StringPtrOutput)
}

// Gets the group type for the assessment.
func (o AvsAssessmentsOperationOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.GroupType }).(pulumi.StringOutput)
}

// Is Stretch Cluster Enabled.
func (o AvsAssessmentsOperationOutput) IsStretchClusterEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.BoolPtrOutput { return v.IsStretchClusterEnabled }).(pulumi.BoolPtrOutput)
}

// Limiting factor.
func (o AvsAssessmentsOperationOutput) LimitingFactor() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.LimitingFactor }).(pulumi.StringOutput)
}

// Memory overcommit.
func (o AvsAssessmentsOperationOutput) MemOvercommit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64PtrOutput { return v.MemOvercommit }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o AvsAssessmentsOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AVS node type.
func (o AvsAssessmentsOperationOutput) NodeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.NodeType }).(pulumi.StringPtrOutput)
}

// Number of machines part of the assessment.
func (o AvsAssessmentsOperationOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.IntOutput { return v.NumberOfMachines }).(pulumi.IntOutput)
}

// Recommended number of nodes.
func (o AvsAssessmentsOperationOutput) NumberOfNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.IntOutput { return v.NumberOfNodes }).(pulumi.IntOutput)
}

// Percentile of the utilization data values to be considered while assessing
// machines.
func (o AvsAssessmentsOperationOutput) Percentile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.Percentile }).(pulumi.StringPtrOutput)
}

// Gets or sets the end time to consider performance data for assessment.
func (o AvsAssessmentsOperationOutput) PerfDataEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.PerfDataEndTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the start time to consider performance data for assessment.
func (o AvsAssessmentsOperationOutput) PerfDataStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.PerfDataStartTime }).(pulumi.StringPtrOutput)
}

// Time when the Azure Prices were queried. Date-Time represented in ISO-8601
// format.
func (o AvsAssessmentsOperationOutput) PricesTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.PricesTimestamp }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o AvsAssessmentsOperationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Predicted RAM utilization.
func (o AvsAssessmentsOperationOutput) RamUtilization() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.RamUtilization }).(pulumi.Float64Output)
}

// Reserved instance.
func (o AvsAssessmentsOperationOutput) ReservedInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.ReservedInstance }).(pulumi.StringPtrOutput)
}

// Percentage of buffer that user wants on performance metrics when recommending
// Azure sizes.
func (o AvsAssessmentsOperationOutput) ScalingFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64PtrOutput { return v.ScalingFactor }).(pulumi.Float64PtrOutput)
}

// Schema version.
func (o AvsAssessmentsOperationOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.SchemaVersion }).(pulumi.StringOutput)
}

// Assessment sizing criterion.
func (o AvsAssessmentsOperationOutput) SizingCriterion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.SizingCriterion }).(pulumi.StringPtrOutput)
}

// User configurable setting to display the Stage of Assessment.
func (o AvsAssessmentsOperationOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.Stage }).(pulumi.StringOutput)
}

// Whether assessment is in valid state and all machines have been assessed.
func (o AvsAssessmentsOperationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Predicted storage utilization.
func (o AvsAssessmentsOperationOutput) StorageUtilization() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.StorageUtilization }).(pulumi.Float64Output)
}

// Gets or sets the Assessment cloud suitability.
func (o AvsAssessmentsOperationOutput) Suitability() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.Suitability }).(pulumi.StringOutput)
}

// Gets or sets the Assessment suitability explanation.
func (o AvsAssessmentsOperationOutput) SuitabilityExplanation() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.SuitabilityExplanation }).(pulumi.StringOutput)
}

// Cloud suitability summary for all the machines in the assessment.
func (o AvsAssessmentsOperationOutput) SuitabilitySummary() pulumi.IntMapOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.IntMapOutput { return v.SuitabilitySummary }).(pulumi.IntMapOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AvsAssessmentsOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Time Range for which the historic utilization data should be considered for
// assessment.
func (o AvsAssessmentsOperationOutput) TimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringPtrOutput { return v.TimeRange }).(pulumi.StringPtrOutput)
}

// Predicted total CPU cores used.
func (o AvsAssessmentsOperationOutput) TotalCpuCores() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.TotalCpuCores }).(pulumi.Float64Output)
}

// Total monthly cost.
func (o AvsAssessmentsOperationOutput) TotalMonthlyCost() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.TotalMonthlyCost }).(pulumi.Float64Output)
}

// Predicted total RAM used in GB.
func (o AvsAssessmentsOperationOutput) TotalRamInGB() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.TotalRamInGB }).(pulumi.Float64Output)
}

// Predicted total Storage used in GB.
func (o AvsAssessmentsOperationOutput) TotalStorageInGB() pulumi.Float64Output {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64Output { return v.TotalStorageInGB }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AvsAssessmentsOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Date and Time when assessment was last updated.
func (o AvsAssessmentsOperationOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

// VCPU over subscription.
func (o AvsAssessmentsOperationOutput) VcpuOversubscription() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AvsAssessmentsOperation) pulumi.Float64PtrOutput { return v.VcpuOversubscription }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvsAssessmentsOperationOutput{})
}
