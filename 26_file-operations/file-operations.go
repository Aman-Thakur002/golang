/*
=============================================================================
                        📁 GO FILE OPERATIONS TUTORIAL
=============================================================================

📚 CORE CONCEPT:
File operations in Go are handled through the os and io packages. Go provides
both simple and advanced ways to read, write, and manipulate files.

🔑 KEY FEATURES:
• Simple file read/write functions
• Streaming operations for large files
• File metadata and permissions
• Directory operations
• Cross-platform file handling

💡 REAL-WORLD ANALOGY:
File Operations = Library Management
- Reading = Checking out books
- Writing = Adding new books
- Appending = Adding notes to existing books
- Directory = Organizing books by category

🎯 WHY LEARN FILE OPERATIONS?
• Configuration file handling
• Log file management
• Data processing and ETL
• Backup and archival systems

=============================================================================
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("📁 FILE OPERATIONS TUTORIAL")
	fmt.Println("============================")

	// 🎯 DEMO 1: Basic File Writing
	fmt.Println("\n🎯 DEMO 1: Basic File Writing")
	fmt.Println("=============================")

	// Simple write to file
	filename := "demo.txt"
	content := "Hello, Go file operations!\nThis is line 2.\nThis is line 3."

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("❌ Write error: %v\n", err)
		return
	}
	fmt.Printf("✅ Written to %s\n", filename)

	// 🎯 DEMO 2: Basic File Reading
	fmt.Println("\n🎯 DEMO 2: Basic File Reading")
	fmt.Println("=============================")

	// Simple read from file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("❌ Read error: %v\n", err)
		return
	}
	fmt.Printf("📖 File content:\n%s\n", string(data))

	// 🎯 DEMO 3: File Operations with os.File
	fmt.Println("\n🎯 DEMO 3: Advanced File Operations")
	fmt.Println("===================================")

	// Open file for writing
	file, err := os.Create("advanced.txt")
	if err != nil {
		fmt.Printf("❌ Create error: %v\n", err)
		return
	}
	defer file.Close() // Always close files!

	// Write to file
	_, err = file.WriteString("Line 1: Created with os.Create\n")
	if err != nil {
		fmt.Printf("❌ Write error: %v\n", err)
		return
	}

	// Write bytes
	_, err = file.Write([]byte("Line 2: Written as bytes\n"))
	if err != nil {
		fmt.Printf("❌ Write error: %v\n", err)
		return
	}

	// Sync to disk
	err = file.Sync()
	if err != nil {
		fmt.Printf("❌ Sync error: %v\n", err)
		return
	}

	fmt.Println("✅ Advanced file operations completed")

	// 🎯 DEMO 4: Reading File Line by Line
	fmt.Println("\n🎯 DEMO 4: Line-by-Line Reading")
	fmt.Println("===============================")

	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("❌ Open error: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("📝 Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Scanner error: %v\n", err)
		return
	}

	// 🎯 DEMO 5: Appending to Files
	fmt.Println("\n🎯 DEMO 5: Appending to Files")
	fmt.Println("=============================")

	file, err = os.OpenFile("append.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("❌ Open for append error: %v\n", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] Log entry appended\n", timestamp)
	
	_, err = file.WriteString(logEntry)
	if err != nil {
		fmt.Printf("❌ Append error: %v\n", err)
		return
	}

	fmt.Println("✅ Appended log entry")

	// Read and display the appended file
	appendData, err := os.ReadFile("append.txt")
	if err != nil {
		fmt.Printf("❌ Read append file error: %v\n", err)
		return
	}
	fmt.Printf("📖 Append file content:\n%s", string(appendData))

	// 🎯 DEMO 6: File Information and Metadata
	fmt.Println("\n🎯 DEMO 6: File Information")
	fmt.Println("===========================")

	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("❌ Stat error: %v\n", err)
		return
	}

	fmt.Printf("📊 File: %s\n", fileInfo.Name())
	fmt.Printf("📊 Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("📊 Mode: %v\n", fileInfo.Mode())
	fmt.Printf("📊 Modified: %v\n", fileInfo.ModTime())
	fmt.Printf("📊 Is Directory: %v\n", fileInfo.IsDir())

	// 🎯 DEMO 7: Directory Operations
	fmt.Println("\n🎯 DEMO 7: Directory Operations")
	fmt.Println("===============================")

	// Create directory
	dirName := "test_directory"
	err = os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("❌ Mkdir error: %v\n", err)
		return
	}
	fmt.Printf("📁 Created directory: %s\n", dirName)

	// Create nested directories
	nestedDir := filepath.Join(dirName, "nested", "deep")
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("❌ MkdirAll error: %v\n", err)
		return
	}
	fmt.Printf("📁 Created nested directory: %s\n", nestedDir)

	// Create file in directory
	nestedFile := filepath.Join(nestedDir, "nested_file.txt")
	err = os.WriteFile(nestedFile, []byte("File in nested directory"), 0644)
	if err != nil {
		fmt.Printf("❌ Write nested file error: %v\n", err)
		return
	}
	fmt.Printf("📄 Created nested file: %s\n", nestedFile)

	// List directory contents
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("❌ ReadDir error: %v\n", err)
		return
	}

	fmt.Printf("📋 Directory contents of %s:\n", dirName)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  📁 %s/\n", entry.Name())
		} else {
			fmt.Printf("  📄 %s\n", entry.Name())
		}
	}

	// 🎯 DEMO 8: Walking Directory Tree
	fmt.Println("\n🎯 DEMO 8: Walking Directory Tree")
	fmt.Println("=================================")

	err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		indent := strings.Repeat("  ", strings.Count(path, string(os.PathSeparator)))
		if info.IsDir() {
			fmt.Printf("%s📁 %s/\n", indent, info.Name())
		} else {
			fmt.Printf("%s📄 %s (%d bytes)\n", indent, info.Name(), info.Size())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("❌ Walk error: %v\n", err)
		return
	}

	// 🎯 DEMO 9: Copying Files
	fmt.Println("\n🎯 DEMO 9: Copying Files")
	fmt.Println("========================")

	sourceFile := filename
	destFile := "copied_" + filename

	err = copyFile(sourceFile, destFile)
	if err != nil {
		fmt.Printf("❌ Copy error: %v\n", err)
		return
	}
	fmt.Printf("✅ Copied %s to %s\n", sourceFile, destFile)

	// 🎯 DEMO 10: Working with Temporary Files
	fmt.Println("\n🎯 DEMO 10: Temporary Files")
	fmt.Println("===========================")

	// Create temporary file
	tempFile, err := os.CreateTemp("", "go_tutorial_*.txt")
	if err != nil {
		fmt.Printf("❌ CreateTemp error: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up
	defer tempFile.Close()

	fmt.Printf("🗂️ Temporary file: %s\n", tempFile.Name())

	// Write to temporary file
	_, err = tempFile.WriteString("This is temporary data")
	if err != nil {
		fmt.Printf("❌ Write temp error: %v\n", err)
		return
	}

	// Read from temporary file
	_, err = tempFile.Seek(0, 0) // Go back to beginning
	if err != nil {
		fmt.Printf("❌ Seek error: %v\n", err)
		return
	}

	tempData, err := io.ReadAll(tempFile)
	if err != nil {
		fmt.Printf("❌ Read temp error: %v\n", err)
		return
	}
	fmt.Printf("🗂️ Temp file content: %s\n", string(tempData))

	// 🧹 CLEANUP: Remove created files and directories
	fmt.Println("\n🧹 CLEANUP")
	fmt.Println("==========")

	filesToRemove := []string{filename, "advanced.txt", "append.txt", destFile}
	for _, file := range filesToRemove {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			fmt.Printf("⚠️ Failed to remove %s: %v\n", file, err)
		} else {
			fmt.Printf("🗑️ Removed %s\n", file)
		}
	}

	// Remove directory tree
	if err := os.RemoveAll(dirName); err != nil {
		fmt.Printf("⚠️ Failed to remove directory %s: %v\n", dirName, err)
	} else {
		fmt.Printf("🗑️ Removed directory %s\n", dirName)
	}

	fmt.Println("\n✨ All file operations completed!")
}

// 📋 HELPER FUNCTION: Copy file from source to destination
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Sync to ensure data is written to disk
	return destFile.Sync()
}

/*
=============================================================================
                              📝 LEARNING NOTES
=============================================================================

📁 FILE OPERATION FUNCTIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Simple operations                                                    │
│ os.ReadFile(name)           // Read entire file                         │
│ os.WriteFile(name, data, perm) // Write entire file                     │
│                                                                         │
│ // Advanced operations                                                  │
│ os.Open(name)               // Open for reading                         │
│ os.Create(name)             // Create/truncate for writing              │
│ os.OpenFile(name, flag, perm) // Open with specific flags               │
│                                                                         │
│ // File info                                                            │
│ os.Stat(name)               // Get file info                            │
│ file.Stat()                 // Get info from open file                  │
└─────────────────────────────────────────────────────────────────────────┘

🚩 FILE OPEN FLAGS:
┌─────────────────┬─────────────────────────────────────────────────────────┐
│     Flag        │                Description                              │
├─────────────────┼─────────────────────────────────────────────────────────┤
│ os.O_RDONLY     │ Read only                                               │
│ os.O_WRONLY     │ Write only                                              │
│ os.O_RDWR       │ Read and write                                          │
│ os.O_APPEND     │ Append to file                                          │
│ os.O_CREATE     │ Create if doesn't exist                                 │
│ os.O_EXCL       │ Fail if file exists (with O_CREATE)                    │
│ os.O_SYNC       │ Synchronous I/O                                         │
│ os.O_TRUNC      │ Truncate file when opening                              │
└─────────────────┴─────────────────────────────────────────────────────────┘

📊 FILE PERMISSIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Octal notation                                                       │
│ 0644  // rw-r--r-- (owner: read/write, group/others: read)              │
│ 0755  // rwxr-xr-x (owner: all, group/others: read/execute)             │
│ 0600  // rw------- (owner: read/write only)                             │
│                                                                         │
│ // Constants                                                            │
│ os.FileMode(0644)                                                       │
│ os.ModePerm      // 0777                                                │
└─────────────────────────────────────────────────────────────────────────┘

📁 DIRECTORY OPERATIONS:
┌─────────────────────────────────────────────────────────────────────────┐
│ os.Mkdir(name, perm)        // Create single directory                  │
│ os.MkdirAll(path, perm)     // Create directory tree                    │
│ os.Remove(name)             // Remove file or empty directory           │
│ os.RemoveAll(path)          // Remove directory tree                    │
│ os.ReadDir(dirname)         // List directory contents                  │
│ filepath.Walk(root, fn)     // Walk directory tree                      │
└─────────────────────────────────────────────────────────────────────────┘

🔄 READING PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Read entire file (small files)                                       │
│ data, err := os.ReadFile("file.txt")                                    │
│                                                                         │
│ // Read line by line (large files)                                      │
│ file, err := os.Open("file.txt")                                        │
│ scanner := bufio.NewScanner(file)                                       │
│ for scanner.Scan() {                                                    │
│     line := scanner.Text()                                              │
│ }                                                                       │
│                                                                         │
│ // Read with buffer                                                     │
│ reader := bufio.NewReader(file)                                         │
│ buffer := make([]byte, 1024)                                            │
│ n, err := reader.Read(buffer)                                           │
└─────────────────────────────────────────────────────────────────────────┘

✍️ WRITING PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Write entire file (small files)                                      │
│ err := os.WriteFile("file.txt", data, 0644)                             │
│                                                                         │
│ // Write with control                                                   │
│ file, err := os.Create("file.txt")                                      │
│ defer file.Close()                                                      │
│ file.WriteString("content")                                             │
│ file.Write([]byte("content"))                                           │
│                                                                         │
│ // Buffered writing                                                     │
│ writer := bufio.NewWriter(file)                                         │
│ writer.WriteString("content")                                           │
│ writer.Flush()                                                          │
└─────────────────────────────────────────────────────────────────────────┘

🚨 COMMON MISTAKES:
❌ Forgetting to close files (use defer)
❌ Not checking errors
❌ Reading large files entirely into memory
❌ Not handling file permissions correctly
❌ Platform-specific path separators (use filepath package)

💡 BEST PRACTICES:
• Always use defer file.Close() after opening
• Check errors from all file operations
• Use bufio for large files
• Use filepath package for cross-platform paths
• Set appropriate file permissions
• Use os.CreateTemp for temporary files
• Clean up temporary files and directories

🎯 REAL-WORLD PATTERNS:
┌─────────────────────────────────────────────────────────────────────────┐
│ // Configuration file reading                                           │
│ func loadConfig(filename string) (*Config, error) {                     │
│     data, err := os.ReadFile(filename)                                  │
│     if err != nil {                                                     │
│         return nil, err                                                 │
│     }                                                                   │
│     var config Config                                                   │
│     return &config, json.Unmarshal(data, &config)                       │
│ }                                                                       │
│                                                                         │
│ // Log file writing                                                     │
│ func writeLog(message string) error {                                   │
│     file, err := os.OpenFile("app.log",                                 │
│         os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)                      │
│     if err != nil {                                                     │
│         return err                                                      │
│     }                                                                   │
│     defer file.Close()                                                  │
│     timestamp := time.Now().Format(time.RFC3339)                        │
│     _, err = fmt.Fprintf(file, "[%s] %s\n", timestamp, message)         │
│     return err                                                          │
│ }                                                                       │
└─────────────────────────────────────────────────────────────────────────┘

⚡ PERFORMANCE TIPS:
• Use buffered I/O for frequent operations
• Read/write in chunks for large files
• Use io.Copy for efficient file copying
• Consider memory mapping for very large files
• Use sync.Pool for buffer reuse

=============================================================================
*/