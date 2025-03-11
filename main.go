<<<<<<< HEAD
package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
=======
package main

import (
	"fmt"
	"os"
)

// Struktur data sederhana untuk menyimpan item
type Item struct {
	ID    int
	Title string
	Body  string
}

// Slice untuk menyimpan data secara sementara
var items []Item
var nextID int = 1

func main() {
	for {
		fmt.Println("\nCRUD Menu:")
		fmt.Println("1. Create Item")
		fmt.Println("2. Read Items")
		fmt.Println("3. Update Item")
		fmt.Println("4. Delete Item")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createItem()
		case 2:
			readItems()
		case 3:
			updateItem()
		case 4:
			deleteItem()
		case 5:
			fmt.Println("Exiting program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// Create: Tambahkan item baru
func createItem() {
	var title, body string
	fmt.Print("Enter title: ")
	fmt.Scan(&title)
	fmt.Print("Enter body: ")
	fmt.Scan(&body)

	item := Item{
		ID:    nextID,
		Title: title,
		Body:  body,
	}
	nextID++
	items = append(items, item)

	fmt.Println("Item created successfully!")
}

// Read: Tampilkan semua item
func readItems() {
	if len(items) == 0 {
		fmt.Println("No items found.")
		return
	}

	fmt.Println("\nList of Items:")
	for _, item := range items {
		fmt.Printf("ID: %d, Title: %s, Body: %s\n", item.ID, item.Title, item.Body)
	}
}

// Update: Perbarui item berdasarkan ID
func updateItem() {
	var id int
	fmt.Print("Enter ID of the item to update: ")
	fmt.Scan(&id)

	for i, item := range items {
		if item.ID == id {
			var title, body string
			fmt.Print("Enter new title: ")
			fmt.Scan(&title)
			fmt.Print("Enter new body: ")
			fmt.Scan(&body)

			items[i].Title = title
			items[i].Body = body
			fmt.Println("Item updated successfully!")
			return
		}
	}
	fmt.Println("Item not found.")
}

// Delete: Hapus item berdasarkan ID
func deleteItem() {
	var id int
	fmt.Print("Enter ID of the item to delete: ")
	fmt.Scan(&id)

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			fmt.Println("Item deleted successfully!")
			return
		}
	}
	fmt.Println("Item not found.")
}
>>>>>>> 98d8fbf7fc1d2d2f0f2d2b4c568518bd0c896f61
