// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Access Review Schedule Definition.
// Azure REST API version: 2021-12-01-preview. Prior API version in Azure Native 1.x: 2021-03-01-preview.
type AccessReviewScheduleDefinitionById struct {
	pulumi.CustomResourceState

	// The role assignment state eligible/active to review
	AssignmentState pulumi.StringOutput `pulumi:"assignmentState"`
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled pulumi.BoolPtrOutput `pulumi:"autoApplyDecisionsEnabled"`
	// This is the collection of backup reviewers.
	BackupReviewers AccessReviewReviewerResponseArrayOutput `pulumi:"backupReviewers"`
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision pulumi.StringPtrOutput `pulumi:"defaultDecision"`
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled pulumi.BoolPtrOutput `pulumi:"defaultDecisionEnabled"`
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins pulumi.StringPtrOutput `pulumi:"descriptionForAdmins"`
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers pulumi.StringPtrOutput `pulumi:"descriptionForReviewers"`
	// The display name for the schedule definition.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// This is used to indicate the resource id(s) to exclude
	ExcludeResourceId pulumi.StringPtrOutput `pulumi:"excludeResourceId"`
	// This is used to indicate the role definition id(s) to exclude
	ExcludeRoleDefinitionId pulumi.StringPtrOutput `pulumi:"excludeRoleDefinitionId"`
	// Flag to indicate whether to expand nested memberships or not.
	ExpandNestedMemberships pulumi.BoolPtrOutput `pulumi:"expandNestedMemberships"`
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration pulumi.StringPtrOutput `pulumi:"inactiveDuration"`
	// Flag to indicate whether to expand nested memberships or not.
	IncludeAccessBelowResource pulumi.BoolPtrOutput `pulumi:"includeAccessBelowResource"`
	// Flag to indicate whether to expand nested memberships or not.
	IncludeInheritedAccess pulumi.BoolPtrOutput `pulumi:"includeInheritedAccess"`
	// The duration in days for an instance.
	InstanceDurationInDays pulumi.IntPtrOutput `pulumi:"instanceDurationInDays"`
	// This is the collection of instances returned when one does an expand on it.
	Instances AccessReviewInstanceResponseArrayOutput `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval pulumi.BoolPtrOutput `pulumi:"justificationRequiredOnApproval"`
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled pulumi.BoolPtrOutput `pulumi:"mailNotificationsEnabled"`
	// The access review schedule definition unique id.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrOutput `pulumi:"numberOfOccurrences"`
	// The identity id
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The identity display name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// The identity type user/servicePrincipal to review
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// Recommendations for access reviews are calculated by looking back at 30 days of data(w.r.t the start date of the review) by default. However, in some scenarios, customers want to change how far back to look at and want to configure 60 days, 90 days, etc. instead. This setting allows customers to configure this duration. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	RecommendationLookBackDuration pulumi.StringPtrOutput `pulumi:"recommendationLookBackDuration"`
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled pulumi.BoolPtrOutput `pulumi:"recommendationsEnabled"`
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled pulumi.BoolPtrOutput `pulumi:"reminderNotificationsEnabled"`
	// ResourceId in which this review is getting created
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// This is the collection of reviewers.
	Reviewers AccessReviewReviewerResponseArrayOutput `pulumi:"reviewers"`
	// This field specifies the type of reviewers for a review. Usually for a review, reviewers are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen dynamically. For example managers review or self review.
	ReviewersType pulumi.StringOutput `pulumi:"reviewersType"`
	// This is used to indicate the role being reviewed
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// This read-only field specifies the status of an accessReview.
	Status pulumi.StringOutput `pulumi:"status"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The user principal name(if valid)
	UserPrincipalName pulumi.StringOutput `pulumi:"userPrincipalName"`
}

// NewAccessReviewScheduleDefinitionById registers a new resource with the given unique name, arguments, and options.
func NewAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, args *AccessReviewScheduleDefinitionByIdArgs, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	if args == nil {
		args = &AccessReviewScheduleDefinitionByIdArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20180501preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210301preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210701preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211116preview:AccessReviewScheduleDefinitionById"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20211201preview:AccessReviewScheduleDefinitionById"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AccessReviewScheduleDefinitionById
	err := ctx.RegisterResource("azure-native:authorization:AccessReviewScheduleDefinitionById", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessReviewScheduleDefinitionById gets an existing AccessReviewScheduleDefinitionById resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessReviewScheduleDefinitionById(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessReviewScheduleDefinitionByIdState, opts ...pulumi.ResourceOption) (*AccessReviewScheduleDefinitionById, error) {
	var resource AccessReviewScheduleDefinitionById
	err := ctx.ReadResource("azure-native:authorization:AccessReviewScheduleDefinitionById", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessReviewScheduleDefinitionById resources.
type accessReviewScheduleDefinitionByIdState struct {
}

type AccessReviewScheduleDefinitionByIdState struct {
}

func (AccessReviewScheduleDefinitionByIdState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdState)(nil)).Elem()
}

type accessReviewScheduleDefinitionByIdArgs struct {
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled *bool `pulumi:"autoApplyDecisionsEnabled"`
	// This is the collection of backup reviewers.
	BackupReviewers []AccessReviewReviewer `pulumi:"backupReviewers"`
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision *string `pulumi:"defaultDecision"`
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled *bool `pulumi:"defaultDecisionEnabled"`
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins *string `pulumi:"descriptionForAdmins"`
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers *string `pulumi:"descriptionForReviewers"`
	// The display name for the schedule definition.
	DisplayName *string `pulumi:"displayName"`
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate *string `pulumi:"endDate"`
	// This is used to indicate the resource id(s) to exclude
	ExcludeResourceId *string `pulumi:"excludeResourceId"`
	// This is used to indicate the role definition id(s) to exclude
	ExcludeRoleDefinitionId *string `pulumi:"excludeRoleDefinitionId"`
	// Flag to indicate whether to expand nested memberships or not.
	ExpandNestedMemberships *bool `pulumi:"expandNestedMemberships"`
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration *string `pulumi:"inactiveDuration"`
	// Flag to indicate whether to expand nested memberships or not.
	IncludeAccessBelowResource *bool `pulumi:"includeAccessBelowResource"`
	// Flag to indicate whether to expand nested memberships or not.
	IncludeInheritedAccess *bool `pulumi:"includeInheritedAccess"`
	// The duration in days for an instance.
	InstanceDurationInDays *int `pulumi:"instanceDurationInDays"`
	// This is the collection of instances returned when one does an expand on it.
	Instances []AccessReviewInstance `pulumi:"instances"`
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval *int `pulumi:"interval"`
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval *bool `pulumi:"justificationRequiredOnApproval"`
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled *bool `pulumi:"mailNotificationsEnabled"`
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences *int `pulumi:"numberOfOccurrences"`
	// Recommendations for access reviews are calculated by looking back at 30 days of data(w.r.t the start date of the review) by default. However, in some scenarios, customers want to change how far back to look at and want to configure 60 days, 90 days, etc. instead. This setting allows customers to configure this duration. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	RecommendationLookBackDuration *string `pulumi:"recommendationLookBackDuration"`
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled *bool `pulumi:"recommendationsEnabled"`
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled *bool `pulumi:"reminderNotificationsEnabled"`
	// This is the collection of reviewers.
	Reviewers []AccessReviewReviewer `pulumi:"reviewers"`
	// The id of the access review schedule definition.
	ScheduleDefinitionId *string `pulumi:"scheduleDefinitionId"`
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate *string `pulumi:"startDate"`
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AccessReviewScheduleDefinitionById resource.
type AccessReviewScheduleDefinitionByIdArgs struct {
	// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
	AutoApplyDecisionsEnabled pulumi.BoolPtrInput
	// This is the collection of backup reviewers.
	BackupReviewers AccessReviewReviewerArrayInput
	// This specifies the behavior for the autoReview feature when an access review completes.
	DefaultDecision pulumi.StringPtrInput
	// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
	DefaultDecisionEnabled pulumi.BoolPtrInput
	// The description provided by the access review creator and visible to admins.
	DescriptionForAdmins pulumi.StringPtrInput
	// The description provided by the access review creator to be shown to reviewers.
	DescriptionForReviewers pulumi.StringPtrInput
	// The display name for the schedule definition.
	DisplayName pulumi.StringPtrInput
	// The DateTime when the review is scheduled to end. Required if type is endDate
	EndDate pulumi.StringPtrInput
	// This is used to indicate the resource id(s) to exclude
	ExcludeResourceId pulumi.StringPtrInput
	// This is used to indicate the role definition id(s) to exclude
	ExcludeRoleDefinitionId pulumi.StringPtrInput
	// Flag to indicate whether to expand nested memberships or not.
	ExpandNestedMemberships pulumi.BoolPtrInput
	// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	InactiveDuration pulumi.StringPtrInput
	// Flag to indicate whether to expand nested memberships or not.
	IncludeAccessBelowResource pulumi.BoolPtrInput
	// Flag to indicate whether to expand nested memberships or not.
	IncludeInheritedAccess pulumi.BoolPtrInput
	// The duration in days for an instance.
	InstanceDurationInDays pulumi.IntPtrInput
	// This is the collection of instances returned when one does an expand on it.
	Instances AccessReviewInstanceArrayInput
	// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
	Interval pulumi.IntPtrInput
	// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
	JustificationRequiredOnApproval pulumi.BoolPtrInput
	// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
	MailNotificationsEnabled pulumi.BoolPtrInput
	// The number of times to repeat the access review. Required and must be positive if type is numbered.
	NumberOfOccurrences pulumi.IntPtrInput
	// Recommendations for access reviews are calculated by looking back at 30 days of data(w.r.t the start date of the review) by default. However, in some scenarios, customers want to change how far back to look at and want to configure 60 days, 90 days, etc. instead. This setting allows customers to configure this duration. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
	RecommendationLookBackDuration pulumi.StringPtrInput
	// Flag to indicate whether showing recommendations to reviewers is enabled.
	RecommendationsEnabled pulumi.BoolPtrInput
	// Flag to indicate whether sending reminder emails to reviewers are enabled.
	ReminderNotificationsEnabled pulumi.BoolPtrInput
	// This is the collection of reviewers.
	Reviewers AccessReviewReviewerArrayInput
	// The id of the access review schedule definition.
	ScheduleDefinitionId pulumi.StringPtrInput
	// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
	StartDate pulumi.StringPtrInput
	// The recurrence range type. The possible values are: endDate, noEnd, numbered.
	Type pulumi.StringPtrInput
}

func (AccessReviewScheduleDefinitionByIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessReviewScheduleDefinitionByIdArgs)(nil)).Elem()
}

type AccessReviewScheduleDefinitionByIdInput interface {
	pulumi.Input

	ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput
	ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput
}

func (*AccessReviewScheduleDefinitionById) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return i.ToAccessReviewScheduleDefinitionByIdOutputWithContext(context.Background())
}

func (i *AccessReviewScheduleDefinitionById) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewScheduleDefinitionByIdOutput)
}

type AccessReviewScheduleDefinitionByIdOutput struct{ *pulumi.OutputState }

func (AccessReviewScheduleDefinitionByIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewScheduleDefinitionById)(nil)).Elem()
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutput() AccessReviewScheduleDefinitionByIdOutput {
	return o
}

func (o AccessReviewScheduleDefinitionByIdOutput) ToAccessReviewScheduleDefinitionByIdOutputWithContext(ctx context.Context) AccessReviewScheduleDefinitionByIdOutput {
	return o
}

// The role assignment state eligible/active to review
func (o AccessReviewScheduleDefinitionByIdOutput) AssignmentState() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.AssignmentState }).(pulumi.StringOutput)
}

// Flag to indicate whether auto-apply capability, to automatically change the target object access resource, is enabled. If not enabled, a user must, after the review completes, apply the access review.
func (o AccessReviewScheduleDefinitionByIdOutput) AutoApplyDecisionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.AutoApplyDecisionsEnabled }).(pulumi.BoolPtrOutput)
}

// This is the collection of backup reviewers.
func (o AccessReviewScheduleDefinitionByIdOutput) BackupReviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.BackupReviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

// This specifies the behavior for the autoReview feature when an access review completes.
func (o AccessReviewScheduleDefinitionByIdOutput) DefaultDecision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DefaultDecision }).(pulumi.StringPtrOutput)
}

// Flag to indicate whether reviewers are required to provide a justification when reviewing access.
func (o AccessReviewScheduleDefinitionByIdOutput) DefaultDecisionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.DefaultDecisionEnabled }).(pulumi.BoolPtrOutput)
}

// The description provided by the access review creator and visible to admins.
func (o AccessReviewScheduleDefinitionByIdOutput) DescriptionForAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DescriptionForAdmins }).(pulumi.StringPtrOutput)
}

// The description provided by the access review creator to be shown to reviewers.
func (o AccessReviewScheduleDefinitionByIdOutput) DescriptionForReviewers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DescriptionForReviewers }).(pulumi.StringPtrOutput)
}

// The display name for the schedule definition.
func (o AccessReviewScheduleDefinitionByIdOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The DateTime when the review is scheduled to end. Required if type is endDate
func (o AccessReviewScheduleDefinitionByIdOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

// This is used to indicate the resource id(s) to exclude
func (o AccessReviewScheduleDefinitionByIdOutput) ExcludeResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.ExcludeResourceId }).(pulumi.StringPtrOutput)
}

// This is used to indicate the role definition id(s) to exclude
func (o AccessReviewScheduleDefinitionByIdOutput) ExcludeRoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.ExcludeRoleDefinitionId }).(pulumi.StringPtrOutput)
}

// Flag to indicate whether to expand nested memberships or not.
func (o AccessReviewScheduleDefinitionByIdOutput) ExpandNestedMemberships() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.ExpandNestedMemberships }).(pulumi.BoolPtrOutput)
}

// Duration users are inactive for. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
func (o AccessReviewScheduleDefinitionByIdOutput) InactiveDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.InactiveDuration }).(pulumi.StringPtrOutput)
}

// Flag to indicate whether to expand nested memberships or not.
func (o AccessReviewScheduleDefinitionByIdOutput) IncludeAccessBelowResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.IncludeAccessBelowResource }).(pulumi.BoolPtrOutput)
}

// Flag to indicate whether to expand nested memberships or not.
func (o AccessReviewScheduleDefinitionByIdOutput) IncludeInheritedAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.IncludeInheritedAccess }).(pulumi.BoolPtrOutput)
}

// The duration in days for an instance.
func (o AccessReviewScheduleDefinitionByIdOutput) InstanceDurationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.InstanceDurationInDays }).(pulumi.IntPtrOutput)
}

// This is the collection of instances returned when one does an expand on it.
func (o AccessReviewScheduleDefinitionByIdOutput) Instances() AccessReviewInstanceResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewInstanceResponseArrayOutput {
		return v.Instances
	}).(AccessReviewInstanceResponseArrayOutput)
}

// The interval for recurrence. For a quarterly review, the interval is 3 for type : absoluteMonthly.
func (o AccessReviewScheduleDefinitionByIdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Flag to indicate whether the reviewer is required to pass justification when recording a decision.
func (o AccessReviewScheduleDefinitionByIdOutput) JustificationRequiredOnApproval() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.JustificationRequiredOnApproval
	}).(pulumi.BoolPtrOutput)
}

// Flag to indicate whether sending mails to reviewers and the review creator is enabled.
func (o AccessReviewScheduleDefinitionByIdOutput) MailNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.MailNotificationsEnabled }).(pulumi.BoolPtrOutput)
}

// The access review schedule definition unique id.
func (o AccessReviewScheduleDefinitionByIdOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of times to repeat the access review. Required and must be positive if type is numbered.
func (o AccessReviewScheduleDefinitionByIdOutput) NumberOfOccurrences() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.IntPtrOutput { return v.NumberOfOccurrences }).(pulumi.IntPtrOutput)
}

// The identity id
func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The identity display name
func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// The identity type user/servicePrincipal to review
func (o AccessReviewScheduleDefinitionByIdOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// Recommendations for access reviews are calculated by looking back at 30 days of data(w.r.t the start date of the review) by default. However, in some scenarios, customers want to change how far back to look at and want to configure 60 days, 90 days, etc. instead. This setting allows customers to configure this duration. The value should be in ISO  8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimeSpan to a valid interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
func (o AccessReviewScheduleDefinitionByIdOutput) RecommendationLookBackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput {
		return v.RecommendationLookBackDuration
	}).(pulumi.StringPtrOutput)
}

// Flag to indicate whether showing recommendations to reviewers is enabled.
func (o AccessReviewScheduleDefinitionByIdOutput) RecommendationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput { return v.RecommendationsEnabled }).(pulumi.BoolPtrOutput)
}

// Flag to indicate whether sending reminder emails to reviewers are enabled.
func (o AccessReviewScheduleDefinitionByIdOutput) ReminderNotificationsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.BoolPtrOutput {
		return v.ReminderNotificationsEnabled
	}).(pulumi.BoolPtrOutput)
}

// ResourceId in which this review is getting created
func (o AccessReviewScheduleDefinitionByIdOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// This is the collection of reviewers.
func (o AccessReviewScheduleDefinitionByIdOutput) Reviewers() AccessReviewReviewerResponseArrayOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) AccessReviewReviewerResponseArrayOutput {
		return v.Reviewers
	}).(AccessReviewReviewerResponseArrayOutput)
}

// This field specifies the type of reviewers for a review. Usually for a review, reviewers are explicitly assigned. However, in some cases, the reviewers may not be assigned and instead be chosen dynamically. For example managers review or self review.
func (o AccessReviewScheduleDefinitionByIdOutput) ReviewersType() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.ReviewersType }).(pulumi.StringOutput)
}

// This is used to indicate the role being reviewed
func (o AccessReviewScheduleDefinitionByIdOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

// The DateTime when the review is scheduled to be start. This could be a date in the future. Required on create.
func (o AccessReviewScheduleDefinitionByIdOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringPtrOutput { return v.StartDate }).(pulumi.StringPtrOutput)
}

// This read-only field specifies the status of an accessReview.
func (o AccessReviewScheduleDefinitionByIdOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The resource type.
func (o AccessReviewScheduleDefinitionByIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The user principal name(if valid)
func (o AccessReviewScheduleDefinitionByIdOutput) UserPrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessReviewScheduleDefinitionById) pulumi.StringOutput { return v.UserPrincipalName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewScheduleDefinitionByIdOutput{})
}
