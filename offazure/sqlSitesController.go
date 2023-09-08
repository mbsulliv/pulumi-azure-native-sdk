// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SQL site web model.
// Azure REST API version: 2023-06-06.
type SqlSitesController struct {
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

// NewSqlSitesController registers a new resource with the given unique name, arguments, and options.
func NewSqlSitesController(ctx *pulumi.Context,
	name string, args *SqlSitesControllerArgs, opts ...pulumi.ResourceOption) (*SqlSitesController, error) {
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
			Type: pulumi.String("azure-native:offazure/v20230606:SqlSitesController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlSitesController
	err := ctx.RegisterResource("azure-native:offazure:SqlSitesController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlSitesController gets an existing SqlSitesController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlSitesController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlSitesControllerState, opts ...pulumi.ResourceOption) (*SqlSitesController, error) {
	var resource SqlSitesController
	err := ctx.ReadResource("azure-native:offazure:SqlSitesController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlSitesController resources.
type sqlSitesControllerState struct {
}

type SqlSitesControllerState struct {
}

func (SqlSitesControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlSitesControllerState)(nil)).Elem()
}

type sqlSitesControllerArgs struct {
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
	// SQL site name.
	SqlSiteName *string `pulumi:"sqlSiteName"`
}

// The set of arguments for constructing a SqlSitesController resource.
type SqlSitesControllerArgs struct {
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
	// SQL site name.
	SqlSiteName pulumi.StringPtrInput
}

func (SqlSitesControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlSitesControllerArgs)(nil)).Elem()
}

type SqlSitesControllerInput interface {
	pulumi.Input

	ToSqlSitesControllerOutput() SqlSitesControllerOutput
	ToSqlSitesControllerOutputWithContext(ctx context.Context) SqlSitesControllerOutput
}

func (*SqlSitesController) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlSitesController)(nil)).Elem()
}

func (i *SqlSitesController) ToSqlSitesControllerOutput() SqlSitesControllerOutput {
	return i.ToSqlSitesControllerOutputWithContext(context.Background())
}

func (i *SqlSitesController) ToSqlSitesControllerOutputWithContext(ctx context.Context) SqlSitesControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlSitesControllerOutput)
}

type SqlSitesControllerOutput struct{ *pulumi.OutputState }

func (SqlSitesControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlSitesController)(nil)).Elem()
}

func (o SqlSitesControllerOutput) ToSqlSitesControllerOutput() SqlSitesControllerOutput {
	return o
}

func (o SqlSitesControllerOutput) ToSqlSitesControllerOutputWithContext(ctx context.Context) SqlSitesControllerOutput {
	return o
}

// Gets or sets the discovery scenario.
func (o SqlSitesControllerOutput) DiscoveryScenario() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlSitesController) pulumi.StringPtrOutput { return v.DiscoveryScenario }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SqlSitesControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlSitesController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o SqlSitesControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlSitesController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the service endpoint.
func (o SqlSitesControllerOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlSitesController) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Gets or sets the appliance details used by service to communicate
//
// to the appliance.
func (o SqlSitesControllerOutput) SiteAppliancePropertiesCollection() SiteAppliancePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *SqlSitesController) SiteAppliancePropertiesResponseArrayOutput {
		return v.SiteAppliancePropertiesCollection
	}).(SiteAppliancePropertiesResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SqlSitesControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlSitesController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqlSitesControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlSitesController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlSitesControllerOutput{})
}