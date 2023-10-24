// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Configuration Assignment
// Azure REST API version: 2022-11-01-preview. Prior API version in Azure Native 1.x: 2021-04-01-preview.
//
// Other available API versions: 2023-04-01, 2023-09-01-preview.
type ConfigurationAssignment struct {
	pulumi.CustomResourceState

	// Location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique resourceId
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationAssignment registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance/v20210401preview:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210901preview:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20220701preview:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20221101preview:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20230401:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20230901preview:ConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationAssignment
	err := ctx.RegisterResource("azure-native:maintenance:ConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAssignment gets an existing ConfigurationAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationAssignment, error) {
	var resource ConfigurationAssignment
	err := ctx.ReadResource("azure-native:maintenance:ConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAssignment resources.
type configurationAssignmentState struct {
}

type ConfigurationAssignmentState struct {
}

func (ConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentState)(nil)).Elem()
}

type configurationAssignmentArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Resource provider name
	ProviderName string `pulumi:"providerName"`
	// Resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
	// Resource identifier
	ResourceName string `pulumi:"resourceName"`
	// Resource type
	ResourceType string `pulumi:"resourceType"`
}

// The set of arguments for constructing a ConfigurationAssignment resource.
type ConfigurationAssignmentArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringPtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Resource provider name
	ProviderName pulumi.StringInput
	// Resource group name
	ResourceGroupName pulumi.StringInput
	// The unique resourceId
	ResourceId pulumi.StringPtrInput
	// Resource identifier
	ResourceName pulumi.StringInput
	// Resource type
	ResourceType pulumi.StringInput
}

func (ConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentArgs)(nil)).Elem()
}

type ConfigurationAssignmentInput interface {
	pulumi.Input

	ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput
	ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput
}

func (*ConfigurationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignment)(nil)).Elem()
}

func (i *ConfigurationAssignment) ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput {
	return i.ToConfigurationAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationAssignment) ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentOutput)
}

func (i *ConfigurationAssignment) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationAssignment] {
	return pulumix.Output[*ConfigurationAssignment]{
		OutputState: i.ToConfigurationAssignmentOutputWithContext(ctx).OutputState,
	}
}

type ConfigurationAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignment)(nil)).Elem()
}

func (o ConfigurationAssignmentOutput) ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput {
	return o
}

func (o ConfigurationAssignmentOutput) ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput {
	return o
}

func (o ConfigurationAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationAssignment] {
	return pulumix.Output[*ConfigurationAssignment]{
		OutputState: o.OutputState,
	}
}

// Location of the resource
func (o ConfigurationAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o ConfigurationAssignmentOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o ConfigurationAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o ConfigurationAssignmentOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o ConfigurationAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentOutput{})
}
