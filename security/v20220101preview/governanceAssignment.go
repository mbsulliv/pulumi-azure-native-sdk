// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Governance assignment over a given scope
type GovernanceAssignment struct {
	pulumi.CustomResourceState

	// The additional data for the governance assignment - e.g. links to ticket (optional), see example
	AdditionalData GovernanceAssignmentAdditionalDataResponsePtrOutput `pulumi:"additionalData"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification GovernanceEmailNotificationResponsePtrOutput `pulumi:"governanceEmailNotification"`
	// Defines whether there is a grace period on the governance assignment
	IsGracePeriod pulumi.BoolPtrOutput `pulumi:"isGracePeriod"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The Owner for the governance assignment - e.g. user@contoso.com - see example
	Owner pulumi.StringPtrOutput `pulumi:"owner"`
	// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
	RemediationDueDate pulumi.StringOutput `pulumi:"remediationDueDate"`
	// The ETA (estimated time of arrival) for remediation (optional), see example
	RemediationEta RemediationEtaResponsePtrOutput `pulumi:"remediationEta"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGovernanceAssignment registers a new resource with the given unique name, arguments, and options.
func NewGovernanceAssignment(ctx *pulumi.Context,
	name string, args *GovernanceAssignmentArgs, opts ...pulumi.ResourceOption) (*GovernanceAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssessmentName == nil {
		return nil, errors.New("invalid value for required argument 'AssessmentName'")
	}
	if args.RemediationDueDate == nil {
		return nil, errors.New("invalid value for required argument 'RemediationDueDate'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:GovernanceAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GovernanceAssignment
	err := ctx.RegisterResource("azure-native:security/v20220101preview:GovernanceAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGovernanceAssignment gets an existing GovernanceAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGovernanceAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GovernanceAssignmentState, opts ...pulumi.ResourceOption) (*GovernanceAssignment, error) {
	var resource GovernanceAssignment
	err := ctx.ReadResource("azure-native:security/v20220101preview:GovernanceAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GovernanceAssignment resources.
type governanceAssignmentState struct {
}

type GovernanceAssignmentState struct {
}

func (GovernanceAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceAssignmentState)(nil)).Elem()
}

type governanceAssignmentArgs struct {
	// The additional data for the governance assignment - e.g. links to ticket (optional), see example
	AdditionalData *GovernanceAssignmentAdditionalData `pulumi:"additionalData"`
	// The Assessment Key - A unique key for the assessment type
	AssessmentName string `pulumi:"assessmentName"`
	// The governance assignment key - the assessment key of the required governance assignment
	AssignmentKey *string `pulumi:"assignmentKey"`
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification *GovernanceEmailNotification `pulumi:"governanceEmailNotification"`
	// Defines whether there is a grace period on the governance assignment
	IsGracePeriod *bool `pulumi:"isGracePeriod"`
	// The Owner for the governance assignment - e.g. user@contoso.com - see example
	Owner *string `pulumi:"owner"`
	// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
	RemediationDueDate string `pulumi:"remediationDueDate"`
	// The ETA (estimated time of arrival) for remediation (optional), see example
	RemediationEta *RemediationEta `pulumi:"remediationEta"`
	// The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a GovernanceAssignment resource.
type GovernanceAssignmentArgs struct {
	// The additional data for the governance assignment - e.g. links to ticket (optional), see example
	AdditionalData GovernanceAssignmentAdditionalDataPtrInput
	// The Assessment Key - A unique key for the assessment type
	AssessmentName pulumi.StringInput
	// The governance assignment key - the assessment key of the required governance assignment
	AssignmentKey pulumi.StringPtrInput
	// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
	GovernanceEmailNotification GovernanceEmailNotificationPtrInput
	// Defines whether there is a grace period on the governance assignment
	IsGracePeriod pulumi.BoolPtrInput
	// The Owner for the governance assignment - e.g. user@contoso.com - see example
	Owner pulumi.StringPtrInput
	// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
	RemediationDueDate pulumi.StringInput
	// The ETA (estimated time of arrival) for remediation (optional), see example
	RemediationEta RemediationEtaPtrInput
	// The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
	Scope pulumi.StringInput
}

func (GovernanceAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceAssignmentArgs)(nil)).Elem()
}

type GovernanceAssignmentInput interface {
	pulumi.Input

	ToGovernanceAssignmentOutput() GovernanceAssignmentOutput
	ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput
}

func (*GovernanceAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignment)(nil)).Elem()
}

func (i *GovernanceAssignment) ToGovernanceAssignmentOutput() GovernanceAssignmentOutput {
	return i.ToGovernanceAssignmentOutputWithContext(context.Background())
}

func (i *GovernanceAssignment) ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceAssignmentOutput)
}

type GovernanceAssignmentOutput struct{ *pulumi.OutputState }

func (GovernanceAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceAssignment)(nil)).Elem()
}

func (o GovernanceAssignmentOutput) ToGovernanceAssignmentOutput() GovernanceAssignmentOutput {
	return o
}

func (o GovernanceAssignmentOutput) ToGovernanceAssignmentOutputWithContext(ctx context.Context) GovernanceAssignmentOutput {
	return o
}

// The additional data for the governance assignment - e.g. links to ticket (optional), see example
func (o GovernanceAssignmentOutput) AdditionalData() GovernanceAssignmentAdditionalDataResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) GovernanceAssignmentAdditionalDataResponsePtrOutput {
		return v.AdditionalData
	}).(GovernanceAssignmentAdditionalDataResponsePtrOutput)
}

// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
func (o GovernanceAssignmentOutput) GovernanceEmailNotification() GovernanceEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) GovernanceEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceEmailNotificationResponsePtrOutput)
}

// Defines whether there is a grace period on the governance assignment
func (o GovernanceAssignmentOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

// Resource name
func (o GovernanceAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Owner for the governance assignment - e.g. user@contoso.com - see example
func (o GovernanceAssignmentOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringPtrOutput { return v.Owner }).(pulumi.StringPtrOutput)
}

// The remediation due-date - after this date Secure Score will be affected (in case of  active grace-period)
func (o GovernanceAssignmentOutput) RemediationDueDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.RemediationDueDate }).(pulumi.StringOutput)
}

// The ETA (estimated time of arrival) for remediation (optional), see example
func (o GovernanceAssignmentOutput) RemediationEta() RemediationEtaResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceAssignment) RemediationEtaResponsePtrOutput { return v.RemediationEta }).(RemediationEtaResponsePtrOutput)
}

// Resource type
func (o GovernanceAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GovernanceAssignmentOutput{})
}
