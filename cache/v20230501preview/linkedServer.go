// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Response to put/get linked server (with properties) for Redis cache.
type LinkedServer struct {
	pulumi.CustomResourceState

	// The unchanging DNS name which will always point to current geo-primary cache among the linked redis caches for seamless Geo Failover experience.
	GeoReplicatedPrimaryHostName pulumi.StringOutput `pulumi:"geoReplicatedPrimaryHostName"`
	// Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId pulumi.StringOutput `pulumi:"linkedRedisCacheId"`
	// Location of the linked redis cache.
	LinkedRedisCacheLocation pulumi.StringOutput `pulumi:"linkedRedisCacheLocation"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The changing DNS name that resolves to the current geo-primary cache among the linked redis caches before or after the Geo Failover.
	PrimaryHostName pulumi.StringOutput `pulumi:"primaryHostName"`
	// Terminal state of the link between primary and secondary redis cache.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Role of the linked server.
	ServerRole pulumi.StringOutput `pulumi:"serverRole"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLinkedServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServer(ctx *pulumi.Context,
	name string, args *LinkedServerArgs, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedRedisCacheId == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheId'")
	}
	if args.LinkedRedisCacheLocation == nil {
		return nil, errors.New("invalid value for required argument 'LinkedRedisCacheLocation'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerRole == nil {
		return nil, errors.New("invalid value for required argument 'ServerRole'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220601:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230401:LinkedServer"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801:LinkedServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedServer
	err := ctx.RegisterResource("azure-native:cache/v20230501preview:LinkedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServer gets an existing LinkedServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServerState, opts ...pulumi.ResourceOption) (*LinkedServer, error) {
	var resource LinkedServer
	err := ctx.ReadResource("azure-native:cache/v20230501preview:LinkedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServer resources.
type linkedServerState struct {
}

type LinkedServerState struct {
}

func (LinkedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerState)(nil)).Elem()
}

type linkedServerArgs struct {
	// Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId string `pulumi:"linkedRedisCacheId"`
	// Location of the linked redis cache.
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	// The name of the linked server that is being added to the Redis cache.
	LinkedServerName *string `pulumi:"linkedServerName"`
	// The name of the Redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Role of the linked server.
	ServerRole ReplicationRole `pulumi:"serverRole"`
}

// The set of arguments for constructing a LinkedServer resource.
type LinkedServerArgs struct {
	// Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId pulumi.StringInput
	// Location of the linked redis cache.
	LinkedRedisCacheLocation pulumi.StringInput
	// The name of the linked server that is being added to the Redis cache.
	LinkedServerName pulumi.StringPtrInput
	// The name of the Redis cache.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Role of the linked server.
	ServerRole ReplicationRoleInput
}

func (LinkedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServerArgs)(nil)).Elem()
}

type LinkedServerInput interface {
	pulumi.Input

	ToLinkedServerOutput() LinkedServerOutput
	ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput
}

func (*LinkedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServer)(nil)).Elem()
}

func (i *LinkedServer) ToLinkedServerOutput() LinkedServerOutput {
	return i.ToLinkedServerOutputWithContext(context.Background())
}

func (i *LinkedServer) ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServerOutput)
}

func (i *LinkedServer) ToOutput(ctx context.Context) pulumix.Output[*LinkedServer] {
	return pulumix.Output[*LinkedServer]{
		OutputState: i.ToLinkedServerOutputWithContext(ctx).OutputState,
	}
}

type LinkedServerOutput struct{ *pulumi.OutputState }

func (LinkedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServer)(nil)).Elem()
}

func (o LinkedServerOutput) ToLinkedServerOutput() LinkedServerOutput {
	return o
}

func (o LinkedServerOutput) ToLinkedServerOutputWithContext(ctx context.Context) LinkedServerOutput {
	return o
}

func (o LinkedServerOutput) ToOutput(ctx context.Context) pulumix.Output[*LinkedServer] {
	return pulumix.Output[*LinkedServer]{
		OutputState: o.OutputState,
	}
}

// The unchanging DNS name which will always point to current geo-primary cache among the linked redis caches for seamless Geo Failover experience.
func (o LinkedServerOutput) GeoReplicatedPrimaryHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.GeoReplicatedPrimaryHostName }).(pulumi.StringOutput)
}

// Fully qualified resourceId of the linked redis cache.
func (o LinkedServerOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

// Location of the linked redis cache.
func (o LinkedServerOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

// The name of the resource
func (o LinkedServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The changing DNS name that resolves to the current geo-primary cache among the linked redis caches before or after the Geo Failover.
func (o LinkedServerOutput) PrimaryHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.PrimaryHostName }).(pulumi.StringOutput)
}

// Terminal state of the link between primary and secondary redis cache.
func (o LinkedServerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Role of the linked server.
func (o LinkedServerOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.ServerRole }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LinkedServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServerOutput{})
}
