# go-boilerplate

> A starter Golang project with a Dockerfile and a Makefile

## Getting started

```bash
# Clone project
git clone https://github.com/l-lin/go-boilerplate <your-project-name>
# Initialize your go module
go mod init github.com/l-lin/your-project-name
# Change the bin name in the Makefile if you want
sed -i 's/BIN_NAME=app/BIN_NAME=<your-project-name>/g' Makefile

# Use make watch to format, test and build your projet on file changes
make watch

# Use make to build the binary
make

# Use make release to generate binaries for all platforms
make release
```
