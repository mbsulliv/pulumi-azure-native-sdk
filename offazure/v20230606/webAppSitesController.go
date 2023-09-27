// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// WebApp site web model.
type WebAppSitesController struct {
	pulumi.CustomResourceState

	// Gets or sets the discovery scenario.
	DiscoveryScenario pulumi.StringPtrOutput `pulumi:"discoveryScenario"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// provisioning state enum
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets the service endpoint.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// Gets or sets the appliance details used by service to communicate
	//
	// to the appliance.
	SiteAppliancePropertiesCollection SiteAppliancePropertiesResponseArrayOutput `pulumi:"siteAppliancePropertiesCollection"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSitesController registers a new resource with the given unique name, arguments, and options.
func NewWebAppSitesController(ctx *pulumi.Context,
	name string, args *WebAppSitesControllerArgs, opts ...pulumi.ResourceOption) (*WebAppSitesController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SiteName == nil {
		return nil, errors.New("invalid value for required argument 'SiteName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:WebAppSitesController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSitesController
	err := ctx.RegisterResource("azure-native:offazure/v20230606:WebAppSitesController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSitesController gets an existing WebAppSitesController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSitesController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSitesControllerState, opts ...pulumi.ResourceOption) (*WebAppSitesController, error) {
	var resource WebAppSitesController
	err := ctx.ReadResource("azure-native:offazure/v20230606:WebAppSitesController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSitesController resources.
type webAppSitesControllerState struct {
}

type WebAppSitesControllerState struct {
}

func (WebAppSitesControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitesControllerState)(nil)).Elem()
}

type webAppSitesControllerArgs struct {
	// Gets or sets the discovery scenario.
	DiscoveryScenario *string `pulumi:"discoveryScenario"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the appliance details used by service to communicate
	//
	// to the appliance.
	SiteAppliancePropertiesCollection []SiteApplianceProperties `pulumi:"siteAppliancePropertiesCollection"`
	// Site name
	SiteName string `pulumi:"siteName"`
	// Web app site name.
	WebAppSiteName *string `pulumi:"webAppSiteName"`
}

// The set of arguments for constructing a WebAppSitesController resource.
type WebAppSitesControllerArgs struct {
	// Gets or sets the discovery scenario.
	DiscoveryScenario pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the appliance details used by service to communicate
	//
	// to the appliance.
	SiteAppliancePropertiesCollection SiteAppliancePropertiesArrayInput
	// Site name
	SiteName pulumi.StringInput
	// Web app site name.
	WebAppSiteName pulumi.StringPtrInput
}

func (WebAppSitesControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitesControllerArgs)(nil)).Elem()
}

type WebAppSitesControllerInput interface {
	pulumi.Input

	ToWebAppSitesControllerOutput() WebAppSitesControllerOutput
	ToWebAppSitesControllerOutputWithContext(ctx context.Context) WebAppSitesControllerOutput
}

func (*WebAppSitesController) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitesController)(nil)).Elem()
}

func (i *WebAppSitesController) ToWebAppSitesControllerOutput() WebAppSitesControllerOutput {
	return i.ToWebAppSitesControllerOutputWithContext(context.Background())
}

func (i *WebAppSitesController) ToWebAppSitesControllerOutputWithContext(ctx context.Context) WebAppSitesControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSitesControllerOutput)
}

func (i *WebAppSitesController) ToOutput(ctx context.Context) pulumix.Output[*WebAppSitesController] {
	return pulumix.Output[*WebAppSitesController]{
		OutputState: i.ToWebAppSitesControllerOutputWithContext(ctx).OutputState,
	}
}

type WebAppSitesControllerOutput struct{ *pulumi.OutputState }

func (WebAppSitesControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitesController)(nil)).Elem()
}

func (o WebAppSitesControllerOutput) ToWebAppSitesControllerOutput() WebAppSitesControllerOutput {
	return o
}

func (o WebAppSitesControllerOutput) ToWebAppSitesControllerOutputWithContext(ctx context.Context) WebAppSitesControllerOutput {
	return o
}

func (o WebAppSitesControllerOutput) ToOutput(ctx context.Context) pulumix.Output[*WebAppSitesController] {
	return pulumix.Output[*WebAppSitesController]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the discovery scenario.
func (o WebAppSitesControllerOutput) DiscoveryScenario() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitesController) pulumi.StringPtrOutput { return v.DiscoveryScenario }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WebAppSitesControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitesController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o WebAppSitesControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitesController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o WebAppSitesControllerOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitesController) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the appliance details used by service to communicate
//
// to the appliance.
func (o WebAppSitesControllerOutput) SiteAppliancePropertiesCollection() SiteAppliancePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *WebAppSitesController) SiteAppliancePropertiesResponseArrayOutput {
		return v.SiteAppliancePropertiesCollection
	}).(SiteAppliancePropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WebAppSitesControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppSitesController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WebAppSitesControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitesController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSitesControllerOutput{})
}
