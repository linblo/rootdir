rootdir
=======

Utility to get the parent directory that has either a file or directory. Useful
for determining project root dir by looking for files like go.mod and 
package.json or .git directory.

## Usage

```root dir [options] file [file2 file3 ...]```

Options:

`d` Set the max depth to recurse. Default is 10.


