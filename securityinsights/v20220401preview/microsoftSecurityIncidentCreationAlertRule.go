// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents MicrosoftSecurityIncidentCreation rule.
type MicrosoftSecurityIncidentCreationAlertRule struct {
	pulumi.CustomResourceState

	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrOutput `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// the alerts' displayNames on which the cases will not be generated
	DisplayNamesExcludeFilter pulumi.StringArrayOutput `pulumi:"displayNamesExcludeFilter"`
	// the alerts' displayNames on which the cases will be generated
	DisplayNamesFilter pulumi.StringArrayOutput `pulumi:"displayNamesFilter"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the alert rule
	// Expected value is 'MicrosoftSecurityIncidentCreation'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last time that this alert has been modified.
	LastModifiedUtc pulumi.StringOutput `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The alerts' productName on which the cases will be generated
	ProductFilter pulumi.StringOutput `pulumi:"productFilter"`
	// the alerts' severities on which the cases will be generated
	SeveritiesFilter pulumi.StringArrayOutput `pulumi:"severitiesFilter"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMicrosoftSecurityIncidentCreationAlertRule registers a new resource with the given unique name, arguments, and options.
func NewMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context,
	name string, args *MicrosoftSecurityIncidentCreationAlertRuleArgs, opts ...pulumi.ResourceOption) (*MicrosoftSecurityIncidentCreationAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ProductFilter == nil {
		return nil, errors.New("invalid value for required argument 'ProductFilter'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("MicrosoftSecurityIncidentCreation")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:MicrosoftSecurityIncidentCreationAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource MicrosoftSecurityIncidentCreationAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:MicrosoftSecurityIncidentCreationAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMicrosoftSecurityIncidentCreationAlertRule gets an existing MicrosoftSecurityIncidentCreationAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMicrosoftSecurityIncidentCreationAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MicrosoftSecurityIncidentCreationAlertRuleState, opts ...pulumi.ResourceOption) (*MicrosoftSecurityIncidentCreationAlertRule, error) {
	var resource MicrosoftSecurityIncidentCreationAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:MicrosoftSecurityIncidentCreationAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MicrosoftSecurityIncidentCreationAlertRule resources.
type microsoftSecurityIncidentCreationAlertRuleState struct {
}

type MicrosoftSecurityIncidentCreationAlertRuleState struct {
}

func (MicrosoftSecurityIncidentCreationAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftSecurityIncidentCreationAlertRuleState)(nil)).Elem()
}

type microsoftSecurityIncidentCreationAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName *string `pulumi:"alertRuleTemplateName"`
	// The description of the alert rule.
	Description *string `pulumi:"description"`
	// The display name for alerts created by this alert rule.
	DisplayName string `pulumi:"displayName"`
	// the alerts' displayNames on which the cases will not be generated
	DisplayNamesExcludeFilter []string `pulumi:"displayNamesExcludeFilter"`
	// the alerts' displayNames on which the cases will be generated
	DisplayNamesFilter []string `pulumi:"displayNamesFilter"`
	// Determines whether this alert rule is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// The kind of the alert rule
	// Expected value is 'MicrosoftSecurityIncidentCreation'.
	Kind string `pulumi:"kind"`
	// The alerts' productName on which the cases will be generated
	ProductFilter string `pulumi:"productFilter"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Alert rule ID
	RuleId *string `pulumi:"ruleId"`
	// the alerts' severities on which the cases will be generated
	SeveritiesFilter []string `pulumi:"severitiesFilter"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MicrosoftSecurityIncidentCreationAlertRule resource.
type MicrosoftSecurityIncidentCreationAlertRuleArgs struct {
	// The Name of the alert rule template used to create this rule.
	AlertRuleTemplateName pulumi.StringPtrInput
	// The description of the alert rule.
	Description pulumi.StringPtrInput
	// The display name for alerts created by this alert rule.
	DisplayName pulumi.StringInput
	// the alerts' displayNames on which the cases will not be generated
	DisplayNamesExcludeFilter pulumi.StringArrayInput
	// the alerts' displayNames on which the cases will be generated
	DisplayNamesFilter pulumi.StringArrayInput
	// Determines whether this alert rule is enabled or disabled.
	Enabled pulumi.BoolInput
	// The kind of the alert rule
	// Expected value is 'MicrosoftSecurityIncidentCreation'.
	Kind pulumi.StringInput
	// The alerts' productName on which the cases will be generated
	ProductFilter pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Alert rule ID
	RuleId pulumi.StringPtrInput
	// the alerts' severities on which the cases will be generated
	SeveritiesFilter pulumi.StringArrayInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MicrosoftSecurityIncidentCreationAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*microsoftSecurityIncidentCreationAlertRuleArgs)(nil)).Elem()
}

type MicrosoftSecurityIncidentCreationAlertRuleInput interface {
	pulumi.Input

	ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput
	ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput
}

func (*MicrosoftSecurityIncidentCreationAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftSecurityIncidentCreationAlertRule)(nil)).Elem()
}

func (i *MicrosoftSecurityIncidentCreationAlertRule) ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return i.ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(context.Background())
}

func (i *MicrosoftSecurityIncidentCreationAlertRule) ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftSecurityIncidentCreationAlertRuleOutput)
}

type MicrosoftSecurityIncidentCreationAlertRuleOutput struct{ *pulumi.OutputState }

func (MicrosoftSecurityIncidentCreationAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MicrosoftSecurityIncidentCreationAlertRule)(nil)).Elem()
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ToMicrosoftSecurityIncidentCreationAlertRuleOutput() MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return o
}

func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ToMicrosoftSecurityIncidentCreationAlertRuleOutputWithContext(ctx context.Context) MicrosoftSecurityIncidentCreationAlertRuleOutput {
	return o
}

// The Name of the alert rule template used to create this rule.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput {
		return v.AlertRuleTemplateName
	}).(pulumi.StringPtrOutput)
}

// The description of the alert rule.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for alerts created by this alert rule.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// the alerts' displayNames on which the cases will not be generated
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayNamesExcludeFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.DisplayNamesExcludeFilter
	}).(pulumi.StringArrayOutput)
}

// the alerts' displayNames on which the cases will be generated
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) DisplayNamesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.DisplayNamesFilter
	}).(pulumi.StringArrayOutput)
}

// Determines whether this alert rule is enabled or disabled.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the alert rule
// Expected value is 'MicrosoftSecurityIncidentCreation'.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this alert has been modified.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The alerts' productName on which the cases will be generated
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) ProductFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.ProductFilter }).(pulumi.StringOutput)
}

// the alerts' severities on which the cases will be generated
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) SeveritiesFilter() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringArrayOutput {
		return v.SeveritiesFilter
	}).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MicrosoftSecurityIncidentCreationAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MicrosoftSecurityIncidentCreationAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MicrosoftSecurityIncidentCreationAlertRuleOutput{})
}
