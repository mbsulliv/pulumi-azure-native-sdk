// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Static Site User Provided Function App ARM resource.
type StaticSiteUserProvidedFunctionAppForStaticSite struct {
	pulumi.CustomResourceState

	// The date and time on which the function app was registered with the static site.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The region of the function app registered with the static site
	FunctionAppRegion pulumi.StringPtrOutput `pulumi:"functionAppRegion"`
	// The resource id of the function app registered with the static site
	FunctionAppResourceId pulumi.StringPtrOutput `pulumi:"functionAppResourceId"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStaticSiteUserProvidedFunctionAppForStaticSite registers a new resource with the given unique name, arguments, and options.
func NewStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteUserProvidedFunctionAppForStaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:StaticSiteUserProvidedFunctionAppForStaticSite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StaticSiteUserProvidedFunctionAppForStaticSite
	err := ctx.RegisterResource("azure-native:web/v20230101:StaticSiteUserProvidedFunctionAppForStaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSiteUserProvidedFunctionAppForStaticSite gets an existing StaticSiteUserProvidedFunctionAppForStaticSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteUserProvidedFunctionAppForStaticSiteState, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSite, error) {
	var resource StaticSiteUserProvidedFunctionAppForStaticSite
	err := ctx.ReadResource("azure-native:web/v20230101:StaticSiteUserProvidedFunctionAppForStaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSiteUserProvidedFunctionAppForStaticSite resources.
type staticSiteUserProvidedFunctionAppForStaticSiteState struct {
}

type StaticSiteUserProvidedFunctionAppForStaticSiteState struct {
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteState)(nil)).Elem()
}

type staticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	// Name of the function app to register with the static site.
	FunctionAppName *string `pulumi:"functionAppName"`
	// The region of the function app registered with the static site
	FunctionAppRegion *string `pulumi:"functionAppRegion"`
	// The resource id of the function app registered with the static site
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	// Specify <code>true</code> to force the update of the auth configuration on the function app even if an AzureStaticWebApps provider is already configured on the function app. The default is <code>false</code>.
	IsForced *bool `pulumi:"isForced"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a StaticSiteUserProvidedFunctionAppForStaticSite resource.
type StaticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	// Name of the function app to register with the static site.
	FunctionAppName pulumi.StringPtrInput
	// The region of the function app registered with the static site
	FunctionAppRegion pulumi.StringPtrInput
	// The resource id of the function app registered with the static site
	FunctionAppResourceId pulumi.StringPtrInput
	// Specify <code>true</code> to force the update of the auth configuration on the function app even if an AzureStaticWebApps provider is already configured on the function app. The default is <code>false</code>.
	IsForced pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the static site.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteArgs)(nil)).Elem()
}

type StaticSiteUserProvidedFunctionAppForStaticSiteInput interface {
	pulumi.Input

	ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput
	ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput
}

func (*StaticSiteUserProvidedFunctionAppForStaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSite)(nil)).Elem()
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSite) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return i.ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSite) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserProvidedFunctionAppForStaticSiteOutput)
}

type StaticSiteUserProvidedFunctionAppForStaticSiteOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSite)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutput() StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteOutput {
	return o
}

// The date and time on which the function app was registered with the static site.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The region of the function app registered with the static site
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput {
		return v.FunctionAppRegion
	}).(pulumi.StringPtrOutput)
}

// The resource id of the function app registered with the static site
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppForStaticSiteOutput{})
}