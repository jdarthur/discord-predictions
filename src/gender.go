package src

import (
	"time"
)

// Gender is the predicted gender of this particular user
type Gender struct {
	UserId                       string    `json:"user_id"`
	PredictedGender              string    `json:"predicted_gender"`
	Probability                  float64   `json:"probability"`
	ProbMale                     float64   `json:"prob_male"`
	ProbFemale                   float64   `json:"prob_female"`
	ProbNonBinaryGenderExpansive float64   `json:"prob_non_binary_gender_expansive"`
	ModelVersion                 time.Time `json:"model_version"`
	DayPt                        string    `json:"day_pt"`
}

func (g Gender) GetModelVersion() string {
	return g.ModelVersion.String()
}

// ProbabilityMale is the predicted probability that this user is male
type ProbabilityMale struct {
	V float64
}

func (p ProbabilityMale) Value() float64 {
	return p.V
}

func (p ProbabilityMale) Label() string {
	return "Male"
}

// ProbabilityFemale is the predicted probability that this user is female
type ProbabilityFemale struct {
	V float64
}

func (p ProbabilityFemale) Value() float64 {
	return p.V
}

func (p ProbabilityFemale) Label() string {
	return "Female"
}

// ProbabilityNonBinary is the predicted probability that this user is non-binary or gender expansive
type ProbabilityNonBinary struct {
	V float64
}

func (p ProbabilityNonBinary) Value() float64 {
	return p.V
}

func (p ProbabilityNonBinary) Label() string {
	return "Nonbinary"
}

// ListOfGender is a list of Gender values that implements the GraphableProbability interface
type ListOfGender []*Gender

func (l ListOfGender) Title() string {
	return "Predicted Gender"
}

func (l ListOfGender) XAxis() []time.Time {
	output := make([]time.Time, 0)
	for _, g := range l {
		output = append(output, g.ModelVersion)
	}
	return output
}

func (l ListOfGender) GetProbabilityMale() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, g := range l {
		output = append(output, ProbabilityMale{
			V: g.ProbMale,
		})
	}

	return output
}

func (l ListOfGender) GetProbabilityFemale() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, g := range l {
		output = append(output, ProbabilityFemale{
			V: g.ProbFemale,
		})
	}

	return output
}

func (l ListOfGender) GetProbabilityNonBinary() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, g := range l {
		output = append(output, ProbabilityNonBinary{
			V: g.ProbNonBinaryGenderExpansive,
		})
	}

	return output
}

func (l ListOfGender) YAxis() [][]YAxisItem {

	output := make([][]YAxisItem, 0)

	m := l.GetProbabilityMale()
	output = append(output, m)

	f := l.GetProbabilityFemale()
	output = append(output, f)

	nb := l.GetProbabilityNonBinary()
	output = append(output, nb)

	return output
}

func (g Gender) UniqueKey() string {
	return "predicted_gender"
}

func (g Gender) ParseInto() Parseable {
	return new(Gender)
}
