// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure SQL Database sync member.
type SyncMember struct {
	pulumi.CustomResourceState

	// Database name of the member database in the sync member.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// Database type of the sync member.
	DatabaseType pulumi.StringPtrOutput `pulumi:"databaseType"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint name of the sync member if use private link connection is enabled, for sync members in Azure.
	PrivateEndpointName pulumi.StringOutput `pulumi:"privateEndpointName"`
	// Server name of the member database in the sync member
	ServerName pulumi.StringPtrOutput `pulumi:"serverName"`
	// SQL Server database id of the sync member.
	SqlServerDatabaseId pulumi.StringPtrOutput `pulumi:"sqlServerDatabaseId"`
	// ARM resource id of the sync agent in the sync member.
	SyncAgentId pulumi.StringPtrOutput `pulumi:"syncAgentId"`
	// Sync direction of the sync member.
	SyncDirection pulumi.StringPtrOutput `pulumi:"syncDirection"`
	// ARM resource id of the sync member logical database, for sync members in Azure.
	SyncMemberAzureDatabaseResourceId pulumi.StringPtrOutput `pulumi:"syncMemberAzureDatabaseResourceId"`
	// Sync state of the sync member.
	SyncState pulumi.StringOutput `pulumi:"syncState"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether to use private link connection.
	UsePrivateLinkConnection pulumi.BoolPtrOutput `pulumi:"usePrivateLinkConnection"`
	// User name of the member database in the sync member.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewSyncMember registers a new resource with the given unique name, arguments, and options.
func NewSyncMember(ctx *pulumi.Context,
	name string, args *SyncMemberArgs, opts ...pulumi.ResourceOption) (*SyncMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.SyncGroupName == nil {
		return nil, errors.New("invalid value for required argument 'SyncGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20190601preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:SyncMember"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:SyncMember"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SyncMember
	err := ctx.RegisterResource("azure-native:sql/v20230801preview:SyncMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncMember gets an existing SyncMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncMemberState, opts ...pulumi.ResourceOption) (*SyncMember, error) {
	var resource SyncMember
	err := ctx.ReadResource("azure-native:sql/v20230801preview:SyncMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncMember resources.
type syncMemberState struct {
}

type SyncMemberState struct {
}

func (SyncMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncMemberState)(nil)).Elem()
}

type syncMemberArgs struct {
	// Database name of the member database in the sync member.
	DatabaseName string `pulumi:"databaseName"`
	// Database type of the sync member.
	DatabaseType *string `pulumi:"databaseType"`
	// Password of the member database in the sync member.
	Password *string `pulumi:"password"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Server name of the member database in the sync member
	ServerName string `pulumi:"serverName"`
	// SQL Server database id of the sync member.
	SqlServerDatabaseId *string `pulumi:"sqlServerDatabaseId"`
	// ARM resource id of the sync agent in the sync member.
	SyncAgentId *string `pulumi:"syncAgentId"`
	// Sync direction of the sync member.
	SyncDirection *string `pulumi:"syncDirection"`
	// The name of the sync group on which the sync member is hosted.
	SyncGroupName string `pulumi:"syncGroupName"`
	// ARM resource id of the sync member logical database, for sync members in Azure.
	SyncMemberAzureDatabaseResourceId *string `pulumi:"syncMemberAzureDatabaseResourceId"`
	// The name of the sync member.
	SyncMemberName *string `pulumi:"syncMemberName"`
	// Whether to use private link connection.
	UsePrivateLinkConnection *bool `pulumi:"usePrivateLinkConnection"`
	// User name of the member database in the sync member.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a SyncMember resource.
type SyncMemberArgs struct {
	// Database name of the member database in the sync member.
	DatabaseName pulumi.StringInput
	// Database type of the sync member.
	DatabaseType pulumi.StringPtrInput
	// Password of the member database in the sync member.
	Password pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Server name of the member database in the sync member
	ServerName pulumi.StringInput
	// SQL Server database id of the sync member.
	SqlServerDatabaseId pulumi.StringPtrInput
	// ARM resource id of the sync agent in the sync member.
	SyncAgentId pulumi.StringPtrInput
	// Sync direction of the sync member.
	SyncDirection pulumi.StringPtrInput
	// The name of the sync group on which the sync member is hosted.
	SyncGroupName pulumi.StringInput
	// ARM resource id of the sync member logical database, for sync members in Azure.
	SyncMemberAzureDatabaseResourceId pulumi.StringPtrInput
	// The name of the sync member.
	SyncMemberName pulumi.StringPtrInput
	// Whether to use private link connection.
	UsePrivateLinkConnection pulumi.BoolPtrInput
	// User name of the member database in the sync member.
	UserName pulumi.StringPtrInput
}

func (SyncMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncMemberArgs)(nil)).Elem()
}

type SyncMemberInput interface {
	pulumi.Input

	ToSyncMemberOutput() SyncMemberOutput
	ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput
}

func (*SyncMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncMember)(nil)).Elem()
}

func (i *SyncMember) ToSyncMemberOutput() SyncMemberOutput {
	return i.ToSyncMemberOutputWithContext(context.Background())
}

func (i *SyncMember) ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncMemberOutput)
}

type SyncMemberOutput struct{ *pulumi.OutputState }

func (SyncMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncMember)(nil)).Elem()
}

func (o SyncMemberOutput) ToSyncMemberOutput() SyncMemberOutput {
	return o
}

func (o SyncMemberOutput) ToSyncMemberOutputWithContext(ctx context.Context) SyncMemberOutput {
	return o
}

// Database name of the member database in the sync member.
func (o SyncMemberOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// Database type of the sync member.
func (o SyncMemberOutput) DatabaseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.DatabaseType }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o SyncMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint name of the sync member if use private link connection is enabled, for sync members in Azure.
func (o SyncMemberOutput) PrivateEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringOutput { return v.PrivateEndpointName }).(pulumi.StringOutput)
}

// Server name of the member database in the sync member
func (o SyncMemberOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.ServerName }).(pulumi.StringPtrOutput)
}

// SQL Server database id of the sync member.
func (o SyncMemberOutput) SqlServerDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.SqlServerDatabaseId }).(pulumi.StringPtrOutput)
}

// ARM resource id of the sync agent in the sync member.
func (o SyncMemberOutput) SyncAgentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.SyncAgentId }).(pulumi.StringPtrOutput)
}

// Sync direction of the sync member.
func (o SyncMemberOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.SyncDirection }).(pulumi.StringPtrOutput)
}

// ARM resource id of the sync member logical database, for sync members in Azure.
func (o SyncMemberOutput) SyncMemberAzureDatabaseResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.SyncMemberAzureDatabaseResourceId }).(pulumi.StringPtrOutput)
}

// Sync state of the sync member.
func (o SyncMemberOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringOutput { return v.SyncState }).(pulumi.StringOutput)
}

// Resource type.
func (o SyncMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether to use private link connection.
func (o SyncMemberOutput) UsePrivateLinkConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.BoolPtrOutput { return v.UsePrivateLinkConnection }).(pulumi.BoolPtrOutput)
}

// User name of the member database in the sync member.
func (o SyncMemberOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncMember) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SyncMemberOutput{})
}
