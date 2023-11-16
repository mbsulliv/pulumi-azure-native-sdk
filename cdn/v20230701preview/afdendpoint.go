// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Front Door endpoint is the entity within a Azure Front Door profile containing configuration information such as origin, protocol, content caching and delivery behavior. The AzureFrontDoor endpoint uses the URL format <endpointname>.azureedge.net.
type AFDEndpoint struct {
	pulumi.CustomResourceState

	// Indicates the endpoint name reuse scope. The default value is TenantReuse.
	AutoGeneratedDomainNameLabelScope pulumi.StringPtrOutput `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeploymentStatus                  pulumi.StringOutput    `pulumi:"deploymentStatus"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the profile which holds the endpoint.
	ProfileName pulumi.StringOutput `pulumi:"profileName"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAFDEndpoint registers a new resource with the given unique name, arguments, and options.
func NewAFDEndpoint(ctx *pulumi.Context,
	name string, args *AFDEndpointArgs, opts ...pulumi.ResourceOption) (*AFDEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:AFDEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:AFDEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AFDEndpoint
	err := ctx.RegisterResource("azure-native:cdn/v20230701preview:AFDEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAFDEndpoint gets an existing AFDEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAFDEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDEndpointState, opts ...pulumi.ResourceOption) (*AFDEndpoint, error) {
	var resource AFDEndpoint
	err := ctx.ReadResource("azure-native:cdn/v20230701preview:AFDEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AFDEndpoint resources.
type afdendpointState struct {
}

type AFDEndpointState struct {
}

func (AFDEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdendpointState)(nil)).Elem()
}

type afdendpointArgs struct {
	// Indicates the endpoint name reuse scope. The default value is TenantReuse.
	AutoGeneratedDomainNameLabelScope *string `pulumi:"autoGeneratedDomainNameLabelScope"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName *string `pulumi:"endpointName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AFDEndpoint resource.
type AFDEndpointArgs struct {
	// Indicates the endpoint name reuse scope. The default value is TenantReuse.
	AutoGeneratedDomainNameLabelScope pulumi.StringPtrInput
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AFDEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdendpointArgs)(nil)).Elem()
}

type AFDEndpointInput interface {
	pulumi.Input

	ToAFDEndpointOutput() AFDEndpointOutput
	ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput
}

func (*AFDEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDEndpoint)(nil)).Elem()
}

func (i *AFDEndpoint) ToAFDEndpointOutput() AFDEndpointOutput {
	return i.ToAFDEndpointOutputWithContext(context.Background())
}

func (i *AFDEndpoint) ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDEndpointOutput)
}

type AFDEndpointOutput struct{ *pulumi.OutputState }

func (AFDEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDEndpoint)(nil)).Elem()
}

func (o AFDEndpointOutput) ToAFDEndpointOutput() AFDEndpointOutput {
	return o
}

func (o AFDEndpointOutput) ToAFDEndpointOutputWithContext(ctx context.Context) AFDEndpointOutput {
	return o
}

// Indicates the endpoint name reuse scope. The default value is TenantReuse.
func (o AFDEndpointOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringPtrOutput { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

func (o AFDEndpointOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
func (o AFDEndpointOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
func (o AFDEndpointOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Resource location.
func (o AFDEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o AFDEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the profile which holds the endpoint.
func (o AFDEndpointOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.ProfileName }).(pulumi.StringOutput)
}

// Provisioning status
func (o AFDEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Read only system data
func (o AFDEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AFDEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o AFDEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDEndpointOutput{})
}
