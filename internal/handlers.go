package internal

import (
	"blog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gomarkdown/markdown"
	"gorm.io/gorm"
	"html/template"
	"log"
	"math"
	"strconv"
	"time"
)

func Route(app *fiber.App, db *gorm.DB) {
	app.Use(RateLimiter())
	app.Get("/", Index(db))
	app.Get("/post/:id", GetPost(db))
	app.Get("/sign_in", RenderPage("sign_in"))
	app.Get("/admin", RenderPage("admin"))
	app.Post("/check_login", CheckLogin())
	app.Post("/save_post", CreatePost(db))
	app.Use(NotFound())

}

func Index(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var page float64
		q, _ := db.Table("posts").Select("COUNT(*)").Rows()
		for q.Next() {
			err := q.Scan(&page)
			if err != nil {
				log.Fatalf("error occured while querying: %v", err)
			}
		}
		page = math.Ceil(page / 3)
		var pages []int
		for i := 1; i <= int(page); i++ {
			pages = append(pages, i)
		}
		num := 1
		temp := c.Query("page")
		if temp != "" {
			num, _ = strconv.Atoi(string(temp[0]))
		}
		var posts []database.Post
		q, _ = db.Table("posts").Limit(3).Offset(num*3 - 3).Order("id").Rows()
		for q.Next() {
			var post database.Post
			err := q.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
			if err != nil {
				log.Fatalf("error occured while querying: %v", err)
			}
			posts = append(posts, post)
		}
		return c.Render("index", fiber.Map{
			"Posts": posts,
			"Pages": pages,
		})
	}
}

func GetPost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post database.Post
		id := c.Params("id")
		q, _ := db.Table("posts").Where("id = ?", id).Rows()
		for q.Next() {
			err := q.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
			if err != nil {
				log.Fatalf("error occured while querying: %v", err)
			}
		}
		post.Text = template.HTML(markdown.ToHTML([]byte(post.Text), nil, nil))
		return c.Render("show", post)
	}
}

func RenderPage(path string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render(path, fiber.Map{})
	}
}

func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
	}
}

func CreatePost(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		announce := c.FormValue("announce")
		text := c.FormValue("full_text")
		if title == "" || announce == "" || text == "" {
			return c.SendString("error: all fields must be filled")
		}
		db.Create(&database.Post{
			Title:    title,
			Announce: announce,
			Text:     template.HTML(text),
		})
		log.Printf("Post created at: %v", time.Now())
		return c.Redirect("/", fiber.StatusSeeOther)
	}
}

func CheckLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.FormValue("login") == "admin" && c.FormValue("password") == "admin" {
			return c.Redirect("/admin", fiber.StatusSeeOther)
		} else {
			return c.Redirect("/sign_in", fiber.StatusSeeOther)
		}
	}
}

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: 10 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests!")
		},
	})
}
