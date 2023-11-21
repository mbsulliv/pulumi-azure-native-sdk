// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for App Attach Package properties.
// Azure REST API version: 2023-10-04-preview.
type AppAttachPackage struct {
	pulumi.CustomResourceState

	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag     pulumi.StringOutput                                          `pulumi:"etag"`
	Identity ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput                                      `pulumi:"name"`
	Plan ResourceModelWithAllowedPropertySetResponsePlanPtrOutput `pulumi:"plan"`
	// Detailed properties for App Attach Package
	Properties AppAttachPackagePropertiesResponseOutput                `pulumi:"properties"`
	Sku        ResourceModelWithAllowedPropertySetResponseSkuPtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppAttachPackage registers a new resource with the given unique name, arguments, and options.
func NewAppAttachPackage(ctx *pulumi.Context,
	name string, args *AppAttachPackageArgs, opts ...pulumi.ResourceOption) (*AppAttachPackage, error) {
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
			Type: pulumi.String("azure-native:desktopvirtualization/v20231004preview:AppAttachPackage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AppAttachPackage
	err := ctx.RegisterResource("azure-native:desktopvirtualization:AppAttachPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppAttachPackage gets an existing AppAttachPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppAttachPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppAttachPackageState, opts ...pulumi.ResourceOption) (*AppAttachPackage, error) {
	var resource AppAttachPackage
	err := ctx.ReadResource("azure-native:desktopvirtualization:AppAttachPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppAttachPackage resources.
type appAttachPackageState struct {
}

type AppAttachPackageState struct {
}

func (AppAttachPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*appAttachPackageState)(nil)).Elem()
}

type appAttachPackageArgs struct {
	// The name of the App Attach package arm object
	AppAttachPackageName *string                                      `pulumi:"appAttachPackageName"`
	Identity             *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string                                  `pulumi:"managedBy"`
	Plan      *ResourceModelWithAllowedPropertySetPlan `pulumi:"plan"`
	// Detailed properties for App Attach Package
	Properties AppAttachPackageProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string                                  `pulumi:"resourceGroupName"`
	Sku               *ResourceModelWithAllowedPropertySetSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AppAttachPackage resource.
type AppAttachPackageArgs struct {
	// The name of the App Attach package arm object
	AppAttachPackageName pulumi.StringPtrInput
	Identity             ResourceModelWithAllowedPropertySetIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrInput
	Plan      ResourceModelWithAllowedPropertySetPlanPtrInput
	// Detailed properties for App Attach Package
	Properties AppAttachPackagePropertiesInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	Sku               ResourceModelWithAllowedPropertySetSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AppAttachPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appAttachPackageArgs)(nil)).Elem()
}

type AppAttachPackageInput interface {
	pulumi.Input

	ToAppAttachPackageOutput() AppAttachPackageOutput
	ToAppAttachPackageOutputWithContext(ctx context.Context) AppAttachPackageOutput
}

func (*AppAttachPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAttachPackage)(nil)).Elem()
}

func (i *AppAttachPackage) ToAppAttachPackageOutput() AppAttachPackageOutput {
	return i.ToAppAttachPackageOutputWithContext(context.Background())
}

func (i *AppAttachPackage) ToAppAttachPackageOutputWithContext(ctx context.Context) AppAttachPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppAttachPackageOutput)
}

type AppAttachPackageOutput struct{ *pulumi.OutputState }

func (AppAttachPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppAttachPackage)(nil)).Elem()
}

func (o AppAttachPackageOutput) ToAppAttachPackageOutput() AppAttachPackageOutput {
	return o
}

func (o AppAttachPackageOutput) ToAppAttachPackageOutputWithContext(ctx context.Context) AppAttachPackageOutput {
	return o
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o AppAttachPackageOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o AppAttachPackageOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o AppAttachPackageOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o AppAttachPackageOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o AppAttachPackageOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o AppAttachPackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppAttachPackageOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// Detailed properties for App Attach Package
func (o AppAttachPackageOutput) Properties() AppAttachPackagePropertiesResponseOutput {
	return o.ApplyT(func(v *AppAttachPackage) AppAttachPackagePropertiesResponseOutput { return v.Properties }).(AppAttachPackagePropertiesResponseOutput)
}

func (o AppAttachPackageOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *AppAttachPackage) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AppAttachPackageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AppAttachPackage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AppAttachPackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AppAttachPackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppAttachPackage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppAttachPackageOutput{})
}