// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type Creator struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Creator resource properties.
	Properties CreatorPropertiesResponseOutput `pulumi:"properties"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCreator registers a new resource with the given unique name, arguments, and options.
func NewCreator(ctx *pulumi.Context,
	name string, args *CreatorArgs, opts ...pulumi.ResourceOption) (*Creator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maps:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20200201preview:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210201:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20210701preview:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20211201preview:Creator"),
		},
		{
			Type: pulumi.String("azure-native:maps/v20230801preview:Creator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Creator
	err := ctx.RegisterResource("azure-native:maps/v20230601:Creator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCreator gets an existing Creator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCreator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CreatorState, opts ...pulumi.ResourceOption) (*Creator, error) {
	var resource Creator
	err := ctx.ReadResource("azure-native:maps/v20230601:Creator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Creator resources.
type creatorState struct {
}

type CreatorState struct {
}

func (CreatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorState)(nil)).Elem()
}

type creatorArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the Maps Creator instance.
	CreatorName *string `pulumi:"creatorName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The Creator resource properties.
	Properties CreatorProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Creator resource.
type CreatorArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput
	// The name of the Maps Creator instance.
	CreatorName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The Creator resource properties.
	Properties CreatorPropertiesInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CreatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorArgs)(nil)).Elem()
}

type CreatorInput interface {
	pulumi.Input

	ToCreatorOutput() CreatorOutput
	ToCreatorOutputWithContext(ctx context.Context) CreatorOutput
}

func (*Creator) ElementType() reflect.Type {
	return reflect.TypeOf((**Creator)(nil)).Elem()
}

func (i *Creator) ToCreatorOutput() CreatorOutput {
	return i.ToCreatorOutputWithContext(context.Background())
}

func (i *Creator) ToCreatorOutputWithContext(ctx context.Context) CreatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatorOutput)
}

type CreatorOutput struct{ *pulumi.OutputState }

func (CreatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Creator)(nil)).Elem()
}

func (o CreatorOutput) ToCreatorOutput() CreatorOutput {
	return o
}

func (o CreatorOutput) ToCreatorOutputWithContext(ctx context.Context) CreatorOutput {
	return o
}

// The geo-location where the resource lives
func (o CreatorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Creator) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CreatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Creator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Creator resource properties.
func (o CreatorOutput) Properties() CreatorPropertiesResponseOutput {
	return o.ApplyT(func(v *Creator) CreatorPropertiesResponseOutput { return v.Properties }).(CreatorPropertiesResponseOutput)
}

// The system meta data relating to this resource.
func (o CreatorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Creator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CreatorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Creator) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CreatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Creator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CreatorOutput{})
}
