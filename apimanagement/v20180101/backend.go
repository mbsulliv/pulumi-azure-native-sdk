// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backend details.
type Backend struct {
	pulumi.CustomResourceState

	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractResponsePtrOutput `pulumi:"credentials"`
	// Backend Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Backend Properties contract
	Properties BackendPropertiesResponseOutput `pulumi:"properties"`
	// Backend communication protocol.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractResponsePtrOutput `pulumi:"proxy"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Backend Title.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// Backend TLS Properties
	Tls BackendTlsPropertiesResponsePtrOutput `pulumi:"tls"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Runtime Url of the Backend.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewBackend registers a new resource with the given unique name, arguments, and options.
func NewBackend(ctx *pulumi.Context,
	name string, args *BackendArgs, opts ...pulumi.ResourceOption) (*Backend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Tls != nil {
		args.Tls = args.Tls.ToBackendTlsPropertiesPtrOutput().ApplyT(func(v *BackendTlsProperties) *BackendTlsProperties { return v.Defaults() }).(BackendTlsPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Backend"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Backend"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Backend
	err := ctx.RegisterResource("azure-native:apimanagement/v20180101:Backend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackend gets an existing Backend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackendState, opts ...pulumi.ResourceOption) (*Backend, error) {
	var resource Backend
	err := ctx.ReadResource("azure-native:apimanagement/v20180101:Backend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backend resources.
type backendState struct {
}

type BackendState struct {
}

func (BackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*backendState)(nil)).Elem()
}

type backendArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Backendid *string `pulumi:"backendid"`
	// Backend Credentials Contract Properties
	Credentials *BackendCredentialsContract `pulumi:"credentials"`
	// Backend Description.
	Description *string `pulumi:"description"`
	// Backend Properties contract
	Properties *BackendProperties `pulumi:"properties"`
	// Backend communication protocol.
	Protocol string `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy *BackendProxyContract `pulumi:"proxy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Backend Title.
	Title *string `pulumi:"title"`
	// Backend TLS Properties
	Tls *BackendTlsProperties `pulumi:"tls"`
	// Runtime Url of the Backend.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Backend resource.
type BackendArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Backendid pulumi.StringPtrInput
	// Backend Credentials Contract Properties
	Credentials BackendCredentialsContractPtrInput
	// Backend Description.
	Description pulumi.StringPtrInput
	// Backend Properties contract
	Properties BackendPropertiesPtrInput
	// Backend communication protocol.
	Protocol pulumi.StringInput
	// Backend Proxy Contract Properties
	Proxy BackendProxyContractPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Backend Title.
	Title pulumi.StringPtrInput
	// Backend TLS Properties
	Tls BackendTlsPropertiesPtrInput
	// Runtime Url of the Backend.
	Url pulumi.StringInput
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backendArgs)(nil)).Elem()
}

type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(ctx context.Context) BackendOutput
}

func (*Backend) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (i *Backend) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i *Backend) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

// Backend Credentials Contract Properties
func (o BackendOutput) Credentials() BackendCredentialsContractResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendCredentialsContractResponsePtrOutput { return v.Credentials }).(BackendCredentialsContractResponsePtrOutput)
}

// Backend Description.
func (o BackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o BackendOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Backend Properties contract
func (o BackendOutput) Properties() BackendPropertiesResponseOutput {
	return o.ApplyT(func(v *Backend) BackendPropertiesResponseOutput { return v.Properties }).(BackendPropertiesResponseOutput)
}

// Backend communication protocol.
func (o BackendOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Backend Proxy Contract Properties
func (o BackendOutput) Proxy() BackendProxyContractResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendProxyContractResponsePtrOutput { return v.Proxy }).(BackendProxyContractResponsePtrOutput)
}

// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
func (o BackendOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Backend Title.
func (o BackendOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// Backend TLS Properties
func (o BackendOutput) Tls() BackendTlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Backend) BackendTlsPropertiesResponsePtrOutput { return v.Tls }).(BackendTlsPropertiesResponsePtrOutput)
}

// Resource type for API Management resource.
func (o BackendOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Runtime Url of the Backend.
func (o BackendOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Backend) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackendOutput{})
}
