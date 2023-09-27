// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Information about the connection monitor.
type ConnectionMonitor struct {
	pulumi.CustomResourceState

	// Determines if the connection monitor will start automatically once created.
	AutoStart pulumi.BoolPtrOutput `pulumi:"autoStart"`
	// Type of connection monitor.
	ConnectionMonitorType pulumi.StringOutput `pulumi:"connectionMonitorType"`
	// Describes the destination of connection monitor.
	Destination ConnectionMonitorDestinationResponsePtrOutput `pulumi:"destination"`
	// List of connection monitor endpoints.
	Endpoints ConnectionMonitorEndpointResponseArrayOutput `pulumi:"endpoints"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Connection monitor location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds pulumi.IntPtrOutput `pulumi:"monitoringIntervalInSeconds"`
	// The monitoring status of the connection monitor.
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// Name of the connection monitor.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional notes to be associated with the connection monitor.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// List of connection monitor outputs.
	Outputs ConnectionMonitorOutputResponseArrayOutput `pulumi:"outputs"`
	// The provisioning state of the connection monitor.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the source of connection monitor.
	Source ConnectionMonitorSourceResponsePtrOutput `pulumi:"source"`
	// The date and time when the connection monitor was started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Connection monitor tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// List of connection monitor test configurations.
	TestConfigurations ConnectionMonitorTestConfigurationResponseArrayOutput `pulumi:"testConfigurations"`
	// List of connection monitor test groups.
	TestGroups ConnectionMonitorTestGroupResponseArrayOutput `pulumi:"testGroups"`
	// Connection monitor type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionMonitor registers a new resource with the given unique name, arguments, and options.
func NewConnectionMonitor(ctx *pulumi.Context,
	name string, args *ConnectionMonitorArgs, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkWatcherName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkWatcherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AutoStart == nil {
		args.AutoStart = pulumi.BoolPtr(true)
	}
	if args.MonitoringIntervalInSeconds == nil {
		args.MonitoringIntervalInSeconds = pulumi.IntPtr(60)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ConnectionMonitor"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ConnectionMonitor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectionMonitor
	err := ctx.RegisterResource("azure-native:network/v20230201:ConnectionMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionMonitor gets an existing ConnectionMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionMonitorState, opts ...pulumi.ResourceOption) (*ConnectionMonitor, error) {
	var resource ConnectionMonitor
	err := ctx.ReadResource("azure-native:network/v20230201:ConnectionMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionMonitor resources.
type connectionMonitorState struct {
}

type ConnectionMonitorState struct {
}

func (ConnectionMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorState)(nil)).Elem()
}

type connectionMonitorArgs struct {
	// Determines if the connection monitor will start automatically once created.
	AutoStart *bool `pulumi:"autoStart"`
	// The name of the connection monitor.
	ConnectionMonitorName *string `pulumi:"connectionMonitorName"`
	// Describes the destination of connection monitor.
	Destination *ConnectionMonitorDestination `pulumi:"destination"`
	// List of connection monitor endpoints.
	Endpoints []ConnectionMonitorEndpoint `pulumi:"endpoints"`
	// Connection monitor location.
	Location *string `pulumi:"location"`
	// Value indicating whether connection monitor V1 should be migrated to V2 format.
	Migrate *string `pulumi:"migrate"`
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds *int `pulumi:"monitoringIntervalInSeconds"`
	// The name of the Network Watcher resource.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// Optional notes to be associated with the connection monitor.
	Notes *string `pulumi:"notes"`
	// List of connection monitor outputs.
	Outputs []ConnectionMonitorOutputType `pulumi:"outputs"`
	// The name of the resource group containing Network Watcher.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the source of connection monitor.
	Source *ConnectionMonitorSource `pulumi:"source"`
	// Connection monitor tags.
	Tags map[string]string `pulumi:"tags"`
	// List of connection monitor test configurations.
	TestConfigurations []ConnectionMonitorTestConfiguration `pulumi:"testConfigurations"`
	// List of connection monitor test groups.
	TestGroups []ConnectionMonitorTestGroup `pulumi:"testGroups"`
}

// The set of arguments for constructing a ConnectionMonitor resource.
type ConnectionMonitorArgs struct {
	// Determines if the connection monitor will start automatically once created.
	AutoStart pulumi.BoolPtrInput
	// The name of the connection monitor.
	ConnectionMonitorName pulumi.StringPtrInput
	// Describes the destination of connection monitor.
	Destination ConnectionMonitorDestinationPtrInput
	// List of connection monitor endpoints.
	Endpoints ConnectionMonitorEndpointArrayInput
	// Connection monitor location.
	Location pulumi.StringPtrInput
	// Value indicating whether connection monitor V1 should be migrated to V2 format.
	Migrate pulumi.StringPtrInput
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds pulumi.IntPtrInput
	// The name of the Network Watcher resource.
	NetworkWatcherName pulumi.StringInput
	// Optional notes to be associated with the connection monitor.
	Notes pulumi.StringPtrInput
	// List of connection monitor outputs.
	Outputs ConnectionMonitorOutputTypeArrayInput
	// The name of the resource group containing Network Watcher.
	ResourceGroupName pulumi.StringInput
	// Describes the source of connection monitor.
	Source ConnectionMonitorSourcePtrInput
	// Connection monitor tags.
	Tags pulumi.StringMapInput
	// List of connection monitor test configurations.
	TestConfigurations ConnectionMonitorTestConfigurationArrayInput
	// List of connection monitor test groups.
	TestGroups ConnectionMonitorTestGroupArrayInput
}

func (ConnectionMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionMonitorArgs)(nil)).Elem()
}

type ConnectionMonitorInput interface {
	pulumi.Input

	ToConnectionMonitorOutput() ConnectionMonitorOutput
	ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput
}

func (*ConnectionMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitor)(nil)).Elem()
}

func (i *ConnectionMonitor) ToConnectionMonitorOutput() ConnectionMonitorOutput {
	return i.ToConnectionMonitorOutputWithContext(context.Background())
}

func (i *ConnectionMonitor) ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionMonitorOutput)
}

func (i *ConnectionMonitor) ToOutput(ctx context.Context) pulumix.Output[*ConnectionMonitor] {
	return pulumix.Output[*ConnectionMonitor]{
		OutputState: i.ToConnectionMonitorOutputWithContext(ctx).OutputState,
	}
}

type ConnectionMonitorOutput struct{ *pulumi.OutputState }

func (ConnectionMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionMonitor)(nil)).Elem()
}

func (o ConnectionMonitorOutput) ToConnectionMonitorOutput() ConnectionMonitorOutput {
	return o
}

func (o ConnectionMonitorOutput) ToConnectionMonitorOutputWithContext(ctx context.Context) ConnectionMonitorOutput {
	return o
}

func (o ConnectionMonitorOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectionMonitor] {
	return pulumix.Output[*ConnectionMonitor]{
		OutputState: o.OutputState,
	}
}

// Determines if the connection monitor will start automatically once created.
func (o ConnectionMonitorOutput) AutoStart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.BoolPtrOutput { return v.AutoStart }).(pulumi.BoolPtrOutput)
}

// Type of connection monitor.
func (o ConnectionMonitorOutput) ConnectionMonitorType() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.ConnectionMonitorType }).(pulumi.StringOutput)
}

// Describes the destination of connection monitor.
func (o ConnectionMonitorOutput) Destination() ConnectionMonitorDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorDestinationResponsePtrOutput { return v.Destination }).(ConnectionMonitorDestinationResponsePtrOutput)
}

// List of connection monitor endpoints.
func (o ConnectionMonitorOutput) Endpoints() ConnectionMonitorEndpointResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorEndpointResponseArrayOutput { return v.Endpoints }).(ConnectionMonitorEndpointResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ConnectionMonitorOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Connection monitor location.
func (o ConnectionMonitorOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Monitoring interval in seconds.
func (o ConnectionMonitorOutput) MonitoringIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.IntPtrOutput { return v.MonitoringIntervalInSeconds }).(pulumi.IntPtrOutput)
}

// The monitoring status of the connection monitor.
func (o ConnectionMonitorOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.MonitoringStatus }).(pulumi.StringOutput)
}

// Name of the connection monitor.
func (o ConnectionMonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional notes to be associated with the connection monitor.
func (o ConnectionMonitorOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// List of connection monitor outputs.
func (o ConnectionMonitorOutput) Outputs() ConnectionMonitorOutputResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorOutputResponseArrayOutput { return v.Outputs }).(ConnectionMonitorOutputResponseArrayOutput)
}

// The provisioning state of the connection monitor.
func (o ConnectionMonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the source of connection monitor.
func (o ConnectionMonitorOutput) Source() ConnectionMonitorSourceResponsePtrOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorSourceResponsePtrOutput { return v.Source }).(ConnectionMonitorSourceResponsePtrOutput)
}

// The date and time when the connection monitor was started.
func (o ConnectionMonitorOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Connection monitor tags.
func (o ConnectionMonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// List of connection monitor test configurations.
func (o ConnectionMonitorOutput) TestConfigurations() ConnectionMonitorTestConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorTestConfigurationResponseArrayOutput {
		return v.TestConfigurations
	}).(ConnectionMonitorTestConfigurationResponseArrayOutput)
}

// List of connection monitor test groups.
func (o ConnectionMonitorOutput) TestGroups() ConnectionMonitorTestGroupResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionMonitor) ConnectionMonitorTestGroupResponseArrayOutput { return v.TestGroups }).(ConnectionMonitorTestGroupResponseArrayOutput)
}

// Connection monitor type.
func (o ConnectionMonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionMonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionMonitorOutput{})
}
