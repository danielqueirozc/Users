package main;

import (
	"fmt"
)

type User struct {
	ID    int;
	Name  string;
	Email string;
}

var users []User;
var nextID = 1;

// cria e adiciona um novo usuario na lista (slice)
func createUser(name, email string) User {
	user := User {
		ID:    nextID,
		Name: name,
		Email: email,
	}
	 users = append(users, user);
	 nextID++;

	 return user;
}

// retorna todos os usuarios
func getAllUsers() []User {
	return users;
}

func getUserById(id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true;
		}
	}

	return User{}, false;
}

func updateUser(id int, name, email string) bool {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = name;
			users[i].Email = email;

			return true;
		}
	}

	return false;
}

func deleteUser(id int) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...);
			return true;
		}
	}

	return false;	
}

func main() {
	user1 := createUser("Daniel", "danielqueiroz.job@gmia.com");

	fmt.Println(getAllUsers());

	user, found := getUserById(user1.ID);

	if found {
			 fmt.Println("user found", user) 
		} else {
			fmt.Println("user not found")
		}

		updated := updateUser(user1.ID, "slava", "russo@gmail.com");

		if updated {
			userUpdated, _ := getUserById(user1.ID);

			fmt.Println("user updated", userUpdated);
		}

		deleted := deleteUser(user1.ID);

		if deleted {
			userDeleted, _ := getUserById(user1.ID);
			fmt.Println("user deleted", userDeleted);
			fmt.Println("user deleted");
		}
}