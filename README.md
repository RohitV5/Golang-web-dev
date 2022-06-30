# Golang-web-dev

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
