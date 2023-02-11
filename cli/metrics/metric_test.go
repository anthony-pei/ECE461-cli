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

	assertEquals(t, "Correctness (0, 10)", correctnessMetric.CalculateScore(m), 0.99999000001)
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
func TestBusFactorZeroOrOneCOntributor(t *testing.T) {
	m1 := MockModule{Contributors: 0}
	m2 := MockModule{Contributors: 1}
	busfactorMetric = BusFactorMetric{}

	assertEquals(t, "", busfactorMetric.CalculateScore(m1), 0.0)
	assertEquals(t, "", busfactorMetric.CalculateScore(m2), 0.0)
}
func TestBusFactorMoreThanOneContributor(t *testing.T) {
	m1 := MockModule{Contributors: 2}
	m2 := MockModule{Contributors: 10}
	m3 := MockModule{Contributors: 500}
	m4 := MockModule{Contributors: 1750}
	busfactorMetric = BusFactorMetric{}

	assertEquals(t, "", busfactorMetric.CalculateScore(m1), 0.5)
	assertEquals(t, "", busfactorMetric.CalculateScore(m2), 0.9)
	assertEquals(t, "", busfactorMetric.CalculateScore(m3), 1.0-1.0/500)
	assertEquals(t, "", busfactorMetric.CalculateScore(m4), 1.0-1.0/1750)
}
func assertEquals(t *testing.T, desc string, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("%v Got: %v (%T), Want:%v (%T)", desc, got, got, want, want)
	}
}
