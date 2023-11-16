// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security assessment metadata response
// Azure REST API version: 2021-06-01. Prior API version in Azure Native 1.x: 2020-01-01.
type AssessmentMetadataInSubscription struct {
	pulumi.CustomResourceState

	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringOutput      `pulumi:"assessmentType"`
	Categories     pulumi.StringArrayOutput `pulumi:"categories"`
	// Human readable description of the assessment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrOutput `pulumi:"implementationEffort"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the partner that created the assessment
	PartnerData            SecurityAssessmentMetadataPartnerDataResponsePtrOutput `pulumi:"partnerData"`
	PlannedDeprecationDate pulumi.StringPtrOutput                                 `pulumi:"plannedDeprecationDate"`
	// Azure resource ID of the policy definition that turns this assessment calculation on
	PolicyDefinitionId pulumi.StringOutput `pulumi:"policyDefinitionId"`
	// True if this assessment is in preview release status
	Preview      pulumi.BoolPtrOutput                                                      `pulumi:"preview"`
	PublishDates SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput `pulumi:"publishDates"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrOutput `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity   pulumi.StringOutput      `pulumi:"severity"`
	Tactics    pulumi.StringArrayOutput `pulumi:"tactics"`
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	Threats    pulumi.StringArrayOutput `pulumi:"threats"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The user impact of the assessment
	UserImpact pulumi.StringPtrOutput `pulumi:"userImpact"`
}

// NewAssessmentMetadataInSubscription registers a new resource with the given unique name, arguments, and options.
func NewAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, args *AssessmentMetadataInSubscriptionArgs, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentType == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentType'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20190101preview:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AssessmentMetadataInSubscription"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210601:AssessmentMetadataInSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AssessmentMetadataInSubscription
	err := ctx.RegisterResource("azure-native:security:AssessmentMetadataInSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentMetadataInSubscription gets an existing AssessmentMetadataInSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentMetadataInSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentMetadataInSubscriptionState, opts ...pulumi.ResourceOption) (*AssessmentMetadataInSubscription, error) {
	var resource AssessmentMetadataInSubscription
	err := ctx.ReadResource("azure-native:security:AssessmentMetadataInSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentMetadataInSubscription resources.
type assessmentMetadataInSubscriptionState struct {
}

type AssessmentMetadataInSubscriptionState struct {
}

func (AssessmentMetadataInSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionState)(nil)).Elem()
}

type assessmentMetadataInSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName *string `pulumi:"assessmentMetadataName"`
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType string   `pulumi:"assessmentType"`
	Categories     []string `pulumi:"categories"`
	// Human readable description of the assessment
	Description *string `pulumi:"description"`
	// User friendly display name of the assessment
	DisplayName string `pulumi:"displayName"`
	// The implementation effort required to remediate this assessment
	ImplementationEffort *string `pulumi:"implementationEffort"`
	// Describes the partner that created the assessment
	PartnerData            *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	PlannedDeprecationDate *string                                `pulumi:"plannedDeprecationDate"`
	// True if this assessment is in preview release status
	Preview      *bool                                                     `pulumi:"preview"`
	PublishDates *SecurityAssessmentMetadataPropertiesResponsePublishDates `pulumi:"publishDates"`
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription *string `pulumi:"remediationDescription"`
	// The severity level of the assessment
	Severity   string   `pulumi:"severity"`
	Tactics    []string `pulumi:"tactics"`
	Techniques []string `pulumi:"techniques"`
	Threats    []string `pulumi:"threats"`
	// The user impact of the assessment
	UserImpact *string `pulumi:"userImpact"`
}

// The set of arguments for constructing a AssessmentMetadataInSubscription resource.
type AssessmentMetadataInSubscriptionArgs struct {
	// The Assessment Key - Unique key for the assessment type
	AssessmentMetadataName pulumi.StringPtrInput
	// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
	AssessmentType pulumi.StringInput
	Categories     pulumi.StringArrayInput
	// Human readable description of the assessment
	Description pulumi.StringPtrInput
	// User friendly display name of the assessment
	DisplayName pulumi.StringInput
	// The implementation effort required to remediate this assessment
	ImplementationEffort pulumi.StringPtrInput
	// Describes the partner that created the assessment
	PartnerData            SecurityAssessmentMetadataPartnerDataPtrInput
	PlannedDeprecationDate pulumi.StringPtrInput
	// True if this assessment is in preview release status
	Preview      pulumi.BoolPtrInput
	PublishDates SecurityAssessmentMetadataPropertiesResponsePublishDatesPtrInput
	// Human readable description of what you should do to mitigate this security issue
	RemediationDescription pulumi.StringPtrInput
	// The severity level of the assessment
	Severity   pulumi.StringInput
	Tactics    pulumi.StringArrayInput
	Techniques pulumi.StringArrayInput
	Threats    pulumi.StringArrayInput
	// The user impact of the assessment
	UserImpact pulumi.StringPtrInput
}

func (AssessmentMetadataInSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentMetadataInSubscriptionArgs)(nil)).Elem()
}

type AssessmentMetadataInSubscriptionInput interface {
	pulumi.Input

	ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput
	ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput
}

func (*AssessmentMetadataInSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentMetadataInSubscription)(nil)).Elem()
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return i.ToAssessmentMetadataInSubscriptionOutputWithContext(context.Background())
}

func (i *AssessmentMetadataInSubscription) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentMetadataInSubscriptionOutput)
}

type AssessmentMetadataInSubscriptionOutput struct{ *pulumi.OutputState }

func (AssessmentMetadataInSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentMetadataInSubscription)(nil)).Elem()
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutput() AssessmentMetadataInSubscriptionOutput {
	return o
}

func (o AssessmentMetadataInSubscriptionOutput) ToAssessmentMetadataInSubscriptionOutputWithContext(ctx context.Context) AssessmentMetadataInSubscriptionOutput {
	return o
}

// BuiltIn if the assessment based on built-in Azure Policy definition, Custom if the assessment based on custom Azure Policy definition
func (o AssessmentMetadataInSubscriptionOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

// Human readable description of the assessment
func (o AssessmentMetadataInSubscriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name of the assessment
func (o AssessmentMetadataInSubscriptionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The implementation effort required to remediate this assessment
func (o AssessmentMetadataInSubscriptionOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

// Resource name
func (o AssessmentMetadataInSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes the partner that created the assessment
func (o AssessmentMetadataInSubscriptionOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) PlannedDeprecationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.PlannedDeprecationDate }).(pulumi.StringPtrOutput)
}

// Azure resource ID of the policy definition that turns this assessment calculation on
func (o AssessmentMetadataInSubscriptionOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

// True if this assessment is in preview release status
func (o AssessmentMetadataInSubscriptionOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) PublishDates() SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput {
		return v.PublishDates
	}).(SecurityAssessmentMetadataPropertiesResponseResponsePublishDatesPtrOutput)
}

// Human readable description of what you should do to mitigate this security issue
func (o AssessmentMetadataInSubscriptionOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

// The severity level of the assessment
func (o AssessmentMetadataInSubscriptionOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o AssessmentMetadataInSubscriptionOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringArrayOutput { return v.Threats }).(pulumi.StringArrayOutput)
}

// Resource type
func (o AssessmentMetadataInSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user impact of the assessment
func (o AssessmentMetadataInSubscriptionOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentMetadataInSubscription) pulumi.StringPtrOutput { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentMetadataInSubscriptionOutput{})
}
