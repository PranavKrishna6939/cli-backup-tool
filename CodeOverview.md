# CLI Backup Tool

This tool is a Go program designed to facilitate file backup operations with optional encryption. It provides functionality to copy files from a source directory to a backup directory while offering encryption options using the AES algorithm.

## Code Overview:

### `main` Function:

- The main entry point of the program.
- Uses while loop to provide a command-line interface for users to interact with the tool.
- Reads user commands, parses them, and invokes the appropriate functions based on the command and flags provided.

### `copyCMD` Function:

- This function handles the copy command and its variations (`copy`, `copy -encrypt`, `copy -encrypt-gen`).
- It prompts the user for source and backup directory paths.
- Depending on the command flag, it either performs a simple file copy (`copy`), encrypts files before copying (`copy -encrypt`), or generates a random encryption key before encrypting and copying files (`copy -encrypt-gen`).

### `copyFiles` Function:

- This function is responsible for copying files from a source directory to a backup directory without encryption.
- It iterates over all files in the source directory, opens each file, creates a corresponding file in the backup directory, and copies the contents from the source file to the destination file.

### `copyEncryptFiles` Function:

- This function handles the encryption and copying of files from a source directory to a backup directory.
- It performs similar operations to `copyFiles`, but before copying, it encrypts the contents of each file using the AES algorithm in CTR mode.
- The function generates a random nonce for each file and writes it to the beginning of the encrypted file. It then encrypts the file's contents using the nonce and a key provided by the user or generated randomly.

### `generateRandomKey` Function:

- This function generates a random AES encryption key of 32 bytes (256 bits).
- It is used when the `copy -encrypt-gen` command is invoked to generate a key for encrypting files.

## Important Note:

- Ensure that you have proper permissions to read from the source directory and write to the backup directory.
- Take caution when using encryption features, as losing the encryption key may result in data loss.

