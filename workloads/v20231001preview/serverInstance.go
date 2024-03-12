// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the Server Instance resource.
type ServerInstance struct {
	pulumi.CustomResourceState

	// Configuration data for this server instance.
	ConfigurationData ConfigurationDataResponseOutput `pulumi:"configurationData"`
	// Defines the errors related to SAP Instance resource.
	Errors SAPMigrateErrorResponseOutput `pulumi:"errors"`
	// This is the Instance SID for ASCS/AP/DB instance.  An SAP system with HANA database for example could have a different SID for database Instance than that of ASCS instance.
	InstanceSid pulumi.StringOutput `pulumi:"instanceSid"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// This is Operating System on which the host server is running.
	OperatingSystem pulumi.StringOutput `pulumi:"operatingSystem"`
	// Configuration data for this server instance.
	PerformanceData pulumi.AnyOutput `pulumi:"performanceData"`
	// Defines the provisioning states.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the type SAP instance on this server instance.
	SapInstanceType pulumi.StringOutput `pulumi:"sapInstanceType"`
	// This is the SAP Application Component; e.g. SAP S/4HANA 2022, SAP ERP ENHANCE PACKAGE.
	SapProduct pulumi.StringOutput `pulumi:"sapProduct"`
	// Provide the product version of the SAP product.
	SapProductVersion pulumi.StringOutput `pulumi:"sapProductVersion"`
	// This is the Virtual Machine Name of the SAP system. Add all the virtual machines attached to an SAP system which you wish to migrate to Azure. Keeping this not equal to ID as for single tier all InstanceTypes would be on same server, leading to multiple resources with same servername.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerInstance registers a new resource with the given unique name, arguments, and options.
func NewServerInstance(ctx *pulumi.Context,
	name string, args *ServerInstanceArgs, opts ...pulumi.ResourceOption) (*ServerInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapDiscoverySiteName == nil {
		return nil, errors.New("invalid value for required argument 'SapDiscoverySiteName'")
	}
	if args.SapInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SapInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:ServerInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerInstance
	err := ctx.RegisterResource("azure-native:workloads/v20231001preview:ServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerInstance gets an existing ServerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerInstanceState, opts ...pulumi.ResourceOption) (*ServerInstance, error) {
	var resource ServerInstance
	err := ctx.ReadResource("azure-native:workloads/v20231001preview:ServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerInstance resources.
type serverInstanceState struct {
}

type ServerInstanceState struct {
}

func (ServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverInstanceState)(nil)).Elem()
}

type serverInstanceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName string `pulumi:"sapDiscoverySiteName"`
	// The name of SAP Instance resource for SAP Migration.
	SapInstanceName string `pulumi:"sapInstanceName"`
	// The name of the Server instance resource for SAP Migration.
	ServerInstanceName *string `pulumi:"serverInstanceName"`
}

// The set of arguments for constructing a ServerInstance resource.
type ServerInstanceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName pulumi.StringInput
	// The name of SAP Instance resource for SAP Migration.
	SapInstanceName pulumi.StringInput
	// The name of the Server instance resource for SAP Migration.
	ServerInstanceName pulumi.StringPtrInput
}

func (ServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverInstanceArgs)(nil)).Elem()
}

type ServerInstanceInput interface {
	pulumi.Input

	ToServerInstanceOutput() ServerInstanceOutput
	ToServerInstanceOutputWithContext(ctx context.Context) ServerInstanceOutput
}

func (*ServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerInstance)(nil)).Elem()
}

func (i *ServerInstance) ToServerInstanceOutput() ServerInstanceOutput {
	return i.ToServerInstanceOutputWithContext(context.Background())
}

func (i *ServerInstance) ToServerInstanceOutputWithContext(ctx context.Context) ServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerInstanceOutput)
}

type ServerInstanceOutput struct{ *pulumi.OutputState }

func (ServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerInstance)(nil)).Elem()
}

func (o ServerInstanceOutput) ToServerInstanceOutput() ServerInstanceOutput {
	return o
}

func (o ServerInstanceOutput) ToServerInstanceOutputWithContext(ctx context.Context) ServerInstanceOutput {
	return o
}

// Configuration data for this server instance.
func (o ServerInstanceOutput) ConfigurationData() ConfigurationDataResponseOutput {
	return o.ApplyT(func(v *ServerInstance) ConfigurationDataResponseOutput { return v.ConfigurationData }).(ConfigurationDataResponseOutput)
}

// Defines the errors related to SAP Instance resource.
func (o ServerInstanceOutput) Errors() SAPMigrateErrorResponseOutput {
	return o.ApplyT(func(v *ServerInstance) SAPMigrateErrorResponseOutput { return v.Errors }).(SAPMigrateErrorResponseOutput)
}

// This is the Instance SID for ASCS/AP/DB instance.  An SAP system with HANA database for example could have a different SID for database Instance than that of ASCS instance.
func (o ServerInstanceOutput) InstanceSid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.InstanceSid }).(pulumi.StringOutput)
}

// The name of the resource
func (o ServerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This is Operating System on which the host server is running.
func (o ServerInstanceOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.OperatingSystem }).(pulumi.StringOutput)
}

// Configuration data for this server instance.
func (o ServerInstanceOutput) PerformanceData() pulumi.AnyOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.AnyOutput { return v.PerformanceData }).(pulumi.AnyOutput)
}

// Defines the provisioning states.
func (o ServerInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the type SAP instance on this server instance.
func (o ServerInstanceOutput) SapInstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.SapInstanceType }).(pulumi.StringOutput)
}

// This is the SAP Application Component; e.g. SAP S/4HANA 2022, SAP ERP ENHANCE PACKAGE.
func (o ServerInstanceOutput) SapProduct() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.SapProduct }).(pulumi.StringOutput)
}

// Provide the product version of the SAP product.
func (o ServerInstanceOutput) SapProductVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.SapProductVersion }).(pulumi.StringOutput)
}

// This is the Virtual Machine Name of the SAP system. Add all the virtual machines attached to an SAP system which you wish to migrate to Azure. Keeping this not equal to ID as for single tier all InstanceTypes would be on same server, leading to multiple resources with same servername.
func (o ServerInstanceOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.ServerName }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ServerInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ServerInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ServerInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerInstanceOutput{})
}
