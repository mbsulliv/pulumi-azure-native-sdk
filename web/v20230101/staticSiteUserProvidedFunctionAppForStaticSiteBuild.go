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
type StaticSiteUserProvidedFunctionAppForStaticSiteBuild struct {
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

// NewStaticSiteUserProvidedFunctionAppForStaticSiteBuild registers a new resource with the given unique name, arguments, and options.
func NewStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context,
	name string, args *StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StaticSiteUserProvidedFunctionAppForStaticSiteBuild
	err := ctx.RegisterResource("azure-native:web/v20230101:StaticSiteUserProvidedFunctionAppForStaticSiteBuild", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticSiteUserProvidedFunctionAppForStaticSiteBuild gets an existing StaticSiteUserProvidedFunctionAppForStaticSiteBuild resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteUserProvidedFunctionAppForStaticSiteBuildState, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild, error) {
	var resource StaticSiteUserProvidedFunctionAppForStaticSiteBuild
	err := ctx.ReadResource("azure-native:web/v20230101:StaticSiteUserProvidedFunctionAppForStaticSiteBuild", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticSiteUserProvidedFunctionAppForStaticSiteBuild resources.
type staticSiteUserProvidedFunctionAppForStaticSiteBuildState struct {
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildState struct {
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteBuildState)(nil)).Elem()
}

type staticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	// The stage site identifier.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the function app to register with the static site build.
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

// The set of arguments for constructing a StaticSiteUserProvidedFunctionAppForStaticSiteBuild resource.
type StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	// The stage site identifier.
	EnvironmentName pulumi.StringInput
	// Name of the function app to register with the static site build.
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

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteBuildArgs)(nil)).Elem()
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildInput interface {
	pulumi.Input

	ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput
	ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput
}

func (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSiteBuild)(nil)).Elem()
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return i.ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(context.Background())
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput)
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSiteBuild)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return o
}

// The date and time on which the function app was registered with the static site.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The region of the function app registered with the static site
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput {
		return v.FunctionAppRegion
	}).(pulumi.StringPtrOutput)
}

// The resource id of the function app registered with the static site
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput{})
}
