// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the admin rule collection.
type AdminRuleCollection struct {
	pulumi.CustomResourceState

	// Groups for configuration
	AppliesToGroups NetworkManagerSecurityGroupItemResponseArrayOutput `pulumi:"appliesToGroups"`
	// A description of the admin rule collection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAdminRuleCollection registers a new resource with the given unique name, arguments, and options.
func NewAdminRuleCollection(ctx *pulumi.Context,
	name string, args *AdminRuleCollectionArgs, opts ...pulumi.ResourceOption) (*AdminRuleCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppliesToGroups == nil {
		return nil, errors.New("invalid value for required argument 'AppliesToGroups'")
	}
	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:AdminRuleCollection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:AdminRuleCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AdminRuleCollection
	err := ctx.RegisterResource("azure-native:network/v20230201:AdminRuleCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdminRuleCollection gets an existing AdminRuleCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdminRuleCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdminRuleCollectionState, opts ...pulumi.ResourceOption) (*AdminRuleCollection, error) {
	var resource AdminRuleCollection
	err := ctx.ReadResource("azure-native:network/v20230201:AdminRuleCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdminRuleCollection resources.
type adminRuleCollectionState struct {
}

type AdminRuleCollectionState struct {
}

func (AdminRuleCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleCollectionState)(nil)).Elem()
}

type adminRuleCollectionArgs struct {
	// Groups for configuration
	AppliesToGroups []NetworkManagerSecurityGroupItem `pulumi:"appliesToGroups"`
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// A description of the admin rule collection.
	Description *string `pulumi:"description"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName *string `pulumi:"ruleCollectionName"`
}

// The set of arguments for constructing a AdminRuleCollection resource.
type AdminRuleCollectionArgs struct {
	// Groups for configuration
	AppliesToGroups NetworkManagerSecurityGroupItemArrayInput
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput
	// A description of the admin rule collection.
	Description pulumi.StringPtrInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringPtrInput
}

func (AdminRuleCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adminRuleCollectionArgs)(nil)).Elem()
}

type AdminRuleCollectionInput interface {
	pulumi.Input

	ToAdminRuleCollectionOutput() AdminRuleCollectionOutput
	ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput
}

func (*AdminRuleCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRuleCollection)(nil)).Elem()
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return i.ToAdminRuleCollectionOutputWithContext(context.Background())
}

func (i *AdminRuleCollection) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdminRuleCollectionOutput)
}

func (i *AdminRuleCollection) ToOutput(ctx context.Context) pulumix.Output[*AdminRuleCollection] {
	return pulumix.Output[*AdminRuleCollection]{
		OutputState: i.ToAdminRuleCollectionOutputWithContext(ctx).OutputState,
	}
}

type AdminRuleCollectionOutput struct{ *pulumi.OutputState }

func (AdminRuleCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdminRuleCollection)(nil)).Elem()
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutput() AdminRuleCollectionOutput {
	return o
}

func (o AdminRuleCollectionOutput) ToAdminRuleCollectionOutputWithContext(ctx context.Context) AdminRuleCollectionOutput {
	return o
}

func (o AdminRuleCollectionOutput) ToOutput(ctx context.Context) pulumix.Output[*AdminRuleCollection] {
	return pulumix.Output[*AdminRuleCollection]{
		OutputState: o.OutputState,
	}
}

// Groups for configuration
func (o AdminRuleCollectionOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v *AdminRuleCollection) NetworkManagerSecurityGroupItemResponseArrayOutput {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

// A description of the admin rule collection.
func (o AdminRuleCollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o AdminRuleCollectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o AdminRuleCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o AdminRuleCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o AdminRuleCollectionOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o AdminRuleCollectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AdminRuleCollection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o AdminRuleCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AdminRuleCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AdminRuleCollectionOutput{})
}
