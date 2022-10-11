package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeDayAgo_Min_Date(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-10T00:00:00+07:00"), 7)
	assert.Equal(t, newTime("2022-08-03T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Max_Date(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T23:59:59+07:00"), 7)
	assert.Equal(t, newTime("2022-07-30T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MinPlus_Date(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-10T00:00:01+07:00"), 7)
	assert.Equal(t, newTime("2022-08-04T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MaxMinus_Date(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T23:59:58+07:00"), 7)
	assert.Equal(t, newTime("2022-07-30T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Date(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T10:00:00+07:00"), 7)
	assert.Equal(t, newTime("2022-07-30T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Min_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T12:00:00+07:00"), 1)
	assert.Equal(t, newTime("2022-08-05T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Max_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T10:00:00+07:00"), 30)
	assert.Equal(t, newTime("2022-07-07T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MinPlus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T10:00:00+07:00"), 2)
	assert.Equal(t, newTime("2022-08-04T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_MaxMinus_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T10:00:00+07:00"), 29)
	assert.Equal(t, newTime("2022-07-08T00:00:00+07:00").Unix(), result)
}

func TestComputeDayAgo_Nom_Limit(t *testing.T) {
	result := computeDayAgo(newTime("2022-08-05T10:00:00+07:00"), 15)
	assert.Equal(t, newTime("2022-07-22T00:00:00+07:00").Unix(), result)
}
