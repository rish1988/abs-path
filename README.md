# Abs-path
A small utility for generating absolute directory/file path from relative path

# Usage 

Import the package `github.com/rish1988/abs-path` which provides access to two APIs:

* `GetAbsoluteDirPath` - Returns absolute path to a directory
* `GetAbsoluteFilePath` - Returns absolute path to a file

# Example

```go
package main

import abs "github.com/rish1988/abs-path"

func main() {
    // Get absolute path to user ssh directory
    absSshDirPath := abs.GetAbsoluteDirPath("~/.ssh")
    
    // Get absolute path to user ssh key
    absSshKeyPath := abs.GetAbsoluteDirPath("~/.ssh", "id_rsa.pub")
}
```
