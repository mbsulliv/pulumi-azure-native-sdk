// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The NSP access rule resource
type NspAccessRule struct {
	pulumi.CustomResourceState

	// Inbound address prefixes (IPv4/IPv6)
	AddressPrefixes pulumi.StringArrayOutput `pulumi:"addressPrefixes"`
	// Direction that specifies whether the access rules is inbound/outbound.
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	// Outbound rules email address format.
	EmailAddresses pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	// Outbound rules fully qualified domain name format.
	FullyQualifiedDomainNames pulumi.StringArrayOutput `pulumi:"fullyQualifiedDomainNames"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Rule specified by the perimeter id.
	NetworkSecurityPerimeters PerimeterBasedAccessRuleResponseArrayOutput `pulumi:"networkSecurityPerimeters"`
	// Outbound rules phone number format.
	PhoneNumbers pulumi.StringArrayOutput `pulumi:"phoneNumbers"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Inbound rules service tag names.
	ServiceTags pulumi.StringArrayOutput `pulumi:"serviceTags"`
	// List of subscription ids
	Subscriptions SubscriptionIdResponseArrayOutput `pulumi:"subscriptions"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNspAccessRule registers a new resource with the given unique name, arguments, and options.
func NewNspAccessRule(ctx *pulumi.Context,
	name string, args *NspAccessRuleArgs, opts ...pulumi.ResourceOption) (*NspAccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NspAccessRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NspAccessRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230701preview:NspAccessRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NspAccessRule
	err := ctx.RegisterResource("azure-native:network/v20230801preview:NspAccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNspAccessRule gets an existing NspAccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNspAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspAccessRuleState, opts ...pulumi.ResourceOption) (*NspAccessRule, error) {
	var resource NspAccessRule
	err := ctx.ReadResource("azure-native:network/v20230801preview:NspAccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NspAccessRule resources.
type nspAccessRuleState struct {
}

type NspAccessRuleState struct {
}

func (NspAccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAccessRuleState)(nil)).Elem()
}

type nspAccessRuleArgs struct {
	// The name of the NSP access rule.
	AccessRuleName *string `pulumi:"accessRuleName"`
	// Inbound address prefixes (IPv4/IPv6)
	AddressPrefixes []string `pulumi:"addressPrefixes"`
	// Direction that specifies whether the access rules is inbound/outbound.
	Direction *string `pulumi:"direction"`
	// Outbound rules email address format.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Outbound rules fully qualified domain name format.
	FullyQualifiedDomainNames []string `pulumi:"fullyQualifiedDomainNames"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the access rule that is unique within a profile. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// Outbound rules phone number format.
	PhoneNumbers []string `pulumi:"phoneNumbers"`
	// The name of the NSP profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Inbound rules service tag names.
	ServiceTags []string `pulumi:"serviceTags"`
	// List of subscription ids
	Subscriptions []SubscriptionId `pulumi:"subscriptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NspAccessRule resource.
type NspAccessRuleArgs struct {
	// The name of the NSP access rule.
	AccessRuleName pulumi.StringPtrInput
	// Inbound address prefixes (IPv4/IPv6)
	AddressPrefixes pulumi.StringArrayInput
	// Direction that specifies whether the access rules is inbound/outbound.
	Direction pulumi.StringPtrInput
	// Outbound rules email address format.
	EmailAddresses pulumi.StringArrayInput
	// Outbound rules fully qualified domain name format.
	FullyQualifiedDomainNames pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the access rule that is unique within a profile. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput
	// Outbound rules phone number format.
	PhoneNumbers pulumi.StringArrayInput
	// The name of the NSP profile.
	ProfileName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Inbound rules service tag names.
	ServiceTags pulumi.StringArrayInput
	// List of subscription ids
	Subscriptions SubscriptionIdArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NspAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAccessRuleArgs)(nil)).Elem()
}

type NspAccessRuleInput interface {
	pulumi.Input

	ToNspAccessRuleOutput() NspAccessRuleOutput
	ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput
}

func (*NspAccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAccessRule)(nil)).Elem()
}

func (i *NspAccessRule) ToNspAccessRuleOutput() NspAccessRuleOutput {
	return i.ToNspAccessRuleOutputWithContext(context.Background())
}

func (i *NspAccessRule) ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspAccessRuleOutput)
}

func (i *NspAccessRule) ToOutput(ctx context.Context) pulumix.Output[*NspAccessRule] {
	return pulumix.Output[*NspAccessRule]{
		OutputState: i.ToNspAccessRuleOutputWithContext(ctx).OutputState,
	}
}

type NspAccessRuleOutput struct{ *pulumi.OutputState }

func (NspAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAccessRule)(nil)).Elem()
}

func (o NspAccessRuleOutput) ToNspAccessRuleOutput() NspAccessRuleOutput {
	return o
}

func (o NspAccessRuleOutput) ToNspAccessRuleOutputWithContext(ctx context.Context) NspAccessRuleOutput {
	return o
}

func (o NspAccessRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*NspAccessRule] {
	return pulumix.Output[*NspAccessRule]{
		OutputState: o.OutputState,
	}
}

// Inbound address prefixes (IPv4/IPv6)
func (o NspAccessRuleOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

// Direction that specifies whether the access rules is inbound/outbound.
func (o NspAccessRuleOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringPtrOutput { return v.Direction }).(pulumi.StringPtrOutput)
}

// Outbound rules email address format.
func (o NspAccessRuleOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Outbound rules fully qualified domain name format.
func (o NspAccessRuleOutput) FullyQualifiedDomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.FullyQualifiedDomainNames }).(pulumi.StringArrayOutput)
}

// Resource location.
func (o NspAccessRuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NspAccessRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Rule specified by the perimeter id.
func (o NspAccessRuleOutput) NetworkSecurityPerimeters() PerimeterBasedAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) PerimeterBasedAccessRuleResponseArrayOutput { return v.NetworkSecurityPerimeters }).(PerimeterBasedAccessRuleResponseArrayOutput)
}

// Outbound rules phone number format.
func (o NspAccessRuleOutput) PhoneNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.PhoneNumbers }).(pulumi.StringArrayOutput)
}

// The provisioning state of the scope assignment resource.
func (o NspAccessRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Inbound rules service tag names.
func (o NspAccessRuleOutput) ServiceTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringArrayOutput { return v.ServiceTags }).(pulumi.StringArrayOutput)
}

// List of subscription ids
func (o NspAccessRuleOutput) Subscriptions() SubscriptionIdResponseArrayOutput {
	return o.ApplyT(func(v *NspAccessRule) SubscriptionIdResponseArrayOutput { return v.Subscriptions }).(SubscriptionIdResponseArrayOutput)
}

// Resource tags.
func (o NspAccessRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NspAccessRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAccessRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspAccessRuleOutput{})
}
