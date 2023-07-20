// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The IPPrefix resource definition.
type IpPrefix struct {
	pulumi.CustomResourceState

	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// IpPrefix contains the list of IP PrefixRules objects.
	IpPrefixRules IpPrefixPropertiesResponseIpPrefixRulesArrayOutput `pulumi:"ipPrefixRules"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpPrefix registers a new resource with the given unique name, arguments, and options.
func NewIpPrefix(ctx *pulumi.Context,
	name string, args *IpPrefixArgs, opts ...pulumi.ResourceOption) (*IpPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpPrefixRules == nil {
		return nil, errors.New("invalid value for required argument 'IpPrefixRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:IpPrefix"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:IpPrefix"),
		},
	})
	opts = append(opts, aliases)
	var resource IpPrefix
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230201preview:IpPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpPrefix gets an existing IpPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpPrefixState, opts ...pulumi.ResourceOption) (*IpPrefix, error) {
	var resource IpPrefix
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230201preview:IpPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpPrefix resources.
type ipPrefixState struct {
}

type IpPrefixState struct {
}

func (IpPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipPrefixState)(nil)).Elem()
}

type ipPrefixArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the IP Prefix
	IpPrefixName *string `pulumi:"ipPrefixName"`
	// IpPrefix contains the list of IP PrefixRules objects.
	IpPrefixRules []IpPrefixPropertiesIpPrefixRules `pulumi:"ipPrefixRules"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IpPrefix resource.
type IpPrefixArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the IP Prefix
	IpPrefixName pulumi.StringPtrInput
	// IpPrefix contains the list of IP PrefixRules objects.
	IpPrefixRules IpPrefixPropertiesIpPrefixRulesArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (IpPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipPrefixArgs)(nil)).Elem()
}

type IpPrefixInput interface {
	pulumi.Input

	ToIpPrefixOutput() IpPrefixOutput
	ToIpPrefixOutputWithContext(ctx context.Context) IpPrefixOutput
}

func (*IpPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**IpPrefix)(nil)).Elem()
}

func (i *IpPrefix) ToIpPrefixOutput() IpPrefixOutput {
	return i.ToIpPrefixOutputWithContext(context.Background())
}

func (i *IpPrefix) ToIpPrefixOutputWithContext(ctx context.Context) IpPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpPrefixOutput)
}

type IpPrefixOutput struct{ *pulumi.OutputState }

func (IpPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpPrefix)(nil)).Elem()
}

func (o IpPrefixOutput) ToIpPrefixOutput() IpPrefixOutput {
	return o
}

func (o IpPrefixOutput) ToIpPrefixOutputWithContext(ctx context.Context) IpPrefixOutput {
	return o
}

// Switch configuration description.
func (o IpPrefixOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// IpPrefix contains the list of IP PrefixRules objects.
func (o IpPrefixOutput) IpPrefixRules() IpPrefixPropertiesResponseIpPrefixRulesArrayOutput {
	return o.ApplyT(func(v *IpPrefix) IpPrefixPropertiesResponseIpPrefixRulesArrayOutput { return v.IpPrefixRules }).(IpPrefixPropertiesResponseIpPrefixRulesArrayOutput)
}

// The geo-location where the resource lives
func (o IpPrefixOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o IpPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o IpPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IpPrefixOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IpPrefix) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o IpPrefixOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IpPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IpPrefixOutput{})
}
