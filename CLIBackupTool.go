package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func copyCMD(flag string) {

	// Taking in input of source and backup directory
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Source Directory: ")
	sourceDir, _ := reader.ReadString('\n')
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Printf("Error: Source directory '%s' does not exist.\n", sourceDir)
	}
	fmt.Printf("Enter Backup Directory: ")
	backupDir, _ := reader.ReadString('\n')
	sourceDir = strings.TrimSpace(sourceDir)
	backupDir = strings.TrimSpace(backupDir)

	// Execute command according to the flag

	// Just copies the directory if no flag is used
	if flag == "none" {
		err := copyFiles(sourceDir, backupDir)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
		} else {
			fmt.Println("All files copied successfully!")
		}

		// Encrypts and copies the directory requiring a key file for encryption
	} else if flag == "encrypt" {

		var key []byte
		var keyFile string

		fmt.Printf("Enter Directory of the Key File: ")
		keyFile, _ = reader.ReadString('\n')
		keyFile = strings.TrimSpace(keyFile)
		key, err := os.ReadFile(keyFile)
		if err != nil {
			fmt.Printf("Error reading key file: %v", err)
		}

		err = copyEncryptFiles(sourceDir, backupDir, key)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
		} else {
			fmt.Println("All files copied successfully!")
		}

		// Generates a key file for encryption and encrypts and copies the directory
	} else if flag == "encrypt-gen" {

		var key []byte
		var keyFile string

		key, err := generateRandomKey()
		if err != nil {
			fmt.Printf("Error generating random key: %v", err)
		}

		keyFile = filepath.Join(backupDir, "key.txt")
		err = os.WriteFile(keyFile, key, 0644)
		if err != nil {
			fmt.Printf("Error saving key to file: %v", err)
		}
		fmt.Printf("Random key generated and saved to %s\n", keyFile)

		err = copyEncryptFiles(sourceDir, backupDir, key)
		if err != nil {
			fmt.Printf("Error copying file: %v\n", err)
		} else {
			fmt.Println("All files copied successfully!")
		}
	}
}

func copyFiles(sourceDir string, backupDir string) error {

	// Code to copy all files from source directory to backup directory
	// Open the source directory
	source, err := os.Open(sourceDir)
	if err != nil {
		return fmt.Errorf("Error opening source directory: %v", err)
	}
	defer source.Close()

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("Error creating destination directory: %v", err)
	}

	// Read all files in the source directory
	files, err := source.Readdir(-1)
	if err != nil {
		return fmt.Errorf("Error reading source directory: %v", err)
	}

	// Copy each file from source directory to destination directory
	for _, file := range files {

		srcFile := filepath.Join(sourceDir, file.Name())
		destFile := filepath.Join(backupDir, file.Name())

		// Open source file
		src, err := os.Open(srcFile)
		if err != nil {
			return fmt.Errorf("Error opening source file: %v", err)
		}
		defer src.Close()

		// Create destination file
		dst, err := os.Create(destFile)
		if err != nil {
			return fmt.Errorf("Error creating destination file: %v", err)
		}
		defer dst.Close()

		// Copy contents from source to destination
		_, err = io.Copy(dst, src)
		if err != nil {
			return fmt.Errorf("Error copying file: %v", err)
		}

		fmt.Printf("File %s copied successfully to %s\n", srcFile, destFile)
	}
	return nil
}

func copyEncryptFiles(sourceDir, backupDir string, key []byte) error {

	// Code to copy and encrypt all files from source directory to backup directory
	// Open the source directory
	source, err := os.Open(sourceDir)
	if err != nil {
		return fmt.Errorf("Error opening source directory: %v", err)
	}
	defer source.Close()

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("Error creating destination directory: %v", err)
	}

	// Read all files in the source directory
	files, err := source.Readdir(-1)
	if err != nil {
		return fmt.Errorf("Error reading source directory: %v", err)
	}

	// Create AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("Error creating block cipher: %v", err)
	}

	// Get the block size
	blockSize := block.BlockSize()

	// Iterate over each file in the source directory
	for _, file := range files {
		// Construct source and destination file paths
		srcFile := filepath.Join(sourceDir, file.Name())
		destFile := filepath.Join(backupDir, file.Name())

		// Open source file
		src, err := os.Open(srcFile)
		if err != nil {
			return fmt.Errorf("Error opening source file: %v", err)
		}
		defer src.Close()

		// Create destination file
		dst, err := os.Create(destFile)
		if err != nil {
			return fmt.Errorf("Error creating destination file: %v", err)
		}
		defer dst.Close()

		// Generate nonce and write nonce to the beginning of the file
		nonce := make([]byte, blockSize)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return fmt.Errorf("Error generating nonce: %v", err)
		}
		if _, err := dst.Write(nonce); err != nil {
			return fmt.Errorf("Error writing nonce: %v", err)
		}

		// Create cipher stream using CTR mode
		stream := cipher.NewCTR(block, nonce)

		// Encrypt and copy data
		buffer := make([]byte, 4096)
		for {
			n, err := src.Read(buffer)
			if err != nil && err != io.EOF {
				return fmt.Errorf("Error encrypting and copying file: %v", err)
			}
			if n == 0 {
				break
			}
			stream.XORKeyStream(buffer[:n], buffer[:n])
			if _, err := dst.Write(buffer[:n]); err != nil {
				return fmt.Errorf("Error encrypting and copying file: %v", err)
			}
		}
		fmt.Printf("File %s encrypted and copied successfully to %s\n", srcFile, destFile)
	}

	return nil
}

func generateRandomKey() ([]byte, error) {
	// AES-256 requires a 32-byte key

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func main() {
	x := 1
	for x == 1 {

		// Taking Command input
		var input string
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("User:JJBigDub > ")
		input, _ = reader.ReadString('\n')

		// Split the input command into parts using spaces as separators
		parts := strings.Fields(input)
		if parts[0] == "exit" {
			os.Exit(0)
		}

		// Understanding the input command
		if len(parts) == 1 {
			if parts[0] == "copy" {
				copyCMD("none")
			} else {
				fmt.Println("Invalid command: copy -flag ")
			}
		} else if len(parts) == 2 {
			if parts[0] == "copy" && parts[1][0] == '-' {
				copyCMD(parts[1][1:])
			} else {
				fmt.Println("Invalid command: copy -flag ")
			}
		} else {
			fmt.Println("Invalid command: copy -flag ")
		}

	}
}
