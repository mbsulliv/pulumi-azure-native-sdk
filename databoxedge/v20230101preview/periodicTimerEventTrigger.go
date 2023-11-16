// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Trigger details.
type PeriodicTimerEventTrigger struct {
	pulumi.CustomResourceState

	// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
	CustomContextTag pulumi.StringPtrOutput `pulumi:"customContextTag"`
	// Trigger Kind.
	// Expected value is 'PeriodicTimerEvent'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Role Sink information.
	SinkInfo RoleSinkInfoResponseOutput `pulumi:"sinkInfo"`
	// Periodic timer details.
	SourceInfo PeriodicTimerSourceInfoResponseOutput `pulumi:"sourceInfo"`
	// Metadata pertaining to creation and last modification of Trigger
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeriodicTimerEventTrigger registers a new resource with the given unique name, arguments, and options.
func NewPeriodicTimerEventTrigger(ctx *pulumi.Context,
	name string, args *PeriodicTimerEventTriggerArgs, opts ...pulumi.ResourceOption) (*PeriodicTimerEventTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SinkInfo == nil {
		return nil, errors.New("invalid value for required argument 'SinkInfo'")
	}
	if args.SourceInfo == nil {
		return nil, errors.New("invalid value for required argument 'SourceInfo'")
	}
	args.Kind = pulumi.String("PeriodicTimerEvent")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:PeriodicTimerEventTrigger"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:PeriodicTimerEventTrigger"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PeriodicTimerEventTrigger
	err := ctx.RegisterResource("azure-native:databoxedge/v20230101preview:PeriodicTimerEventTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeriodicTimerEventTrigger gets an existing PeriodicTimerEventTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeriodicTimerEventTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeriodicTimerEventTriggerState, opts ...pulumi.ResourceOption) (*PeriodicTimerEventTrigger, error) {
	var resource PeriodicTimerEventTrigger
	err := ctx.ReadResource("azure-native:databoxedge/v20230101preview:PeriodicTimerEventTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeriodicTimerEventTrigger resources.
type periodicTimerEventTriggerState struct {
}

type PeriodicTimerEventTriggerState struct {
}

func (PeriodicTimerEventTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*periodicTimerEventTriggerState)(nil)).Elem()
}

type periodicTimerEventTriggerArgs struct {
	// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
	CustomContextTag *string `pulumi:"customContextTag"`
	// Creates or updates a trigger
	DeviceName string `pulumi:"deviceName"`
	// Trigger Kind.
	// Expected value is 'PeriodicTimerEvent'.
	Kind string `pulumi:"kind"`
	// The trigger name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Role Sink information.
	SinkInfo RoleSinkInfo `pulumi:"sinkInfo"`
	// Periodic timer details.
	SourceInfo PeriodicTimerSourceInfo `pulumi:"sourceInfo"`
}

// The set of arguments for constructing a PeriodicTimerEventTrigger resource.
type PeriodicTimerEventTriggerArgs struct {
	// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
	CustomContextTag pulumi.StringPtrInput
	// Creates or updates a trigger
	DeviceName pulumi.StringInput
	// Trigger Kind.
	// Expected value is 'PeriodicTimerEvent'.
	Kind pulumi.StringInput
	// The trigger name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Role Sink information.
	SinkInfo RoleSinkInfoInput
	// Periodic timer details.
	SourceInfo PeriodicTimerSourceInfoInput
}

func (PeriodicTimerEventTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*periodicTimerEventTriggerArgs)(nil)).Elem()
}

type PeriodicTimerEventTriggerInput interface {
	pulumi.Input

	ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput
	ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput
}

func (*PeriodicTimerEventTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerEventTrigger)(nil)).Elem()
}

func (i *PeriodicTimerEventTrigger) ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput {
	return i.ToPeriodicTimerEventTriggerOutputWithContext(context.Background())
}

func (i *PeriodicTimerEventTrigger) ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicTimerEventTriggerOutput)
}

type PeriodicTimerEventTriggerOutput struct{ *pulumi.OutputState }

func (PeriodicTimerEventTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicTimerEventTrigger)(nil)).Elem()
}

func (o PeriodicTimerEventTriggerOutput) ToPeriodicTimerEventTriggerOutput() PeriodicTimerEventTriggerOutput {
	return o
}

func (o PeriodicTimerEventTriggerOutput) ToPeriodicTimerEventTriggerOutputWithContext(ctx context.Context) PeriodicTimerEventTriggerOutput {
	return o
}

// A custom context tag typically used to correlate the trigger against its usage. For example, if a periodic timer trigger is intended for certain specific IoT modules in the device, the tag can be the name or the image URL of the module.
func (o PeriodicTimerEventTriggerOutput) CustomContextTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringPtrOutput { return v.CustomContextTag }).(pulumi.StringPtrOutput)
}

// Trigger Kind.
// Expected value is 'PeriodicTimerEvent'.
func (o PeriodicTimerEventTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o PeriodicTimerEventTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Role Sink information.
func (o PeriodicTimerEventTriggerOutput) SinkInfo() RoleSinkInfoResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) RoleSinkInfoResponseOutput { return v.SinkInfo }).(RoleSinkInfoResponseOutput)
}

// Periodic timer details.
func (o PeriodicTimerEventTriggerOutput) SourceInfo() PeriodicTimerSourceInfoResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) PeriodicTimerSourceInfoResponseOutput { return v.SourceInfo }).(PeriodicTimerSourceInfoResponseOutput)
}

// Metadata pertaining to creation and last modification of Trigger
func (o PeriodicTimerEventTriggerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o PeriodicTimerEventTriggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeriodicTimerEventTrigger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeriodicTimerEventTriggerOutput{})
}
