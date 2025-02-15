// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

//go:generate mockgen -package=mocks -destination=../mocks/frauddetector.go -source=frauddetector.go FrauddetectorClient
type FrauddetectorClient interface {
	BatchGetVariable(context.Context, *frauddetector.BatchGetVariableInput, ...func(*frauddetector.Options)) (*frauddetector.BatchGetVariableOutput, error)
	DescribeDetector(context.Context, *frauddetector.DescribeDetectorInput, ...func(*frauddetector.Options)) (*frauddetector.DescribeDetectorOutput, error)
	DescribeModelVersions(context.Context, *frauddetector.DescribeModelVersionsInput, ...func(*frauddetector.Options)) (*frauddetector.DescribeModelVersionsOutput, error)
	GetBatchImportJobs(context.Context, *frauddetector.GetBatchImportJobsInput, ...func(*frauddetector.Options)) (*frauddetector.GetBatchImportJobsOutput, error)
	GetBatchPredictionJobs(context.Context, *frauddetector.GetBatchPredictionJobsInput, ...func(*frauddetector.Options)) (*frauddetector.GetBatchPredictionJobsOutput, error)
	GetDeleteEventsByEventTypeStatus(context.Context, *frauddetector.GetDeleteEventsByEventTypeStatusInput, ...func(*frauddetector.Options)) (*frauddetector.GetDeleteEventsByEventTypeStatusOutput, error)
	GetDetectorVersion(context.Context, *frauddetector.GetDetectorVersionInput, ...func(*frauddetector.Options)) (*frauddetector.GetDetectorVersionOutput, error)
	GetDetectors(context.Context, *frauddetector.GetDetectorsInput, ...func(*frauddetector.Options)) (*frauddetector.GetDetectorsOutput, error)
	GetEntityTypes(context.Context, *frauddetector.GetEntityTypesInput, ...func(*frauddetector.Options)) (*frauddetector.GetEntityTypesOutput, error)
	GetEvent(context.Context, *frauddetector.GetEventInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventOutput, error)
	GetEventPrediction(context.Context, *frauddetector.GetEventPredictionInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionOutput, error)
	GetEventPredictionMetadata(context.Context, *frauddetector.GetEventPredictionMetadataInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventPredictionMetadataOutput, error)
	GetEventTypes(context.Context, *frauddetector.GetEventTypesInput, ...func(*frauddetector.Options)) (*frauddetector.GetEventTypesOutput, error)
	GetExternalModels(context.Context, *frauddetector.GetExternalModelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetExternalModelsOutput, error)
	GetKMSEncryptionKey(context.Context, *frauddetector.GetKMSEncryptionKeyInput, ...func(*frauddetector.Options)) (*frauddetector.GetKMSEncryptionKeyOutput, error)
	GetLabels(context.Context, *frauddetector.GetLabelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetLabelsOutput, error)
	GetModelVersion(context.Context, *frauddetector.GetModelVersionInput, ...func(*frauddetector.Options)) (*frauddetector.GetModelVersionOutput, error)
	GetModels(context.Context, *frauddetector.GetModelsInput, ...func(*frauddetector.Options)) (*frauddetector.GetModelsOutput, error)
	GetOutcomes(context.Context, *frauddetector.GetOutcomesInput, ...func(*frauddetector.Options)) (*frauddetector.GetOutcomesOutput, error)
	GetRules(context.Context, *frauddetector.GetRulesInput, ...func(*frauddetector.Options)) (*frauddetector.GetRulesOutput, error)
	GetVariables(context.Context, *frauddetector.GetVariablesInput, ...func(*frauddetector.Options)) (*frauddetector.GetVariablesOutput, error)
	ListEventPredictions(context.Context, *frauddetector.ListEventPredictionsInput, ...func(*frauddetector.Options)) (*frauddetector.ListEventPredictionsOutput, error)
	ListTagsForResource(context.Context, *frauddetector.ListTagsForResourceInput, ...func(*frauddetector.Options)) (*frauddetector.ListTagsForResourceOutput, error)
}
