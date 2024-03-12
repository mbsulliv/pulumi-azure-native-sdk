// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Connects an edge site to an orbital gateway and describes what layer 2 traffic to forward between them.
type L2Connection struct {
	pulumi.CustomResourceState

	// Globally-unique identifier for this connection that is to be used as a circuit ID.
	CircuitId pulumi.StringOutput `pulumi:"circuitId"`
	// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
	EdgeSite L2ConnectionsPropertiesResponseEdgeSiteOutput `pulumi:"edgeSite"`
	// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
	GroundStation L2ConnectionsPropertiesResponseGroundStationOutput `pulumi:"groundStation"`
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
	// The VLAN ID for the L2 connection.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewL2Connection registers a new resource with the given unique name, arguments, and options.
func NewL2Connection(ctx *pulumi.Context,
	name string, args *L2ConnectionArgs, opts ...pulumi.ResourceOption) (*L2Connection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeSite == nil {
		return nil, errors.New("invalid value for required argument 'EdgeSite'")
	}
	if args.GroundStation == nil {
		return nil, errors.New("invalid value for required argument 'GroundStation'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:orbital:L2Connection"),
		},
		{
			Type: pulumi.String("azure-native:orbital/v20240301:L2Connection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource L2Connection
	err := ctx.RegisterResource("azure-native:orbital/v20240301preview:L2Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL2Connection gets an existing L2Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL2Connection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L2ConnectionState, opts ...pulumi.ResourceOption) (*L2Connection, error) {
	var resource L2Connection
	err := ctx.ReadResource("azure-native:orbital/v20240301preview:L2Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L2Connection resources.
type l2connectionState struct {
}

type L2ConnectionState struct {
}

func (L2ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*l2connectionState)(nil)).Elem()
}

type l2connectionArgs struct {
	// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
	EdgeSite L2ConnectionsPropertiesEdgeSite `pulumi:"edgeSite"`
	// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
	GroundStation L2ConnectionsPropertiesGroundStation `pulumi:"groundStation"`
	// L2 Connection name.
	L2ConnectionName *string `pulumi:"l2ConnectionName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The unique name of the partner router that cross-connects with the Orbital Edge Router at the ground station site.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VLAN ID for the L2 connection.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a L2Connection resource.
type L2ConnectionArgs struct {
	// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
	EdgeSite L2ConnectionsPropertiesEdgeSiteInput
	// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
	GroundStation L2ConnectionsPropertiesGroundStationInput
	// L2 Connection name.
	L2ConnectionName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The unique name of the partner router that cross-connects with the Orbital Edge Router at the ground station site.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VLAN ID for the L2 connection.
	VlanId pulumi.IntInput
}

func (L2ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l2connectionArgs)(nil)).Elem()
}

type L2ConnectionInput interface {
	pulumi.Input

	ToL2ConnectionOutput() L2ConnectionOutput
	ToL2ConnectionOutputWithContext(ctx context.Context) L2ConnectionOutput
}

func (*L2Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**L2Connection)(nil)).Elem()
}

func (i *L2Connection) ToL2ConnectionOutput() L2ConnectionOutput {
	return i.ToL2ConnectionOutputWithContext(context.Background())
}

func (i *L2Connection) ToL2ConnectionOutputWithContext(ctx context.Context) L2ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2ConnectionOutput)
}

type L2ConnectionOutput struct{ *pulumi.OutputState }

func (L2ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L2Connection)(nil)).Elem()
}

func (o L2ConnectionOutput) ToL2ConnectionOutput() L2ConnectionOutput {
	return o
}

func (o L2ConnectionOutput) ToL2ConnectionOutputWithContext(ctx context.Context) L2ConnectionOutput {
	return o
}

// Globally-unique identifier for this connection that is to be used as a circuit ID.
func (o L2ConnectionOutput) CircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.StringOutput { return v.CircuitId }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
func (o L2ConnectionOutput) EdgeSite() L2ConnectionsPropertiesResponseEdgeSiteOutput {
	return o.ApplyT(func(v *L2Connection) L2ConnectionsPropertiesResponseEdgeSiteOutput { return v.EdgeSite }).(L2ConnectionsPropertiesResponseEdgeSiteOutput)
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
func (o L2ConnectionOutput) GroundStation() L2ConnectionsPropertiesResponseGroundStationOutput {
	return o.ApplyT(func(v *L2Connection) L2ConnectionsPropertiesResponseGroundStationOutput { return v.GroundStation }).(L2ConnectionsPropertiesResponseGroundStationOutput)
}

// The geo-location where the resource lives
func (o L2ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o L2ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o L2ConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *L2Connection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o L2ConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o L2ConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VLAN ID for the L2 connection.
func (o L2ConnectionOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *L2Connection) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(L2ConnectionOutput{})
}
