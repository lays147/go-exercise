package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := `[{"name": "Lays", "age": 30}, {"name": "Alice", "age": 19}, {"name": "Bob", "age": 35}, {"name": "Charlie", "age": 40}, {"name": "Diana", "age": 20}]`
	users, err := loadUsersFromData(data)

	if err != nil {
		fmt.Println("%w", err)
		panic("failed to parse data")
	}

	filteredUsers := filterUsersAboveAge(users, 30)
	fmt.Println("Filtered Users older than 30:", filteredUsers)

	sortedUsers := sortUsersByName(users)
	fmt.Println("Users sorted by name:", sortedUsers)

	groupedUsers := groupUsersByAge(users)
	fmt.Println("Users grouped by age:", groupedUsers)

	updatedUser := updateUsersAge(users, 1)
	fmt.Println("Users after incrementing age by 1:", updatedUser)

	count := countUsersAboveAge(users, 30)
	fmt.Println("Count of users older than 30:", count)
}

func loadUsersFromData(data string) ([]User, error) {
	var users []User
	err := json.Unmarshal([]byte(data), &users)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input data: %w", err)
	}
	return users, nil
}

func filterUsersAboveAge(users []User, age int) []User {
	var filtered []User
	for _, user := range users {
		if user.Age > age {
			filtered = append(filtered, user)
		}
	}
	return filtered
}

func sortUsersByName(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	return users
}

func groupUsersByAge(users []User) map[string][]User {
	grouped := make(map[string][]User)
	for _, user := range users {
		if user.Age < 20 {
			grouped["<20"] = append(grouped["<20"], user)
		} else if user.Age >= 20 && user.Age < 30 {
			grouped["20-29"] = append(grouped["20-29"], user)
		} else if user.Age >= 30 && user.Age < 40 {
			grouped["30-39"] = append(grouped["30-39"], user)
		} else {
			grouped["40+"] = append(grouped["40+"], user)
		}
	}
	return grouped
}

func updateUsersAge(users []User, increment int) []User {
	// Golang treats slices with pointers instead of copying values
	// Thanks chatgpt =)
	newUsers := make([]User, len(users))
	for i := range newUsers {
		newUsers[i].Age += increment
	}
	return newUsers
}

func countUsersAboveAge(users []User, age int) int {
	count := 0
	for i := range users {
		if users[i].Age > age {
			count++
		}
	}
	return count
}
