package rdb

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Do somethings before all test cases
func BeforeTest() {

}

func TestJsonEncoding(t *testing.T) {
	expected := []byte(`{"ID":{"$like":"ABC%"},"status":1}`)
	s := QueryAppPrefixPayload{
		ID: LikeCriterion{
			Like: "ABC%",
		},
		Status: 1,
	}
	result, _ := json.Marshal(s)
	assert.Equal(t, expected, result)
}
func TestGetStage(t *testing.T) {
	s := Application{
		StageFlag: 1,
	}
	assert.Equal(t, "Sandbox", s.GetStage())
	s.StageFlag = 2
	assert.Equal(t, "Production", s.GetStage())
}
func TestGetStatus(t *testing.T) {
	s := Application{
		Status: 1,
	}
	assert.Equal(t, "Valid", s.GetStatus())
	s.Status = 2
	assert.Equal(t, "Deleted", s.GetStatus())
}

// Do somethings after all test cases
func AfterTest() {

}

func TestMain(m *testing.M) {
	BeforeTest()
	retCode := m.Run()
	AfterTest()
	os.Exit(retCode)
}
