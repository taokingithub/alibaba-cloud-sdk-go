package aegis

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeStatistics invokes the aegis.DescribeStatistics API synchronously
// api document: https://help.aliyun.com/api/aegis/describestatistics.html
func (client *Client) DescribeStatistics(request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
	response = CreateDescribeStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStatisticsWithChan invokes the aegis.DescribeStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStatisticsWithChan(request *DescribeStatisticsRequest) (<-chan *DescribeStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStatistics(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeStatisticsWithCallback invokes the aegis.DescribeStatistics API asynchronously
// api document: https://help.aliyun.com/api/aegis/describestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStatisticsWithCallback(request *DescribeStatisticsRequest, callback func(response *DescribeStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeStatistics(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeStatisticsRequest is the request struct for api DescribeStatistics
type DescribeStatisticsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeStatisticsResponse is the response struct for api DescribeStatistics
type DescribeStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Statistics Statistics `json:"Statistics" xml:"Statistics"`
}

// CreateDescribeStatisticsRequest creates a request to invoke DescribeStatistics API
func CreateDescribeStatisticsRequest() (request *DescribeStatisticsRequest) {
	request = &DescribeStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeStatistics", "vipaegis", "openAPI")
	return
}

// CreateDescribeStatisticsResponse creates a response to parse from DescribeStatistics response
func CreateDescribeStatisticsResponse() (response *DescribeStatisticsResponse) {
	response = &DescribeStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}