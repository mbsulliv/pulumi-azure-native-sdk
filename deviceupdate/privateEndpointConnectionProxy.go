// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceupdate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Private endpoint connection proxy details.
// Azure REST API version: 2023-07-01. Prior API version in Azure Native 1.x: 2020-03-01-preview.
type PrivateEndpointConnectionProxy struct {
	pulumi.CustomResourceState

	// ETag from NRP.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the private endpoint connection proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Remote private endpoint details.
	RemotePrivateEndpoint RemotePrivateEndpointResponsePtrOutput `pulumi:"remotePrivateEndpoint"`
	// Operation status.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionProxy registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionProxyArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20220401preview:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221001:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20221201preview:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20230701:PrivateEndpointConnectionProxy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionProxy
	err := ctx.RegisterResource("azure-native:deviceupdate:PrivateEndpointConnectionProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionProxy gets an existing PrivateEndpointConnectionProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionProxyState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	var resource PrivateEndpointConnectionProxy
	err := ctx.ReadResource("azure-native:deviceupdate:PrivateEndpointConnectionProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionProxy resources.
type privateEndpointConnectionProxyState struct {
}

type PrivateEndpointConnectionProxyState struct {
}

func (PrivateEndpointConnectionProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyState)(nil)).Elem()
}

type privateEndpointConnectionProxyArgs struct {
	// Account name.
	AccountName string `pulumi:"accountName"`
	// The ID of the private endpoint connection proxy object.
	PrivateEndpointConnectionProxyId *string `pulumi:"privateEndpointConnectionProxyId"`
	// Remote private endpoint details.
	RemotePrivateEndpoint *RemotePrivateEndpoint `pulumi:"remotePrivateEndpoint"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Operation status.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a PrivateEndpointConnectionProxy resource.
type PrivateEndpointConnectionProxyArgs struct {
	// Account name.
	AccountName pulumi.StringInput
	// The ID of the private endpoint connection proxy object.
	PrivateEndpointConnectionProxyId pulumi.StringPtrInput
	// Remote private endpoint details.
	RemotePrivateEndpoint RemotePrivateEndpointPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Operation status.
	Status pulumi.StringPtrInput
}

func (PrivateEndpointConnectionProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyArgs)(nil)).Elem()
}

type PrivateEndpointConnectionProxyInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput
	ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput
}

func (*PrivateEndpointConnectionProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxy)(nil)).Elem()
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return i.ToPrivateEndpointConnectionProxyOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionProxyOutput)
}

func (i *PrivateEndpointConnectionProxy) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionProxy] {
	return pulumix.Output[*PrivateEndpointConnectionProxy]{
		OutputState: i.ToPrivateEndpointConnectionProxyOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointConnectionProxyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxy)(nil)).Elem()
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return o
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return o
}

func (o PrivateEndpointConnectionProxyOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionProxy] {
	return pulumix.Output[*PrivateEndpointConnectionProxy]{
		OutputState: o.OutputState,
	}
}

// ETag from NRP.
func (o PrivateEndpointConnectionProxyOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the private endpoint connection proxy resource.
func (o PrivateEndpointConnectionProxyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Remote private endpoint details.
func (o PrivateEndpointConnectionProxyOutput) RemotePrivateEndpoint() RemotePrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) RemotePrivateEndpointResponsePtrOutput {
		return v.RemotePrivateEndpoint
	}).(RemotePrivateEndpointResponsePtrOutput)
}

// Operation status.
func (o PrivateEndpointConnectionProxyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrivateEndpointConnectionProxyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionProxyOutput{})
}
