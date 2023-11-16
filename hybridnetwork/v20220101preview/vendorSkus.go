// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sku sub resource.
type VendorSkus struct {
	pulumi.CustomResourceState

	// The sku deployment mode.
	DeploymentMode pulumi.StringPtrOutput `pulumi:"deploymentMode"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters pulumi.AnyOutput `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.AnyOutput `pulumi:"managedApplicationTemplate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The template definition of the network function.
	NetworkFunctionTemplate NetworkFunctionTemplateResponsePtrOutput `pulumi:"networkFunctionTemplate"`
	// The network function type.
	NetworkFunctionType pulumi.StringPtrOutput `pulumi:"networkFunctionType"`
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku type.
	SkuType pulumi.StringPtrOutput `pulumi:"skuType"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVendorSkus registers a new resource with the given unique name, arguments, and options.
func NewVendorSkus(ctx *pulumi.Context,
	name string, args *VendorSkusArgs, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VendorName == nil {
		return nil, errors.New("invalid value for required argument 'VendorName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20200101preview:VendorSkus"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20210501:VendorSkus"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VendorSkus
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20220101preview:VendorSkus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVendorSkus gets an existing VendorSkus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVendorSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkusState, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	var resource VendorSkus
	err := ctx.ReadResource("azure-native:hybridnetwork/v20220101preview:VendorSkus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VendorSkus resources.
type vendorSkusState struct {
}

type VendorSkusState struct {
}

func (VendorSkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusState)(nil)).Elem()
}

type vendorSkusArgs struct {
	// The sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate interface{} `pulumi:"managedApplicationTemplate"`
	// The template definition of the network function.
	NetworkFunctionTemplate *NetworkFunctionTemplate `pulumi:"networkFunctionTemplate"`
	// The network function type.
	NetworkFunctionType *string `pulumi:"networkFunctionType"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The name of the sku.
	SkuName *string `pulumi:"skuName"`
	// The sku type.
	SkuType *string `pulumi:"skuType"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// The set of arguments for constructing a VendorSkus resource.
type VendorSkusArgs struct {
	// The sku deployment mode.
	DeploymentMode pulumi.StringPtrInput
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters pulumi.Input
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.Input
	// The template definition of the network function.
	NetworkFunctionTemplate NetworkFunctionTemplatePtrInput
	// The network function type.
	NetworkFunctionType pulumi.StringPtrInput
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrInput
	// The name of the sku.
	SkuName pulumi.StringPtrInput
	// The sku type.
	SkuType pulumi.StringPtrInput
	// The name of the vendor.
	VendorName pulumi.StringInput
}

func (VendorSkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusArgs)(nil)).Elem()
}

type VendorSkusInput interface {
	pulumi.Input

	ToVendorSkusOutput() VendorSkusOutput
	ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput
}

func (*VendorSkus) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkus)(nil)).Elem()
}

func (i *VendorSkus) ToVendorSkusOutput() VendorSkusOutput {
	return i.ToVendorSkusOutputWithContext(context.Background())
}

func (i *VendorSkus) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VendorSkusOutput)
}

type VendorSkusOutput struct{ *pulumi.OutputState }

func (VendorSkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VendorSkus)(nil)).Elem()
}

func (o VendorSkusOutput) ToVendorSkusOutput() VendorSkusOutput {
	return o
}

func (o VendorSkusOutput) ToVendorSkusOutputWithContext(ctx context.Context) VendorSkusOutput {
	return o
}

// The sku deployment mode.
func (o VendorSkusOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringPtrOutput { return v.DeploymentMode }).(pulumi.StringPtrOutput)
}

// The parameters for the managed application to be supplied by the vendor.
func (o VendorSkusOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.AnyOutput { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

// The template for the managed application deployment.
func (o VendorSkusOutput) ManagedApplicationTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.AnyOutput { return v.ManagedApplicationTemplate }).(pulumi.AnyOutput)
}

// The name of the resource
func (o VendorSkusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The template definition of the network function.
func (o VendorSkusOutput) NetworkFunctionTemplate() NetworkFunctionTemplateResponsePtrOutput {
	return o.ApplyT(func(v *VendorSkus) NetworkFunctionTemplateResponsePtrOutput { return v.NetworkFunctionTemplate }).(NetworkFunctionTemplateResponsePtrOutput)
}

// The network function type.
func (o VendorSkusOutput) NetworkFunctionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringPtrOutput { return v.NetworkFunctionType }).(pulumi.StringPtrOutput)
}

// Indicates if the vendor sku is in preview mode.
func (o VendorSkusOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.BoolPtrOutput { return v.Preview }).(pulumi.BoolPtrOutput)
}

// The provisioning state of the vendor sku sub resource.
func (o VendorSkusOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku type.
func (o VendorSkusOutput) SkuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringPtrOutput { return v.SkuType }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o VendorSkusOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VendorSkus) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VendorSkusOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VendorSkus) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VendorSkusOutput{})
}
