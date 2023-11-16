// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Database, Server or Elastic Pool Advisor.
type ServerAdvisor struct {
	pulumi.CustomResourceState

	// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
	AdvisorStatus pulumi.StringOutput `pulumi:"advisorStatus"`
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteStatus pulumi.StringOutput `pulumi:"autoExecuteStatus"`
	// Gets the resource from which current value of auto-execute status is inherited. Auto-execute status can be set on (and inherited from) different levels in the resource hierarchy. Possible values are 'Subscription', 'Server', 'ElasticPool', 'Database' and 'Default' (when status is not explicitly set on any level).
	AutoExecuteStatusInheritedFrom pulumi.StringOutput `pulumi:"autoExecuteStatusInheritedFrom"`
	// Resource kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Gets the time when the current resource was analyzed for recommendations by this advisor.
	LastChecked pulumi.StringOutput `pulumi:"lastChecked"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available),LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
	RecommendationsStatus pulumi.StringOutput `pulumi:"recommendationsStatus"`
	// Gets the recommended actions for this advisor.
	RecommendedActions RecommendedActionResponseArrayOutput `pulumi:"recommendedActions"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerAdvisor registers a new resource with the given unique name, arguments, and options.
func NewServerAdvisor(ctx *pulumi.Context,
	name string, args *ServerAdvisorArgs, opts ...pulumi.ResourceOption) (*ServerAdvisor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoExecuteStatus == nil {
		return nil, errors.New("invalid value for required argument 'AutoExecuteStatus'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20140401:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ServerAdvisor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerAdvisor
	err := ctx.RegisterResource("azure-native:sql/v20230501preview:ServerAdvisor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAdvisor gets an existing ServerAdvisor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAdvisor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAdvisorState, opts ...pulumi.ResourceOption) (*ServerAdvisor, error) {
	var resource ServerAdvisor
	err := ctx.ReadResource("azure-native:sql/v20230501preview:ServerAdvisor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAdvisor resources.
type serverAdvisorState struct {
}

type ServerAdvisorState struct {
}

func (ServerAdvisorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdvisorState)(nil)).Elem()
}

type serverAdvisorArgs struct {
	// The name of the Server Advisor.
	AdvisorName *string `pulumi:"advisorName"`
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteStatus AutoExecuteStatus `pulumi:"autoExecuteStatus"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a ServerAdvisor resource.
type ServerAdvisorArgs struct {
	// The name of the Server Advisor.
	AdvisorName pulumi.StringPtrInput
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteStatus AutoExecuteStatusInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (ServerAdvisorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAdvisorArgs)(nil)).Elem()
}

type ServerAdvisorInput interface {
	pulumi.Input

	ToServerAdvisorOutput() ServerAdvisorOutput
	ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput
}

func (*ServerAdvisor) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdvisor)(nil)).Elem()
}

func (i *ServerAdvisor) ToServerAdvisorOutput() ServerAdvisorOutput {
	return i.ToServerAdvisorOutputWithContext(context.Background())
}

func (i *ServerAdvisor) ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdvisorOutput)
}

type ServerAdvisorOutput struct{ *pulumi.OutputState }

func (ServerAdvisorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdvisor)(nil)).Elem()
}

func (o ServerAdvisorOutput) ToServerAdvisorOutput() ServerAdvisorOutput {
	return o
}

func (o ServerAdvisorOutput) ToServerAdvisorOutputWithContext(ctx context.Context) ServerAdvisorOutput {
	return o
}

// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
func (o ServerAdvisorOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.AdvisorStatus }).(pulumi.StringOutput)
}

// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
func (o ServerAdvisorOutput) AutoExecuteStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.AutoExecuteStatus }).(pulumi.StringOutput)
}

// Gets the resource from which current value of auto-execute status is inherited. Auto-execute status can be set on (and inherited from) different levels in the resource hierarchy. Possible values are 'Subscription', 'Server', 'ElasticPool', 'Database' and 'Default' (when status is not explicitly set on any level).
func (o ServerAdvisorOutput) AutoExecuteStatusInheritedFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.AutoExecuteStatusInheritedFrom }).(pulumi.StringOutput)
}

// Resource kind.
func (o ServerAdvisorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Gets the time when the current resource was analyzed for recommendations by this advisor.
func (o ServerAdvisorOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.LastChecked }).(pulumi.StringOutput)
}

// Resource location.
func (o ServerAdvisorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o ServerAdvisorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available),LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
func (o ServerAdvisorOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

// Gets the recommended actions for this advisor.
func (o ServerAdvisorOutput) RecommendedActions() RecommendedActionResponseArrayOutput {
	return o.ApplyT(func(v *ServerAdvisor) RecommendedActionResponseArrayOutput { return v.RecommendedActions }).(RecommendedActionResponseArrayOutput)
}

// Resource type.
func (o ServerAdvisorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerAdvisor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerAdvisorOutput{})
}
