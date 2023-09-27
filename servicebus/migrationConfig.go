// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Single item in List or Get Migration Config operation
// Azure REST API version: 2022-01-01-preview. Prior API version in Azure Native 1.x: 2017-04-01
type MigrationConfig struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// State in which Standard to Premium Migration is, possible values : Unknown, Reverting, Completing, Initiating, Syncing, Active
	MigrationState pulumi.StringOutput `pulumi:"migrationState"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount pulumi.Float64Output `pulumi:"pendingReplicationOperationsCount"`
	// Name to access Standard Namespace after migration
	PostMigrationName pulumi.StringOutput `pulumi:"postMigrationName"`
	// Provisioning state of Migration Configuration
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Existing premium Namespace ARM Id name which has no entities, will be used for migration
	TargetNamespace pulumi.StringOutput `pulumi:"targetNamespace"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMigrationConfig registers a new resource with the given unique name, arguments, and options.
func NewMigrationConfig(ctx *pulumi.Context,
	name string, args *MigrationConfigArgs, opts ...pulumi.ResourceOption) (*MigrationConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.PostMigrationName == nil {
		return nil, errors.New("invalid value for required argument 'PostMigrationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetNamespace == nil {
		return nil, errors.New("invalid value for required argument 'TargetNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20211101:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20220101preview:MigrationConfig"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20221001preview:MigrationConfig"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MigrationConfig
	err := ctx.RegisterResource("azure-native:servicebus:MigrationConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigrationConfig gets an existing MigrationConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigrationConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationConfigState, opts ...pulumi.ResourceOption) (*MigrationConfig, error) {
	var resource MigrationConfig
	err := ctx.ReadResource("azure-native:servicebus:MigrationConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MigrationConfig resources.
type migrationConfigState struct {
}

type MigrationConfigState struct {
}

func (MigrationConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationConfigState)(nil)).Elem()
}

type migrationConfigArgs struct {
	// The configuration name. Should always be "$default".
	ConfigName *string `pulumi:"configName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name to access Standard Namespace after migration
	PostMigrationName string `pulumi:"postMigrationName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Existing premium Namespace ARM Id name which has no entities, will be used for migration
	TargetNamespace string `pulumi:"targetNamespace"`
}

// The set of arguments for constructing a MigrationConfig resource.
type MigrationConfigArgs struct {
	// The configuration name. Should always be "$default".
	ConfigName pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Name to access Standard Namespace after migration
	PostMigrationName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Existing premium Namespace ARM Id name which has no entities, will be used for migration
	TargetNamespace pulumi.StringInput
}

func (MigrationConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationConfigArgs)(nil)).Elem()
}

type MigrationConfigInput interface {
	pulumi.Input

	ToMigrationConfigOutput() MigrationConfigOutput
	ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput
}

func (*MigrationConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationConfig)(nil)).Elem()
}

func (i *MigrationConfig) ToMigrationConfigOutput() MigrationConfigOutput {
	return i.ToMigrationConfigOutputWithContext(context.Background())
}

func (i *MigrationConfig) ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationConfigOutput)
}

func (i *MigrationConfig) ToOutput(ctx context.Context) pulumix.Output[*MigrationConfig] {
	return pulumix.Output[*MigrationConfig]{
		OutputState: i.ToMigrationConfigOutputWithContext(ctx).OutputState,
	}
}

type MigrationConfigOutput struct{ *pulumi.OutputState }

func (MigrationConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationConfig)(nil)).Elem()
}

func (o MigrationConfigOutput) ToMigrationConfigOutput() MigrationConfigOutput {
	return o
}

func (o MigrationConfigOutput) ToMigrationConfigOutputWithContext(ctx context.Context) MigrationConfigOutput {
	return o
}

func (o MigrationConfigOutput) ToOutput(ctx context.Context) pulumix.Output[*MigrationConfig] {
	return pulumix.Output[*MigrationConfig]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o MigrationConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// State in which Standard to Premium Migration is, possible values : Unknown, Reverting, Completing, Initiating, Syncing, Active
func (o MigrationConfigOutput) MigrationState() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.MigrationState }).(pulumi.StringOutput)
}

// The name of the resource
func (o MigrationConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of entities pending to be replicated.
func (o MigrationConfigOutput) PendingReplicationOperationsCount() pulumi.Float64Output {
	return o.ApplyT(func(v *MigrationConfig) pulumi.Float64Output { return v.PendingReplicationOperationsCount }).(pulumi.Float64Output)
}

// Name to access Standard Namespace after migration
func (o MigrationConfigOutput) PostMigrationName() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.PostMigrationName }).(pulumi.StringOutput)
}

// Provisioning state of Migration Configuration
func (o MigrationConfigOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system meta data relating to this resource.
func (o MigrationConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MigrationConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Existing premium Namespace ARM Id name which has no entities, will be used for migration
func (o MigrationConfigOutput) TargetNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.TargetNamespace }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o MigrationConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrationConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrationConfigOutput{})
}
