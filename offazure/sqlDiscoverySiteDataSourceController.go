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

// A SQL discovery site data source resource.
// Azure REST API version: 2023-06-06.
type SqlDiscoverySiteDataSourceController struct {
	pulumi.CustomResourceState

	// Gets or sets the discovery site Id.
	DiscoverySiteId pulumi.StringPtrOutput `pulumi:"discoverySiteId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// provisioning state enum
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlDiscoverySiteDataSourceController registers a new resource with the given unique name, arguments, and options.
func NewSqlDiscoverySiteDataSourceController(ctx *pulumi.Context,
	name string, args *SqlDiscoverySiteDataSourceControllerArgs, opts ...pulumi.ResourceOption) (*SqlDiscoverySiteDataSourceController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SiteName == nil {
		return nil, errors.New("invalid value for required argument 'SiteName'")
	}
	if args.SqlSiteName == nil {
		return nil, errors.New("invalid value for required argument 'SqlSiteName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure/v20230606:SqlDiscoverySiteDataSourceController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqlDiscoverySiteDataSourceController
	err := ctx.RegisterResource("azure-native:offazure:SqlDiscoverySiteDataSourceController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlDiscoverySiteDataSourceController gets an existing SqlDiscoverySiteDataSourceController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlDiscoverySiteDataSourceController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDiscoverySiteDataSourceControllerState, opts ...pulumi.ResourceOption) (*SqlDiscoverySiteDataSourceController, error) {
	var resource SqlDiscoverySiteDataSourceController
	err := ctx.ReadResource("azure-native:offazure:SqlDiscoverySiteDataSourceController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlDiscoverySiteDataSourceController resources.
type sqlDiscoverySiteDataSourceControllerState struct {
}

type SqlDiscoverySiteDataSourceControllerState struct {
}

func (SqlDiscoverySiteDataSourceControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDiscoverySiteDataSourceControllerState)(nil)).Elem()
}

type sqlDiscoverySiteDataSourceControllerArgs struct {
	// SQL Discovery site data source name.
	DiscoverySiteDataSourceName *string `pulumi:"discoverySiteDataSourceName"`
	// Gets or sets the discovery site Id.
	DiscoverySiteId *string `pulumi:"discoverySiteId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
	// SQL site name.
	SqlSiteName string `pulumi:"sqlSiteName"`
}

// The set of arguments for constructing a SqlDiscoverySiteDataSourceController resource.
type SqlDiscoverySiteDataSourceControllerArgs struct {
	// SQL Discovery site data source name.
	DiscoverySiteDataSourceName pulumi.StringPtrInput
	// Gets or sets the discovery site Id.
	DiscoverySiteId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name
	SiteName pulumi.StringInput
	// SQL site name.
	SqlSiteName pulumi.StringInput
}

func (SqlDiscoverySiteDataSourceControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDiscoverySiteDataSourceControllerArgs)(nil)).Elem()
}

type SqlDiscoverySiteDataSourceControllerInput interface {
	pulumi.Input

	ToSqlDiscoverySiteDataSourceControllerOutput() SqlDiscoverySiteDataSourceControllerOutput
	ToSqlDiscoverySiteDataSourceControllerOutputWithContext(ctx context.Context) SqlDiscoverySiteDataSourceControllerOutput
}

func (*SqlDiscoverySiteDataSourceController) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDiscoverySiteDataSourceController)(nil)).Elem()
}

func (i *SqlDiscoverySiteDataSourceController) ToSqlDiscoverySiteDataSourceControllerOutput() SqlDiscoverySiteDataSourceControllerOutput {
	return i.ToSqlDiscoverySiteDataSourceControllerOutputWithContext(context.Background())
}

func (i *SqlDiscoverySiteDataSourceController) ToSqlDiscoverySiteDataSourceControllerOutputWithContext(ctx context.Context) SqlDiscoverySiteDataSourceControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDiscoverySiteDataSourceControllerOutput)
}

type SqlDiscoverySiteDataSourceControllerOutput struct{ *pulumi.OutputState }

func (SqlDiscoverySiteDataSourceControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDiscoverySiteDataSourceController)(nil)).Elem()
}

func (o SqlDiscoverySiteDataSourceControllerOutput) ToSqlDiscoverySiteDataSourceControllerOutput() SqlDiscoverySiteDataSourceControllerOutput {
	return o
}

func (o SqlDiscoverySiteDataSourceControllerOutput) ToSqlDiscoverySiteDataSourceControllerOutputWithContext(ctx context.Context) SqlDiscoverySiteDataSourceControllerOutput {
	return o
}

// Gets or sets the discovery site Id.
func (o SqlDiscoverySiteDataSourceControllerOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDiscoverySiteDataSourceController) pulumi.StringPtrOutput { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SqlDiscoverySiteDataSourceControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDiscoverySiteDataSourceController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// provisioning state enum
func (o SqlDiscoverySiteDataSourceControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDiscoverySiteDataSourceController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SqlDiscoverySiteDataSourceControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDiscoverySiteDataSourceController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqlDiscoverySiteDataSourceControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDiscoverySiteDataSourceController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDiscoverySiteDataSourceControllerOutput{})
}