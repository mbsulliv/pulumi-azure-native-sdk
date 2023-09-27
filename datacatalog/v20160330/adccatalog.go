// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160330

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Data Catalog.
type ADCCatalog struct {
	pulumi.CustomResourceState

	// Azure data catalog admin list.
	Admins PrincipalsResponseArrayOutput `pulumi:"admins"`
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment pulumi.BoolPtrOutput `pulumi:"enableAutomaticUnitAdjustment"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure data catalog SKU.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// Azure data catalog provision status.
	SuccessfullyProvisioned pulumi.BoolPtrOutput `pulumi:"successfullyProvisioned"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Azure data catalog units.
	Units pulumi.IntPtrOutput `pulumi:"units"`
	// Azure data catalog user list.
	Users PrincipalsResponseArrayOutput `pulumi:"users"`
}

// NewADCCatalog registers a new resource with the given unique name, arguments, and options.
func NewADCCatalog(ctx *pulumi.Context,
	name string, args *ADCCatalogArgs, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datacatalog:ADCCatalog"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ADCCatalog
	err := ctx.RegisterResource("azure-native:datacatalog/v20160330:ADCCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADCCatalog gets an existing ADCCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADCCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADCCatalogState, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	var resource ADCCatalog
	err := ctx.ReadResource("azure-native:datacatalog/v20160330:ADCCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADCCatalog resources.
type adccatalogState struct {
}

type ADCCatalogState struct {
}

func (ADCCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogState)(nil)).Elem()
}

type adccatalogArgs struct {
	// Azure data catalog admin list.
	Admins []Principals `pulumi:"admins"`
	// The name of the data catalog in the specified subscription and resource group.
	CatalogName *string `pulumi:"catalogName"`
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment *bool `pulumi:"enableAutomaticUnitAdjustment"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure data catalog SKU.
	Sku *string `pulumi:"sku"`
	// Azure data catalog provision status.
	SuccessfullyProvisioned *bool `pulumi:"successfullyProvisioned"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure data catalog units.
	Units *int `pulumi:"units"`
	// Azure data catalog user list.
	Users []Principals `pulumi:"users"`
}

// The set of arguments for constructing a ADCCatalog resource.
type ADCCatalogArgs struct {
	// Azure data catalog admin list.
	Admins PrincipalsArrayInput
	// The name of the data catalog in the specified subscription and resource group.
	CatalogName pulumi.StringPtrInput
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment pulumi.BoolPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure data catalog SKU.
	Sku pulumi.StringPtrInput
	// Azure data catalog provision status.
	SuccessfullyProvisioned pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure data catalog units.
	Units pulumi.IntPtrInput
	// Azure data catalog user list.
	Users PrincipalsArrayInput
}

func (ADCCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogArgs)(nil)).Elem()
}

type ADCCatalogInput interface {
	pulumi.Input

	ToADCCatalogOutput() ADCCatalogOutput
	ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput
}

func (*ADCCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**ADCCatalog)(nil)).Elem()
}

func (i *ADCCatalog) ToADCCatalogOutput() ADCCatalogOutput {
	return i.ToADCCatalogOutputWithContext(context.Background())
}

func (i *ADCCatalog) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADCCatalogOutput)
}

func (i *ADCCatalog) ToOutput(ctx context.Context) pulumix.Output[*ADCCatalog] {
	return pulumix.Output[*ADCCatalog]{
		OutputState: i.ToADCCatalogOutputWithContext(ctx).OutputState,
	}
}

type ADCCatalogOutput struct{ *pulumi.OutputState }

func (ADCCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADCCatalog)(nil)).Elem()
}

func (o ADCCatalogOutput) ToADCCatalogOutput() ADCCatalogOutput {
	return o
}

func (o ADCCatalogOutput) ToADCCatalogOutputWithContext(ctx context.Context) ADCCatalogOutput {
	return o
}

func (o ADCCatalogOutput) ToOutput(ctx context.Context) pulumix.Output[*ADCCatalog] {
	return pulumix.Output[*ADCCatalog]{
		OutputState: o.OutputState,
	}
}

// Azure data catalog admin list.
func (o ADCCatalogOutput) Admins() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v *ADCCatalog) PrincipalsResponseArrayOutput { return v.Admins }).(PrincipalsResponseArrayOutput)
}

// Automatic unit adjustment enabled or not.
func (o ADCCatalogOutput) EnableAutomaticUnitAdjustment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.BoolPtrOutput { return v.EnableAutomaticUnitAdjustment }).(pulumi.BoolPtrOutput)
}

// Resource etag
func (o ADCCatalogOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ADCCatalogOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o ADCCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure data catalog SKU.
func (o ADCCatalogOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

// Azure data catalog provision status.
func (o ADCCatalogOutput) SuccessfullyProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.BoolPtrOutput { return v.SuccessfullyProvisioned }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o ADCCatalogOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ADCCatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Azure data catalog units.
func (o ADCCatalogOutput) Units() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ADCCatalog) pulumi.IntPtrOutput { return v.Units }).(pulumi.IntPtrOutput)
}

// Azure data catalog user list.
func (o ADCCatalogOutput) Users() PrincipalsResponseArrayOutput {
	return o.ApplyT(func(v *ADCCatalog) PrincipalsResponseArrayOutput { return v.Users }).(PrincipalsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ADCCatalogOutput{})
}
