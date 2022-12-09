// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceGuardProxy struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBaseResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceGuardProxy registers a new resource with the given unique name, arguments, and options.
func NewResourceGuardProxy(ctx *pulumi.Context,
	name string, args *ResourceGuardProxyArgs, opts ...pulumi.ResourceOption) (*ResourceGuardProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:ResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:ResourceGuardProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGuardProxy
	err := ctx.RegisterResource("azure-native:recoveryservices/v20211201:ResourceGuardProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGuardProxy gets an existing ResourceGuardProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGuardProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGuardProxyState, opts ...pulumi.ResourceOption) (*ResourceGuardProxy, error) {
	var resource ResourceGuardProxy
	err := ctx.ReadResource("azure-native:recoveryservices/v20211201:ResourceGuardProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGuardProxy resources.
type resourceGuardProxyState struct {
}

type ResourceGuardProxyState struct {
}

func (ResourceGuardProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardProxyState)(nil)).Elem()
}

type resourceGuardProxyArgs struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// ResourceGuardProxyBaseResource properties
	Properties *ResourceGuardProxyBase `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ResourceGuardProxyName *string `pulumi:"resourceGuardProxyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a ResourceGuardProxy resource.
type ResourceGuardProxyArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBasePtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      pulumi.StringInput
	ResourceGuardProxyName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (ResourceGuardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGuardProxyArgs)(nil)).Elem()
}

type ResourceGuardProxyInput interface {
	pulumi.Input

	ToResourceGuardProxyOutput() ResourceGuardProxyOutput
	ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput
}

func (*ResourceGuardProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxy)(nil)).Elem()
}

func (i *ResourceGuardProxy) ToResourceGuardProxyOutput() ResourceGuardProxyOutput {
	return i.ToResourceGuardProxyOutputWithContext(context.Background())
}

func (i *ResourceGuardProxy) ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyOutput)
}

type ResourceGuardProxyOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxy)(nil)).Elem()
}

func (o ResourceGuardProxyOutput) ToResourceGuardProxyOutput() ResourceGuardProxyOutput {
	return o
}

func (o ResourceGuardProxyOutput) ToResourceGuardProxyOutputWithContext(ctx context.Context) ResourceGuardProxyOutput {
	return o
}

// Optional ETag.
func (o ResourceGuardProxyOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ResourceGuardProxyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o ResourceGuardProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardProxyBaseResource properties
func (o ResourceGuardProxyOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) ResourceGuardProxyBaseResponseOutput { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

// Resource tags.
func (o ResourceGuardProxyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o ResourceGuardProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGuardProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGuardProxyOutput{})
}
