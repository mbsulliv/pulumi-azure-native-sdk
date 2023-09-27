// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Represents Activity timeline item.
type ActivityTimelineItemResponse struct {
	// The grouping bucket end time.
	BucketEndTimeUTC string `pulumi:"bucketEndTimeUTC"`
	// The grouping bucket start time.
	BucketStartTimeUTC string `pulumi:"bucketStartTimeUTC"`
	// The activity timeline content.
	Content string `pulumi:"content"`
	// The time of the first activity in the grouping bucket.
	FirstActivityTimeUTC string `pulumi:"firstActivityTimeUTC"`
	// The entity query kind
	// Expected value is 'Activity'.
	Kind string `pulumi:"kind"`
	// The time of the last activity in the grouping bucket.
	LastActivityTimeUTC string `pulumi:"lastActivityTimeUTC"`
	// The activity query id.
	QueryId string `pulumi:"queryId"`
	// The activity timeline title.
	Title string `pulumi:"title"`
}

// Represents anomaly timeline item.
type AnomalyTimelineItemResponse struct {
	// The anomaly azure resource id.
	AzureResourceId string `pulumi:"azureResourceId"`
	// The anomaly description.
	Description *string `pulumi:"description"`
	// The anomaly name.
	DisplayName string `pulumi:"displayName"`
	// The anomaly end time.
	EndTimeUtc string `pulumi:"endTimeUtc"`
	// The intent of the anomaly.
	Intent *string `pulumi:"intent"`
	// The entity query kind
	// Expected value is 'Anomaly'.
	Kind string `pulumi:"kind"`
	// The anomaly product name.
	ProductName *string `pulumi:"productName"`
	// The reasons that cause the anomaly.
	Reasons []string `pulumi:"reasons"`
	// The anomaly start time.
	StartTimeUtc string `pulumi:"startTimeUtc"`
	// The techniques of the anomaly.
	Techniques []string `pulumi:"techniques"`
	// The anomaly generated time.
	TimeGenerated string `pulumi:"timeGenerated"`
	// The name of the anomaly vendor.
	Vendor *string `pulumi:"vendor"`
}

// Represents bookmark timeline item.
type BookmarkTimelineItemResponse struct {
	// The bookmark azure resource id.
	AzureResourceId string `pulumi:"azureResourceId"`
	// Describes a user that created the bookmark
	CreatedBy *UserInfoResponse `pulumi:"createdBy"`
	// The bookmark display name.
	DisplayName *string `pulumi:"displayName"`
	// The bookmark end time.
	EndTimeUtc *string `pulumi:"endTimeUtc"`
	// The bookmark event time.
	EventTime *string `pulumi:"eventTime"`
	// The entity query kind
	// Expected value is 'Bookmark'.
	Kind string `pulumi:"kind"`
	// List of labels relevant to this bookmark
	Labels []string `pulumi:"labels"`
	// The notes of the bookmark
	Notes *string `pulumi:"notes"`
	// The bookmark start time.
	StartTimeUtc *string `pulumi:"startTimeUtc"`
}

// Entity insight Item.
type EntityInsightItemResponse struct {
	// Query results for table insights query.
	ChartQueryResults []InsightsTableResultResponse `pulumi:"chartQueryResults"`
	// The query id of the insight
	QueryId *string `pulumi:"queryId"`
	// The Time interval that the query actually executed on.
	QueryTimeInterval *EntityInsightItemResponseQueryTimeInterval `pulumi:"queryTimeInterval"`
	// Query results for table insights query.
	TableQueryResults *InsightsTableResultResponse `pulumi:"tableQueryResults"`
}

// Entity insight Item.
type EntityInsightItemResponseOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponse)(nil)).Elem()
}

func (o EntityInsightItemResponseOutput) ToEntityInsightItemResponseOutput() EntityInsightItemResponseOutput {
	return o
}

func (o EntityInsightItemResponseOutput) ToEntityInsightItemResponseOutputWithContext(ctx context.Context) EntityInsightItemResponseOutput {
	return o
}

func (o EntityInsightItemResponseOutput) ToOutput(ctx context.Context) pulumix.Output[EntityInsightItemResponse] {
	return pulumix.Output[EntityInsightItemResponse]{
		OutputState: o.OutputState,
	}
}

// Query results for table insights query.
func (o EntityInsightItemResponseOutput) ChartQueryResults() InsightsTableResultResponseArrayOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) []InsightsTableResultResponse { return v.ChartQueryResults }).(InsightsTableResultResponseArrayOutput)
}

// The query id of the insight
func (o EntityInsightItemResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

// The Time interval that the query actually executed on.
func (o EntityInsightItemResponseOutput) QueryTimeInterval() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *EntityInsightItemResponseQueryTimeInterval {
		return v.QueryTimeInterval
	}).(EntityInsightItemResponseQueryTimeIntervalPtrOutput)
}

// Query results for table insights query.
func (o EntityInsightItemResponseOutput) TableQueryResults() InsightsTableResultResponsePtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *InsightsTableResultResponse { return v.TableQueryResults }).(InsightsTableResultResponsePtrOutput)
}

type EntityInsightItemResponseArrayOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityInsightItemResponse)(nil)).Elem()
}

func (o EntityInsightItemResponseArrayOutput) ToEntityInsightItemResponseArrayOutput() EntityInsightItemResponseArrayOutput {
	return o
}

func (o EntityInsightItemResponseArrayOutput) ToEntityInsightItemResponseArrayOutputWithContext(ctx context.Context) EntityInsightItemResponseArrayOutput {
	return o
}

func (o EntityInsightItemResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]EntityInsightItemResponse] {
	return pulumix.Output[[]EntityInsightItemResponse]{
		OutputState: o.OutputState,
	}
}

func (o EntityInsightItemResponseArrayOutput) Index(i pulumi.IntInput) EntityInsightItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityInsightItemResponse {
		return vs[0].([]EntityInsightItemResponse)[vs[1].(int)]
	}).(EntityInsightItemResponseOutput)
}

// The Time interval that the query actually executed on.
type EntityInsightItemResponseQueryTimeInterval struct {
	// Insight query end time
	EndTime *string `pulumi:"endTime"`
	// Insight query start time
	StartTime *string `pulumi:"startTime"`
}

// The Time interval that the query actually executed on.
type EntityInsightItemResponseQueryTimeIntervalOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseQueryTimeIntervalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalOutput() EntityInsightItemResponseQueryTimeIntervalOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToOutput(ctx context.Context) pulumix.Output[EntityInsightItemResponseQueryTimeInterval] {
	return pulumix.Output[EntityInsightItemResponseQueryTimeInterval]{
		OutputState: o.OutputState,
	}
}

// Insight query end time
func (o EntityInsightItemResponseQueryTimeIntervalOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponseQueryTimeInterval) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Insight query start time
func (o EntityInsightItemResponseQueryTimeIntervalOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponseQueryTimeInterval) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type EntityInsightItemResponseQueryTimeIntervalPtrOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseQueryTimeIntervalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*EntityInsightItemResponseQueryTimeInterval] {
	return pulumix.Output[*EntityInsightItemResponseQueryTimeInterval]{
		OutputState: o.OutputState,
	}
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) Elem() EntityInsightItemResponseQueryTimeIntervalOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) EntityInsightItemResponseQueryTimeInterval {
		if v != nil {
			return *v
		}
		var ret EntityInsightItemResponseQueryTimeInterval
		return ret
	}).(EntityInsightItemResponseQueryTimeIntervalOutput)
}

// Insight query end time
func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

// Insight query start time
func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

// GetInsights Query Errors.
type GetInsightsErrorKindResponse struct {
	// the error message
	ErrorMessage string `pulumi:"errorMessage"`
	// the query kind
	Kind string `pulumi:"kind"`
	// the query id
	QueryId *string `pulumi:"queryId"`
}

// GetInsights Query Errors.
type GetInsightsErrorKindResponseOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorKindResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsErrorKindResponse)(nil)).Elem()
}

func (o GetInsightsErrorKindResponseOutput) ToGetInsightsErrorKindResponseOutput() GetInsightsErrorKindResponseOutput {
	return o
}

func (o GetInsightsErrorKindResponseOutput) ToGetInsightsErrorKindResponseOutputWithContext(ctx context.Context) GetInsightsErrorKindResponseOutput {
	return o
}

func (o GetInsightsErrorKindResponseOutput) ToOutput(ctx context.Context) pulumix.Output[GetInsightsErrorKindResponse] {
	return pulumix.Output[GetInsightsErrorKindResponse]{
		OutputState: o.OutputState,
	}
}

// the error message
func (o GetInsightsErrorKindResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// the query kind
func (o GetInsightsErrorKindResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) string { return v.Kind }).(pulumi.StringOutput)
}

// the query id
func (o GetInsightsErrorKindResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

type GetInsightsErrorKindResponseArrayOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorKindResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInsightsErrorKindResponse)(nil)).Elem()
}

func (o GetInsightsErrorKindResponseArrayOutput) ToGetInsightsErrorKindResponseArrayOutput() GetInsightsErrorKindResponseArrayOutput {
	return o
}

func (o GetInsightsErrorKindResponseArrayOutput) ToGetInsightsErrorKindResponseArrayOutputWithContext(ctx context.Context) GetInsightsErrorKindResponseArrayOutput {
	return o
}

func (o GetInsightsErrorKindResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]GetInsightsErrorKindResponse] {
	return pulumix.Output[[]GetInsightsErrorKindResponse]{
		OutputState: o.OutputState,
	}
}

func (o GetInsightsErrorKindResponseArrayOutput) Index(i pulumi.IntInput) GetInsightsErrorKindResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInsightsErrorKindResponse {
		return vs[0].([]GetInsightsErrorKindResponse)[vs[1].(int)]
	}).(GetInsightsErrorKindResponseOutput)
}

// Get Insights result metadata.
type GetInsightsResultsMetadataResponse struct {
	// information about the failed queries
	Errors []GetInsightsErrorKindResponse `pulumi:"errors"`
	// the total items found for the insights request
	TotalCount int `pulumi:"totalCount"`
}

// Get Insights result metadata.
type GetInsightsResultsMetadataResponseOutput struct{ *pulumi.OutputState }

func (GetInsightsResultsMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponseOutput() GetInsightsResultsMetadataResponseOutput {
	return o
}

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponseOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponseOutput {
	return o
}

func (o GetInsightsResultsMetadataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[GetInsightsResultsMetadataResponse] {
	return pulumix.Output[GetInsightsResultsMetadataResponse]{
		OutputState: o.OutputState,
	}
}

// information about the failed queries
func (o GetInsightsResultsMetadataResponseOutput) Errors() GetInsightsErrorKindResponseArrayOutput {
	return o.ApplyT(func(v GetInsightsResultsMetadataResponse) []GetInsightsErrorKindResponse { return v.Errors }).(GetInsightsErrorKindResponseArrayOutput)
}

// the total items found for the insights request
func (o GetInsightsResultsMetadataResponseOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInsightsResultsMetadataResponse) int { return v.TotalCount }).(pulumi.IntOutput)
}

type GetInsightsResultsMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (GetInsightsResultsMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (o GetInsightsResultsMetadataResponsePtrOutput) ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput {
	return o
}

func (o GetInsightsResultsMetadataResponsePtrOutput) ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponsePtrOutput {
	return o
}

func (o GetInsightsResultsMetadataResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*GetInsightsResultsMetadataResponse] {
	return pulumix.Output[*GetInsightsResultsMetadataResponse]{
		OutputState: o.OutputState,
	}
}

func (o GetInsightsResultsMetadataResponsePtrOutput) Elem() GetInsightsResultsMetadataResponseOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) GetInsightsResultsMetadataResponse {
		if v != nil {
			return *v
		}
		var ret GetInsightsResultsMetadataResponse
		return ret
	}).(GetInsightsResultsMetadataResponseOutput)
}

// information about the failed queries
func (o GetInsightsResultsMetadataResponsePtrOutput) Errors() GetInsightsErrorKindResponseArrayOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) []GetInsightsErrorKindResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(GetInsightsErrorKindResponseArrayOutput)
}

// the total items found for the insights request
func (o GetInsightsResultsMetadataResponsePtrOutput) TotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalCount
	}).(pulumi.IntPtrOutput)
}

// Query results for table insights query.
type InsightsTableResultResponse struct {
	// Columns Metadata of the table
	Columns []InsightsTableResultResponseColumns `pulumi:"columns"`
	// Rows data of the table
	Rows [][]string `pulumi:"rows"`
}

// Query results for table insights query.
type InsightsTableResultResponseOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponseOutput() InsightsTableResultResponseOutput {
	return o
}

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponseOutputWithContext(ctx context.Context) InsightsTableResultResponseOutput {
	return o
}

func (o InsightsTableResultResponseOutput) ToOutput(ctx context.Context) pulumix.Output[InsightsTableResultResponse] {
	return pulumix.Output[InsightsTableResultResponse]{
		OutputState: o.OutputState,
	}
}

// Columns Metadata of the table
func (o InsightsTableResultResponseOutput) Columns() InsightsTableResultResponseColumnsArrayOutput {
	return o.ApplyT(func(v InsightsTableResultResponse) []InsightsTableResultResponseColumns { return v.Columns }).(InsightsTableResultResponseColumnsArrayOutput)
}

// Rows data of the table
func (o InsightsTableResultResponseOutput) Rows() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v InsightsTableResultResponse) [][]string { return v.Rows }).(pulumi.StringArrayArrayOutput)
}

type InsightsTableResultResponsePtrOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponsePtrOutput) ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput {
	return o
}

func (o InsightsTableResultResponsePtrOutput) ToInsightsTableResultResponsePtrOutputWithContext(ctx context.Context) InsightsTableResultResponsePtrOutput {
	return o
}

func (o InsightsTableResultResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*InsightsTableResultResponse] {
	return pulumix.Output[*InsightsTableResultResponse]{
		OutputState: o.OutputState,
	}
}

func (o InsightsTableResultResponsePtrOutput) Elem() InsightsTableResultResponseOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) InsightsTableResultResponse {
		if v != nil {
			return *v
		}
		var ret InsightsTableResultResponse
		return ret
	}).(InsightsTableResultResponseOutput)
}

// Columns Metadata of the table
func (o InsightsTableResultResponsePtrOutput) Columns() InsightsTableResultResponseColumnsArrayOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) []InsightsTableResultResponseColumns {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(InsightsTableResultResponseColumnsArrayOutput)
}

// Rows data of the table
func (o InsightsTableResultResponsePtrOutput) Rows() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) [][]string {
		if v == nil {
			return nil
		}
		return v.Rows
	}).(pulumi.StringArrayArrayOutput)
}

type InsightsTableResultResponseArrayOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponseArrayOutput) ToInsightsTableResultResponseArrayOutput() InsightsTableResultResponseArrayOutput {
	return o
}

func (o InsightsTableResultResponseArrayOutput) ToInsightsTableResultResponseArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseArrayOutput {
	return o
}

func (o InsightsTableResultResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]InsightsTableResultResponse] {
	return pulumix.Output[[]InsightsTableResultResponse]{
		OutputState: o.OutputState,
	}
}

func (o InsightsTableResultResponseArrayOutput) Index(i pulumi.IntInput) InsightsTableResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InsightsTableResultResponse {
		return vs[0].([]InsightsTableResultResponse)[vs[1].(int)]
	}).(InsightsTableResultResponseOutput)
}

type InsightsTableResultResponseColumns struct {
	// the name of the colum
	Name *string `pulumi:"name"`
	// the type of the colum
	Type *string `pulumi:"type"`
}

type InsightsTableResultResponseColumnsOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseColumnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponseColumns)(nil)).Elem()
}

func (o InsightsTableResultResponseColumnsOutput) ToInsightsTableResultResponseColumnsOutput() InsightsTableResultResponseColumnsOutput {
	return o
}

func (o InsightsTableResultResponseColumnsOutput) ToInsightsTableResultResponseColumnsOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsOutput {
	return o
}

func (o InsightsTableResultResponseColumnsOutput) ToOutput(ctx context.Context) pulumix.Output[InsightsTableResultResponseColumns] {
	return pulumix.Output[InsightsTableResultResponseColumns]{
		OutputState: o.OutputState,
	}
}

// the name of the colum
func (o InsightsTableResultResponseColumnsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InsightsTableResultResponseColumns) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// the type of the colum
func (o InsightsTableResultResponseColumnsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InsightsTableResultResponseColumns) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type InsightsTableResultResponseColumnsArrayOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseColumnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponseColumns)(nil)).Elem()
}

func (o InsightsTableResultResponseColumnsArrayOutput) ToInsightsTableResultResponseColumnsArrayOutput() InsightsTableResultResponseColumnsArrayOutput {
	return o
}

func (o InsightsTableResultResponseColumnsArrayOutput) ToInsightsTableResultResponseColumnsArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsArrayOutput {
	return o
}

func (o InsightsTableResultResponseColumnsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]InsightsTableResultResponseColumns] {
	return pulumix.Output[[]InsightsTableResultResponseColumns]{
		OutputState: o.OutputState,
	}
}

func (o InsightsTableResultResponseColumnsArrayOutput) Index(i pulumi.IntInput) InsightsTableResultResponseColumnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InsightsTableResultResponseColumns {
		return vs[0].([]InsightsTableResultResponseColumns)[vs[1].(int)]
	}).(InsightsTableResultResponseColumnsOutput)
}

// Represents a repository.
type RepoResponse struct {
	// Array of branches.
	Branches []string `pulumi:"branches"`
	// The name of the repository.
	FullName *string `pulumi:"fullName"`
	// The url to access the repository.
	Url *string `pulumi:"url"`
}

// Represents a repository.
type RepoResponseOutput struct{ *pulumi.OutputState }

func (RepoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoResponse)(nil)).Elem()
}

func (o RepoResponseOutput) ToRepoResponseOutput() RepoResponseOutput {
	return o
}

func (o RepoResponseOutput) ToRepoResponseOutputWithContext(ctx context.Context) RepoResponseOutput {
	return o
}

func (o RepoResponseOutput) ToOutput(ctx context.Context) pulumix.Output[RepoResponse] {
	return pulumix.Output[RepoResponse]{
		OutputState: o.OutputState,
	}
}

// Array of branches.
func (o RepoResponseOutput) Branches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepoResponse) []string { return v.Branches }).(pulumi.StringArrayOutput)
}

// The name of the repository.
func (o RepoResponseOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepoResponse) *string { return v.FullName }).(pulumi.StringPtrOutput)
}

// The url to access the repository.
func (o RepoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type RepoResponseArrayOutput struct{ *pulumi.OutputState }

func (RepoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepoResponse)(nil)).Elem()
}

func (o RepoResponseArrayOutput) ToRepoResponseArrayOutput() RepoResponseArrayOutput {
	return o
}

func (o RepoResponseArrayOutput) ToRepoResponseArrayOutputWithContext(ctx context.Context) RepoResponseArrayOutput {
	return o
}

func (o RepoResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]RepoResponse] {
	return pulumix.Output[[]RepoResponse]{
		OutputState: o.OutputState,
	}
}

func (o RepoResponseArrayOutput) Index(i pulumi.IntInput) RepoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepoResponse {
		return vs[0].([]RepoResponse)[vs[1].(int)]
	}).(RepoResponseOutput)
}

// Represents security alert timeline item.
type SecurityAlertTimelineItemResponse struct {
	// The name of the alert type.
	AlertType string `pulumi:"alertType"`
	// The alert azure resource id.
	AzureResourceId string `pulumi:"azureResourceId"`
	// The alert description.
	Description *string `pulumi:"description"`
	// The alert name.
	DisplayName string `pulumi:"displayName"`
	// The alert end time.
	EndTimeUtc string `pulumi:"endTimeUtc"`
	// The entity query kind
	// Expected value is 'SecurityAlert'.
	Kind string `pulumi:"kind"`
	// The alert product name.
	ProductName *string `pulumi:"productName"`
	// The alert severity.
	Severity string `pulumi:"severity"`
	// The alert start time.
	StartTimeUtc string `pulumi:"startTimeUtc"`
	// The alert generated time.
	TimeGenerated string `pulumi:"timeGenerated"`
}

// timeline aggregation information per kind
type TimelineAggregationResponse struct {
	// the total items found for a kind
	Count int `pulumi:"count"`
	// the query kind
	Kind string `pulumi:"kind"`
}

// timeline aggregation information per kind
type TimelineAggregationResponseOutput struct{ *pulumi.OutputState }

func (TimelineAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineAggregationResponse)(nil)).Elem()
}

func (o TimelineAggregationResponseOutput) ToTimelineAggregationResponseOutput() TimelineAggregationResponseOutput {
	return o
}

func (o TimelineAggregationResponseOutput) ToTimelineAggregationResponseOutputWithContext(ctx context.Context) TimelineAggregationResponseOutput {
	return o
}

func (o TimelineAggregationResponseOutput) ToOutput(ctx context.Context) pulumix.Output[TimelineAggregationResponse] {
	return pulumix.Output[TimelineAggregationResponse]{
		OutputState: o.OutputState,
	}
}

// the total items found for a kind
func (o TimelineAggregationResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v TimelineAggregationResponse) int { return v.Count }).(pulumi.IntOutput)
}

// the query kind
func (o TimelineAggregationResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineAggregationResponse) string { return v.Kind }).(pulumi.StringOutput)
}

type TimelineAggregationResponseArrayOutput struct{ *pulumi.OutputState }

func (TimelineAggregationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineAggregationResponse)(nil)).Elem()
}

func (o TimelineAggregationResponseArrayOutput) ToTimelineAggregationResponseArrayOutput() TimelineAggregationResponseArrayOutput {
	return o
}

func (o TimelineAggregationResponseArrayOutput) ToTimelineAggregationResponseArrayOutputWithContext(ctx context.Context) TimelineAggregationResponseArrayOutput {
	return o
}

func (o TimelineAggregationResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]TimelineAggregationResponse] {
	return pulumix.Output[[]TimelineAggregationResponse]{
		OutputState: o.OutputState,
	}
}

func (o TimelineAggregationResponseArrayOutput) Index(i pulumi.IntInput) TimelineAggregationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimelineAggregationResponse {
		return vs[0].([]TimelineAggregationResponse)[vs[1].(int)]
	}).(TimelineAggregationResponseOutput)
}

// Timeline Query Errors.
type TimelineErrorResponse struct {
	// the error message
	ErrorMessage string `pulumi:"errorMessage"`
	// the query kind
	Kind string `pulumi:"kind"`
	// the query id
	QueryId *string `pulumi:"queryId"`
}

// Timeline Query Errors.
type TimelineErrorResponseOutput struct{ *pulumi.OutputState }

func (TimelineErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineErrorResponse)(nil)).Elem()
}

func (o TimelineErrorResponseOutput) ToTimelineErrorResponseOutput() TimelineErrorResponseOutput {
	return o
}

func (o TimelineErrorResponseOutput) ToTimelineErrorResponseOutputWithContext(ctx context.Context) TimelineErrorResponseOutput {
	return o
}

func (o TimelineErrorResponseOutput) ToOutput(ctx context.Context) pulumix.Output[TimelineErrorResponse] {
	return pulumix.Output[TimelineErrorResponse]{
		OutputState: o.OutputState,
	}
}

// the error message
func (o TimelineErrorResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineErrorResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// the query kind
func (o TimelineErrorResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineErrorResponse) string { return v.Kind }).(pulumi.StringOutput)
}

// the query id
func (o TimelineErrorResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimelineErrorResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

type TimelineErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (TimelineErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineErrorResponse)(nil)).Elem()
}

func (o TimelineErrorResponseArrayOutput) ToTimelineErrorResponseArrayOutput() TimelineErrorResponseArrayOutput {
	return o
}

func (o TimelineErrorResponseArrayOutput) ToTimelineErrorResponseArrayOutputWithContext(ctx context.Context) TimelineErrorResponseArrayOutput {
	return o
}

func (o TimelineErrorResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]TimelineErrorResponse] {
	return pulumix.Output[[]TimelineErrorResponse]{
		OutputState: o.OutputState,
	}
}

func (o TimelineErrorResponseArrayOutput) Index(i pulumi.IntInput) TimelineErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimelineErrorResponse {
		return vs[0].([]TimelineErrorResponse)[vs[1].(int)]
	}).(TimelineErrorResponseOutput)
}

// Expansion result metadata.
type TimelineResultsMetadataResponse struct {
	// timeline aggregation per kind
	Aggregations []TimelineAggregationResponse `pulumi:"aggregations"`
	// information about the failure queries
	Errors []TimelineErrorResponse `pulumi:"errors"`
	// the total items found for the timeline request
	TotalCount int `pulumi:"totalCount"`
}

// Expansion result metadata.
type TimelineResultsMetadataResponseOutput struct{ *pulumi.OutputState }

func (TimelineResultsMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineResultsMetadataResponse)(nil)).Elem()
}

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponseOutput() TimelineResultsMetadataResponseOutput {
	return o
}

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponseOutputWithContext(ctx context.Context) TimelineResultsMetadataResponseOutput {
	return o
}

func (o TimelineResultsMetadataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[TimelineResultsMetadataResponse] {
	return pulumix.Output[TimelineResultsMetadataResponse]{
		OutputState: o.OutputState,
	}
}

// timeline aggregation per kind
func (o TimelineResultsMetadataResponseOutput) Aggregations() TimelineAggregationResponseArrayOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) []TimelineAggregationResponse { return v.Aggregations }).(TimelineAggregationResponseArrayOutput)
}

// information about the failure queries
func (o TimelineResultsMetadataResponseOutput) Errors() TimelineErrorResponseArrayOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) []TimelineErrorResponse { return v.Errors }).(TimelineErrorResponseArrayOutput)
}

// the total items found for the timeline request
func (o TimelineResultsMetadataResponseOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) int { return v.TotalCount }).(pulumi.IntOutput)
}

type TimelineResultsMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (TimelineResultsMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimelineResultsMetadataResponse)(nil)).Elem()
}

func (o TimelineResultsMetadataResponsePtrOutput) ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput {
	return o
}

func (o TimelineResultsMetadataResponsePtrOutput) ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx context.Context) TimelineResultsMetadataResponsePtrOutput {
	return o
}

func (o TimelineResultsMetadataResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*TimelineResultsMetadataResponse] {
	return pulumix.Output[*TimelineResultsMetadataResponse]{
		OutputState: o.OutputState,
	}
}

func (o TimelineResultsMetadataResponsePtrOutput) Elem() TimelineResultsMetadataResponseOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) TimelineResultsMetadataResponse {
		if v != nil {
			return *v
		}
		var ret TimelineResultsMetadataResponse
		return ret
	}).(TimelineResultsMetadataResponseOutput)
}

// timeline aggregation per kind
func (o TimelineResultsMetadataResponsePtrOutput) Aggregations() TimelineAggregationResponseArrayOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) []TimelineAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregations
	}).(TimelineAggregationResponseArrayOutput)
}

// information about the failure queries
func (o TimelineResultsMetadataResponsePtrOutput) Errors() TimelineErrorResponseArrayOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) []TimelineErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(TimelineErrorResponseArrayOutput)
}

// the total items found for the timeline request
func (o TimelineResultsMetadataResponsePtrOutput) TotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalCount
	}).(pulumi.IntPtrOutput)
}

// User information that made some action
type UserInfoResponse struct {
	// The email of the user.
	Email string `pulumi:"email"`
	// The name of the user.
	Name string `pulumi:"name"`
	// The object id of the user.
	ObjectId *string `pulumi:"objectId"`
}

func init() {
	pulumi.RegisterOutputType(EntityInsightItemResponseOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalPtrOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorKindResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorKindResponseArrayOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponsePtrOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseArrayOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseColumnsOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseColumnsArrayOutput{})
	pulumi.RegisterOutputType(RepoResponseOutput{})
	pulumi.RegisterOutputType(RepoResponseArrayOutput{})
	pulumi.RegisterOutputType(TimelineAggregationResponseOutput{})
	pulumi.RegisterOutputType(TimelineAggregationResponseArrayOutput{})
	pulumi.RegisterOutputType(TimelineErrorResponseOutput{})
	pulumi.RegisterOutputType(TimelineErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(TimelineResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(TimelineResultsMetadataResponsePtrOutput{})
}
