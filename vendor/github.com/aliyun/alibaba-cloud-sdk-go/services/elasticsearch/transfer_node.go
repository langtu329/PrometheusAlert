package elasticsearch

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

// TransferNode invokes the elasticsearch.TransferNode API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/transfernode.html
func (client *Client) TransferNode(request *TransferNodeRequest) (response *TransferNodeResponse, err error) {
	response = CreateTransferNodeResponse()
	err = client.DoAction(request, response)
	return
}

// TransferNodeWithChan invokes the elasticsearch.TransferNode API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/transfernode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferNodeWithChan(request *TransferNodeRequest) (<-chan *TransferNodeResponse, <-chan error) {
	responseChan := make(chan *TransferNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferNode(request)
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

// TransferNodeWithCallback invokes the elasticsearch.TransferNode API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/transfernode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TransferNodeWithCallback(request *TransferNodeRequest, callback func(response *TransferNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferNodeResponse
		var err error
		defer close(result)
		response, err = client.TransferNode(request)
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

// TransferNodeRequest is the request struct for api TransferNode
type TransferNodeRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	NodeType    string `position:"Query" name:"nodeType"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// TransferNodeResponse is the response struct for api TransferNode
type TransferNodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateTransferNodeRequest creates a request to invoke TransferNode API
func CreateTransferNodeRequest() (request *TransferNodeRequest) {
	request = &TransferNodeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "TransferNode", "/openapi/instances/[InstanceId]/actions/transfer", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTransferNodeResponse creates a response to parse from TransferNode response
func CreateTransferNodeResponse() (response *TransferNodeResponse) {
	response = &TransferNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}