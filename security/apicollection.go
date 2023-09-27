// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An API collection as represented by Defender for APIs.
// Azure REST API version: 2022-11-20-preview.
type APICollection struct {
	pulumi.CustomResourceState

	// Additional data regarding the API collection.
	AdditionalData pulumi.StringMapOutput `pulumi:"additionalData"`
	// The display name of the Azure API Management API.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAPICollection registers a new resource with the given unique name, arguments, and options.
func NewAPICollection(ctx *pulumi.Context,
	name string, args *APICollectionArgs, opts ...pulumi.ResourceOption) (*APICollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20221120preview:APICollection"),
		},
		{
			Type: pulumi.String("azure-native:security/v20231115:APICollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource APICollection
	err := ctx.RegisterResource("azure-native:security:APICollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPICollection gets an existing APICollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPICollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APICollectionState, opts ...pulumi.ResourceOption) (*APICollection, error) {
	var resource APICollection
	err := ctx.ReadResource("azure-native:security:APICollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APICollection resources.
type apicollectionState struct {
}

type APICollectionState struct {
}

func (APICollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionState)(nil)).Elem()
}

type apicollectionArgs struct {
	// A string representing the apiCollections resource within the Microsoft.Security provider namespace. This string matches the Azure API Management API name.
	ApiCollectionId *string `pulumi:"apiCollectionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a APICollection resource.
type APICollectionArgs struct {
	// A string representing the apiCollections resource within the Microsoft.Security provider namespace. This string matches the Azure API Management API name.
	ApiCollectionId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (APICollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apicollectionArgs)(nil)).Elem()
}

type APICollectionInput interface {
	pulumi.Input

	ToAPICollectionOutput() APICollectionOutput
	ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput
}

func (*APICollection) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollection)(nil)).Elem()
}

func (i *APICollection) ToAPICollectionOutput() APICollectionOutput {
	return i.ToAPICollectionOutputWithContext(context.Background())
}

func (i *APICollection) ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APICollectionOutput)
}

func (i *APICollection) ToOutput(ctx context.Context) pulumix.Output[*APICollection] {
	return pulumix.Output[*APICollection]{
		OutputState: i.ToAPICollectionOutputWithContext(ctx).OutputState,
	}
}

type APICollectionOutput struct{ *pulumi.OutputState }

func (APICollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APICollection)(nil)).Elem()
}

func (o APICollectionOutput) ToAPICollectionOutput() APICollectionOutput {
	return o
}

func (o APICollectionOutput) ToAPICollectionOutputWithContext(ctx context.Context) APICollectionOutput {
	return o
}

func (o APICollectionOutput) ToOutput(ctx context.Context) pulumix.Output[*APICollection] {
	return pulumix.Output[*APICollection]{
		OutputState: o.OutputState,
	}
}

// Additional data regarding the API collection.
func (o APICollectionOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringMapOutput { return v.AdditionalData }).(pulumi.StringMapOutput)
}

// The display name of the Azure API Management API.
func (o APICollectionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name
func (o APICollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o APICollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *APICollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(APICollectionOutput{})
}
