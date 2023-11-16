// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ResourceGuardProxyBaseResource object, used for response and request bodies for ResourceGuardProxy APIs
type DppResourceGuardProxy struct {
	pulumi.CustomResourceState

	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBaseResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDppResourceGuardProxy registers a new resource with the given unique name, arguments, and options.
func NewDppResourceGuardProxy(ctx *pulumi.Context,
	name string, args *DppResourceGuardProxyArgs, opts ...pulumi.ResourceOption) (*DppResourceGuardProxy, error) {
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
			Type: pulumi.String("azure-native:dataprotection:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221101preview:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230101:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230401preview:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230501:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230601preview:DppResourceGuardProxy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230801preview:DppResourceGuardProxy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DppResourceGuardProxy
	err := ctx.RegisterResource("azure-native:dataprotection/v20230801:DppResourceGuardProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDppResourceGuardProxy gets an existing DppResourceGuardProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDppResourceGuardProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DppResourceGuardProxyState, opts ...pulumi.ResourceOption) (*DppResourceGuardProxy, error) {
	var resource DppResourceGuardProxy
	err := ctx.ReadResource("azure-native:dataprotection/v20230801:DppResourceGuardProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DppResourceGuardProxy resources.
type dppResourceGuardProxyState struct {
}

type DppResourceGuardProxyState struct {
}

func (DppResourceGuardProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dppResourceGuardProxyState)(nil)).Elem()
}

type dppResourceGuardProxyArgs struct {
	// ResourceGuardProxyBaseResource properties
	Properties *ResourceGuardProxyBase `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// name of the resource guard proxy
	ResourceGuardProxyName *string `pulumi:"resourceGuardProxyName"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a DppResourceGuardProxy resource.
type DppResourceGuardProxyArgs struct {
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBasePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// name of the resource guard proxy
	ResourceGuardProxyName pulumi.StringPtrInput
	// The name of the backup vault.
	VaultName pulumi.StringInput
}

func (DppResourceGuardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dppResourceGuardProxyArgs)(nil)).Elem()
}

type DppResourceGuardProxyInput interface {
	pulumi.Input

	ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput
	ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput
}

func (*DppResourceGuardProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**DppResourceGuardProxy)(nil)).Elem()
}

func (i *DppResourceGuardProxy) ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput {
	return i.ToDppResourceGuardProxyOutputWithContext(context.Background())
}

func (i *DppResourceGuardProxy) ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DppResourceGuardProxyOutput)
}

type DppResourceGuardProxyOutput struct{ *pulumi.OutputState }

func (DppResourceGuardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DppResourceGuardProxy)(nil)).Elem()
}

func (o DppResourceGuardProxyOutput) ToDppResourceGuardProxyOutput() DppResourceGuardProxyOutput {
	return o
}

func (o DppResourceGuardProxyOutput) ToDppResourceGuardProxyOutputWithContext(ctx context.Context) DppResourceGuardProxyOutput {
	return o
}

// Resource name associated with the resource.
func (o DppResourceGuardProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardProxyBaseResource properties
func (o DppResourceGuardProxyOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) ResourceGuardProxyBaseResponseOutput { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DppResourceGuardProxyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o DppResourceGuardProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DppResourceGuardProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DppResourceGuardProxyOutput{})
}
