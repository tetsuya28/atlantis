// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: IGithubClient)

package mocks

import (
	github "github.com/google/go-github/v31/github"
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockIGithubClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockIGithubClient(options ...pegomock.Option) *MockIGithubClient {
	mock := &MockIGithubClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockIGithubClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockIGithubClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockIGithubClient) CreateComment(_param0 models.Repo, _param1 int, _param2 string, _param3 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateComment", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockIGithubClient) DownloadRepoConfigFile(_param0 models.PullRequest) (bool, []byte, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DownloadRepoConfigFile", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 []byte
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].([]byte)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockIGithubClient) GetContents(_param0 string, _param1 string, _param2 string, _param3 string) ([]byte, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetContents", params, []reflect.Type{reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []byte
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]byte)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) GetModifiedFiles(_param0 models.Repo, _param1 models.PullRequest) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetModifiedFiles", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) GetPullRequest(_param0 models.Repo, _param1 int) (*github.PullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequest", params, []reflect.Type{reflect.TypeOf((**github.PullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *github.PullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*github.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) GetPullRequestFromName(_param0 string, _param1 string, _param2 int) (*github.PullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequestFromName", params, []reflect.Type{reflect.TypeOf((**github.PullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *github.PullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*github.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) GetRepoStatuses(_param0 models.Repo, _param1 models.PullRequest) ([]*github.RepoStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetRepoStatuses", params, []reflect.Type{reflect.TypeOf((*[]*github.RepoStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []*github.RepoStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]*github.RepoStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) HidePrevCommandComments(_param0 models.Repo, _param1 int, _param2 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HidePrevCommandComments", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockIGithubClient) MarkdownPullLink(_param0 models.PullRequest) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MarkdownPullLink", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) MergePull(_param0 models.PullRequest, _param1 models.PullRequestOptions) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MergePull", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockIGithubClient) PullIsApproved(_param0 models.Repo, _param1 models.PullRequest) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) PullIsMergeable(_param0 models.Repo, _param1 models.PullRequest) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsMergeable", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) PullIsSQLocked(_param0 models.Repo, _param1 models.PullRequest, _param2 []*github.RepoStatus) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsSQLocked", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) PullIsSQMergeable(_param0 models.Repo, _param1 models.PullRequest, _param2 []*github.RepoStatus) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsSQMergeable", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockIGithubClient) SupportsSingleFileDownload(_param0 models.Repo) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SupportsSingleFileDownload", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockIGithubClient) UpdateStatus(_param0 models.Repo, _param1 models.PullRequest, _param2 models.CommitStatus, _param3 string, _param4 string, _param5 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockIGithubClient().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockIGithubClient) VerifyWasCalledOnce() *VerifierMockIGithubClient {
	return &VerifierMockIGithubClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockIGithubClient) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockIGithubClient {
	return &VerifierMockIGithubClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockIGithubClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockIGithubClient {
	return &VerifierMockIGithubClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockIGithubClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockIGithubClient {
	return &VerifierMockIGithubClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockIGithubClient struct {
	mock                   *MockIGithubClient
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockIGithubClient) CreateComment(_param0 models.Repo, _param1 int, _param2 string, _param3 string) *MockIGithubClient_CreateComment_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateComment", params, verifier.timeout)
	return &MockIGithubClient_CreateComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_CreateComment_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_CreateComment_OngoingVerification) GetCapturedArguments() (models.Repo, int, string, string) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockIGithubClient_CreateComment_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) DownloadRepoConfigFile(_param0 models.PullRequest) *MockIGithubClient_DownloadRepoConfigFile_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DownloadRepoConfigFile", params, verifier.timeout)
	return &MockIGithubClient_DownloadRepoConfigFile_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_DownloadRepoConfigFile_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_DownloadRepoConfigFile_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockIGithubClient_DownloadRepoConfigFile_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) GetContents(_param0 string, _param1 string, _param2 string, _param3 string) *MockIGithubClient_GetContents_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetContents", params, verifier.timeout)
	return &MockIGithubClient_GetContents_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_GetContents_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_GetContents_OngoingVerification) GetCapturedArguments() (string, string, string, string) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockIGithubClient_GetContents_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) GetModifiedFiles(_param0 models.Repo, _param1 models.PullRequest) *MockIGithubClient_GetModifiedFiles_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetModifiedFiles", params, verifier.timeout)
	return &MockIGithubClient_GetModifiedFiles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_GetModifiedFiles_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_GetModifiedFiles_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_GetModifiedFiles_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) GetPullRequest(_param0 models.Repo, _param1 int) *MockIGithubClient_GetPullRequest_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequest", params, verifier.timeout)
	return &MockIGithubClient_GetPullRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_GetPullRequest_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_GetPullRequest_OngoingVerification) GetCapturedArguments() (models.Repo, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_GetPullRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) GetPullRequestFromName(_param0 string, _param1 string, _param2 int) *MockIGithubClient_GetPullRequestFromName_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequestFromName", params, verifier.timeout)
	return &MockIGithubClient_GetPullRequestFromName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_GetPullRequestFromName_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_GetPullRequestFromName_OngoingVerification) GetCapturedArguments() (string, string, int) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockIGithubClient_GetPullRequestFromName_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]int, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) GetRepoStatuses(_param0 models.Repo, _param1 models.PullRequest) *MockIGithubClient_GetRepoStatuses_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetRepoStatuses", params, verifier.timeout)
	return &MockIGithubClient_GetRepoStatuses_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_GetRepoStatuses_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_GetRepoStatuses_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_GetRepoStatuses_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) HidePrevCommandComments(_param0 models.Repo, _param1 int, _param2 string) *MockIGithubClient_HidePrevCommandComments_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HidePrevCommandComments", params, verifier.timeout)
	return &MockIGithubClient_HidePrevCommandComments_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_HidePrevCommandComments_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_HidePrevCommandComments_OngoingVerification) GetCapturedArguments() (models.Repo, int, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockIGithubClient_HidePrevCommandComments_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) MarkdownPullLink(_param0 models.PullRequest) *MockIGithubClient_MarkdownPullLink_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MarkdownPullLink", params, verifier.timeout)
	return &MockIGithubClient_MarkdownPullLink_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_MarkdownPullLink_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_MarkdownPullLink_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockIGithubClient_MarkdownPullLink_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) MergePull(_param0 models.PullRequest, _param1 models.PullRequestOptions) *MockIGithubClient_MergePull_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MergePull", params, verifier.timeout)
	return &MockIGithubClient_MergePull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_MergePull_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_MergePull_OngoingVerification) GetCapturedArguments() (models.PullRequest, models.PullRequestOptions) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_MergePull_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []models.PullRequestOptions) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]models.PullRequestOptions, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequestOptions)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) PullIsApproved(_param0 models.Repo, _param1 models.PullRequest) *MockIGithubClient_PullIsApproved_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", params, verifier.timeout)
	return &MockIGithubClient_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_PullIsApproved_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_PullIsApproved_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) PullIsMergeable(_param0 models.Repo, _param1 models.PullRequest) *MockIGithubClient_PullIsMergeable_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsMergeable", params, verifier.timeout)
	return &MockIGithubClient_PullIsMergeable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_PullIsMergeable_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_PullIsMergeable_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockIGithubClient_PullIsMergeable_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) PullIsSQLocked(_param0 models.Repo, _param1 models.PullRequest, _param2 []*github.RepoStatus) *MockIGithubClient_PullIsSQLocked_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsSQLocked", params, verifier.timeout)
	return &MockIGithubClient_PullIsSQLocked_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_PullIsSQLocked_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_PullIsSQLocked_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, []*github.RepoStatus) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockIGithubClient_PullIsSQLocked_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 [][]*github.RepoStatus) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([][]*github.RepoStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.([]*github.RepoStatus)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) PullIsSQMergeable(_param0 models.Repo, _param1 models.PullRequest, _param2 []*github.RepoStatus) *MockIGithubClient_PullIsSQMergeable_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsSQMergeable", params, verifier.timeout)
	return &MockIGithubClient_PullIsSQMergeable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_PullIsSQMergeable_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_PullIsSQMergeable_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, []*github.RepoStatus) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockIGithubClient_PullIsSQMergeable_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 [][]*github.RepoStatus) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([][]*github.RepoStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.([]*github.RepoStatus)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) SupportsSingleFileDownload(_param0 models.Repo) *MockIGithubClient_SupportsSingleFileDownload_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SupportsSingleFileDownload", params, verifier.timeout)
	return &MockIGithubClient_SupportsSingleFileDownload_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_SupportsSingleFileDownload_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_SupportsSingleFileDownload_OngoingVerification) GetCapturedArguments() models.Repo {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockIGithubClient_SupportsSingleFileDownload_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
	}
	return
}

func (verifier *VerifierMockIGithubClient) UpdateStatus(_param0 models.Repo, _param1 models.PullRequest, _param2 models.CommitStatus, _param3 string, _param4 string, _param5 string) *MockIGithubClient_UpdateStatus_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateStatus", params, verifier.timeout)
	return &MockIGithubClient_UpdateStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockIGithubClient_UpdateStatus_OngoingVerification struct {
	mock              *MockIGithubClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockIGithubClient_UpdateStatus_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, string, string, string) {
	_param0, _param1, _param2, _param3, _param4, _param5 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1], _param5[len(_param5)-1]
}

func (c *MockIGithubClient_UpdateStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []string, _param4 []string, _param5 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
		_param5 = make([]string, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(string)
		}
	}
	return
}