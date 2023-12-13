// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated by protoveneer. DO NOT EDIT.

package genai

import (
	"fmt"

	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta/generativelanguagepb"
	"github.com/google/generative-ai-go/internal/support"
)

// Blob contains raw media bytes.
//
// Text should not be sent as raw bytes, use the 'text' field.
type Blob struct {
	// The IANA standard MIME type of the source data.
	// Accepted types include: "image/png", "image/jpeg", "image/heic",
	// "image/heif", "image/webp".
	MIMEType string
	// Raw bytes for media formats.
	Data []byte
}

func (v *Blob) toProto() *pb.Blob {
	if v == nil {
		return nil
	}
	return &pb.Blob{
		MimeType: v.MIMEType,
		Data:     v.Data,
	}
}

func (Blob) fromProto(p *pb.Blob) *Blob {
	if p == nil {
		return nil
	}
	return &Blob{
		MIMEType: p.MimeType,
		Data:     p.Data,
	}
}

// BlockReason is specifies what was the reason why prompt was blocked.
type BlockReason int32

const (
	// BlockReasonUnspecified means default value. This value is unused.
	BlockReasonUnspecified BlockReason = 0
	// BlockReasonSafety means prompt was blocked due to safety reasons. You can inspect
	// `safety_ratings` to understand which safety category blocked it.
	BlockReasonSafety BlockReason = 1
	// BlockReasonOther means prompt was blocked due to unknown reaasons.
	BlockReasonOther BlockReason = 2
)

var namesForBlockReason = map[BlockReason]string{
	BlockReasonUnspecified: "BlockReasonUnspecified",
	BlockReasonSafety:      "BlockReasonSafety",
	BlockReasonOther:       "BlockReasonOther",
}

func (v BlockReason) String() string {
	if n, ok := namesForBlockReason[v]; ok {
		return n
	}
	return fmt.Sprintf("BlockReason(%d)", v)
}

// Candidate is a response candidate generated from the model.
type Candidate struct {
	// Output only. Index of the candidate in the list of candidates.
	Index int32
	// Output only. Generated content returned from the model.
	Content *Content
	// Optional. Output only. The reason why the model stopped generating tokens.
	//
	// If empty, the model has not stopped generating the tokens.
	FinishReason FinishReason
	// List of ratings for the safety of a response candidate.
	//
	// There is at most one rating per category.
	SafetyRatings []*SafetyRating
	// Output only. Citation information for model-generated candidate.
	//
	// This field may be populated with recitation information for any text
	// included in the `content`. These are passages that are "recited" from
	// copyrighted material in the foundational LLM's training data.
	CitationMetadata *CitationMetadata
	// Output only. Token count for this candidate.
	TokenCount int32
}

func (v *Candidate) toProto() *pb.Candidate {
	if v == nil {
		return nil
	}
	return &pb.Candidate{
		Index:            support.AddrOrNil(v.Index),
		Content:          v.Content.toProto(),
		FinishReason:     pb.Candidate_FinishReason(v.FinishReason),
		SafetyRatings:    support.TransformSlice(v.SafetyRatings, (*SafetyRating).toProto),
		CitationMetadata: v.CitationMetadata.toProto(),
		TokenCount:       v.TokenCount,
	}
}

func (Candidate) fromProto(p *pb.Candidate) *Candidate {
	if p == nil {
		return nil
	}
	return &Candidate{
		Index:            support.DerefOrZero(p.Index),
		Content:          (Content{}).fromProto(p.Content),
		FinishReason:     FinishReason(p.FinishReason),
		SafetyRatings:    support.TransformSlice(p.SafetyRatings, (SafetyRating{}).fromProto),
		CitationMetadata: (CitationMetadata{}).fromProto(p.CitationMetadata),
		TokenCount:       p.TokenCount,
	}
}

// CitationMetadata is a collection of source attributions for a piece of content.
type CitationMetadata struct {
	// Citations to sources for a specific response.
	CitationSources []*CitationSource
}

func (v *CitationMetadata) toProto() *pb.CitationMetadata {
	if v == nil {
		return nil
	}
	return &pb.CitationMetadata{
		CitationSources: support.TransformSlice(v.CitationSources, (*CitationSource).toProto),
	}
}

func (CitationMetadata) fromProto(p *pb.CitationMetadata) *CitationMetadata {
	if p == nil {
		return nil
	}
	return &CitationMetadata{
		CitationSources: support.TransformSlice(p.CitationSources, (CitationSource{}).fromProto),
	}
}

// CitationSource contains a citation to a source for a portion of a specific response.
type CitationSource struct {
	// Optional. Start of segment of the response that is attributed to this
	// source.
	//
	// Index indicates the start of the segment, measured in bytes.
	StartIndex *int32
	// Optional. End of the attributed segment, exclusive.
	EndIndex *int32
	// Optional. URI that is attributed as a source for a portion of the text.
	URI *string
	// Optional. License for the GitHub project that is attributed as a source for
	// segment.
	//
	// License info is required for code citations.
	License string
}

func (v *CitationSource) toProto() *pb.CitationSource {
	if v == nil {
		return nil
	}
	return &pb.CitationSource{
		StartIndex: v.StartIndex,
		EndIndex:   v.EndIndex,
		Uri:        v.URI,
		License:    support.AddrOrNil(v.License),
	}
}

func (CitationSource) fromProto(p *pb.CitationSource) *CitationSource {
	if p == nil {
		return nil
	}
	return &CitationSource{
		StartIndex: p.StartIndex,
		EndIndex:   p.EndIndex,
		URI:        p.Uri,
		License:    support.DerefOrZero(p.License),
	}
}

// Content is the base structured datatype containing multi-part content of a message.
//
// A `Content` includes a `role` field designating the producer of the `Content`
// and a `parts` field containing multi-part data that contains the content of
// the message turn.
type Content struct {
	// Ordered `Parts` that constitute a single message. Parts may have different
	// MIME types.
	Parts []Part
	// Optional. The producer of the content. Must be either 'user' or 'model'.
	//
	// Useful to set for multi-turn conversations, otherwise can be left blank
	// or unset.
	Role string
}

func (v *Content) toProto() *pb.Content {
	if v == nil {
		return nil
	}
	return &pb.Content{
		Parts: support.TransformSlice(v.Parts, partToProto),
		Role:  v.Role,
	}
}

func (Content) fromProto(p *pb.Content) *Content {
	if p == nil {
		return nil
	}
	return &Content{
		Parts: support.TransformSlice(p.Parts, partFromProto),
		Role:  p.Role,
	}
}

// ContentEmbedding is a list of floats representing an embedding.
type ContentEmbedding struct {
	// The embedding values.
	Values []float32
}

func (v *ContentEmbedding) toProto() *pb.ContentEmbedding {
	if v == nil {
		return nil
	}
	return &pb.ContentEmbedding{
		Values: v.Values,
	}
}

func (ContentEmbedding) fromProto(p *pb.ContentEmbedding) *ContentEmbedding {
	if p == nil {
		return nil
	}
	return &ContentEmbedding{
		Values: p.Values,
	}
}

// CountTokensResponse is a response from `CountTokens`.
//
// It returns the model's `token_count` for the `prompt`.
type CountTokensResponse struct {
	// The number of tokens that the `model` tokenizes the `prompt` into.
	//
	// Always non-negative.
	TotalTokens int32
}

func (v *CountTokensResponse) toProto() *pb.CountTokensResponse {
	if v == nil {
		return nil
	}
	return &pb.CountTokensResponse{
		TotalTokens: v.TotalTokens,
	}
}

func (CountTokensResponse) fromProto(p *pb.CountTokensResponse) *CountTokensResponse {
	if p == nil {
		return nil
	}
	return &CountTokensResponse{
		TotalTokens: p.TotalTokens,
	}
}

// EmbedContentResponse is the response to an `EmbedContentRequest`.
type EmbedContentResponse struct {
	// Output only. The embedding generated from the input content.
	Embedding *ContentEmbedding
}

func (v *EmbedContentResponse) toProto() *pb.EmbedContentResponse {
	if v == nil {
		return nil
	}
	return &pb.EmbedContentResponse{
		Embedding: v.Embedding.toProto(),
	}
}

func (EmbedContentResponse) fromProto(p *pb.EmbedContentResponse) *EmbedContentResponse {
	if p == nil {
		return nil
	}
	return &EmbedContentResponse{
		Embedding: (ContentEmbedding{}).fromProto(p.Embedding),
	}
}

// FinishReason is defines the reason why the model stopped generating tokens.
type FinishReason int32

const (
	// FinishReasonUnspecified means default value. This value is unused.
	FinishReasonUnspecified FinishReason = 0
	// FinishReasonStop means natural stop point of the model or provided stop sequence.
	FinishReasonStop FinishReason = 1
	// FinishReasonMaxTokens means the maximum number of tokens as specified in the request was reached.
	FinishReasonMaxTokens FinishReason = 2
	// FinishReasonSafety means the candidate content was flagged for safety reasons.
	FinishReasonSafety FinishReason = 3
	// FinishReasonRecitation means the candidate content was flagged for recitation reasons.
	FinishReasonRecitation FinishReason = 4
	// FinishReasonOther means unknown reason.
	FinishReasonOther FinishReason = 5
)

var namesForFinishReason = map[FinishReason]string{
	FinishReasonUnspecified: "FinishReasonUnspecified",
	FinishReasonStop:        "FinishReasonStop",
	FinishReasonMaxTokens:   "FinishReasonMaxTokens",
	FinishReasonSafety:      "FinishReasonSafety",
	FinishReasonRecitation:  "FinishReasonRecitation",
	FinishReasonOther:       "FinishReasonOther",
}

func (v FinishReason) String() string {
	if n, ok := namesForFinishReason[v]; ok {
		return n
	}
	return fmt.Sprintf("FinishReason(%d)", v)
}

// FunctionCall is a predicted `FunctionCall` returned from the model that contains
// a string representing the `FunctionDeclaration.name` with the
// arguments and their values.
type FunctionCall struct {
	// Required. The name of the function to call.
	// Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	// length of 63.
	Name string
	// Optional. The function parameters and values in JSON object format.
	Args map[string]any
}

func (v *FunctionCall) toProto() *pb.FunctionCall {
	if v == nil {
		return nil
	}
	return &pb.FunctionCall{
		Name: v.Name,
		Args: support.MapToStructPB(v.Args),
	}
}

func (FunctionCall) fromProto(p *pb.FunctionCall) *FunctionCall {
	if p == nil {
		return nil
	}
	return &FunctionCall{
		Name: p.Name,
		Args: support.MapFromStructPB(p.Args),
	}
}

// FunctionDeclaration is structured representation of a function declaration as defined by the
// [OpenAPI 3.03 specification](https://spec.openapis.org/oas/v3.0.3). Included
// in this declaration are the function name and parameters. This
// FunctionDeclaration is a representation of a block of code that can be used
// as a `Tool` by the model and executed by the client.
type FunctionDeclaration struct {
	// Required. The name of the function.
	// Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	// length of 63.
	Name string
	// Required. A brief description of the function.
	Description string
	// Optional. Describes the parameters to this function. Reflects the Open
	// API 3.03 Parameter Object string Key: the name of the parameter. Parameter
	// names are case sensitive. Schema Value: the Schema defining the type used
	// for the parameter.
	Parameters *Schema
}

func (v *FunctionDeclaration) toProto() *pb.FunctionDeclaration {
	if v == nil {
		return nil
	}
	return &pb.FunctionDeclaration{
		Name:        v.Name,
		Description: v.Description,
		Parameters:  v.Parameters.toProto(),
	}
}

func (FunctionDeclaration) fromProto(p *pb.FunctionDeclaration) *FunctionDeclaration {
	if p == nil {
		return nil
	}
	return &FunctionDeclaration{
		Name:        p.Name,
		Description: p.Description,
		Parameters:  (Schema{}).fromProto(p.Parameters),
	}
}

// FunctionResponse is the result output from a `FunctionCall` that contains a string
// representing the `FunctionDeclaration.name` and a structured JSON
// object containing any output from the function is used as context to
// the model. This should contain the result of a`FunctionCall` made
// based on model prediction.
type FunctionResponse struct {
	// Required. The name of the function to call.
	// Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	// length of 63.
	Name string
	// Required. The function response in JSON object format.
	Response map[string]any
}

func (v *FunctionResponse) toProto() *pb.FunctionResponse {
	if v == nil {
		return nil
	}
	return &pb.FunctionResponse{
		Name:     v.Name,
		Response: support.MapToStructPB(v.Response),
	}
}

func (FunctionResponse) fromProto(p *pb.FunctionResponse) *FunctionResponse {
	if p == nil {
		return nil
	}
	return &FunctionResponse{
		Name:     p.Name,
		Response: support.MapFromStructPB(p.Response),
	}
}

// GenerationConfig is configuration options for model generation and outputs. Not all parameters
// may be configurable for every model.
type GenerationConfig struct {
	// Optional. Number of generated responses to return.
	//
	// This value must be between [1, 8], inclusive. If unset, this will default
	// to 1.
	CandidateCount int32
	// Optional. The set of character sequences (up to 5) that will stop output
	// generation. If specified, the API will stop at the first appearance of a
	// stop sequence. The stop sequence will not be included as part of the
	// response.
	StopSequences []string
	// Optional. The maximum number of tokens to include in a candidate.
	//
	// If unset, this will default to output_token_limit specified in the `Model`
	// specification.
	MaxOutputTokens int32
	// Optional. Controls the randomness of the output.
	// Note: The default value varies by model, see the `Model.temperature`
	// attribute of the `Model` returned the `getModel` function.
	//
	// Values can range from [0.0,1.0],
	// inclusive. A value closer to 1.0 will produce responses that are more
	// varied and creative, while a value closer to 0.0 will typically result in
	// more straightforward responses from the model.
	Temperature float32
	// Optional. The maximum cumulative probability of tokens to consider when
	// sampling.
	//
	// The model uses combined Top-k and nucleus sampling.
	//
	// Tokens are sorted based on their assigned probabilities so that only the
	// most likely tokens are considered. Top-k sampling directly limits the
	// maximum number of tokens to consider, while Nucleus sampling limits number
	// of tokens based on the cumulative probability.
	//
	// Note: The default value varies by model, see the `Model.top_p`
	// attribute of the `Model` returned the `getModel` function.
	TopP float32
	// Optional. The maximum number of tokens to consider when sampling.
	//
	// The model uses combined Top-k and nucleus sampling.
	//
	// Top-k sampling considers the set of `top_k` most probable tokens.
	// Defaults to 40.
	//
	// Note: The default value varies by model, see the `Model.top_k`
	// attribute of the `Model` returned the `getModel` function.
	TopK int32
}

func (v *GenerationConfig) toProto() *pb.GenerationConfig {
	if v == nil {
		return nil
	}
	return &pb.GenerationConfig{
		CandidateCount:  support.AddrOrNil(v.CandidateCount),
		StopSequences:   v.StopSequences,
		MaxOutputTokens: support.AddrOrNil(v.MaxOutputTokens),
		Temperature:     support.AddrOrNil(v.Temperature),
		TopP:            support.AddrOrNil(v.TopP),
		TopK:            support.AddrOrNil(v.TopK),
	}
}

func (GenerationConfig) fromProto(p *pb.GenerationConfig) *GenerationConfig {
	if p == nil {
		return nil
	}
	return &GenerationConfig{
		CandidateCount:  support.DerefOrZero(p.CandidateCount),
		StopSequences:   p.StopSequences,
		MaxOutputTokens: support.DerefOrZero(p.MaxOutputTokens),
		Temperature:     support.DerefOrZero(p.Temperature),
		TopP:            support.DerefOrZero(p.TopP),
		TopK:            support.DerefOrZero(p.TopK),
	}
}

// HarmBlockThreshold specifies block at and beyond a specified harm probability.
type HarmBlockThreshold int32

const (
	// HarmBlockUnspecified means threshold is unspecified.
	HarmBlockUnspecified HarmBlockThreshold = 0
	// HarmBlockLowAndAbove means content with NEGLIGIBLE will be allowed.
	HarmBlockLowAndAbove HarmBlockThreshold = 1
	// HarmBlockMediumAndAbove means content with NEGLIGIBLE and LOW will be allowed.
	HarmBlockMediumAndAbove HarmBlockThreshold = 2
	// HarmBlockOnlyHigh means content with NEGLIGIBLE, LOW, and MEDIUM will be allowed.
	HarmBlockOnlyHigh HarmBlockThreshold = 3
	// HarmBlockNone means all content will be allowed.
	HarmBlockNone HarmBlockThreshold = 4
)

var namesForHarmBlockThreshold = map[HarmBlockThreshold]string{
	HarmBlockUnspecified:    "HarmBlockUnspecified",
	HarmBlockLowAndAbove:    "HarmBlockLowAndAbove",
	HarmBlockMediumAndAbove: "HarmBlockMediumAndAbove",
	HarmBlockOnlyHigh:       "HarmBlockOnlyHigh",
	HarmBlockNone:           "HarmBlockNone",
}

func (v HarmBlockThreshold) String() string {
	if n, ok := namesForHarmBlockThreshold[v]; ok {
		return n
	}
	return fmt.Sprintf("HarmBlockThreshold(%d)", v)
}

// HarmCategory specifies the category of a rating.
//
// These categories cover various kinds of harms that developers
// may wish to adjust.
type HarmCategory int32

const (
	// HarmCategoryUnspecified means category is unspecified.
	HarmCategoryUnspecified HarmCategory = 0
	// HarmCategoryDerogatory means negative or harmful comments targeting identity and/or protected attribute.
	HarmCategoryDerogatory HarmCategory = 1
	// HarmCategoryToxicity means content that is rude, disrepspectful, or profane.
	HarmCategoryToxicity HarmCategory = 2
	// HarmCategoryViolence means describes scenarios depictng violence against an individual or group, or
	// general descriptions of gore.
	HarmCategoryViolence HarmCategory = 3
	// HarmCategorySexual means contains references to sexual acts or other lewd content.
	HarmCategorySexual HarmCategory = 4
	// HarmCategoryMedical means promotes unchecked medical advice.
	HarmCategoryMedical HarmCategory = 5
	// HarmCategoryDangerous means dangerous content that promotes, facilitates, or encourages harmful acts.
	HarmCategoryDangerous HarmCategory = 6
	// HarmCategoryHarassment means harasment content.
	HarmCategoryHarassment HarmCategory = 7
	// HarmCategoryHateSpeech means hate speech and content.
	HarmCategoryHateSpeech HarmCategory = 8
	// HarmCategorySexuallyExplicit means sexually explicit content.
	HarmCategorySexuallyExplicit HarmCategory = 9
	// HarmCategoryDangerousContent means dangerous content.
	HarmCategoryDangerousContent HarmCategory = 10
)

var namesForHarmCategory = map[HarmCategory]string{
	HarmCategoryUnspecified:      "HarmCategoryUnspecified",
	HarmCategoryDerogatory:       "HarmCategoryDerogatory",
	HarmCategoryToxicity:         "HarmCategoryToxicity",
	HarmCategoryViolence:         "HarmCategoryViolence",
	HarmCategorySexual:           "HarmCategorySexual",
	HarmCategoryMedical:          "HarmCategoryMedical",
	HarmCategoryDangerous:        "HarmCategoryDangerous",
	HarmCategoryHarassment:       "HarmCategoryHarassment",
	HarmCategoryHateSpeech:       "HarmCategoryHateSpeech",
	HarmCategorySexuallyExplicit: "HarmCategorySexuallyExplicit",
	HarmCategoryDangerousContent: "HarmCategoryDangerousContent",
}

func (v HarmCategory) String() string {
	if n, ok := namesForHarmCategory[v]; ok {
		return n
	}
	return fmt.Sprintf("HarmCategory(%d)", v)
}

// HarmProbability specifies the probability that a piece of content is harmful.
//
// The classification system gives the probability of the content being
// unsafe. This does not indicate the severity of harm for a piece of content.
type HarmProbability int32

const (
	// HarmProbabilityUnspecified means probability is unspecified.
	HarmProbabilityUnspecified HarmProbability = 0
	// HarmProbabilityNegligible means content has a negligible chance of being unsafe.
	HarmProbabilityNegligible HarmProbability = 1
	// HarmProbabilityLow means content has a low chance of being unsafe.
	HarmProbabilityLow HarmProbability = 2
	// HarmProbabilityMedium means content has a medium chance of being unsafe.
	HarmProbabilityMedium HarmProbability = 3
	// HarmProbabilityHigh means content has a high chance of being unsafe.
	HarmProbabilityHigh HarmProbability = 4
)

var namesForHarmProbability = map[HarmProbability]string{
	HarmProbabilityUnspecified: "HarmProbabilityUnspecified",
	HarmProbabilityNegligible:  "HarmProbabilityNegligible",
	HarmProbabilityLow:         "HarmProbabilityLow",
	HarmProbabilityMedium:      "HarmProbabilityMedium",
	HarmProbabilityHigh:        "HarmProbabilityHigh",
}

func (v HarmProbability) String() string {
	if n, ok := namesForHarmProbability[v]; ok {
		return n
	}
	return fmt.Sprintf("HarmProbability(%d)", v)
}

// Model is information about a Generative Language Model.
type Model struct {
	// Required. The resource name of the `Model`.
	//
	// Format: `models/{model}` with a `{model}` naming convention of:
	//
	// * "{base_model_id}-{version}"
	//
	// Examples:
	//
	// * `models/chat-bison-001`
	Name string
	// Required. The name of the base model, pass this to the generation request.
	//
	// Examples:
	//
	// * `chat-bison`
	BaseModeID string
	// Required. The version number of the model.
	//
	// This represents the major version
	Version string
	// The human-readable name of the model. E.g. "Chat Bison".
	//
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	DisplayName string
	// A short description of the model.
	Description string
	// Maximum number of input tokens allowed for this model.
	InputTokenLimit int32
	// Maximum number of output tokens available for this model.
	OutputTokenLimit int32
	// The model's supported generation methods.
	//
	// The method names are defined as Pascal case
	// strings, such as `generateMessage` which correspond to API methods.
	SupportedGenerationMethods []string
	// Controls the randomness of the output.
	//
	// Values can range over `[0.0,1.0]`, inclusive. A value closer to `1.0` will
	// produce responses that are more varied, while a value closer to `0.0` will
	// typically result in less surprising responses from the model.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	Temperature float32
	// For Nucleus sampling.
	//
	// Nucleus sampling considers the smallest set of tokens whose probability
	// sum is at least `top_p`.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	TopP float32
	// For Top-k sampling.
	//
	// Top-k sampling considers the set of `top_k` most probable tokens.
	// This value specifies default to be used by the backend while making the
	// call to the model.
	TopK int32
}

func (v *Model) toProto() *pb.Model {
	if v == nil {
		return nil
	}
	return &pb.Model{
		Name:                       v.Name,
		BaseModelId:                v.BaseModeID,
		Version:                    v.Version,
		DisplayName:                v.DisplayName,
		Description:                v.Description,
		InputTokenLimit:            v.InputTokenLimit,
		OutputTokenLimit:           v.OutputTokenLimit,
		SupportedGenerationMethods: v.SupportedGenerationMethods,
		Temperature:                support.AddrOrNil(v.Temperature),
		TopP:                       support.AddrOrNil(v.TopP),
		TopK:                       support.AddrOrNil(v.TopK),
	}
}

func (Model) fromProto(p *pb.Model) *Model {
	if p == nil {
		return nil
	}
	return &Model{
		Name:                       p.Name,
		BaseModeID:                 p.BaseModelId,
		Version:                    p.Version,
		DisplayName:                p.DisplayName,
		Description:                p.Description,
		InputTokenLimit:            p.InputTokenLimit,
		OutputTokenLimit:           p.OutputTokenLimit,
		SupportedGenerationMethods: p.SupportedGenerationMethods,
		Temperature:                support.DerefOrZero(p.Temperature),
		TopP:                       support.DerefOrZero(p.TopP),
		TopK:                       support.DerefOrZero(p.TopK),
	}
}

// PromptFeedback contains a set of the feedback metadata the prompt specified in
// `GenerateContentRequest.content`.
type PromptFeedback struct {
	// Optional. If set, the prompt was blocked and no candidates are returned.
	// Rephrase your prompt.
	BlockReason BlockReason
	// Ratings for safety of the prompt.
	// There is at most one rating per category.
	SafetyRatings []*SafetyRating
}

func (v *PromptFeedback) toProto() *pb.GenerateContentResponse_PromptFeedback {
	if v == nil {
		return nil
	}
	return &pb.GenerateContentResponse_PromptFeedback{
		BlockReason:   pb.GenerateContentResponse_PromptFeedback_BlockReason(v.BlockReason),
		SafetyRatings: support.TransformSlice(v.SafetyRatings, (*SafetyRating).toProto),
	}
}

func (PromptFeedback) fromProto(p *pb.GenerateContentResponse_PromptFeedback) *PromptFeedback {
	if p == nil {
		return nil
	}
	return &PromptFeedback{
		BlockReason:   BlockReason(p.BlockReason),
		SafetyRatings: support.TransformSlice(p.SafetyRatings, (SafetyRating{}).fromProto),
	}
}

// SafetyRating is the safety rating for a piece of content.
//
// The safety rating contains the category of harm and the
// harm probability level in that category for a piece of content.
// Content is classified for safety across a number of
// harm categories and the probability of the harm classification is included
// here.
type SafetyRating struct {
	// Required. The category for this rating.
	Category HarmCategory
	// Required. The probability of harm for this content.
	Probability HarmProbability
	// Was this content blocked because of this rating?
	Blocked bool
}

func (v *SafetyRating) toProto() *pb.SafetyRating {
	if v == nil {
		return nil
	}
	return &pb.SafetyRating{
		Category:    pb.HarmCategory(v.Category),
		Probability: pb.SafetyRating_HarmProbability(v.Probability),
		Blocked:     v.Blocked,
	}
}

func (SafetyRating) fromProto(p *pb.SafetyRating) *SafetyRating {
	if p == nil {
		return nil
	}
	return &SafetyRating{
		Category:    HarmCategory(p.Category),
		Probability: HarmProbability(p.Probability),
		Blocked:     p.Blocked,
	}
}

// SafetySetting is safety setting, affecting the safety-blocking behavior.
//
// Passing a safety setting for a category changes the allowed proability that
// content is blocked.
type SafetySetting struct {
	// Required. The category for this setting.
	Category HarmCategory
	// Required. Controls the probability threshold at which harm is blocked.
	Threshold HarmBlockThreshold
}

func (v *SafetySetting) toProto() *pb.SafetySetting {
	if v == nil {
		return nil
	}
	return &pb.SafetySetting{
		Category:  pb.HarmCategory(v.Category),
		Threshold: pb.SafetySetting_HarmBlockThreshold(v.Threshold),
	}
}

func (SafetySetting) fromProto(p *pb.SafetySetting) *SafetySetting {
	if p == nil {
		return nil
	}
	return &SafetySetting{
		Category:  HarmCategory(p.Category),
		Threshold: HarmBlockThreshold(p.Threshold),
	}
}

// Schema is the `Schema` object allows the definition of input and output data types.
// These types can be objects, but also primitives and arrays.
// Represents a select subset of an [OpenAPI 3.0 schema
// object](https://spec.openapis.org/oas/v3.0.3#schema).
type Schema struct {
	// Optional. Data type.
	Type Type
	// Optional. The format of the data. This is used obnly for primative
	// datatypes. Supported formats:
	//  for NUMBER type: float, double
	//  for INTEGER type: int32, int64
	Format string
	// Optional. A brief description of the parameter. This could contain examples
	// of use. Parameter description may be formatted as Markdown.
	Description string
	// Optional. Indicates if the value may be null.
	Nullable bool
	// Optional. Possible values of the element of Type.STRING with enum format.
	// For example we can define an Enum Direction as :
	// {type:STRING, format:enum, enum:["EAST", NORTH", "SOUTH", "WEST"]}
	Enum []string
	// Optional. Schema of the elements of Type.ARRAY.
	Items *Schema
	// Optional. Properties of Type.OBJECT.
	Properties map[string]*Schema
	// Optional. Required properties of Type.OBJECT.
	Required []string
}

func (v *Schema) toProto() *pb.Schema {
	if v == nil {
		return nil
	}
	return &pb.Schema{
		Type:        pb.Type(v.Type),
		Format:      v.Format,
		Description: v.Description,
		Nullable:    v.Nullable,
		Enum:        v.Enum,
		Items:       v.Items.toProto(),
		Properties:  support.TransformMapValues(v.Properties, (*Schema).toProto),
		Required:    v.Required,
	}
}

func (Schema) fromProto(p *pb.Schema) *Schema {
	if p == nil {
		return nil
	}
	return &Schema{
		Type:        Type(p.Type),
		Format:      p.Format,
		Description: p.Description,
		Nullable:    p.Nullable,
		Enum:        p.Enum,
		Items:       (Schema{}).fromProto(p.Items),
		Properties:  support.TransformMapValues(p.Properties, (Schema{}).fromProto),
		Required:    p.Required,
	}
}

// TaskType is type of task for which the embedding will be used.
type TaskType int32

const (
	// TaskTypeUnspecified means unset value, which will default to one of the other enum values.
	TaskTypeUnspecified TaskType = 0
	// TaskTypeRetrievalQuery means specifies the given text is a query in a search/retrieval setting.
	TaskTypeRetrievalQuery TaskType = 1
	// TaskTypeRetrievalDocument means specifies the given text is a document from the corpus being searched.
	TaskTypeRetrievalDocument TaskType = 2
	// TaskTypeSemanticSimilarity means specifies the given text will be used for STS.
	TaskTypeSemanticSimilarity TaskType = 3
	// TaskTypeClassification means specifies that the given text will be classified.
	TaskTypeClassification TaskType = 4
	// TaskTypeClustering means specifies that the embeddings will be used for clustering.
	TaskTypeClustering TaskType = 5
)

var namesForTaskType = map[TaskType]string{
	TaskTypeUnspecified:        "TaskTypeUnspecified",
	TaskTypeRetrievalQuery:     "TaskTypeRetrievalQuery",
	TaskTypeRetrievalDocument:  "TaskTypeRetrievalDocument",
	TaskTypeSemanticSimilarity: "TaskTypeSemanticSimilarity",
	TaskTypeClassification:     "TaskTypeClassification",
	TaskTypeClustering:         "TaskTypeClustering",
}

func (v TaskType) String() string {
	if n, ok := namesForTaskType[v]; ok {
		return n
	}
	return fmt.Sprintf("TaskType(%d)", v)
}

// Tool details that the model may use to generate response.
//
// A `Tool` is a piece of code that enables the system to interact with
// external systems to perform an action, or set of actions, outside of
// knowledge and scope of the model.
type Tool struct {
	// Optional. A list of `FunctionDeclarations` available to the model that can
	// be used for function calling.
	//
	// The model or system does not execute the function. Instead the defined
	// function may be returned as a [FunctionCall][content.part.function_call]
	// with arguments to the client side for execution. The model may decide to
	// call a subset of these functions by populating
	// [FunctionCall][content.part.function_call] in the response. The next
	// conversation turn may contain a
	// [FunctionResponse][content.part.function_response]
	// with the [conent.role] "function" generation context for the next model
	// turn.
	FunctionDeclarations []*FunctionDeclaration
}

func (v *Tool) toProto() *pb.Tool {
	if v == nil {
		return nil
	}
	return &pb.Tool{
		FunctionDeclarations: support.TransformSlice(v.FunctionDeclarations, (*FunctionDeclaration).toProto),
	}
}

func (Tool) fromProto(p *pb.Tool) *Tool {
	if p == nil {
		return nil
	}
	return &Tool{
		FunctionDeclarations: support.TransformSlice(p.FunctionDeclarations, (FunctionDeclaration{}).fromProto),
	}
}

// Type contains the list of OpenAPI data types as defined by
// https://spec.openapis.org/oas/v3.0.3#data-types
type Type int32

const (
	// TypeUnspecified means not specified, should not be used.
	TypeUnspecified Type = 0
	// TypeString means string type.
	TypeString Type = 1
	// TypeNumber means number type.
	TypeNumber Type = 2
	// TypeInteger means integer type.
	TypeInteger Type = 3
	// TypeBoolean means boolean type.
	TypeBoolean Type = 4
	// TypeArray means array type.
	TypeArray Type = 5
	// TypeObject means object type.
	TypeObject Type = 6
)

var namesForType = map[Type]string{
	TypeUnspecified: "TypeUnspecified",
	TypeString:      "TypeString",
	TypeNumber:      "TypeNumber",
	TypeInteger:     "TypeInteger",
	TypeBoolean:     "TypeBoolean",
	TypeArray:       "TypeArray",
	TypeObject:      "TypeObject",
}

func (v Type) String() string {
	if n, ok := namesForType[v]; ok {
		return n
	}
	return fmt.Sprintf("Type(%d)", v)
}
