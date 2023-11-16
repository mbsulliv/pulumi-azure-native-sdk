// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration account map.
// Azure REST API version: 2019-05-01. Prior API version in Azure Native 1.x: 2019-05-01.
//
// Other available API versions: 2015-08-01-preview.
type IntegrationAccountMap struct {
	pulumi.CustomResourceState

	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The content.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponseOutput `pulumi:"contentLink"`
	// The content type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The map type.
	MapType pulumi.StringOutput `pulumi:"mapType"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput `pulumi:"parametersSchema"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountMap registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountMap(ctx *pulumi.Context,
	name string, args *IntegrationAccountMapArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.MapType == nil {
		return nil, errors.New("invalid value for required argument 'MapType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountMap"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationAccountMap
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountMap gets an existing IntegrationAccountMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountMapState, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	var resource IntegrationAccountMap
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountMap resources.
type integrationAccountMapState struct {
}

type IntegrationAccountMapState struct {
}

func (IntegrationAccountMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapState)(nil)).Elem()
}

type integrationAccountMapArgs struct {
	// The content.
	Content *string `pulumi:"content"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration account map name.
	MapName *string `pulumi:"mapName"`
	// The map type.
	MapType string `pulumi:"mapType"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The parameters schema of integration account map.
	ParametersSchema *IntegrationAccountMapPropertiesParametersSchema `pulumi:"parametersSchema"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountMap resource.
type IntegrationAccountMapArgs struct {
	// The content.
	Content pulumi.StringPtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration account map name.
	MapName pulumi.StringPtrInput
	// The map type.
	MapType pulumi.StringInput
	// The metadata.
	Metadata pulumi.Input
	// The parameters schema of integration account map.
	ParametersSchema IntegrationAccountMapPropertiesParametersSchemaPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapArgs)(nil)).Elem()
}

type IntegrationAccountMapInput interface {
	pulumi.Input

	ToIntegrationAccountMapOutput() IntegrationAccountMapOutput
	ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput
}

func (*IntegrationAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return i.ToIntegrationAccountMapOutputWithContext(context.Background())
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapOutput)
}

type IntegrationAccountMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return o
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return o
}

// The changed time.
func (o IntegrationAccountMapOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// The content.
func (o IntegrationAccountMapOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// The content link.
func (o IntegrationAccountMapOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) ContentLinkResponseOutput { return v.ContentLink }).(ContentLinkResponseOutput)
}

// The content type.
func (o IntegrationAccountMapOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o IntegrationAccountMapOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The resource location.
func (o IntegrationAccountMapOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The map type.
func (o IntegrationAccountMapOutput) MapType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.MapType }).(pulumi.StringOutput)
}

// The metadata.
func (o IntegrationAccountMapOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o IntegrationAccountMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parameters schema of integration account map.
func (o IntegrationAccountMapOutput) ParametersSchema() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
		return v.ParametersSchema
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput)
}

// The resource tags.
func (o IntegrationAccountMapOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o IntegrationAccountMapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountMap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountMapOutput{})
}
