package interfaces

// mudler suggested these interfaces may belong up in localai itself, as go interfaces usually live near the consumer.
// For now, comment them out rather than fully remove, but let's see if we actually need them.

// type AIBackend interface {
// 	Free()
// 	GetBackendName() string

// 	// Somethign to do with new / init?
// }

// type LLMTextBackend interface {
// 	AIBackend

// 	Predict(text string, opts ...PredictTextOptionSetter) (string, error)
// 	PredictWithOptions(text string, opts PredictTextOptions) (string, error)

// 	TokenPredict(tokens []int, opts ...PredictTextOptionSetter) (string, error)
// 	TokenPredictWithOptions(tokens []int, opts PredictTextOptions) (string, error)
// }

// // TODO: These are totally untested and here for illustration until LLMTextBackend is wired up!

// type ImageGenerationBackend interface {
// 	AIBackend
// 	GenerateImage(height, width, mode, step, seed int, positive_prompt, negative_prompt, dst string)
// }

// type AudioBackend interface {
// 	AIBackend
// 	Transcript(path, langage string)
// }

// // Brainstorming init below, not in use yet

// // func NewLLMBackendWithOptions[TBackend LLMBackend](modelPath string, opts InitializationOptions) TBackend {

// // }
