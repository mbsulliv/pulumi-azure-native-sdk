// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database sync agent.
type SyncAgent struct {
	pulumi.CustomResourceState

	// Expiration time of the sync agent version.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// If the sync agent version is up to date.
	IsUpToDate pulumi.BoolOutput `pulumi:"isUpToDate"`
	// Last alive time of the sync agent.
	LastAliveTime pulumi.StringOutput `pulumi:"lastAliveTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the sync agent.
	State pulumi.StringOutput `pulumi:"state"`
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId pulumi.StringPtrOutput `pulumi:"syncDatabaseId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the sync agent.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewSyncAgent registers a new resource with the given unique name, arguments, and options.
func NewSyncAgent(ctx *pulumi.Context,
	name string, args *SyncAgentArgs, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:SyncAgent"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:SyncAgent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SyncAgent
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:SyncAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncAgent gets an existing SyncAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncAgentState, opts ...pulumi.ResourceOption) (*SyncAgent, error) {
	var resource SyncAgent
	err := ctx.ReadResource("azure-native:sql/v20221101preview:SyncAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncAgent resources.
type syncAgentState struct {
}

type SyncAgentState struct {
}

func (SyncAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentState)(nil)).Elem()
}

type syncAgentArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server on which the sync agent is hosted.
	ServerName string `pulumi:"serverName"`
	// The name of the sync agent.
	SyncAgentName *string `pulumi:"syncAgentName"`
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId *string `pulumi:"syncDatabaseId"`
}

// The set of arguments for constructing a SyncAgent resource.
type SyncAgentArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server on which the sync agent is hosted.
	ServerName pulumi.StringInput
	// The name of the sync agent.
	SyncAgentName pulumi.StringPtrInput
	// ARM resource id of the sync database in the sync agent.
	SyncDatabaseId pulumi.StringPtrInput
}

func (SyncAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncAgentArgs)(nil)).Elem()
}

type SyncAgentInput interface {
	pulumi.Input

	ToSyncAgentOutput() SyncAgentOutput
	ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput
}

func (*SyncAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAgent)(nil)).Elem()
}

func (i *SyncAgent) ToSyncAgentOutput() SyncAgentOutput {
	return i.ToSyncAgentOutputWithContext(context.Background())
}

func (i *SyncAgent) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncAgentOutput)
}

type SyncAgentOutput struct{ *pulumi.OutputState }

func (SyncAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncAgent)(nil)).Elem()
}

func (o SyncAgentOutput) ToSyncAgentOutput() SyncAgentOutput {
	return o
}

func (o SyncAgentOutput) ToSyncAgentOutputWithContext(ctx context.Context) SyncAgentOutput {
	return o
}

// Expiration time of the sync agent version.
func (o SyncAgentOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.ExpiryTime }).(pulumi.StringOutput)
}

// If the sync agent version is up to date.
func (o SyncAgentOutput) IsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.BoolOutput { return v.IsUpToDate }).(pulumi.BoolOutput)
}

// Last alive time of the sync agent.
func (o SyncAgentOutput) LastAliveTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.LastAliveTime }).(pulumi.StringOutput)
}

// Resource name.
func (o SyncAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the sync agent.
func (o SyncAgentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ARM resource id of the sync database in the sync agent.
func (o SyncAgentOutput) SyncDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringPtrOutput { return v.SyncDatabaseId }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o SyncAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the sync agent.
func (o SyncAgentOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncAgent) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncAgentOutput{})
}
