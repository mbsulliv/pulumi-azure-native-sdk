// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A hostname binding object.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2020-10-01, 2023-01-01.
type WebAppHostNameBindingSlot struct {
	pulumi.CustomResourceState

	// Azure resource name.
	AzureResourceName pulumi.StringPtrOutput `pulumi:"azureResourceName"`
	// Azure resource type.
	AzureResourceType pulumi.StringPtrOutput `pulumi:"azureResourceType"`
	// Custom DNS record type.
	CustomHostNameDnsRecordType pulumi.StringPtrOutput `pulumi:"customHostNameDnsRecordType"`
	// Fully qualified ARM domain resource URI.
	DomainId pulumi.StringPtrOutput `pulumi:"domainId"`
	// Hostname type.
	HostNameType pulumi.StringPtrOutput `pulumi:"hostNameType"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// App Service app name.
	SiteName pulumi.StringPtrOutput `pulumi:"siteName"`
	// SSL type
	SslState pulumi.StringPtrOutput `pulumi:"sslState"`
	// SSL certificate thumbprint
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIP pulumi.StringOutput `pulumi:"virtualIP"`
}

// NewWebAppHostNameBindingSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppHostNameBindingSlot(ctx *pulumi.Context,
	name string, args *WebAppHostNameBindingSlotArgs, opts ...pulumi.ResourceOption) (*WebAppHostNameBindingSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppHostNameBindingSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppHostNameBindingSlot
	err := ctx.RegisterResource("azure-native:web:WebAppHostNameBindingSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppHostNameBindingSlot gets an existing WebAppHostNameBindingSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppHostNameBindingSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHostNameBindingSlotState, opts ...pulumi.ResourceOption) (*WebAppHostNameBindingSlot, error) {
	var resource WebAppHostNameBindingSlot
	err := ctx.ReadResource("azure-native:web:WebAppHostNameBindingSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppHostNameBindingSlot resources.
type webAppHostNameBindingSlotState struct {
}

type WebAppHostNameBindingSlotState struct {
}

func (WebAppHostNameBindingSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingSlotState)(nil)).Elem()
}

type webAppHostNameBindingSlotArgs struct {
	// Azure resource name.
	AzureResourceName *string `pulumi:"azureResourceName"`
	// Azure resource type.
	AzureResourceType *AzureResourceType `pulumi:"azureResourceType"`
	// Custom DNS record type.
	CustomHostNameDnsRecordType *CustomHostNameDnsRecordType `pulumi:"customHostNameDnsRecordType"`
	// Fully qualified ARM domain resource URI.
	DomainId *string `pulumi:"domainId"`
	// Hostname in the hostname binding.
	HostName *string `pulumi:"hostName"`
	// Hostname type.
	HostNameType *HostNameType `pulumi:"hostNameType"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// App Service app name.
	SiteName *string `pulumi:"siteName"`
	// Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
	Slot string `pulumi:"slot"`
	// SSL type
	SslState *SslState `pulumi:"sslState"`
	// SSL certificate thumbprint
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a WebAppHostNameBindingSlot resource.
type WebAppHostNameBindingSlotArgs struct {
	// Azure resource name.
	AzureResourceName pulumi.StringPtrInput
	// Azure resource type.
	AzureResourceType AzureResourceTypePtrInput
	// Custom DNS record type.
	CustomHostNameDnsRecordType CustomHostNameDnsRecordTypePtrInput
	// Fully qualified ARM domain resource URI.
	DomainId pulumi.StringPtrInput
	// Hostname in the hostname binding.
	HostName pulumi.StringPtrInput
	// Hostname type.
	HostNameType HostNameTypePtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// App Service app name.
	SiteName pulumi.StringPtrInput
	// Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.
	Slot pulumi.StringInput
	// SSL type
	SslState SslStatePtrInput
	// SSL certificate thumbprint
	Thumbprint pulumi.StringPtrInput
}

func (WebAppHostNameBindingSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingSlotArgs)(nil)).Elem()
}

type WebAppHostNameBindingSlotInput interface {
	pulumi.Input

	ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput
	ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput
}

func (*WebAppHostNameBindingSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBindingSlot)(nil)).Elem()
}

func (i *WebAppHostNameBindingSlot) ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput {
	return i.ToWebAppHostNameBindingSlotOutputWithContext(context.Background())
}

func (i *WebAppHostNameBindingSlot) ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHostNameBindingSlotOutput)
}

type WebAppHostNameBindingSlotOutput struct{ *pulumi.OutputState }

func (WebAppHostNameBindingSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBindingSlot)(nil)).Elem()
}

func (o WebAppHostNameBindingSlotOutput) ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput {
	return o
}

func (o WebAppHostNameBindingSlotOutput) ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput {
	return o
}

// Azure resource name.
func (o WebAppHostNameBindingSlotOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

// Azure resource type.
func (o WebAppHostNameBindingSlotOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

// Custom DNS record type.
func (o WebAppHostNameBindingSlotOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

// Fully qualified ARM domain resource URI.
func (o WebAppHostNameBindingSlotOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

// Hostname type.
func (o WebAppHostNameBindingSlotOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.HostNameType }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppHostNameBindingSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppHostNameBindingSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// App Service app name.
func (o WebAppHostNameBindingSlotOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.SiteName }).(pulumi.StringPtrOutput)
}

// SSL type
func (o WebAppHostNameBindingSlotOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.SslState }).(pulumi.StringPtrOutput)
}

// SSL certificate thumbprint
func (o WebAppHostNameBindingSlotOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppHostNameBindingSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Virtual IP address assigned to the hostname if IP based SSL is enabled.
func (o WebAppHostNameBindingSlotOutput) VirtualIP() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.VirtualIP }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHostNameBindingSlotOutput{})
}
