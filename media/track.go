// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Asset Track resource.
// Azure REST API version: 2023-01-01. Prior API version in Azure Native 1.x: 2021-11-01.
type Track struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the asset track.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Detailed information about a track in the asset.
	Track pulumi.AnyOutput `pulumi:"track"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTrack registers a new resource with the given unique name, arguments, and options.
func NewTrack(ctx *pulumi.Context,
	name string, args *TrackArgs, opts ...pulumi.ResourceOption) (*Track, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media/v20211101:Track"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:Track"),
		},
		{
			Type: pulumi.String("azure-native:media/v20230101:Track"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Track
	err := ctx.RegisterResource("azure-native:media:Track", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrack gets an existing Track resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrackState, opts ...pulumi.ResourceOption) (*Track, error) {
	var resource Track
	err := ctx.ReadResource("azure-native:media:Track", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Track resources.
type trackState struct {
}

type TrackState struct {
}

func (TrackState) ElementType() reflect.Type {
	return reflect.TypeOf((*trackState)(nil)).Elem()
}

type trackArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Detailed information about a track in the asset.
	Track interface{} `pulumi:"track"`
	// The Asset Track name.
	TrackName *string `pulumi:"trackName"`
}

// The set of arguments for constructing a Track resource.
type TrackArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Asset name.
	AssetName pulumi.StringInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Detailed information about a track in the asset.
	Track pulumi.Input
	// The Asset Track name.
	TrackName pulumi.StringPtrInput
}

func (TrackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trackArgs)(nil)).Elem()
}

type TrackInput interface {
	pulumi.Input

	ToTrackOutput() TrackOutput
	ToTrackOutputWithContext(ctx context.Context) TrackOutput
}

func (*Track) ElementType() reflect.Type {
	return reflect.TypeOf((**Track)(nil)).Elem()
}

func (i *Track) ToTrackOutput() TrackOutput {
	return i.ToTrackOutputWithContext(context.Background())
}

func (i *Track) ToTrackOutputWithContext(ctx context.Context) TrackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackOutput)
}

type TrackOutput struct{ *pulumi.OutputState }

func (TrackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Track)(nil)).Elem()
}

func (o TrackOutput) ToTrackOutput() TrackOutput {
	return o
}

func (o TrackOutput) ToTrackOutputWithContext(ctx context.Context) TrackOutput {
	return o
}

// The name of the resource
func (o TrackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the asset track.
func (o TrackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Detailed information about a track in the asset.
func (o TrackOutput) Track() pulumi.AnyOutput {
	return o.ApplyT(func(v *Track) pulumi.AnyOutput { return v.Track }).(pulumi.AnyOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TrackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Track) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrackOutput{})
}
