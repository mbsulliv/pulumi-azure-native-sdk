// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Security Partner Provider resource.
type SecurityPartnerProvider struct {
	pulumi.CustomResourceState

	// The connection status with the Security Partner Provider.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the Security Partner Provider resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The security provider name.
	SecurityProviderName pulumi.StringPtrOutput `pulumi:"securityProviderName"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The virtualHub to which the Security Partner Provider belongs.
	VirtualHub SubResourceResponsePtrOutput `pulumi:"virtualHub"`
}

// NewSecurityPartnerProvider registers a new resource with the given unique name, arguments, and options.
func NewSecurityPartnerProvider(ctx *pulumi.Context,
	name string, args *SecurityPartnerProviderArgs, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:SecurityPartnerProvider"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:SecurityPartnerProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecurityPartnerProvider
	err := ctx.RegisterResource("azure-native:network/v20230401:SecurityPartnerProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPartnerProvider gets an existing SecurityPartnerProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPartnerProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPartnerProviderState, opts ...pulumi.ResourceOption) (*SecurityPartnerProvider, error) {
	var resource SecurityPartnerProvider
	err := ctx.ReadResource("azure-native:network/v20230401:SecurityPartnerProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPartnerProvider resources.
type securityPartnerProviderState struct {
}

type SecurityPartnerProviderState struct {
}

func (SecurityPartnerProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderState)(nil)).Elem()
}

type securityPartnerProviderArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Security Partner Provider.
	SecurityPartnerProviderName *string `pulumi:"securityPartnerProviderName"`
	// The security provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The virtualHub to which the Security Partner Provider belongs.
	VirtualHub *SubResource `pulumi:"virtualHub"`
}

// The set of arguments for constructing a SecurityPartnerProvider resource.
type SecurityPartnerProviderArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Security Partner Provider.
	SecurityPartnerProviderName pulumi.StringPtrInput
	// The security provider name.
	SecurityProviderName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The virtualHub to which the Security Partner Provider belongs.
	VirtualHub SubResourcePtrInput
}

func (SecurityPartnerProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPartnerProviderArgs)(nil)).Elem()
}

type SecurityPartnerProviderInput interface {
	pulumi.Input

	ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput
	ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput
}

func (*SecurityPartnerProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPartnerProvider)(nil)).Elem()
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return i.ToSecurityPartnerProviderOutputWithContext(context.Background())
}

func (i *SecurityPartnerProvider) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPartnerProviderOutput)
}

type SecurityPartnerProviderOutput struct{ *pulumi.OutputState }

func (SecurityPartnerProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityPartnerProvider)(nil)).Elem()
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutput() SecurityPartnerProviderOutput {
	return o
}

func (o SecurityPartnerProviderOutput) ToSecurityPartnerProviderOutputWithContext(ctx context.Context) SecurityPartnerProviderOutput {
	return o
}

// The connection status with the Security Partner Provider.
func (o SecurityPartnerProviderOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SecurityPartnerProviderOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o SecurityPartnerProviderOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o SecurityPartnerProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the Security Partner Provider resource.
func (o SecurityPartnerProviderOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The security provider name.
func (o SecurityPartnerProviderOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringPtrOutput { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o SecurityPartnerProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o SecurityPartnerProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The virtualHub to which the Security Partner Provider belongs.
func (o SecurityPartnerProviderOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *SecurityPartnerProvider) SubResourceResponsePtrOutput { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityPartnerProviderOutput{})
}
