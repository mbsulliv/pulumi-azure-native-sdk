// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extendedlocation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Sync Rules definition.
// Azure REST API version: 2021-08-31-preview. Prior API version in Azure Native 1.x: 2021-08-31-preview
type ResourceSyncRule struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority represents a priority of the Resource Sync Rule
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Provisioning State for the Resource Sync Rule.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
	Selector ResourceSyncRulePropertiesResponseSelectorPtrOutput `pulumi:"selector"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
	TargetResourceGroup pulumi.StringPtrOutput `pulumi:"targetResourceGroup"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceSyncRule registers a new resource with the given unique name, arguments, and options.
func NewResourceSyncRule(ctx *pulumi.Context,
	name string, args *ResourceSyncRuleArgs, opts ...pulumi.ResourceOption) (*ResourceSyncRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:extendedlocation/v20210831preview:ResourceSyncRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ResourceSyncRule
	err := ctx.RegisterResource("azure-native:extendedlocation:ResourceSyncRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceSyncRule gets an existing ResourceSyncRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceSyncRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceSyncRuleState, opts ...pulumi.ResourceOption) (*ResourceSyncRule, error) {
	var resource ResourceSyncRule
	err := ctx.ReadResource("azure-native:extendedlocation:ResourceSyncRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceSyncRule resources.
type resourceSyncRuleState struct {
}

type ResourceSyncRuleState struct {
}

func (ResourceSyncRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceSyncRuleState)(nil)).Elem()
}

type resourceSyncRuleArgs struct {
	// Resource Sync Rule name.
	ChildResourceName *string `pulumi:"childResourceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Priority represents a priority of the Resource Sync Rule
	Priority *int `pulumi:"priority"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName string `pulumi:"resourceName"`
	// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
	Selector *ResourceSyncRulePropertiesSelector `pulumi:"selector"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
	TargetResourceGroup *string `pulumi:"targetResourceGroup"`
}

// The set of arguments for constructing a ResourceSyncRule resource.
type ResourceSyncRuleArgs struct {
	// Resource Sync Rule name.
	ChildResourceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Priority represents a priority of the Resource Sync Rule
	Priority pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Custom Locations name.
	ResourceName pulumi.StringInput
	// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
	Selector ResourceSyncRulePropertiesSelectorPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
	TargetResourceGroup pulumi.StringPtrInput
}

func (ResourceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceSyncRuleArgs)(nil)).Elem()
}

type ResourceSyncRuleInput interface {
	pulumi.Input

	ToResourceSyncRuleOutput() ResourceSyncRuleOutput
	ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput
}

func (*ResourceSyncRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRule)(nil)).Elem()
}

func (i *ResourceSyncRule) ToResourceSyncRuleOutput() ResourceSyncRuleOutput {
	return i.ToResourceSyncRuleOutputWithContext(context.Background())
}

func (i *ResourceSyncRule) ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRuleOutput)
}

func (i *ResourceSyncRule) ToOutput(ctx context.Context) pulumix.Output[*ResourceSyncRule] {
	return pulumix.Output[*ResourceSyncRule]{
		OutputState: i.ToResourceSyncRuleOutputWithContext(ctx).OutputState,
	}
}

type ResourceSyncRuleOutput struct{ *pulumi.OutputState }

func (ResourceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRule)(nil)).Elem()
}

func (o ResourceSyncRuleOutput) ToResourceSyncRuleOutput() ResourceSyncRuleOutput {
	return o
}

func (o ResourceSyncRuleOutput) ToResourceSyncRuleOutputWithContext(ctx context.Context) ResourceSyncRuleOutput {
	return o
}

func (o ResourceSyncRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceSyncRule] {
	return pulumix.Output[*ResourceSyncRule]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o ResourceSyncRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ResourceSyncRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority represents a priority of the Resource Sync Rule
func (o ResourceSyncRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Provisioning State for the Resource Sync Rule.
func (o ResourceSyncRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
func (o ResourceSyncRuleOutput) Selector() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) ResourceSyncRulePropertiesResponseSelectorPtrOutput { return v.Selector }).(ResourceSyncRulePropertiesResponseSelectorPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o ResourceSyncRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ResourceSyncRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ResourceSyncRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// For an unmapped custom resource, its labels will be used to find matching resource sync rules. If this resource sync rule is one of the matching rules with highest priority, then the unmapped custom resource will be projected to the target resource group associated with this resource sync rule. The user creating this resource sync rule should have write permissions on the target resource group and this write permission will be validated when creating the resource sync rule.
func (o ResourceSyncRuleOutput) TargetResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringPtrOutput { return v.TargetResourceGroup }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ResourceSyncRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceSyncRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceSyncRuleOutput{})
}
