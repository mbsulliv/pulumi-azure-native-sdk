// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231004preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a ApplicationGroup definition.
type ApplicationGroup struct {
	pulumi.CustomResourceState

	// Resource Type of ApplicationGroup.
	ApplicationGroupType pulumi.StringOutput `pulumi:"applicationGroupType"`
	// Is cloud pc resource.
	CloudPcResource pulumi.BoolOutput `pulumi:"cloudPcResource"`
	// Description of ApplicationGroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Friendly name of ApplicationGroup.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// HostPool arm path of ApplicationGroup.
	HostPoolArmPath pulumi.StringOutput                                          `pulumi:"hostPoolArmPath"`
	Identity        ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ObjectId of ApplicationGroup. (internal use)
	ObjectId pulumi.StringOutput                                      `pulumi:"objectId"`
	Plan     ResourceModelWithAllowedPropertySetResponsePlanPtrOutput `pulumi:"plan"`
	// Boolean representing whether the applicationGroup is show in the feed.
	ShowInFeed pulumi.BoolPtrOutput                                    `pulumi:"showInFeed"`
	Sku        ResourceModelWithAllowedPropertySetResponseSkuPtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Workspace arm path of ApplicationGroup.
	WorkspaceArmPath pulumi.StringOutput `pulumi:"workspaceArmPath"`
}

// NewApplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupType == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupType'")
	}
	if args.HostPoolArmPath == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolArmPath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230905:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20231101preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20240116preview:ApplicationGroup"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20240306preview:ApplicationGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationGroup
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20231004preview:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGroup gets an existing ApplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20231004preview:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGroup resources.
type applicationGroupState struct {
}

type ApplicationGroupState struct {
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	// The name of the application group
	ApplicationGroupName *string `pulumi:"applicationGroupName"`
	// Resource Type of ApplicationGroup.
	ApplicationGroupType string `pulumi:"applicationGroupType"`
	// Description of ApplicationGroup.
	Description *string `pulumi:"description"`
	// Friendly name of ApplicationGroup.
	FriendlyName *string `pulumi:"friendlyName"`
	// HostPool arm path of ApplicationGroup.
	HostPoolArmPath string                                       `pulumi:"hostPoolArmPath"`
	Identity        *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy *string                                  `pulumi:"managedBy"`
	Plan      *ResourceModelWithAllowedPropertySetPlan `pulumi:"plan"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Boolean representing whether the applicationGroup is show in the feed.
	ShowInFeed *bool                                   `pulumi:"showInFeed"`
	Sku        *ResourceModelWithAllowedPropertySetSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationGroup resource.
type ApplicationGroupArgs struct {
	// The name of the application group
	ApplicationGroupName pulumi.StringPtrInput
	// Resource Type of ApplicationGroup.
	ApplicationGroupType pulumi.StringInput
	// Description of ApplicationGroup.
	Description pulumi.StringPtrInput
	// Friendly name of ApplicationGroup.
	FriendlyName pulumi.StringPtrInput
	// HostPool arm path of ApplicationGroup.
	HostPoolArmPath pulumi.StringInput
	Identity        ResourceModelWithAllowedPropertySetIdentityPtrInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
	ManagedBy pulumi.StringPtrInput
	Plan      ResourceModelWithAllowedPropertySetPlanPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Boolean representing whether the applicationGroup is show in the feed.
	ShowInFeed pulumi.BoolPtrInput
	Sku        ResourceModelWithAllowedPropertySetSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

// Resource Type of ApplicationGroup.
func (o ApplicationGroupOutput) ApplicationGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ApplicationGroupType }).(pulumi.StringOutput)
}

// Is cloud pc resource.
func (o ApplicationGroupOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.BoolOutput { return v.CloudPcResource }).(pulumi.BoolOutput)
}

// Description of ApplicationGroup.
func (o ApplicationGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o ApplicationGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Friendly name of ApplicationGroup.
func (o ApplicationGroupOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// HostPool arm path of ApplicationGroup.
func (o ApplicationGroupOutput) HostPoolArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.HostPoolArmPath }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
		return v.Identity
	}).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o ApplicationGroupOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o ApplicationGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The fully qualified resource ID of the resource that manages this resource. Indicates if this resource is managed by another Azure resource. If this is present, complete mode deployment will not delete the resource if it is removed from the template since it is managed by another resource.
func (o ApplicationGroupOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ApplicationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ObjectId of ApplicationGroup. (internal use)
func (o ApplicationGroupOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o ApplicationGroupOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

// Boolean representing whether the applicationGroup is show in the feed.
func (o ApplicationGroupOutput) ShowInFeed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.BoolPtrOutput { return v.ShowInFeed }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGroupOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ApplicationGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApplicationGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ApplicationGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApplicationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Workspace arm path of ApplicationGroup.
func (o ApplicationGroupOutput) WorkspaceArmPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.WorkspaceArmPath }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
}
