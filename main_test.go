package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeDayAgo_Nom_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-11T12:00:00+07:00"), 7)
	assert.Equal(t, newTime("2022-10-05T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Date_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-11T12:00:00+07:00"), 1)
	assert.Equal(t, newTime("2022-10-11T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Date_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-11T12:00:00+07:00"), 2)
	assert.Equal(t, newTime("2022-10-10T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Date_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-11T12:00:00+07:00"), 29)
	assert.Equal(t, newTime("2022-09-13T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Date_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-11T12:00:00+07:00"), 30)
	assert.Equal(t, newTime("2022-09-12T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Min_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T00:00:00+07:00"), 15)
	assert.Equal(t, newTime("2022-09-30T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MinPlus_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T00:00:01+07:00"), 15)
	assert.Equal(t, newTime("2022-10-01T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Max_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T23:59:59+07:00"), 15)
	assert.Equal(t, newTime("2022-10-01T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MaxMinus_Date_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T23:59:58+07:00"), 15)
	assert.Equal(t, newTime("2022-10-01T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Limit_Overbound(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T23:59:58+07:00"), 31)
	assert.Equal(t, int64(0), result)
}

func TestComputeDayAgo_Limit_Underbound(t *testing.T) {
	result := computeDayAgo(newTime("2022-10-15T23:59:58+07:00"), 0)
	assert.Equal(t, int64(0), result)
}
