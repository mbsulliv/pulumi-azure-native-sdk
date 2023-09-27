// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the connected registry.
func LookupConnectedRegistry(ctx *pulumi.Context, args *LookupConnectedRegistryArgs, opts ...pulumi.InvokeOption) (*LookupConnectedRegistryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectedRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20230101preview:getConnectedRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedRegistryArgs struct {
	// The name of the connected registry.
	ConnectedRegistryName string `pulumi:"connectedRegistryName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a connected registry for a container registry.
type LookupConnectedRegistryResult struct {
	// The activation properties of the connected registry.
	Activation ActivationPropertiesResponse `pulumi:"activation"`
	// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
	ClientTokenIds []string `pulumi:"clientTokenIds"`
	// The current connection state of the connected registry.
	ConnectionState string `pulumi:"connectionState"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The last activity time of the connected registry.
	LastActivityTime string `pulumi:"lastActivityTime"`
	// The logging properties of the connected registry.
	Logging *LoggingPropertiesResponse `pulumi:"logging"`
	// The login server properties of the connected registry.
	LoginServer *LoginServerPropertiesResponse `pulumi:"loginServer"`
	// The mode of the connected registry resource that indicates the permissions of the registry.
	Mode string `pulumi:"mode"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The list of notifications subscription information for the connected registry.
	NotificationsList []string `pulumi:"notificationsList"`
	// The parent of the connected registry.
	Parent ParentPropertiesResponse `pulumi:"parent"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of current statuses of the connected registry.
	StatusDetails []StatusDetailPropertiesResponse `pulumi:"statusDetails"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The current version of ACR runtime on the connected registry.
	Version string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupConnectedRegistryResult
func (val *LookupConnectedRegistryResult) Defaults() *LookupConnectedRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Logging = tmp.Logging.Defaults()

	return &tmp
}

func LookupConnectedRegistryOutput(ctx *pulumi.Context, args LookupConnectedRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedRegistryResult, error) {
			args := v.(LookupConnectedRegistryArgs)
			r, err := LookupConnectedRegistry(ctx, &args, opts...)
			var s LookupConnectedRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedRegistryResultOutput)
}

type LookupConnectedRegistryOutputArgs struct {
	// The name of the connected registry.
	ConnectedRegistryName pulumi.StringInput `pulumi:"connectedRegistryName"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedRegistryArgs)(nil)).Elem()
}

// An object that represents a connected registry for a container registry.
type LookupConnectedRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedRegistryResult)(nil)).Elem()
}

func (o LookupConnectedRegistryResultOutput) ToLookupConnectedRegistryResultOutput() LookupConnectedRegistryResultOutput {
	return o
}

func (o LookupConnectedRegistryResultOutput) ToLookupConnectedRegistryResultOutputWithContext(ctx context.Context) LookupConnectedRegistryResultOutput {
	return o
}

func (o LookupConnectedRegistryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectedRegistryResult] {
	return pulumix.Output[LookupConnectedRegistryResult]{
		OutputState: o.OutputState,
	}
}

// The activation properties of the connected registry.
func (o LookupConnectedRegistryResultOutput) Activation() ActivationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) ActivationPropertiesResponse { return v.Activation }).(ActivationPropertiesResponseOutput)
}

// The list of the ACR token resource IDs used to authenticate clients to the connected registry.
func (o LookupConnectedRegistryResultOutput) ClientTokenIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) []string { return v.ClientTokenIds }).(pulumi.StringArrayOutput)
}

// The current connection state of the connected registry.
func (o LookupConnectedRegistryResultOutput) ConnectionState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.ConnectionState }).(pulumi.StringOutput)
}

// The resource ID.
func (o LookupConnectedRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last activity time of the connected registry.
func (o LookupConnectedRegistryResultOutput) LastActivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.LastActivityTime }).(pulumi.StringOutput)
}

// The logging properties of the connected registry.
func (o LookupConnectedRegistryResultOutput) Logging() LoggingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) *LoggingPropertiesResponse { return v.Logging }).(LoggingPropertiesResponsePtrOutput)
}

// The login server properties of the connected registry.
func (o LookupConnectedRegistryResultOutput) LoginServer() LoginServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) *LoginServerPropertiesResponse { return v.LoginServer }).(LoginServerPropertiesResponsePtrOutput)
}

// The mode of the connected registry resource that indicates the permissions of the registry.
func (o LookupConnectedRegistryResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Mode }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupConnectedRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of notifications subscription information for the connected registry.
func (o LookupConnectedRegistryResultOutput) NotificationsList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) []string { return v.NotificationsList }).(pulumi.StringArrayOutput)
}

// The parent of the connected registry.
func (o LookupConnectedRegistryResultOutput) Parent() ParentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) ParentPropertiesResponse { return v.Parent }).(ParentPropertiesResponseOutput)
}

// Provisioning state of the resource.
func (o LookupConnectedRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of current statuses of the connected registry.
func (o LookupConnectedRegistryResultOutput) StatusDetails() StatusDetailPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) []StatusDetailPropertiesResponse { return v.StatusDetails }).(StatusDetailPropertiesResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupConnectedRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupConnectedRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

// The current version of ACR runtime on the connected registry.
func (o LookupConnectedRegistryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedRegistryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedRegistryResultOutput{})
}
