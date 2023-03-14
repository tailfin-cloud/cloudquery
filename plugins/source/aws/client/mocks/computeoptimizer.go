// Code generated by MockGen. DO NOT EDIT.
// Source: computeoptimizer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	computeoptimizer "github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	gomock "github.com/golang/mock/gomock"
)

// MockComputeoptimizerClient is a mock of ComputeoptimizerClient interface.
type MockComputeoptimizerClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeoptimizerClientMockRecorder
}

// MockComputeoptimizerClientMockRecorder is the mock recorder for MockComputeoptimizerClient.
type MockComputeoptimizerClientMockRecorder struct {
	mock *MockComputeoptimizerClient
}

// NewMockComputeoptimizerClient creates a new mock instance.
func NewMockComputeoptimizerClient(ctrl *gomock.Controller) *MockComputeoptimizerClient {
	mock := &MockComputeoptimizerClient{ctrl: ctrl}
	mock.recorder = &MockComputeoptimizerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeoptimizerClient) EXPECT() *MockComputeoptimizerClientMockRecorder {
	return m.recorder
}

// DescribeRecommendationExportJobs mocks base method.
func (m *MockComputeoptimizerClient) DescribeRecommendationExportJobs(arg0 context.Context, arg1 *computeoptimizer.DescribeRecommendationExportJobsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRecommendationExportJobs", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.DescribeRecommendationExportJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRecommendationExportJobs indicates an expected call of DescribeRecommendationExportJobs.
func (mr *MockComputeoptimizerClientMockRecorder) DescribeRecommendationExportJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRecommendationExportJobs", reflect.TypeOf((*MockComputeoptimizerClient)(nil).DescribeRecommendationExportJobs), varargs...)
}

// GetAutoScalingGroupRecommendations mocks base method.
func (m *MockComputeoptimizerClient) GetAutoScalingGroupRecommendations(arg0 context.Context, arg1 *computeoptimizer.GetAutoScalingGroupRecommendationsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAutoScalingGroupRecommendations", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutoScalingGroupRecommendations indicates an expected call of GetAutoScalingGroupRecommendations.
func (mr *MockComputeoptimizerClientMockRecorder) GetAutoScalingGroupRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutoScalingGroupRecommendations", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetAutoScalingGroupRecommendations), varargs...)
}

// GetEBSVolumeRecommendations mocks base method.
func (m *MockComputeoptimizerClient) GetEBSVolumeRecommendations(arg0 context.Context, arg1 *computeoptimizer.GetEBSVolumeRecommendationsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEBSVolumeRecommendations", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEBSVolumeRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEBSVolumeRecommendations indicates an expected call of GetEBSVolumeRecommendations.
func (mr *MockComputeoptimizerClientMockRecorder) GetEBSVolumeRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEBSVolumeRecommendations", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEBSVolumeRecommendations), varargs...)
}

// GetEC2InstanceRecommendations mocks base method.
func (m *MockComputeoptimizerClient) GetEC2InstanceRecommendations(arg0 context.Context, arg1 *computeoptimizer.GetEC2InstanceRecommendationsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEC2InstanceRecommendations", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2InstanceRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2InstanceRecommendations indicates an expected call of GetEC2InstanceRecommendations.
func (mr *MockComputeoptimizerClientMockRecorder) GetEC2InstanceRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2InstanceRecommendations", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEC2InstanceRecommendations), varargs...)
}

// GetEC2RecommendationProjectedMetrics mocks base method.
func (m *MockComputeoptimizerClient) GetEC2RecommendationProjectedMetrics(arg0 context.Context, arg1 *computeoptimizer.GetEC2RecommendationProjectedMetricsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEC2RecommendationProjectedMetrics", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEC2RecommendationProjectedMetrics indicates an expected call of GetEC2RecommendationProjectedMetrics.
func (mr *MockComputeoptimizerClientMockRecorder) GetEC2RecommendationProjectedMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEC2RecommendationProjectedMetrics", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEC2RecommendationProjectedMetrics), varargs...)
}

// GetECSServiceRecommendationProjectedMetrics mocks base method.
func (m *MockComputeoptimizerClient) GetECSServiceRecommendationProjectedMetrics(arg0 context.Context, arg1 *computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetECSServiceRecommendationProjectedMetrics", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetECSServiceRecommendationProjectedMetrics indicates an expected call of GetECSServiceRecommendationProjectedMetrics.
func (mr *MockComputeoptimizerClientMockRecorder) GetECSServiceRecommendationProjectedMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetECSServiceRecommendationProjectedMetrics", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetECSServiceRecommendationProjectedMetrics), varargs...)
}

// GetECSServiceRecommendations mocks base method.
func (m *MockComputeoptimizerClient) GetECSServiceRecommendations(arg0 context.Context, arg1 *computeoptimizer.GetECSServiceRecommendationsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetECSServiceRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetECSServiceRecommendations", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetECSServiceRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetECSServiceRecommendations indicates an expected call of GetECSServiceRecommendations.
func (mr *MockComputeoptimizerClientMockRecorder) GetECSServiceRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetECSServiceRecommendations", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetECSServiceRecommendations), varargs...)
}

// GetEffectiveRecommendationPreferences mocks base method.
func (m *MockComputeoptimizerClient) GetEffectiveRecommendationPreferences(arg0 context.Context, arg1 *computeoptimizer.GetEffectiveRecommendationPreferencesInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEffectiveRecommendationPreferencesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEffectiveRecommendationPreferences", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEffectiveRecommendationPreferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEffectiveRecommendationPreferences indicates an expected call of GetEffectiveRecommendationPreferences.
func (mr *MockComputeoptimizerClientMockRecorder) GetEffectiveRecommendationPreferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEffectiveRecommendationPreferences", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEffectiveRecommendationPreferences), varargs...)
}

// GetEnrollmentStatus mocks base method.
func (m *MockComputeoptimizerClient) GetEnrollmentStatus(arg0 context.Context, arg1 *computeoptimizer.GetEnrollmentStatusInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnrollmentStatus", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEnrollmentStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnrollmentStatus indicates an expected call of GetEnrollmentStatus.
func (mr *MockComputeoptimizerClientMockRecorder) GetEnrollmentStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollmentStatus", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEnrollmentStatus), varargs...)
}

// GetEnrollmentStatusesForOrganization mocks base method.
func (m *MockComputeoptimizerClient) GetEnrollmentStatusesForOrganization(arg0 context.Context, arg1 *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnrollmentStatusesForOrganization", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnrollmentStatusesForOrganization indicates an expected call of GetEnrollmentStatusesForOrganization.
func (mr *MockComputeoptimizerClientMockRecorder) GetEnrollmentStatusesForOrganization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollmentStatusesForOrganization", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetEnrollmentStatusesForOrganization), varargs...)
}

// GetLambdaFunctionRecommendations mocks base method.
func (m *MockComputeoptimizerClient) GetLambdaFunctionRecommendations(arg0 context.Context, arg1 *computeoptimizer.GetLambdaFunctionRecommendationsInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLambdaFunctionRecommendations", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetLambdaFunctionRecommendationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLambdaFunctionRecommendations indicates an expected call of GetLambdaFunctionRecommendations.
func (mr *MockComputeoptimizerClientMockRecorder) GetLambdaFunctionRecommendations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLambdaFunctionRecommendations", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetLambdaFunctionRecommendations), varargs...)
}

// GetRecommendationPreferences mocks base method.
func (m *MockComputeoptimizerClient) GetRecommendationPreferences(arg0 context.Context, arg1 *computeoptimizer.GetRecommendationPreferencesInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetRecommendationPreferencesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecommendationPreferences", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetRecommendationPreferencesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationPreferences indicates an expected call of GetRecommendationPreferences.
func (mr *MockComputeoptimizerClientMockRecorder) GetRecommendationPreferences(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationPreferences", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetRecommendationPreferences), varargs...)
}

// GetRecommendationSummaries mocks base method.
func (m *MockComputeoptimizerClient) GetRecommendationSummaries(arg0 context.Context, arg1 *computeoptimizer.GetRecommendationSummariesInput, arg2 ...func(*computeoptimizer.Options)) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecommendationSummaries", varargs...)
	ret0, _ := ret[0].(*computeoptimizer.GetRecommendationSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendationSummaries indicates an expected call of GetRecommendationSummaries.
func (mr *MockComputeoptimizerClientMockRecorder) GetRecommendationSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationSummaries", reflect.TypeOf((*MockComputeoptimizerClient)(nil).GetRecommendationSummaries), varargs...)
}
