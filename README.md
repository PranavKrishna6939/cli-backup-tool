# CLI Backup Tool

This tool allows users to copy files from a source directory to a backup directory with different encryption options.
This README provides a brief overview of the tool's usage, along with examples of the commands in action. Users can refer to this document for guidance on how to use the tool effectively.

## Commands:

### 1. `copy`

```shell
$ go run main.go
User:JJBigDub > copy
Enter Source Directory: /path/to/source
Enter Backup Directory: /path/to/backup
File /path/to/source/test1.txt copied successfully to /path/to/backup/test1.txt
File /path/to/source/test2.txt copied successfully to /path/to/backup/test2.txt
All files copied successfully!
User:JJBigDub >
```
This command copies all files from a source directory to a backup directory without encryption.

### 2. `copy -encrypt`

```shell
$ go run main.go
User:JJBigDub > copy -encrypt
Enter Source Directory: /path/to/source
Enter Backup Directory: /path/to/backup
Enter Directory of the Key File: /path/to/key.txt
File /path/to/source/test1.txt encrypted and copied successfully to /path/to/backup/test1.txt
File /path/to/source/test2.txt encrypted and copied successfully to /path/to/backup/test2.txt
All files copied successfully!
User:JJBigDub >
```
This command encrypts and copies all files from a source directory to a backup directory. It requires a key file for encryption. Users will be prompted to enter the directory of the key file.

### 3. `copy -encrypt-gen`

```shell
$ go run main.go
User:JJBigDub > copy -encrypt-gen
Enter Source Directory: /path/to/source
Enter Backup Directory: /path/to/backup
Random key generated and saved to /path/to/backup/key.txt
File /path/to/source/test1.txt encrypted and copied successfully to /path/to/backup/test1.txt
File /path/to/source/test2.txt encrypted and copied successfully to /path/to/backup/test2.txt
All files copied successfully!
User:JJBigDub >
```

This command generates a random key, saves it to a file in the backup directory (`key.txt`), and then encrypts and copies all files from a source directory to a backup directory using the generated key.

### 4. `exit`

```shell
User:JJBigDub > exit
```
This command exits the command line interface (CLI).

## Usage:

1. Run the program and enter the desired command (`copy`, `copy -encrypt`, or `copy -encrypt-gen`).
2. Follow the prompts to provide the necessary input such as source directory, backup directory, and key file location (if applicable).
3. The tool will perform the specified action and provide feedback on the progress and any errors encountered.


