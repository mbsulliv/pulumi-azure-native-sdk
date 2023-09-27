// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A server trust group.
type ServerTrustGroup struct {
	pulumi.CustomResourceState

	// Group members information for the server trust group.
	GroupMembers ServerInfoResponseArrayOutput `pulumi:"groupMembers"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Trust scope of the server trust group.
	TrustScopes pulumi.StringArrayOutput `pulumi:"trustScopes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerTrustGroup registers a new resource with the given unique name, arguments, and options.
func NewServerTrustGroup(ctx *pulumi.Context,
	name string, args *ServerTrustGroupArgs, opts ...pulumi.ResourceOption) (*ServerTrustGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupMembers == nil {
		return nil, errors.New("invalid value for required argument 'GroupMembers'")
	}
	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrustScopes == nil {
		return nil, errors.New("invalid value for required argument 'TrustScopes'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerTrustGroup"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ServerTrustGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerTrustGroup
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:ServerTrustGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerTrustGroup gets an existing ServerTrustGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerTrustGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerTrustGroupState, opts ...pulumi.ResourceOption) (*ServerTrustGroup, error) {
	var resource ServerTrustGroup
	err := ctx.ReadResource("azure-native:sql/v20221101preview:ServerTrustGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerTrustGroup resources.
type serverTrustGroupState struct {
}

type ServerTrustGroupState struct {
}

func (ServerTrustGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustGroupState)(nil)).Elem()
}

type serverTrustGroupArgs struct {
	// Group members information for the server trust group.
	GroupMembers []ServerInfo `pulumi:"groupMembers"`
	// The name of the region where the resource is located.
	LocationName string `pulumi:"locationName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server trust group.
	ServerTrustGroupName *string `pulumi:"serverTrustGroupName"`
	// Trust scope of the server trust group.
	TrustScopes []string `pulumi:"trustScopes"`
}

// The set of arguments for constructing a ServerTrustGroup resource.
type ServerTrustGroupArgs struct {
	// Group members information for the server trust group.
	GroupMembers ServerInfoArrayInput
	// The name of the region where the resource is located.
	LocationName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server trust group.
	ServerTrustGroupName pulumi.StringPtrInput
	// Trust scope of the server trust group.
	TrustScopes pulumi.StringArrayInput
}

func (ServerTrustGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustGroupArgs)(nil)).Elem()
}

type ServerTrustGroupInput interface {
	pulumi.Input

	ToServerTrustGroupOutput() ServerTrustGroupOutput
	ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput
}

func (*ServerTrustGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustGroup)(nil)).Elem()
}

func (i *ServerTrustGroup) ToServerTrustGroupOutput() ServerTrustGroupOutput {
	return i.ToServerTrustGroupOutputWithContext(context.Background())
}

func (i *ServerTrustGroup) ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTrustGroupOutput)
}

func (i *ServerTrustGroup) ToOutput(ctx context.Context) pulumix.Output[*ServerTrustGroup] {
	return pulumix.Output[*ServerTrustGroup]{
		OutputState: i.ToServerTrustGroupOutputWithContext(ctx).OutputState,
	}
}

type ServerTrustGroupOutput struct{ *pulumi.OutputState }

func (ServerTrustGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustGroup)(nil)).Elem()
}

func (o ServerTrustGroupOutput) ToServerTrustGroupOutput() ServerTrustGroupOutput {
	return o
}

func (o ServerTrustGroupOutput) ToServerTrustGroupOutputWithContext(ctx context.Context) ServerTrustGroupOutput {
	return o
}

func (o ServerTrustGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerTrustGroup] {
	return pulumix.Output[*ServerTrustGroup]{
		OutputState: o.OutputState,
	}
}

// Group members information for the server trust group.
func (o ServerTrustGroupOutput) GroupMembers() ServerInfoResponseArrayOutput {
	return o.ApplyT(func(v *ServerTrustGroup) ServerInfoResponseArrayOutput { return v.GroupMembers }).(ServerInfoResponseArrayOutput)
}

// Resource name.
func (o ServerTrustGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Trust scope of the server trust group.
func (o ServerTrustGroupOutput) TrustScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringArrayOutput { return v.TrustScopes }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o ServerTrustGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerTrustGroupOutput{})
}
