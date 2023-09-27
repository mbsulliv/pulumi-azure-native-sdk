// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automanage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Configuration profile assignment is an association between a VM and automanage profile configuration.
// Azure REST API version: 2022-05-04.
type ConfigurationProfileHCIAssignment struct {
	pulumi.CustomResourceState

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfileHCIAssignment registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfileHCIAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileHCIAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCIAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfileHCIAssignment"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfileHCIAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationProfileHCIAssignment
	err := ctx.RegisterResource("azure-native:automanage:ConfigurationProfileHCIAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfileHCIAssignment gets an existing ConfigurationProfileHCIAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfileHCIAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileHCIAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCIAssignment, error) {
	var resource ConfigurationProfileHCIAssignment
	err := ctx.ReadResource("azure-native:automanage:ConfigurationProfileHCIAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfileHCIAssignment resources.
type configurationProfileHCIAssignmentState struct {
}

type ConfigurationProfileHCIAssignmentState struct {
}

func (ConfigurationProfileHCIAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCIAssignmentState)(nil)).Elem()
}

type configurationProfileHCIAssignmentArgs struct {
	// The name of the Arc machine.
	ClusterName string `pulumi:"clusterName"`
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName *string `pulumi:"configurationProfileAssignmentName"`
	// Properties of the configuration profile assignment.
	Properties *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConfigurationProfileHCIAssignment resource.
type ConfigurationProfileHCIAssignmentArgs struct {
	// The name of the Arc machine.
	ClusterName pulumi.StringInput
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName pulumi.StringPtrInput
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ConfigurationProfileHCIAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCIAssignmentArgs)(nil)).Elem()
}

type ConfigurationProfileHCIAssignmentInput interface {
	pulumi.Input

	ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput
	ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput
}

func (*ConfigurationProfileHCIAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCIAssignment)(nil)).Elem()
}

func (i *ConfigurationProfileHCIAssignment) ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput {
	return i.ToConfigurationProfileHCIAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationProfileHCIAssignment) ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileHCIAssignmentOutput)
}

func (i *ConfigurationProfileHCIAssignment) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationProfileHCIAssignment] {
	return pulumix.Output[*ConfigurationProfileHCIAssignment]{
		OutputState: i.ToConfigurationProfileHCIAssignmentOutputWithContext(ctx).OutputState,
	}
}

type ConfigurationProfileHCIAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileHCIAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCIAssignment)(nil)).Elem()
}

func (o ConfigurationProfileHCIAssignmentOutput) ToConfigurationProfileHCIAssignmentOutput() ConfigurationProfileHCIAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCIAssignmentOutput) ToConfigurationProfileHCIAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCIAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCIAssignmentOutput) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationProfileHCIAssignment] {
	return pulumix.Output[*ConfigurationProfileHCIAssignment]{
		OutputState: o.OutputState,
	}
}

// Azure resource id. Indicates if this resource is managed by another Azure resource.
func (o ConfigurationProfileHCIAssignmentOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationProfileHCIAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile assignment.
func (o ConfigurationProfileHCIAssignmentOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) ConfigurationProfileAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationProfileHCIAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfileHCIAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCIAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileHCIAssignmentOutput{})
}
