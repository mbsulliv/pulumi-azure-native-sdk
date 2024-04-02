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

// Represents a catalog.
type Catalog struct {
	pulumi.CustomResourceState

	// Properties for an Azure DevOps catalog type.
	AdoGit GitCatalogResponsePtrOutput `pulumi:"adoGit"`
	// The connection state of the catalog.
	ConnectionState pulumi.StringOutput `pulumi:"connectionState"`
	// Properties for a GitHub catalog type.
	GitHub GitCatalogResponsePtrOutput `pulumi:"gitHub"`
	// When the catalog was last connected.
	LastConnectionTime pulumi.StringOutput `pulumi:"lastConnectionTime"`
	// Stats of the latest synchronization.
	LastSyncStats SyncStatsResponseOutput `pulumi:"lastSyncStats"`
	// When the catalog was last synced.
	LastSyncTime pulumi.StringOutput `pulumi:"lastSyncTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The synchronization state of the catalog.
	SyncState pulumi.StringOutput `pulumi:"syncState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCatalog registers a new resource with the given unique name, arguments, and options.
func NewCatalog(ctx *pulumi.Context,
	name string, args *CatalogArgs, opts ...pulumi.ResourceOption) (*Catalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevCenterName == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220901preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230101preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20230401:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20231001preview:Catalog"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20240201:Catalog"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Catalog
	err := ctx.RegisterResource("azure-native:devcenter/v20230801preview:Catalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCatalog gets an existing Catalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatalogState, opts ...pulumi.ResourceOption) (*Catalog, error) {
	var resource Catalog
	err := ctx.ReadResource("azure-native:devcenter/v20230801preview:Catalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Catalog resources.
type catalogState struct {
}

type CatalogState struct {
}

func (CatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogState)(nil)).Elem()
}

type catalogArgs struct {
	// Properties for an Azure DevOps catalog type.
	AdoGit *GitCatalog `pulumi:"adoGit"`
	// The name of the Catalog.
	CatalogName *string `pulumi:"catalogName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// Properties for a GitHub catalog type.
	GitHub *GitCatalog `pulumi:"gitHub"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Catalog resource.
type CatalogArgs struct {
	// Properties for an Azure DevOps catalog type.
	AdoGit GitCatalogPtrInput
	// The name of the Catalog.
	CatalogName pulumi.StringPtrInput
	// The name of the devcenter.
	DevCenterName pulumi.StringInput
	// Properties for a GitHub catalog type.
	GitHub GitCatalogPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (CatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catalogArgs)(nil)).Elem()
}

type CatalogInput interface {
	pulumi.Input

	ToCatalogOutput() CatalogOutput
	ToCatalogOutputWithContext(ctx context.Context) CatalogOutput
}

func (*Catalog) ElementType() reflect.Type {
	return reflect.TypeOf((**Catalog)(nil)).Elem()
}

func (i *Catalog) ToCatalogOutput() CatalogOutput {
	return i.ToCatalogOutputWithContext(context.Background())
}

func (i *Catalog) ToCatalogOutputWithContext(ctx context.Context) CatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatalogOutput)
}

type CatalogOutput struct{ *pulumi.OutputState }

func (CatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Catalog)(nil)).Elem()
}

func (o CatalogOutput) ToCatalogOutput() CatalogOutput {
	return o
}

func (o CatalogOutput) ToCatalogOutputWithContext(ctx context.Context) CatalogOutput {
	return o
}

// Properties for an Azure DevOps catalog type.
func (o CatalogOutput) AdoGit() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v *Catalog) GitCatalogResponsePtrOutput { return v.AdoGit }).(GitCatalogResponsePtrOutput)
}

// The connection state of the catalog.
func (o CatalogOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.ConnectionState }).(pulumi.StringOutput)
}

// Properties for a GitHub catalog type.
func (o CatalogOutput) GitHub() GitCatalogResponsePtrOutput {
	return o.ApplyT(func(v *Catalog) GitCatalogResponsePtrOutput { return v.GitHub }).(GitCatalogResponsePtrOutput)
}

// When the catalog was last connected.
func (o CatalogOutput) LastConnectionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.LastConnectionTime }).(pulumi.StringOutput)
}

// Stats of the latest synchronization.
func (o CatalogOutput) LastSyncStats() SyncStatsResponseOutput {
	return o.ApplyT(func(v *Catalog) SyncStatsResponseOutput { return v.LastSyncStats }).(SyncStatsResponseOutput)
}

// When the catalog was last synced.
func (o CatalogOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.LastSyncTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o CatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o CatalogOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The synchronization state of the catalog.
func (o CatalogOutput) SyncState() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.SyncState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CatalogOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Catalog) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Catalog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CatalogOutput{})
}
