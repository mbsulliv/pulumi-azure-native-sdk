// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Database Advisor.
type DatabaseAdvisor struct {
	pulumi.CustomResourceState

	// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
	AdvisorStatus pulumi.StringOutput `pulumi:"advisorStatus"`
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteValue pulumi.StringOutput `pulumi:"autoExecuteValue"`
	// Resource kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Gets the time when the current resource was analyzed for recommendations by this advisor.
	LastChecked pulumi.StringOutput `pulumi:"lastChecked"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available), LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
	RecommendationsStatus pulumi.StringOutput `pulumi:"recommendationsStatus"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDatabaseAdvisor registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAdvisor(ctx *pulumi.Context,
	name string, args *DatabaseAdvisorArgs, opts ...pulumi.ResourceOption) (*DatabaseAdvisor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoExecuteValue == nil {
		return nil, errors.New("invalid value for required argument 'AutoExecuteValue'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20150501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:DatabaseAdvisor"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:DatabaseAdvisor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DatabaseAdvisor
	err := ctx.RegisterResource("azure-native:sql/v20140401:DatabaseAdvisor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAdvisor gets an existing DatabaseAdvisor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAdvisor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAdvisorState, opts ...pulumi.ResourceOption) (*DatabaseAdvisor, error) {
	var resource DatabaseAdvisor
	err := ctx.ReadResource("azure-native:sql/v20140401:DatabaseAdvisor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAdvisor resources.
type databaseAdvisorState struct {
}

type DatabaseAdvisorState struct {
}

func (DatabaseAdvisorState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAdvisorState)(nil)).Elem()
}

type databaseAdvisorArgs struct {
	// The name of the Database Advisor.
	AdvisorName *string `pulumi:"advisorName"`
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteValue AutoExecuteStatus `pulumi:"autoExecuteValue"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a DatabaseAdvisor resource.
type DatabaseAdvisorArgs struct {
	// The name of the Database Advisor.
	AdvisorName pulumi.StringPtrInput
	// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
	AutoExecuteValue AutoExecuteStatusInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (DatabaseAdvisorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAdvisorArgs)(nil)).Elem()
}

type DatabaseAdvisorInput interface {
	pulumi.Input

	ToDatabaseAdvisorOutput() DatabaseAdvisorOutput
	ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput
}

func (*DatabaseAdvisor) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAdvisor)(nil)).Elem()
}

func (i *DatabaseAdvisor) ToDatabaseAdvisorOutput() DatabaseAdvisorOutput {
	return i.ToDatabaseAdvisorOutputWithContext(context.Background())
}

func (i *DatabaseAdvisor) ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAdvisorOutput)
}

type DatabaseAdvisorOutput struct{ *pulumi.OutputState }

func (DatabaseAdvisorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAdvisor)(nil)).Elem()
}

func (o DatabaseAdvisorOutput) ToDatabaseAdvisorOutput() DatabaseAdvisorOutput {
	return o
}

func (o DatabaseAdvisorOutput) ToDatabaseAdvisorOutputWithContext(ctx context.Context) DatabaseAdvisorOutput {
	return o
}

// Gets the status of availability of this advisor to customers. Possible values are 'GA', 'PublicPreview', 'LimitedPublicPreview' and 'PrivatePreview'.
func (o DatabaseAdvisorOutput) AdvisorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.AdvisorStatus }).(pulumi.StringOutput)
}

// Gets the auto-execute status (whether to let the system execute the recommendations) of this advisor. Possible values are 'Enabled' and 'Disabled'
func (o DatabaseAdvisorOutput) AutoExecuteValue() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.AutoExecuteValue }).(pulumi.StringOutput)
}

// Resource kind.
func (o DatabaseAdvisorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Gets the time when the current resource was analyzed for recommendations by this advisor.
func (o DatabaseAdvisorOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.LastChecked }).(pulumi.StringOutput)
}

// Resource location.
func (o DatabaseAdvisorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o DatabaseAdvisorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets that status of recommendations for this advisor and reason for not having any recommendations. Possible values include, but are not limited to, 'Ok' (Recommendations available), LowActivity (not enough workload to analyze), 'DbSeemsTuned' (Database is doing well), etc.
func (o DatabaseAdvisorOutput) RecommendationsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.RecommendationsStatus }).(pulumi.StringOutput)
}

// Resource type.
func (o DatabaseAdvisorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseAdvisor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseAdvisorOutput{})
}
