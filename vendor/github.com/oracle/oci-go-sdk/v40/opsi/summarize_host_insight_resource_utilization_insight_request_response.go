// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v40/common"
	"net/http"
)

// SummarizeHostInsightResourceUtilizationInsightRequest wrapper for the SummarizeHostInsightResourceUtilizationInsight operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUtilizationInsight.go.html to see an example of how to use SummarizeHostInsightResourceUtilizationInsightRequest.
type SummarizeHostInsightResourceUtilizationInsightRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by host resource metric.
	// Supported values are CPU, MEMORY, and LOGICAL_MEMORY.
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Filter by one or more platform types.
	// Possible value is LINUX.
	PlatformType []SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceUtilizationInsightRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceUtilizationInsightRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceUtilizationInsightRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceUtilizationInsightRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeHostInsightResourceUtilizationInsightResponse wrapper for the SummarizeHostInsightResourceUtilizationInsight operation
type SummarizeHostInsightResourceUtilizationInsightResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceUtilizationInsightAggregation instances
	SummarizeHostInsightResourceUtilizationInsightAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SummarizeHostInsightResourceUtilizationInsightResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceUtilizationInsightResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum
const (
	SummarizeHostInsightResourceUtilizationInsightPlatformTypeLinux SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum = "LINUX"
)

var mappingSummarizeHostInsightResourceUtilizationInsightPlatformType = map[string]SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum{
	"LINUX": SummarizeHostInsightResourceUtilizationInsightPlatformTypeLinux,
}

// GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum
func GetSummarizeHostInsightResourceUtilizationInsightPlatformTypeEnumValues() []SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceUtilizationInsightPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUtilizationInsightPlatformType {
		values = append(values, v)
	}
	return values
}
