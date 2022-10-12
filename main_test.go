package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeDayAgo_Nom_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 15)
	assert.Equal(t, newTime("2022-09-26T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Min_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 15)
	assert.Equal(t, newTime("2022-09-25T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Max_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 15)
	assert.Equal(t, newTime("2022-09-26T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MaxMinus_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 15)
	assert.Equal(t, newTime("2022-09-26T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MinPlus_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 15)
	assert.Equal(t, newTime("2022-09-26T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Nom_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 1)
	assert.Equal(t, newTime("2022-10-10T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Min_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 1)
	assert.Equal(t, newTime("2022-10-09T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Max_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 1)
	assert.Equal(t, newTime("2022-10-10T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MaxMinus_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 1)
	assert.Equal(t, newTime("2022-10-10T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MinPlus_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 1)
	assert.Equal(t, newTime("2022-10-10T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Nom_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 30)
	assert.Equal(t, newTime("2022-09-11T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Min_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 30)
	assert.Equal(t, newTime("2022-09-10T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Max_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 30)
	assert.Equal(t, newTime("2022-09-11T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MaxMinus_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 30)
	assert.Equal(t, newTime("2022-09-11T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MinPlus_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 30)
	assert.Equal(t, newTime("2022-09-11T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Nom_Date_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 0)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_Min_Date_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 0)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_Max_Date_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 0)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_MaxMinus_Date_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 0)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_MinPlus_Date_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 0)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_Nom_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 29)
	assert.Equal(t, newTime("2022-09-12T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Min_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 29)
	assert.Equal(t, newTime("2022-09-11T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Max_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 29)
	assert.Equal(t, newTime("2022-09-12T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MaxMinus_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 29)
	assert.Equal(t, newTime("2022-09-12T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MinPlus_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 29)
	assert.Equal(t, newTime("2022-09-12T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Nom_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 2)
	assert.Equal(t, newTime("2022-10-09T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Min_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 2)
	assert.Equal(t, newTime("2022-10-08T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Max_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 2)
	assert.Equal(t, newTime("2022-10-09T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MaxMinus_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 2)
	assert.Equal(t, newTime("2022-10-09T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_MinPlus_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 2)
	assert.Equal(t, newTime("2022-10-09T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_Nom_Date_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T12:00:00+07:00"), 31)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_Min_Date_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:00+07:00"), 31)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_Max_Date_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:59+07:00"), 31)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_MaxMinus_Date_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T23:59:58+07:00"), 31)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_MinPlus_Date_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T00:00:01+07:00"), 31)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_RandomIndBound_Date_RandomInBound_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T15:33:22+07:00"), 23)
	assert.Equal(t, newTime("2022-09-18T00:00:00+07:00").String(), result)
}

func TestComputeDayAgo_OutOfBound_Date_RandomInBound_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:33:40+07:00"), 13)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_OutOfBound_Date_OutOfBound_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:22:21+07:00"), 49)
	assert.Equal(t, "Error Limit Input", result)
}

func TestComputeDayAgo_MaxPlus_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 15)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:00+07:00"), 1)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 30)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_MinMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 0)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 29)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_MaxPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 31)
	assert.Equal(t, "Invalid test case", result)
}

func TestComputeDayAgo_MaxPlus_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-10T24:00:01+07:00"), 2)
	assert.Equal(t, "Invalid test case", result)
}
