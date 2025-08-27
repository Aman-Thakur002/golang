/*
=============================================================================
                        🌐 GO HTTP CLIENT TUTORIAL
=============================================================================

📚 CORE CONCEPT:
HTTP clients in Go allow you to make requests to web servers and APIs.
The net/http package provides powerful tools for building robust HTTP clients.

🔑 KEY FEATURES:
• Simple HTTP requests (GET, POST, PUT, DELETE)
• Custom headers and authentication
• Request/response handling
• Timeouts and retries
• JSON API communication

💡 REAL-WORLD ANALOGY:
HTTP Client = Postal Service
- Request = Sending a letter with specific instructions
- Response = Getting a reply back
- Headers = Special delivery instructions
- Timeout = Maximum waiting time for reply

🎯 WHY LEARN HTTP CLIENTS?
• API integration and consumption
• Web scraping and data collection
• Microservice communication
• Third-party service integration

=============================================================================
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 📊 DATA STRUCTURES: For API communication
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	fmt.Println("🌐 HTTP CLIENT TUTORIAL")
	fmt.Println("=======================")

	// 🎯 DEMO 1: Simple GET Request
	fmt.Println("\n🎯 DEMO 1: Simple GET Request")
	fmt.Println("=============================")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("❌ GET error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Status: %s\n", resp.Status)
	fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read body error: %v\n", err)
		return
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("📝 Post: %+v\n", post)

	// 🎯 DEMO 2: GET with Query Parameters
	fmt.Println("\n🎯 DEMO 2: GET with Query Parameters")
	fmt.Println("====================================")

	baseURL := "https://jsonplaceholder.typicode.com/posts"
	params := url.Values{}
	params.Add("userId", "1")
	params.Add("_limit", "3")

	fullURL := baseURL + "?" + params.Encode()
	fmt.Printf("🔗 URL: %s\n", fullURL)

	resp, err = http.Get(fullURL)
	if err != nil {
		fmt.Printf("❌ GET with params error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read body error: %v\n", err)
		return
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("📝 Found %d posts:\n", len(posts))
	for _, p := range posts {
		fmt.Printf("  - %s\n", p.Title)
	}

	// 🎯 DEMO 3: POST Request with JSON
	fmt.Println("\n🎯 DEMO 3: POST Request with JSON")
	fmt.Println("=================================")

	newPost := Post{
		UserID: 1,
		Title:  "My New Post",
		Body:   "This is the content of my new post.",
	}

	jsonData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Printf("❌ JSON marshal error: %v\n", err)
		return
	}

	resp, err = http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Printf("❌ POST error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 POST Status: %s\n", resp.Status)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read POST response error: %v\n", err)
		return
	}

	var createdPost Post
	err = json.Unmarshal(body, &createdPost)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("✅ Created post: %+v\n", createdPost)

	// 🎯 DEMO 4: Custom HTTP Client with Timeout
	fmt.Println("\n🎯 DEMO 4: Custom HTTP Client")
	fmt.Println("=============================")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	if err != nil {
		fmt.Printf("❌ Create request error: %v\n", err)
		return
	}

	// Add custom headers
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("❌ Custom client error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Response Headers:\n")
	for key, values := range resp.Header {
		fmt.Printf("  %s: %s\n", key, strings.Join(values, ", "))
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read body error: %v\n", err)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("👤 User: %+v\n", user)

	// 🎯 DEMO 5: PUT Request
	fmt.Println("\n🎯 DEMO 5: PUT Request")
	fmt.Println("======================")

	updatedPost := Post{
		UserID: 1,
		ID:     1,
		Title:  "Updated Post Title",
		Body:   "This post has been updated.",
	}

	jsonData, err = json.Marshal(updatedPost)
	if err != nil {
		fmt.Printf("❌ JSON marshal error: %v\n", err)
		return
	}

	req, err = http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Create PUT request error: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("❌ PUT error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 PUT Status: %s\n", resp.Status)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read PUT response error: %v\n", err)
		return
	}

	var putResponse Post
	err = json.Unmarshal(body, &putResponse)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("✅ Updated post: %+v\n", putResponse)

	// 🎯 DEMO 6: DELETE Request
	fmt.Println("\n🎯 DEMO 6: DELETE Request")
	fmt.Println("=========================")

	req, err = http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Printf("❌ Create DELETE request error: %v\n", err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("❌ DELETE error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 DELETE Status: %s\n", resp.Status)
	fmt.Printf("✅ Post deleted (simulated)\n")

	// 🎯 DEMO 7: Form Data POST
	fmt.Println("\n🎯 DEMO 7: Form Data POST")
	fmt.Println("=========================")

	formData := url.Values{}
	formData.Set("title", "Form Post Title")
	formData.Set("body", "This is a form post")
	formData.Set("userId", "1")

	resp, err = http.PostForm("https://jsonplaceholder.typicode.com/posts", formData)
	if err != nil {
		fmt.Printf("❌ Form POST error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Form POST Status: %s\n", resp.Status)

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Read form response error: %v\n", err)
		return
	}

	var formPost Post
	err = json.Unmarshal(body, &formPost)
	if err != nil {
		fmt.Printf("❌ JSON unmarshal error: %v\n", err)
		return
	}

	fmt.Printf("✅ Form post created: %+v\n", formPost)

	// 🎯 DEMO 8: Error Handling and Status Codes
	fmt.Println("\n🎯 DEMO 8: Error Handling")
	fmt.Println("=========================")

	resp, err = http.Get("https://jsonplaceholder.typicode.com/posts/999999")
	if err != nil {
		fmt.Printf("❌ Request error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("✅ Success!")
	case http.StatusNotFound:
		fmt.Println("❌ Resource not found")
	case http.StatusInternalServerError:
		fmt.Println("❌ Server error")
	default:
		fmt.Printf("⚠️ Unexpected status: %s\n", resp.Status)
	}

	// 🎯 DEMO 9: Authentication Example
	fmt.Println("\n🎯 DEMO 9: Authentication")
	fmt.Println("=========================")

	req, err = http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Printf("❌ Create auth request error: %v\n", err)
		return
	}

	// Basic Auth (example)
	req.SetBasicAuth("username", "password")

	// Bearer Token (example)
	req.Header.Set("Authorization", "Bearer your-token-here")

	// API Key (example)
	req.Header.Set("X-API-Key", "your-api-key-here")

	fmt.Println("🔐 Authentication headers set (examples)")
	fmt.Printf("  Basic Auth: %s\n", req.Header.Get("Authorization"))
	fmt.Printf("  API Key: %s\n", req.Header.Get("X-API-Key"))

	fmt.Println("\n✨ All HTTP client demos completed!")
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

🌐 HTTP METHODS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│    Method       │                Purpose                                  │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ GET             │ Retrieve data from server                               │
│ POST            │ Create new resource                                     │
│ PUT             │ Update entire resource                                  │
│ PATCH           │ Partial update of resource                              │
│ DELETE          │ Remove resource                                         │
│ HEAD            │ Get headers only (no body)                              │
│ OPTIONS         │ Get allowed methods                                     │
└─────────────────┴─────────────────────────────────────────────────────────┘

🔧 HTTP CLIENT PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Simple GET                                                           │
│ resp, err := http.Get("https://api.example.com/data")                    │
│                                                                         │
│ // Custom request                                                       │
│ req, err := http.NewRequest("POST", url, body)                          │
│ req.Header.Set("Content-Type", "application/json")                      │
│ resp, err := client.Do(req)                                             │
│                                                                         │
│ // With timeout                                                         │
│ client := &http.Client{Timeout: 30 * time.Second}                       │
│ resp, err := client.Do(req)                                             │
└─────────────────────────────────────────────────────────────────────────┘

📊 STATUS CODES:
┌─────────────────┬─────────────────┬─────────────────────────────────────┐
│   Code Range    │   Category      │           Meaning                   │
├─────────────────┼─────────────────┼─────────────────────────────────────┤
│ 200-299         │ Success         │ Request successful                  │
│ 300-399         │ Redirection     │ Further action needed               │
│ 400-499         │ Client Error    │ Bad request from client             │
│ 500-599         │ Server Error    │ Server failed to fulfill request    │
└─────────────────┴─────────────────┴─────────────────────────────────────┘

🔐 AUTHENTICATION METHODS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Basic Auth                                                           │
│ req.SetBasicAuth("username", "password")                                 │
│                                                                         │
│ // Bearer Token                                                         │
│ req.Header.Set("Authorization", "Bearer " + token)                       │
│                                                                         │
│ // API Key                                                              │
│ req.Header.Set("X-API-Key", apiKey)                                     │
│                                                                         │
│ // Custom Header                                                        │
│ req.Header.Set("X-Custom-Auth", authValue)                              │
└─────────────────────────────────────────────────────────────────────────┘

📝 REQUEST BODY FORMATS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // JSON                                                                 │
│ jsonData, _ := json.Marshal(data)                                       │
│ req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))        │
│ req.Header.Set("Content-Type", "application/json")                      │
│                                                                         │
│ // Form Data                                                            │
│ formData := url.Values{}                                                │
│ formData.Set("key", "value")                                            │
│ resp, _ := http.PostForm(url, formData)                                  │
│                                                                         │
│ // Raw String                                                           │
│ body := strings.NewReader("raw data")                                   │
│ req, _ := http.NewRequest("POST", url, body)                             │
└─────────────────────────────────────────────────────────────────────────┘

🚨 ERROR HANDLING:
┌─────────────────────────────────────────────────────────────────────────┐
│ resp, err := http.Get(url)                                              │
│ if err != nil {                                                         │
│     // Network error, timeout, etc.                                    │
│     return err                                                          │
│ }                                                                       │
│ defer resp.Body.Close()                                                 │
│                                                                         │
│ if resp.StatusCode != http.StatusOK {                                   │
│     // HTTP error status                                               │
│     return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)      │
│ }                                                                       │
│                                                                         │
│ body, err := io.ReadAll(resp.Body)                                      │
│ if err != nil {                                                         │
│     // Error reading response body                                      │
│     return err                                                          │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

💡 BEST PRACTICES:
• Always close response bodies with defer
• Set appropriate timeouts
• Handle both network and HTTP errors
• Use context for cancellation
• Reuse HTTP clients when possible
• Set proper Content-Type headers
• Validate response status codes

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // API Client struct                                                    │
│ type APIClient struct {                                                 │
│     BaseURL    string                                                   │
│     HTTPClient *http.Client                                             │
│     APIKey     string                                                   │
│ }                                                                       │
│                                                                         │
│ func (c *APIClient) Get(endpoint string) (*Response, error) {           │
│     req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)          │
│     if err != nil {                                                     │
│         return nil, err                                                 │
│     }                                                                   │
│     req.Header.Set("Authorization", "Bearer "+c.APIKey)                 │
│     return c.HTTPClient.Do(req)                                         │
│ }                                                                       │
│                                                                         │
│ // Retry logic                                                          │
│ func (c *APIClient) GetWithRetry(url string, maxRetries int) {          │
│     for i := 0; i < maxRetries; i++ {                                   │
│         resp, err := c.HTTPClient.Get(url)                              │
│         if err == nil && resp.StatusCode == 200 {                       │
│             return resp, nil                                            │
│         }                                                               │
│         time.Sleep(time.Duration(i+1) * time.Second)                    │
│     }                                                                   │
│     return nil, errors.New("max retries exceeded")                      │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE TIPS:
• Reuse http.Client instances
• Use connection pooling
• Set appropriate timeouts
• Use streaming for large responses
• Consider using context for cancellation
• Pool request/response objects when possible

=============================================================================
*/