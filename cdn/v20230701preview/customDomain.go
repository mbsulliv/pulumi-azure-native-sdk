// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type CustomDomain struct {
	pulumi.CustomResourceState

	// Certificate parameters for securing custom HTTPS
	CustomHttpsParameters pulumi.AnyOutput `pulumi:"customHttpsParameters"`
	// Provisioning status of the custom domain.
	CustomHttpsProvisioningState pulumi.StringOutput `pulumi:"customHttpsProvisioningState"`
	// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
	CustomHttpsProvisioningSubstate pulumi.StringOutput `pulumi:"customHttpsProvisioningSubstate"`
	// The host name of the custom domain. Must be a domain name.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning status of Custom Https of the custom domain.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the custom domain.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
	ValidationData pulumi.StringPtrOutput `pulumi:"validationData"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:CustomDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomDomain
	err := ctx.RegisterResource("azure-native:cdn/v20230701preview:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("azure-native:cdn/v20230701preview:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
}

type CustomDomainState struct {
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// Name of the custom domain within an endpoint.
	CustomDomainName *string `pulumi:"customDomainName"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// The host name of the custom domain. Must be a domain name.
	HostName string `pulumi:"hostName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// Name of the custom domain within an endpoint.
	CustomDomainName pulumi.StringPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// The host name of the custom domain. Must be a domain name.
	HostName pulumi.StringInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i *CustomDomain) ToOutput(ctx context.Context) pulumix.Output[*CustomDomain] {
	return pulumix.Output[*CustomDomain]{
		OutputState: i.ToCustomDomainOutputWithContext(ctx).OutputState,
	}
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomDomain] {
	return pulumix.Output[*CustomDomain]{
		OutputState: o.OutputState,
	}
}

// Certificate parameters for securing custom HTTPS
func (o CustomDomainOutput) CustomHttpsParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.AnyOutput { return v.CustomHttpsParameters }).(pulumi.AnyOutput)
}

// Provisioning status of the custom domain.
func (o CustomDomainOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
func (o CustomDomainOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

// The host name of the custom domain. Must be a domain name.
func (o CustomDomainOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Resource name.
func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning status of Custom Https of the custom domain.
func (o CustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the custom domain.
func (o CustomDomainOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// Read only system data
func (o CustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o CustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
func (o CustomDomainOutput) ValidationData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringPtrOutput { return v.ValidationData }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
}
