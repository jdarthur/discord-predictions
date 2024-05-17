package src

import (
	"time"
)

// Age is the predicted age range of this particular user
type Age struct {
	UserId       string    `json:"user_id"`
	PredictedAge string    `json:"predicted_age"`
	Probability  float64   `json:"probability"`
	Prob1317     float64   `json:"prob_13_17"`
	Prob1824     float64   `json:"prob_18_24"`
	Prob2534     float64   `json:"prob_25_34"`
	Prob35Over   float64   `json:"prob_35_over"`
	ModelVersion time.Time `json:"model_version"`
	DayPt        string    `json:"day_pt"`
}

func (a Age) UniqueKey() string {
	return "predicted_age"
}

func (a Age) ParseInto() Parseable {
	return new(Age)
}

// Probability13To17 is the predicted probability that this user is ages 13 to 17
type Probability13To17 struct {
	V float64
}

func (p Probability13To17) Label() string {
	return "Age 13-17"
}

func (p Probability13To17) Value() float64 {
	return p.V
}

// Probability18To24 is the predicted probability that this user is ages 18 to 24
type Probability18To24 struct {
	V float64
}

func (p Probability18To24) Label() string {
	return "Age 18-24"
}

func (p Probability18To24) Value() float64 {
	return p.V
}

// Probability25To34 is the predicted probability that this user is ages 25-34
type Probability25To34 struct {
	V float64
}

func (p Probability25To34) Label() string {
	return "Age 25-34"
}

func (p Probability25To34) Value() float64 {
	return p.V
}

// Probability35Plus is the predicted probability that this user is age 35 or older
type Probability35Plus struct {
	V float64
}

func (p Probability35Plus) Label() string {
	return "Age 35+"
}

func (p Probability35Plus) Value() float64 {
	return p.V
}

// ListOfAge is a list of Age values that implements the GraphableProbability interface
type ListOfAge []*Age

func (l ListOfAge) Title() string {
	return "Predicted Age"
}
func (l ListOfAge) Get1317() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, age := range l {
		output = append(output, Probability13To17{V: age.Prob1317})
	}

	return output
}

func (l ListOfAge) Get1824() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, age := range l {
		output = append(output, Probability18To24{V: age.Prob1824})
	}

	return output
}

func (l ListOfAge) Get2534() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, age := range l {
		output = append(output, Probability25To34{V: age.Prob2534})
	}

	return output
}

func (l ListOfAge) Get35Plus() []YAxisItem {
	output := make([]YAxisItem, 0)

	for _, age := range l {
		output = append(output, Probability35Plus{V: age.Prob35Over})
	}

	return output
}

func (l ListOfAge) XAxis() []time.Time {

	output := make([]time.Time, 0)
	for _, age := range l {
		output = append(output, age.ModelVersion)
	}

	return output
}

func (l ListOfAge) YAxis() [][]YAxisItem {
	output := make([][]YAxisItem, 0)
	m := l.Get1317()
	output = append(output, m)

	m = l.Get1824()
	output = append(output, m)

	m = l.Get2534()
	output = append(output, m)

	m = l.Get35Plus()
	output = append(output, m)

	return output

}
