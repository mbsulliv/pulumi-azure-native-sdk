// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a scope map for a container registry.
// Azure REST API version: 2022-12-01. Prior API version in Azure Native 1.x: 2020-11-01-preview.
//
// Other available API versions: 2023-01-01-preview, 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview.
type ScopeMap struct {
	pulumi.CustomResourceState

	// The list of scoped permissions for registry artifacts.
	// E.g. repositories/repository-name/content/read,
	// repositories/repository-name/metadata/write
	Actions pulumi.StringArrayOutput `pulumi:"actions"`
	// The creation date of scope map.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The user friendly description of the scope map.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewScopeMap registers a new resource with the given unique name, arguments, and options.
func NewScopeMap(ctx *pulumi.Context,
	name string, args *ScopeMapArgs, opts ...pulumi.ResourceOption) (*ScopeMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20221201:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230701:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:ScopeMap"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:ScopeMap"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScopeMap
	err := ctx.RegisterResource("azure-native:containerregistry:ScopeMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScopeMap gets an existing ScopeMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScopeMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScopeMapState, opts ...pulumi.ResourceOption) (*ScopeMap, error) {
	var resource ScopeMap
	err := ctx.ReadResource("azure-native:containerregistry:ScopeMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScopeMap resources.
type scopeMapState struct {
}

type ScopeMapState struct {
}

func (ScopeMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeMapState)(nil)).Elem()
}

type scopeMapArgs struct {
	// The list of scoped permissions for registry artifacts.
	// E.g. repositories/repository-name/content/read,
	// repositories/repository-name/metadata/write
	Actions []string `pulumi:"actions"`
	// The user friendly description of the scope map.
	Description *string `pulumi:"description"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the scope map.
	ScopeMapName *string `pulumi:"scopeMapName"`
}

// The set of arguments for constructing a ScopeMap resource.
type ScopeMapArgs struct {
	// The list of scoped permissions for registry artifacts.
	// E.g. repositories/repository-name/content/read,
	// repositories/repository-name/metadata/write
	Actions pulumi.StringArrayInput
	// The user friendly description of the scope map.
	Description pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the scope map.
	ScopeMapName pulumi.StringPtrInput
}

func (ScopeMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scopeMapArgs)(nil)).Elem()
}

type ScopeMapInput interface {
	pulumi.Input

	ToScopeMapOutput() ScopeMapOutput
	ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput
}

func (*ScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeMap)(nil)).Elem()
}

func (i *ScopeMap) ToScopeMapOutput() ScopeMapOutput {
	return i.ToScopeMapOutputWithContext(context.Background())
}

func (i *ScopeMap) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeMapOutput)
}

type ScopeMapOutput struct{ *pulumi.OutputState }

func (ScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeMap)(nil)).Elem()
}

func (o ScopeMapOutput) ToScopeMapOutput() ScopeMapOutput {
	return o
}

func (o ScopeMapOutput) ToScopeMapOutputWithContext(ctx context.Context) ScopeMapOutput {
	return o
}

// The list of scoped permissions for registry artifacts.
// E.g. repositories/repository-name/content/read,
// repositories/repository-name/metadata/write
func (o ScopeMapOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

// The creation date of scope map.
func (o ScopeMapOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The user friendly description of the scope map.
func (o ScopeMapOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ScopeMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o ScopeMapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ScopeMapOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScopeMap) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o ScopeMapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScopeMap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScopeMapOutput{})
}
