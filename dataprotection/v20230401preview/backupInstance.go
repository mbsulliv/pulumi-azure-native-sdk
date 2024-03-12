// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// BackupInstance Resource
type BackupInstance struct {
	pulumi.CustomResourceState

	// Proxy Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// BackupInstanceResource properties
	Properties BackupInstanceResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Proxy Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Proxy Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackupInstance registers a new resource with the given unique name, arguments, and options.
func NewBackupInstance(ctx *pulumi.Context,
	name string, args *BackupInstanceArgs, opts ...pulumi.ResourceOption) (*BackupInstance, error) {
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
			Type: pulumi.String("azure-native:dataprotection:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210601preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210701:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211001preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20211201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220201preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220301:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220331preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220401:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220501:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20220901preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221001preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221101preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20221201:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230501:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230601preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20230801preview:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20231101:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20231201:BackupInstance"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20240201preview:BackupInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupInstance
	err := ctx.RegisterResource("azure-native:dataprotection/v20230401preview:BackupInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupInstance gets an existing BackupInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupInstanceState, opts ...pulumi.ResourceOption) (*BackupInstance, error) {
	var resource BackupInstance
	err := ctx.ReadResource("azure-native:dataprotection/v20230401preview:BackupInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupInstance resources.
type backupInstanceState struct {
}

type BackupInstanceState struct {
}

func (BackupInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceState)(nil)).Elem()
}

type backupInstanceArgs struct {
	// The name of the backup instance.
	BackupInstanceName *string `pulumi:"backupInstanceName"`
	// BackupInstanceResource properties
	Properties *BackupInstanceType `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Proxy Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a BackupInstance resource.
type BackupInstanceArgs struct {
	// The name of the backup instance.
	BackupInstanceName pulumi.StringPtrInput
	// BackupInstanceResource properties
	Properties BackupInstanceTypePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Proxy Resource tags.
	Tags pulumi.StringMapInput
	// The name of the backup vault.
	VaultName pulumi.StringInput
}

func (BackupInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupInstanceArgs)(nil)).Elem()
}

type BackupInstanceInput interface {
	pulumi.Input

	ToBackupInstanceOutput() BackupInstanceOutput
	ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput
}

func (*BackupInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstance)(nil)).Elem()
}

func (i *BackupInstance) ToBackupInstanceOutput() BackupInstanceOutput {
	return i.ToBackupInstanceOutputWithContext(context.Background())
}

func (i *BackupInstance) ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupInstanceOutput)
}

type BackupInstanceOutput struct{ *pulumi.OutputState }

func (BackupInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupInstance)(nil)).Elem()
}

func (o BackupInstanceOutput) ToBackupInstanceOutput() BackupInstanceOutput {
	return o
}

func (o BackupInstanceOutput) ToBackupInstanceOutputWithContext(ctx context.Context) BackupInstanceOutput {
	return o
}

// Proxy Resource name associated with the resource.
func (o BackupInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// BackupInstanceResource properties
func (o BackupInstanceOutput) Properties() BackupInstanceResponseOutput {
	return o.ApplyT(func(v *BackupInstance) BackupInstanceResponseOutput { return v.Properties }).(BackupInstanceResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o BackupInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BackupInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Proxy Resource tags.
func (o BackupInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BackupInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Proxy Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o BackupInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupInstanceOutput{})
}
