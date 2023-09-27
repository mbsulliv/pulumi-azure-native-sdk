// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response for POST requests that return single SharedAccessAuthorizationRule.
type NotificationHubAuthorizationRule struct {
	pulumi.CustomResourceState

	// Deprecated - only for compatibility.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// SharedAccessAuthorizationRule properties.
	Properties SharedAccessAuthorizationRulePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Deprecated - only for compatibility.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNotificationHubAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *NotificationHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.NotificationHubName == nil {
		return nil, errors.New("invalid value for required argument 'NotificationHubName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:notificationhubs:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20160301:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20170401:NotificationHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:notificationhubs/v20230901:NotificationHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NotificationHubAuthorizationRule
	err := ctx.RegisterResource("azure-native:notificationhubs/v20230101preview:NotificationHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationHubAuthorizationRule gets an existing NotificationHubAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NotificationHubAuthorizationRule, error) {
	var resource NotificationHubAuthorizationRule
	err := ctx.ReadResource("azure-native:notificationhubs/v20230101preview:NotificationHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationHubAuthorizationRule resources.
type notificationHubAuthorizationRuleState struct {
}

type NotificationHubAuthorizationRuleState struct {
}

func (NotificationHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleState)(nil)).Elem()
}

type notificationHubAuthorizationRuleArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// Deprecated - only for compatibility.
	Location *string `pulumi:"location"`
	// Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Notification Hub name
	NotificationHubName string `pulumi:"notificationHubName"`
	// SharedAccessAuthorizationRule properties.
	Properties *SharedAccessAuthorizationRuleProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Deprecated - only for compatibility.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NotificationHubAuthorizationRule resource.
type NotificationHubAuthorizationRuleArgs struct {
	// Authorization Rule Name
	AuthorizationRuleName pulumi.StringPtrInput
	// Deprecated - only for compatibility.
	Location pulumi.StringPtrInput
	// Namespace name
	NamespaceName pulumi.StringInput
	// Notification Hub name
	NotificationHubName pulumi.StringInput
	// SharedAccessAuthorizationRule properties.
	Properties SharedAccessAuthorizationRulePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Deprecated - only for compatibility.
	Tags pulumi.StringMapInput
}

func (NotificationHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationHubAuthorizationRuleArgs)(nil)).Elem()
}

type NotificationHubAuthorizationRuleInput interface {
	pulumi.Input

	ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput
	ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput
}

func (*NotificationHubAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubAuthorizationRule)(nil)).Elem()
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return i.ToNotificationHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NotificationHubAuthorizationRule) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationHubAuthorizationRuleOutput)
}

func (i *NotificationHubAuthorizationRule) ToOutput(ctx context.Context) pulumix.Output[*NotificationHubAuthorizationRule] {
	return pulumix.Output[*NotificationHubAuthorizationRule]{
		OutputState: i.ToNotificationHubAuthorizationRuleOutputWithContext(ctx).OutputState,
	}
}

type NotificationHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NotificationHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationHubAuthorizationRule)(nil)).Elem()
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutput() NotificationHubAuthorizationRuleOutput {
	return o
}

func (o NotificationHubAuthorizationRuleOutput) ToNotificationHubAuthorizationRuleOutputWithContext(ctx context.Context) NotificationHubAuthorizationRuleOutput {
	return o
}

func (o NotificationHubAuthorizationRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*NotificationHubAuthorizationRule] {
	return pulumix.Output[*NotificationHubAuthorizationRule]{
		OutputState: o.OutputState,
	}
}

// Deprecated - only for compatibility.
func (o NotificationHubAuthorizationRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o NotificationHubAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SharedAccessAuthorizationRule properties.
func (o NotificationHubAuthorizationRuleOutput) Properties() SharedAccessAuthorizationRulePropertiesResponseOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) SharedAccessAuthorizationRulePropertiesResponseOutput {
		return v.Properties
	}).(SharedAccessAuthorizationRulePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NotificationHubAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Deprecated - only for compatibility.
func (o NotificationHubAuthorizationRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NotificationHubAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationHubAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NotificationHubAuthorizationRuleOutput{})
}
