package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func usersDataSample(t *testing.T) []User {
	t.Helper()
	return []User{
		{Name: "Alice", Age: 18},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Lays", Age: 32},
		{Name: "Diana", Age: 42},
	}
}

func Test_LoadUserFromDataSuccess(t *testing.T) {
	data := `[{"name": "Lays", "age": 30}, {"name": "Alice", "age": 19}, {"name": "Bob", "age": 35}, {"name": "Charlie", "age": 40}, {"name": "Diana", "age": 20}]`
	users, err := loadUsersFromData(data)
	assert.NoError(t, err)
	assert.Len(t, users, 5)
	assert.Equal(t, users, []User{
		{Name: "Lays", Age: 30},
		{Name: "Alice", Age: 19},
		{Name: "Bob", Age: 35},
		{Name: "Charlie", Age: 40},
		{Name: "Diana", Age: 20},
	})
}

func Test_LoadUserFromDataFailure(t *testing.T) {
	data := `[{"name": "Lays", "age": "30"}, {"name": "Alice", "age": 19}, {"name": "Bob", "age": 35}, {"name": "Charlie", "age": 40}, {"name": "Diana", "age": 20}]`
	users, err := loadUsersFromData(data)
	assert.Error(t, err)
	assert.Equal(t, "failed to parse input data: json: cannot unmarshal string into Go struct field User.age of type int", err.Error())
	assert.Nil(t, users)
}

func Test_FilterUsersAboveAge(t *testing.T) {
	users := usersDataSample(t)
	testCases := []struct {
		age              int
		expectedResponse int
	}{
		{15, 5},
		{20, 4},
		{25, 3},
		{30, 3},
		{35, 1},
		{40, 1},
	}
	for index, tc := range testCases {
		name := fmt.Sprintf("Test Case #%d", index)
		t.Run(name, func(t *testing.T) {
			filteredUsers := filterUsersAboveAge(users, tc.age)
			assert.Len(t, filteredUsers, tc.expectedResponse, "expected %d users to be above age %d", tc.expectedResponse, tc.age)
		})
	}
}

func Test_SortUsersByName(t *testing.T) {
	users := usersDataSample(t)
	sortedUsers := sortUsersByName(users)
	assert.Equal(t, []User{
		{Name: "Alice", Age: 18},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 42},
		{Name: "Lays", Age: 32},
	}, sortedUsers)
}

func Test_GroupUsersByAge(t *testing.T) {
	testCases := []struct {
		name     string
		input    []User
		expected map[string][]User
	}{
		{
			name:     "It should return an empty map with an empty input",
			input:    []User{},
			expected: map[string][]User{},
		},
		{
			name: "It should have at least one user per group.",
			input: []User{
				{Name: "Alice", Age: 18},
				{Name: "Bob", Age: 25},
				{Name: "Charlie", Age: 35},
				{Name: "Diana", Age: 42},
				{Name: "Lays", Age: 32},
			},
			expected: map[string][]User{
				"<20":   {{Name: "Alice", Age: 18}},
				"20-29": {{Name: "Bob", Age: 25}},
				"30-39": {{Name: "Charlie", Age: 35}, {Name: "Lays", Age: 32}},
				"40+":   {{Name: "Diana", Age: 42}},
			},
		},
		{
			name: "It should put all users in one group",
			input: []User{
				{Name: "Charlie", Age: 45},
				{Name: "Diana", Age: 50},
			},
			expected: map[string][]User{
				"40+": {{Name: "Charlie", Age: 45}, {Name: "Diana", Age: 50}},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := groupUsersByAge(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_UpdateUsersAge(t *testing.T) {
	users := usersDataSample(t)
	amountToAdd := 1
	updatedUsers := updateUsersAge(users, amountToAdd)
	for _, updated := range updatedUsers {
		for _, original := range users {
			if original.Name == updated.Name {
				updatedAge := original.Age + amountToAdd
				assert.Equal(t, updatedAge, updated.Age)
			}
		}
	}
}

func Test_CountUsersAboveAge(t *testing.T) {
	users := usersDataSample(t)
	testCases := []struct {
		age           int
		expectedCount int
	}{
		{15, 5},
		{20, 4},
		{25, 3},
		{30, 3},
		{35, 1},
		{40, 1},
	}

	for index, tc := range testCases {
		name := fmt.Sprintf("Test Case #%d", index)
		t.Run(name, func(t *testing.T) {
			counter := countUsersAboveAge(users, tc.age)
			assert.Equal(t, tc.expectedCount, counter, "expected %d users above %d age", tc.expectedCount, tc.age)
		})
	}
}
