// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log profile resource.
type LogProfile struct {
	pulumi.CustomResourceState

	// the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
	Categories pulumi.StringArrayOutput `pulumi:"categories"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// the retention policy for the events in the log.
	RetentionPolicy RetentionPolicyResponseOutput `pulumi:"retentionPolicy"`
	// The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
	ServiceBusRuleId pulumi.StringPtrOutput `pulumi:"serviceBusRuleId"`
	// the resource id of the storage account to which you would like to send the Activity Log.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLogProfile registers a new resource with the given unique name, arguments, and options.
func NewLogProfile(ctx *pulumi.Context,
	name string, args *LogProfileArgs, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Categories == nil {
		return nil, errors.New("invalid value for required argument 'Categories'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.RetentionPolicy == nil {
		return nil, errors.New("invalid value for required argument 'RetentionPolicy'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:LogProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LogProfile
	err := ctx.RegisterResource("azure-native:insights/v20160301:LogProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogProfile gets an existing LogProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogProfileState, opts ...pulumi.ResourceOption) (*LogProfile, error) {
	var resource LogProfile
	err := ctx.ReadResource("azure-native:insights/v20160301:LogProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogProfile resources.
type logProfileState struct {
}

type LogProfileState struct {
}

func (LogProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileState)(nil)).Elem()
}

type logProfileArgs struct {
	// the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
	Categories []string `pulumi:"categories"`
	// Resource location
	Location *string `pulumi:"location"`
	// List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
	Locations []string `pulumi:"locations"`
	// The name of the log profile.
	LogProfileName *string `pulumi:"logProfileName"`
	// the retention policy for the events in the log.
	RetentionPolicy RetentionPolicy `pulumi:"retentionPolicy"`
	// The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
	ServiceBusRuleId *string `pulumi:"serviceBusRuleId"`
	// the resource id of the storage account to which you would like to send the Activity Log.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LogProfile resource.
type LogProfileArgs struct {
	// the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
	Categories pulumi.StringArrayInput
	// Resource location
	Location pulumi.StringPtrInput
	// List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
	Locations pulumi.StringArrayInput
	// The name of the log profile.
	LogProfileName pulumi.StringPtrInput
	// the retention policy for the events in the log.
	RetentionPolicy RetentionPolicyInput
	// The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
	ServiceBusRuleId pulumi.StringPtrInput
	// the resource id of the storage account to which you would like to send the Activity Log.
	StorageAccountId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (LogProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logProfileArgs)(nil)).Elem()
}

type LogProfileInput interface {
	pulumi.Input

	ToLogProfileOutput() LogProfileOutput
	ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput
}

func (*LogProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**LogProfile)(nil)).Elem()
}

func (i *LogProfile) ToLogProfileOutput() LogProfileOutput {
	return i.ToLogProfileOutputWithContext(context.Background())
}

func (i *LogProfile) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogProfileOutput)
}

type LogProfileOutput struct{ *pulumi.OutputState }

func (LogProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogProfile)(nil)).Elem()
}

func (o LogProfileOutput) ToLogProfileOutput() LogProfileOutput {
	return o
}

func (o LogProfileOutput) ToLogProfileOutputWithContext(ctx context.Context) LogProfileOutput {
	return o
}

// the categories of the logs. These categories are created as is convenient to the user. Some values are: 'Write', 'Delete', and/or 'Action.'
func (o LogProfileOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

// Resource location
func (o LogProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// List of regions for which Activity Log events should be stored or streamed. It is a comma separated list of valid ARM locations including the 'global' location.
func (o LogProfileOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

// Azure resource name
func (o LogProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the retention policy for the events in the log.
func (o LogProfileOutput) RetentionPolicy() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *LogProfile) RetentionPolicyResponseOutput { return v.RetentionPolicy }).(RetentionPolicyResponseOutput)
}

// The service bus rule ID of the service bus namespace in which you would like to have Event Hubs created for streaming the Activity Log. The rule ID is of the format: '{service bus resource ID}/authorizationrules/{key name}'.
func (o LogProfileOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

// the resource id of the storage account to which you would like to send the Activity Log.
func (o LogProfileOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LogProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o LogProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LogProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LogProfileOutput{})
}
