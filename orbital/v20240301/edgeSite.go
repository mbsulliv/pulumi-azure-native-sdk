// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A customer's reference to a global communications site site.
type EdgeSite struct {
	pulumi.CustomResourceState

	// A reference to global communications site.
	GlobalCommunicationsSite EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput `pulumi:"globalCommunicationsSite"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEdgeSite registers a new resource with the given unique name, arguments, and options.
func NewEdgeSite(ctx *pulumi.Context,
	name string, args *EdgeSiteArgs, opts ...pulumi.ResourceOption) (*EdgeSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalCommunicationsSite == nil {
		return nil, errors.New("invalid value for required argument 'GlobalCommunicationsSite'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital:EdgeSite"),
		},
		{
			Type: pulumi.String("azure-native:orbital/v20240301preview:EdgeSite"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EdgeSite
	err := ctx.RegisterResource("azure-native:orbital/v20240301:EdgeSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeSite gets an existing EdgeSite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeSiteState, opts ...pulumi.ResourceOption) (*EdgeSite, error) {
	var resource EdgeSite
	err := ctx.ReadResource("azure-native:orbital/v20240301:EdgeSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeSite resources.
type edgeSiteState struct {
}

type EdgeSiteState struct {
}

func (EdgeSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeSiteState)(nil)).Elem()
}

type edgeSiteArgs struct {
	// Edge site name.
	EdgeSiteName *string `pulumi:"edgeSiteName"`
	// A reference to global communications site.
	GlobalCommunicationsSite EdgeSitesPropertiesGlobalCommunicationsSite `pulumi:"globalCommunicationsSite"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EdgeSite resource.
type EdgeSiteArgs struct {
	// Edge site name.
	EdgeSiteName pulumi.StringPtrInput
	// A reference to global communications site.
	GlobalCommunicationsSite EdgeSitesPropertiesGlobalCommunicationsSiteInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (EdgeSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeSiteArgs)(nil)).Elem()
}

type EdgeSiteInput interface {
	pulumi.Input

	ToEdgeSiteOutput() EdgeSiteOutput
	ToEdgeSiteOutputWithContext(ctx context.Context) EdgeSiteOutput
}

func (*EdgeSite) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeSite)(nil)).Elem()
}

func (i *EdgeSite) ToEdgeSiteOutput() EdgeSiteOutput {
	return i.ToEdgeSiteOutputWithContext(context.Background())
}

func (i *EdgeSite) ToEdgeSiteOutputWithContext(ctx context.Context) EdgeSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeSiteOutput)
}

type EdgeSiteOutput struct{ *pulumi.OutputState }

func (EdgeSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeSite)(nil)).Elem()
}

func (o EdgeSiteOutput) ToEdgeSiteOutput() EdgeSiteOutput {
	return o
}

func (o EdgeSiteOutput) ToEdgeSiteOutputWithContext(ctx context.Context) EdgeSiteOutput {
	return o
}

// A reference to global communications site.
func (o EdgeSiteOutput) GlobalCommunicationsSite() EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput {
	return o.ApplyT(func(v *EdgeSite) EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput {
		return v.GlobalCommunicationsSite
	}).(EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput)
}

// The geo-location where the resource lives
func (o EdgeSiteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeSite) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o EdgeSiteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeSite) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EdgeSiteOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EdgeSite) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o EdgeSiteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EdgeSite) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EdgeSiteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeSite) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EdgeSiteOutput{})
}
