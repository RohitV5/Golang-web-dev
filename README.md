# Golang-web-dev

go.dev
pkg.go.dev

Main method is the entry point.

Gohtml is just a convention. You can rename it anything.

Go packages can be imported from github if given proper package name.

To init a new project
go mod init [yourModule.name]

This will create the go.mod file and the go.sum file for you and download the modules you use in your project.

Implementing an interface

interface StringToIntMapper {
  fun map(string: String): Int
}
class DefaultStringToIntMapper : StringToIntMapper {
  override fun map(string: String): Int {
    return string.toInt()
  }
}

Capital letters make a method public else private

Defer executes the statement at the end of the function. In LIFO order.

For files we have io utils package.
file, err := os.Create("./nylcogofile.text");
io.WriteString(file, "content here")


Array is less used as compare to Slice for a list of data.

Making build for windows, unix and mac
go build


https://www.youtube.com/watch?v=vMpBUkIMLzY&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=50



For mongodb in local make sure mongo service is running before troubleshooting.

#### Write go like a senior
##### https://levelup.gitconnected.com/write-go-like-a-senior-engineer-eee7f03a1883

#### Different ways to initialize go struct.
##### https://asankov.dev/blog/2022/01/29/different-ways-to-initialize-go-structs/
