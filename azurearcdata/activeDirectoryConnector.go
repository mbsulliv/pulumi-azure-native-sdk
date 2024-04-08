// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurearcdata

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Active directory connector resource
// Azure REST API version: 2023-01-15-preview. Prior API version in Azure Native 1.x: 2022-03-01-preview.
//
// Other available API versions: 2024-01-01.
type ActiveDirectoryConnector struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// null
	Properties ActiveDirectoryConnectorPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewActiveDirectoryConnector registers a new resource with the given unique name, arguments, and options.
func NewActiveDirectoryConnector(ctx *pulumi.Context,
	name string, args *ActiveDirectoryConnectorArgs, opts ...pulumi.ResourceOption) (*ActiveDirectoryConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataControllerName == nil {
		return nil, errors.New("invalid value for required argument 'DataControllerName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToActiveDirectoryConnectorPropertiesOutput().ApplyT(func(v ActiveDirectoryConnectorProperties) ActiveDirectoryConnectorProperties { return *v.Defaults() }).(ActiveDirectoryConnectorPropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220301preview:ActiveDirectoryConnector"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20220615preview:ActiveDirectoryConnector"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20230115preview:ActiveDirectoryConnector"),
		},
		{
			Type: pulumi.String("azure-native:azurearcdata/v20240101:ActiveDirectoryConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ActiveDirectoryConnector
	err := ctx.RegisterResource("azure-native:azurearcdata:ActiveDirectoryConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveDirectoryConnector gets an existing ActiveDirectoryConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveDirectoryConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveDirectoryConnectorState, opts ...pulumi.ResourceOption) (*ActiveDirectoryConnector, error) {
	var resource ActiveDirectoryConnector
	err := ctx.ReadResource("azure-native:azurearcdata:ActiveDirectoryConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveDirectoryConnector resources.
type activeDirectoryConnectorState struct {
}

type ActiveDirectoryConnectorState struct {
}

func (ActiveDirectoryConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryConnectorState)(nil)).Elem()
}

type activeDirectoryConnectorArgs struct {
	// The name of the Active Directory connector instance
	ActiveDirectoryConnectorName *string `pulumi:"activeDirectoryConnectorName"`
	// The name of the data controller
	DataControllerName string `pulumi:"dataControllerName"`
	// null
	Properties ActiveDirectoryConnectorProperties `pulumi:"properties"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ActiveDirectoryConnector resource.
type ActiveDirectoryConnectorArgs struct {
	// The name of the Active Directory connector instance
	ActiveDirectoryConnectorName pulumi.StringPtrInput
	// The name of the data controller
	DataControllerName pulumi.StringInput
	// null
	Properties ActiveDirectoryConnectorPropertiesInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
}

func (ActiveDirectoryConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeDirectoryConnectorArgs)(nil)).Elem()
}

type ActiveDirectoryConnectorInput interface {
	pulumi.Input

	ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput
	ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput
}

func (*ActiveDirectoryConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnector)(nil)).Elem()
}

func (i *ActiveDirectoryConnector) ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput {
	return i.ToActiveDirectoryConnectorOutputWithContext(context.Background())
}

func (i *ActiveDirectoryConnector) ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActiveDirectoryConnectorOutput)
}

type ActiveDirectoryConnectorOutput struct{ *pulumi.OutputState }

func (ActiveDirectoryConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActiveDirectoryConnector)(nil)).Elem()
}

func (o ActiveDirectoryConnectorOutput) ToActiveDirectoryConnectorOutput() ActiveDirectoryConnectorOutput {
	return o
}

func (o ActiveDirectoryConnectorOutput) ToActiveDirectoryConnectorOutputWithContext(ctx context.Context) ActiveDirectoryConnectorOutput {
	return o
}

// The name of the resource
func (o ActiveDirectoryConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// null
func (o ActiveDirectoryConnectorOutput) Properties() ActiveDirectoryConnectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) ActiveDirectoryConnectorPropertiesResponseOutput {
		return v.Properties
	}).(ActiveDirectoryConnectorPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ActiveDirectoryConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ActiveDirectoryConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ActiveDirectoryConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ActiveDirectoryConnectorOutput{})
}
