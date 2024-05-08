package main

import (
	"fmt"
	"testing"

	"github.com/go-rod/rod"
)

// This example opens https://github.com/, searches for "git",
// and then gets the header element which gives the description for Git.
func Test_Main(t *testing.T) {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://github.com/login").MustWaitStable()

	page.MustScreenshot("./screenshots/github.png")

	page.MustElement("#login_field").MustInput("sumitpantheon")

	page.MustElement("#password").MustInput("test")
	// page.MustElement("#password").MustInput("test").MustType(input.Enter)

	page.MustElement("input[type=submit]").MustClick()

	page.MustWaitRequestIdle()

	page.MustScreenshot("./screenshots/github_login.png")

	text := page.MustElementR("div", "Incorrect username or password").MustText()

	fmt.Println(text)

	page.MustScreenshot("./screenshots/github_login_failed.png")
}
