# CLI Backup Tool

This tool allows users to copy files from a source directory to a backup directory with different encryption options.

## Commands:

### 1. `copy`

```shell
copy
```
This command copies all files from a source directory to a backup directory without encryption.

### 2. `copy -encrypt`

```shell
$ go run main.go
User:JJBigDub > copy
Enter Source Directory: /path/to/source
Enter Backup Directory: /path/to/backup
All files copied successfully!
User:JJBigDub >
```
This command encrypts and copies all files from a source directory to a backup directory. It requires a key file for encryption. Users will be prompted to enter the directory of the key file.

### 3. `copy -encrypt-gen`

This command generates a random key, saves it to a file in the backup directory (`key.txt`), and then encrypts and copies all files from a source directory to a backup directory using the generated key.

## Usage:

1. Run the program and enter the desired command (`copy`, `copy -encrypt`, or `copy -encrypt-gen`).
2. Follow the prompts to provide the necessary input such as source directory, backup directory, and key file location (if applicable).
3. The tool will perform the specified action and provide feedback on the progress and any errors encountered.


