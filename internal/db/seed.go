package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/leonardoaraujodf/social/internal/store"
)

var usernames = []string{
	"alice",
	"bob",
	"charlie",
	"dave",
	"eve",
	"frank",
	"grace",
	"heidi",
	"ivan",
	"judy",
	"mallory",
	"oscar",
	"peggy",
	"trent",
	"victor",
	"walter",
	"xavier",
	"yvonne",
	"zara",
	"quinn",
	"rachel",
	"steve",
	"tina",
	"ursula",
	"viktor",
	"wendy",
	"xander",
	"yara",
	"zane",
}

var titles = []string{
	"Hello World",
	"My First Post",
	"Go Programming",
	"Database Design",
	"Web Development",
	"Microservices Architecture",
	"Cloud Computing",
	"DevOps Best Practices",
	"Testing in Go",
	"Concurrency Patterns",
	"Error Handling in Go",
	"Building REST APIs",
	"GraphQL vs REST",
	"Containerization with Docker",
	"Kubernetes for Beginners",
	"CI/CD Pipelines",
	"Monitoring and Logging",
	"Security in Web Applications",
	"Performance Optimization",
	"Go Modules and Dependency Management",
}

var contents = []string{
	"This is the content of the post.",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	"Go is an open-source programming language designed for simplicity and efficiency.",
	"Database design is crucial for building scalable applications.",
	"Web development involves creating websites and web applications.",
	"Microservices architecture allows for building applications as a collection of small services.",
	"Cloud computing provides on-demand access to computing resources over the internet.",
	"DevOps best practices help teams deliver software faster and more reliably.",
	"Testing in Go can be done using the built-in testing package.",
	"Concurrency patterns in Go include goroutines and channels.",
	"Error handling in Go is done using error values.",
	"Building REST APIs involves designing endpoints and handling HTTP requests.",
	"GraphQL vs REST is a common debate in API design.",
	"Containerization with Docker allows for packaging applications with their dependencies.",
	"Kubernetes for beginners provides an introduction to container orchestration.",
	"CI/CD pipelines automate the process of building, testing, and deploying software.",
	"Monitoring and logging are essential for maintaining application health.",
	"Security in web applications involves protecting against common vulnerabilities.",
	"Performance optimization can improve the speed and responsiveness of applications.",
	"Go modules and dependency management help manage project dependencies effectively.",
}

var tags = []string{
	"programming",
	"go",
	"webdev",
	"database",
	"cloud",
	"devops",
	"testing",
	"concurrency",
	"error-handling",
	"rest-api",
	"graphql",
	"docker",
	"kubernetes",
	"ci-cd",
	"monitoring",
	"security",
	"performance",
}

var comments = []string{
	"Great post!",
	"Thanks for sharing.",
	"Very informative.",
	"I learned a lot from this.",
	"Can you provide more examples?",
	"I have a question about this topic.",
	"This is exactly what I was looking for.",
	"Awesome content!",
	"Keep up the good work.",
	"I disagree with some points, but overall good read.",
	"Can you explain this part in more detail?",
	"This is a must-read for anyone interested in this topic.",
	"Thanks for the insights!",
	"I have a similar experience to share.",
	"Looking forward to more posts like this.",
	"This post has sparked some interesting discussions.",
	"Can you recommend any resources for further reading?",
	"I appreciate the depth of information provided here.",
	"This is a great starting point for beginners.",
	"Your writing style makes complex topics easy to understand.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, posts, users)
	for _, comment := range comments {
		if err := store.Comment.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + "_" + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + "_" + fmt.Sprintf("%d", i) + "@example.com",
			Password: "12345",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	commentsList := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]
		commentsList[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return commentsList
}
