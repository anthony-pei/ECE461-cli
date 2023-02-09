package metrics

import (
	"testing"
)

type MockModule struct {
	URL             string
	License         string
	OpenIssues      int
	StargazersCount int
	Contributors    int
	Name            string
}

func (m MockModule) GetGitHubUrl() string {
	return m.URL
}

func (m MockModule) GetLicense() string {
	return m.License
}

func (m MockModule) GetOpenIssuesCount() int {
	return m.OpenIssues
}

func (m MockModule) GetStargazerCount() int {
	return m.StargazersCount
}

func (m MockModule) GetContributorCount() int {
	return m.Contributors
}

func (m MockModule) GetName() string {
	return m.Name
}

func (m MockModule) Clone(dir string) {
}

func TestRampUpNoComments(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 100, 0, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 0.0)
}
func TestRampUpEqualCommentsCode(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 100, 100, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 1.0)
}

func TestRampUpNoCode(t *testing.T) {
	m := MockModule{}
	rampUpMetric := RampUpMetric{}

	analyzeDirFunction = func(dir string) (int64, int64, int64, int64) {
		return 0, 0, 0, 0
	}
	assertEquals(t, "", rampUpMetric.CalculateScore(m), 0.0)
}
func TestCorrectnessZeroIssues(t *testing.T) {
	m := MockModule{OpenIssues: 0, StargazersCount: 10}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.99999)
}

func TestCorrectnessZeroStargazers(t *testing.T) {
	m := MockModule{OpenIssues: 10, StargazersCount: 0}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.0)
}
func TestCorrectnessEqualStargazersOpenIssues(t *testing.T) {
	m := MockModule{OpenIssues: 10, StargazersCount: 10}
	correctnessMetric := CorrectnessMetric{}

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.0)
}
func assertEquals(t *testing.T, desc string, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("%v Got: %v (%T), Want:%v (%T)", desc, got, got, want, want)
	}
}
