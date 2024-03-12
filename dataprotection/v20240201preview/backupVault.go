// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backup Vault Resource
type BackupVault struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Input Managed Identity Details
	Identity DppIdentityDetailsResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// BackupVaultResource properties
	Properties BackupVaultResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackupVault registers a new resource with the given unique name, arguments, and options.
func NewBackupVault(ctx *pulumi.Context,
	name string, args *BackupVaultArgs, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dataprotection:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210601preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221101preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221201:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230401preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230501:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230601preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230801preview:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20231101:BackupVault"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20231201:BackupVault"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupVault
	err := ctx.RegisterResource("azure-native:dataprotection/v20240201preview:BackupVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupVault gets an existing BackupVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupVaultState, opts ...pulumi.ResourceOption) (*BackupVault, error) {
	var resource BackupVault
	err := ctx.ReadResource("azure-native:dataprotection/v20240201preview:BackupVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupVault resources.
type backupVaultState struct {
}

type BackupVaultState struct {
}

func (BackupVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultState)(nil)).Elem()
}

type backupVaultArgs struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Input Managed Identity Details
	Identity *DppIdentityDetails `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// BackupVaultResource properties
	Properties BackupVaultType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the backup vault.
	VaultName *string `pulumi:"vaultName"`
}

// The set of arguments for constructing a BackupVault resource.
type BackupVaultArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Input Managed Identity Details
	Identity DppIdentityDetailsPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// BackupVaultResource properties
	Properties BackupVaultTypeInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the backup vault.
	VaultName pulumi.StringPtrInput
}

func (BackupVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupVaultArgs)(nil)).Elem()
}

type BackupVaultInput interface {
	pulumi.Input

	ToBackupVaultOutput() BackupVaultOutput
	ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput
}

func (*BackupVault) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (i *BackupVault) ToBackupVaultOutput() BackupVaultOutput {
	return i.ToBackupVaultOutputWithContext(context.Background())
}

func (i *BackupVault) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupVaultOutput)
}

type BackupVaultOutput struct{ *pulumi.OutputState }

func (BackupVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupVault)(nil)).Elem()
}

func (o BackupVaultOutput) ToBackupVaultOutput() BackupVaultOutput {
	return o
}

func (o BackupVaultOutput) ToBackupVaultOutputWithContext(ctx context.Context) BackupVaultOutput {
	return o
}

// Optional ETag.
func (o BackupVaultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Input Managed Identity Details
func (o BackupVaultOutput) Identity() DppIdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v *BackupVault) DppIdentityDetailsResponsePtrOutput { return v.Identity }).(DppIdentityDetailsResponsePtrOutput)
}

// Resource location.
func (o BackupVaultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name associated with the resource.
func (o BackupVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// BackupVaultResource properties
func (o BackupVaultOutput) Properties() BackupVaultResponseOutput {
	return o.ApplyT(func(v *BackupVault) BackupVaultResponseOutput { return v.Properties }).(BackupVaultResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BackupVaultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupVault) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BackupVaultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o BackupVaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupVault) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupVaultOutput{})
}
