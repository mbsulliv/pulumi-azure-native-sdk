// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The alert rule information
type SmartDetectorAlertRule struct {
	pulumi.CustomResourceState

	// The alert rule actions.
	ActionGroups ActionGroupsInformationResponseOutput `pulumi:"actionGroups"`
	// The alert rule description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The alert rule's detector.
	Detector DetectorResponseOutput `pulumi:"detector"`
	// The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 1 minute, depending on the detector.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The alert rule resources scope.
	Scope pulumi.StringArrayOutput `pulumi:"scope"`
	// The alert rule severity.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// The alert rule state.
	State pulumi.StringOutput `pulumi:"state"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The alert rule throttling information.
	Throttling ThrottlingInformationResponsePtrOutput `pulumi:"throttling"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSmartDetectorAlertRule registers a new resource with the given unique name, arguments, and options.
func NewSmartDetectorAlertRule(ctx *pulumi.Context,
	name string, args *SmartDetectorAlertRuleArgs, opts ...pulumi.ResourceOption) (*SmartDetectorAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionGroups == nil {
		return nil, errors.New("invalid value for required argument 'ActionGroups'")
	}
	if args.Detector == nil {
		return nil, errors.New("invalid value for required argument 'Detector'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	if args.Location == nil {
		args.Location = pulumi.StringPtr("global")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:alertsmanagement:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190301:SmartDetectorAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:alertsmanagement/v20190601:SmartDetectorAlertRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SmartDetectorAlertRule
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20210401:SmartDetectorAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmartDetectorAlertRule gets an existing SmartDetectorAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmartDetectorAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmartDetectorAlertRuleState, opts ...pulumi.ResourceOption) (*SmartDetectorAlertRule, error) {
	var resource SmartDetectorAlertRule
	err := ctx.ReadResource("azure-native:alertsmanagement/v20210401:SmartDetectorAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmartDetectorAlertRule resources.
type smartDetectorAlertRuleState struct {
}

type SmartDetectorAlertRuleState struct {
}

func (SmartDetectorAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectorAlertRuleState)(nil)).Elem()
}

type smartDetectorAlertRuleArgs struct {
	// The alert rule actions.
	ActionGroups ActionGroupsInformation `pulumi:"actionGroups"`
	// The name of the alert rule.
	AlertRuleName *string `pulumi:"alertRuleName"`
	// The alert rule description.
	Description *string `pulumi:"description"`
	// The alert rule's detector.
	Detector Detector `pulumi:"detector"`
	// The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 1 minute, depending on the detector.
	Frequency string `pulumi:"frequency"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The alert rule resources scope.
	Scope []string `pulumi:"scope"`
	// The alert rule severity.
	Severity string `pulumi:"severity"`
	// The alert rule state.
	State string `pulumi:"state"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The alert rule throttling information.
	Throttling *ThrottlingInformation `pulumi:"throttling"`
}

// The set of arguments for constructing a SmartDetectorAlertRule resource.
type SmartDetectorAlertRuleArgs struct {
	// The alert rule actions.
	ActionGroups ActionGroupsInformationInput
	// The name of the alert rule.
	AlertRuleName pulumi.StringPtrInput
	// The alert rule description.
	Description pulumi.StringPtrInput
	// The alert rule's detector.
	Detector DetectorInput
	// The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 1 minute, depending on the detector.
	Frequency pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The alert rule resources scope.
	Scope pulumi.StringArrayInput
	// The alert rule severity.
	Severity pulumi.StringInput
	// The alert rule state.
	State pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The alert rule throttling information.
	Throttling ThrottlingInformationPtrInput
}

func (SmartDetectorAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smartDetectorAlertRuleArgs)(nil)).Elem()
}

type SmartDetectorAlertRuleInput interface {
	pulumi.Input

	ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput
	ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput
}

func (*SmartDetectorAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SmartDetectorAlertRule)(nil)).Elem()
}

func (i *SmartDetectorAlertRule) ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput {
	return i.ToSmartDetectorAlertRuleOutputWithContext(context.Background())
}

func (i *SmartDetectorAlertRule) ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmartDetectorAlertRuleOutput)
}

type SmartDetectorAlertRuleOutput struct{ *pulumi.OutputState }

func (SmartDetectorAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmartDetectorAlertRule)(nil)).Elem()
}

func (o SmartDetectorAlertRuleOutput) ToSmartDetectorAlertRuleOutput() SmartDetectorAlertRuleOutput {
	return o
}

func (o SmartDetectorAlertRuleOutput) ToSmartDetectorAlertRuleOutputWithContext(ctx context.Context) SmartDetectorAlertRuleOutput {
	return o
}

// The alert rule actions.
func (o SmartDetectorAlertRuleOutput) ActionGroups() ActionGroupsInformationResponseOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) ActionGroupsInformationResponseOutput { return v.ActionGroups }).(ActionGroupsInformationResponseOutput)
}

// The alert rule description.
func (o SmartDetectorAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The alert rule's detector.
func (o SmartDetectorAlertRuleOutput) Detector() DetectorResponseOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) DetectorResponseOutput { return v.Detector }).(DetectorResponseOutput)
}

// The alert rule frequency in ISO8601 format. The time granularity must be in minutes and minimum value is 1 minute, depending on the detector.
func (o SmartDetectorAlertRuleOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

// The resource location.
func (o SmartDetectorAlertRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o SmartDetectorAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The alert rule resources scope.
func (o SmartDetectorAlertRuleOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringArrayOutput { return v.Scope }).(pulumi.StringArrayOutput)
}

// The alert rule severity.
func (o SmartDetectorAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// The alert rule state.
func (o SmartDetectorAlertRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The resource tags.
func (o SmartDetectorAlertRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The alert rule throttling information.
func (o SmartDetectorAlertRuleOutput) Throttling() ThrottlingInformationResponsePtrOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) ThrottlingInformationResponsePtrOutput { return v.Throttling }).(ThrottlingInformationResponsePtrOutput)
}

// The resource type.
func (o SmartDetectorAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SmartDetectorAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SmartDetectorAlertRuleOutput{})
}
