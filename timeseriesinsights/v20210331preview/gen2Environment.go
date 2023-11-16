// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210331preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource. Gen2 environments do not have set data retention limits.
type Gen2Environment struct {
	pulumi.CustomResourceState

	// The time the resource was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessFqdn pulumi.StringOutput `pulumi:"dataAccessFqdn"`
	// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
	DataAccessId pulumi.StringOutput `pulumi:"dataAccessId"`
	// The kind of the environment.
	// Expected value is 'Gen2'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of private endpoint connections to the environment.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// If 'enabled', public network access is allowed. If 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku SkuResponseOutput `pulumi:"sku"`
	// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
	Status EnvironmentStatusResponseOutput `pulumi:"status"`
	// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
	StorageConfiguration Gen2StorageConfigurationOutputResponseOutput `pulumi:"storageConfiguration"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The list of event properties which will be used to define the environment's time series id.
	TimeSeriesIdProperties TimeSeriesIdPropertyResponseArrayOutput `pulumi:"timeSeriesIdProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
	WarmStoreConfiguration WarmStoreConfigurationPropertiesResponsePtrOutput `pulumi:"warmStoreConfiguration"`
}

// NewGen2Environment registers a new resource with the given unique name, arguments, and options.
func NewGen2Environment(ctx *pulumi.Context,
	name string, args *Gen2EnvironmentArgs, opts ...pulumi.ResourceOption) (*Gen2Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.StorageConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'StorageConfiguration'")
	}
	if args.TimeSeriesIdProperties == nil {
		return nil, errors.New("invalid value for required argument 'TimeSeriesIdProperties'")
	}
	args.Kind = pulumi.String("Gen2")
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:Gen2Environment"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210630preview:Gen2Environment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Gen2Environment
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20210331preview:Gen2Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGen2Environment gets an existing Gen2Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGen2Environment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Gen2EnvironmentState, opts ...pulumi.ResourceOption) (*Gen2Environment, error) {
	var resource Gen2Environment
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20210331preview:Gen2Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gen2Environment resources.
type gen2EnvironmentState struct {
}

type Gen2EnvironmentState struct {
}

func (Gen2EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*gen2EnvironmentState)(nil)).Elem()
}

type gen2EnvironmentArgs struct {
	// Name of the environment
	EnvironmentName *string `pulumi:"environmentName"`
	// The kind of the environment.
	// Expected value is 'Gen2'.
	Kind string `pulumi:"kind"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku Sku `pulumi:"sku"`
	// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
	StorageConfiguration Gen2StorageConfigurationInput `pulumi:"storageConfiguration"`
	// Key-value pairs of additional properties for the resource.
	Tags map[string]string `pulumi:"tags"`
	// The list of event properties which will be used to define the environment's time series id.
	TimeSeriesIdProperties []TimeSeriesIdProperty `pulumi:"timeSeriesIdProperties"`
	// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
	WarmStoreConfiguration *WarmStoreConfigurationProperties `pulumi:"warmStoreConfiguration"`
}

// The set of arguments for constructing a Gen2Environment resource.
type Gen2EnvironmentArgs struct {
	// Name of the environment
	EnvironmentName pulumi.StringPtrInput
	// The kind of the environment.
	// Expected value is 'Gen2'.
	Kind pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
	PublicNetworkAccess pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
	Sku SkuInput
	// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
	StorageConfiguration Gen2StorageConfigurationInputInput
	// Key-value pairs of additional properties for the resource.
	Tags pulumi.StringMapInput
	// The list of event properties which will be used to define the environment's time series id.
	TimeSeriesIdProperties TimeSeriesIdPropertyArrayInput
	// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
	WarmStoreConfiguration WarmStoreConfigurationPropertiesPtrInput
}

func (Gen2EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gen2EnvironmentArgs)(nil)).Elem()
}

type Gen2EnvironmentInput interface {
	pulumi.Input

	ToGen2EnvironmentOutput() Gen2EnvironmentOutput
	ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput
}

func (*Gen2Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Gen2Environment)(nil)).Elem()
}

func (i *Gen2Environment) ToGen2EnvironmentOutput() Gen2EnvironmentOutput {
	return i.ToGen2EnvironmentOutputWithContext(context.Background())
}

func (i *Gen2Environment) ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Gen2EnvironmentOutput)
}

type Gen2EnvironmentOutput struct{ *pulumi.OutputState }

func (Gen2EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gen2Environment)(nil)).Elem()
}

func (o Gen2EnvironmentOutput) ToGen2EnvironmentOutput() Gen2EnvironmentOutput {
	return o
}

func (o Gen2EnvironmentOutput) ToGen2EnvironmentOutputWithContext(ctx context.Context) Gen2EnvironmentOutput {
	return o
}

// The time the resource was created.
func (o Gen2EnvironmentOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o Gen2EnvironmentOutput) DataAccessFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.DataAccessFqdn }).(pulumi.StringOutput)
}

// An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
func (o Gen2EnvironmentOutput) DataAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.DataAccessId }).(pulumi.StringOutput)
}

// The kind of the environment.
// Expected value is 'Gen2'.
func (o Gen2EnvironmentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location
func (o Gen2EnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o Gen2EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections to the environment.
func (o Gen2EnvironmentOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Gen2Environment) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the resource.
func (o Gen2EnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// If 'enabled', public network access is allowed. If 'disabled', traffic over public interface is not allowed, and private endpoint connections would be the exclusive access method.
func (o Gen2EnvironmentOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
func (o Gen2EnvironmentOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Gen2Environment) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// An object that represents the status of the environment, and its internal state in the Time Series Insights service.
func (o Gen2EnvironmentOutput) Status() EnvironmentStatusResponseOutput {
	return o.ApplyT(func(v *Gen2Environment) EnvironmentStatusResponseOutput { return v.Status }).(EnvironmentStatusResponseOutput)
}

// The storage configuration provides the connection details that allows the Time Series Insights service to connect to the customer storage account that is used to store the environment's data.
func (o Gen2EnvironmentOutput) StorageConfiguration() Gen2StorageConfigurationOutputResponseOutput {
	return o.ApplyT(func(v *Gen2Environment) Gen2StorageConfigurationOutputResponseOutput { return v.StorageConfiguration }).(Gen2StorageConfigurationOutputResponseOutput)
}

// Resource tags
func (o Gen2EnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The list of event properties which will be used to define the environment's time series id.
func (o Gen2EnvironmentOutput) TimeSeriesIdProperties() TimeSeriesIdPropertyResponseArrayOutput {
	return o.ApplyT(func(v *Gen2Environment) TimeSeriesIdPropertyResponseArrayOutput { return v.TimeSeriesIdProperties }).(TimeSeriesIdPropertyResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o Gen2EnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gen2Environment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The warm store configuration provides the details to create a warm store cache that will retain a copy of the environment's data available for faster query.
func (o Gen2EnvironmentOutput) WarmStoreConfiguration() WarmStoreConfigurationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Gen2Environment) WarmStoreConfigurationPropertiesResponsePtrOutput {
		return v.WarmStoreConfiguration
	}).(WarmStoreConfigurationPropertiesResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(Gen2EnvironmentOutput{})
}
