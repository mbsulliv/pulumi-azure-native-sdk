// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// network service design version.
type NetworkServiceDesignVersion struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// network service design version properties.
	Properties NetworkServiceDesignVersionPropertiesFormatResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkServiceDesignVersion registers a new resource with the given unique name, arguments, and options.
func NewNetworkServiceDesignVersion(ctx *pulumi.Context,
	name string, args *NetworkServiceDesignVersionArgs, opts ...pulumi.ResourceOption) (*NetworkServiceDesignVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkServiceDesignGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkServiceDesignGroupName'")
	}
	if args.PublisherName == nil {
		return nil, errors.New("invalid value for required argument 'PublisherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork:NetworkServiceDesignVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkServiceDesignVersion
	err := ctx.RegisterResource("azure-native:hybridnetwork/v20230901:NetworkServiceDesignVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkServiceDesignVersion gets an existing NetworkServiceDesignVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkServiceDesignVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkServiceDesignVersionState, opts ...pulumi.ResourceOption) (*NetworkServiceDesignVersion, error) {
	var resource NetworkServiceDesignVersion
	err := ctx.ReadResource("azure-native:hybridnetwork/v20230901:NetworkServiceDesignVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkServiceDesignVersion resources.
type networkServiceDesignVersionState struct {
}

type NetworkServiceDesignVersionState struct {
}

func (NetworkServiceDesignVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServiceDesignVersionState)(nil)).Elem()
}

type networkServiceDesignVersionArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the network service design group.
	NetworkServiceDesignGroupName string `pulumi:"networkServiceDesignGroupName"`
	// The name of the network service design version. The name should conform to the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
	NetworkServiceDesignVersionName *string `pulumi:"networkServiceDesignVersionName"`
	// network service design version properties.
	Properties *NetworkServiceDesignVersionPropertiesFormat `pulumi:"properties"`
	// The name of the publisher.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkServiceDesignVersion resource.
type NetworkServiceDesignVersionArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the network service design group.
	NetworkServiceDesignGroupName pulumi.StringInput
	// The name of the network service design version. The name should conform to the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
	NetworkServiceDesignVersionName pulumi.StringPtrInput
	// network service design version properties.
	Properties NetworkServiceDesignVersionPropertiesFormatPtrInput
	// The name of the publisher.
	PublisherName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkServiceDesignVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServiceDesignVersionArgs)(nil)).Elem()
}

type NetworkServiceDesignVersionInput interface {
	pulumi.Input

	ToNetworkServiceDesignVersionOutput() NetworkServiceDesignVersionOutput
	ToNetworkServiceDesignVersionOutputWithContext(ctx context.Context) NetworkServiceDesignVersionOutput
}

func (*NetworkServiceDesignVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServiceDesignVersion)(nil)).Elem()
}

func (i *NetworkServiceDesignVersion) ToNetworkServiceDesignVersionOutput() NetworkServiceDesignVersionOutput {
	return i.ToNetworkServiceDesignVersionOutputWithContext(context.Background())
}

func (i *NetworkServiceDesignVersion) ToNetworkServiceDesignVersionOutputWithContext(ctx context.Context) NetworkServiceDesignVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkServiceDesignVersionOutput)
}

type NetworkServiceDesignVersionOutput struct{ *pulumi.OutputState }

func (NetworkServiceDesignVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServiceDesignVersion)(nil)).Elem()
}

func (o NetworkServiceDesignVersionOutput) ToNetworkServiceDesignVersionOutput() NetworkServiceDesignVersionOutput {
	return o
}

func (o NetworkServiceDesignVersionOutput) ToNetworkServiceDesignVersionOutputWithContext(ctx context.Context) NetworkServiceDesignVersionOutput {
	return o
}

// The geo-location where the resource lives
func (o NetworkServiceDesignVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NetworkServiceDesignVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// network service design version properties.
func (o NetworkServiceDesignVersionOutput) Properties() NetworkServiceDesignVersionPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) NetworkServiceDesignVersionPropertiesFormatResponseOutput {
		return v.Properties
	}).(NetworkServiceDesignVersionPropertiesFormatResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkServiceDesignVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkServiceDesignVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkServiceDesignVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServiceDesignVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkServiceDesignVersionOutput{})
}