// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network slice resource.
type Slice struct {
	pulumi.CustomResourceState

	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrOutput `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrOutput `pulumi:"createdByType"`
	// An optional description for this network slice.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrOutput `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrOutput `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the network slice resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The S-NSSAI (single network slice selection assistance information). Unique at the scope of a MobileNetwork.
	Snssai SnssaiResponseOutput `pulumi:"snssai"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSlice registers a new resource with the given unique name, arguments, and options.
func NewSlice(ctx *pulumi.Context,
	name string, args *SliceArgs, opts ...pulumi.ResourceOption) (*Slice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MobileNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'MobileNetworkName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Snssai == nil {
		return nil, errors.New("invalid value for required argument 'Snssai'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:mobilenetwork:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20220401preview:Slice"),
		},
		{
			Type: pulumi.String("azure-native:mobilenetwork/v20221101:Slice"),
		},
	})
	opts = append(opts, aliases)
	var resource Slice
	err := ctx.RegisterResource("azure-native:mobilenetwork/v20220301preview:Slice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlice gets an existing Slice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SliceState, opts ...pulumi.ResourceOption) (*Slice, error) {
	var resource Slice
	err := ctx.ReadResource("azure-native:mobilenetwork/v20220301preview:Slice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slice resources.
type sliceState struct {
}

type SliceState struct {
}

func (SliceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceState)(nil)).Elem()
}

type sliceArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// An optional description for this network slice.
	Description *string `pulumi:"description"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the mobile network slice.
	SliceName *string `pulumi:"sliceName"`
	// The S-NSSAI (single network slice selection assistance information). Unique at the scope of a MobileNetwork.
	Snssai Snssai `pulumi:"snssai"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Slice resource.
type SliceArgs struct {
	// The timestamp of resource creation (UTC).
	CreatedAt pulumi.StringPtrInput
	// The identity that created the resource.
	CreatedBy pulumi.StringPtrInput
	// The type of identity that created the resource.
	CreatedByType pulumi.StringPtrInput
	// An optional description for this network slice.
	Description pulumi.StringPtrInput
	// The timestamp of resource last modification (UTC)
	LastModifiedAt pulumi.StringPtrInput
	// The identity that last modified the resource.
	LastModifiedBy pulumi.StringPtrInput
	// The type of identity that last modified the resource.
	LastModifiedByType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the mobile network slice.
	SliceName pulumi.StringPtrInput
	// The S-NSSAI (single network slice selection assistance information). Unique at the scope of a MobileNetwork.
	Snssai SnssaiInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SliceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sliceArgs)(nil)).Elem()
}

type SliceInput interface {
	pulumi.Input

	ToSliceOutput() SliceOutput
	ToSliceOutputWithContext(ctx context.Context) SliceOutput
}

func (*Slice) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (i *Slice) ToSliceOutput() SliceOutput {
	return i.ToSliceOutputWithContext(context.Background())
}

func (i *Slice) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SliceOutput)
}

type SliceOutput struct{ *pulumi.OutputState }

func (SliceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slice)(nil)).Elem()
}

func (o SliceOutput) ToSliceOutput() SliceOutput {
	return o
}

func (o SliceOutput) ToSliceOutputWithContext(ctx context.Context) SliceOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SliceOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SliceOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SliceOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// An optional description for this network slice.
func (o SliceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SliceOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SliceOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SliceOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringPtrOutput { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o SliceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SliceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network slice resource.
func (o SliceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The S-NSSAI (single network slice selection assistance information). Unique at the scope of a MobileNetwork.
func (o SliceOutput) Snssai() SnssaiResponseOutput {
	return o.ApplyT(func(v *Slice) SnssaiResponseOutput { return v.Snssai }).(SnssaiResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SliceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Slice) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SliceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SliceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Slice) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SliceOutput{})
}
