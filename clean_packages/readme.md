## Clean Packages

I've often seen, and have been responsible for, throwing code into packages
without much thought. I've quickly drawn a line in the sand and started putting
code into different folders (which in Go are different packages by definition)
just for the sake of findability. Learning to properly build small and reusable
packages can take your Go career to the next level.

### Rules Of Thumb

#### 1. Hide Internal Logic

If you're familiar with the pillars of OOP, this i a practice in <em>encapsulation</em>.

Oftentimes an application will have complex logic that requires a lot of code.
In almost every case the logic that the application cares about can be exposed
via an API, and most of the dirty work can be kept within a package. For
example, imagine we are building an application that needs to classify images.
We could build a package:

```go
package classifier

// ClassifyImage classifies images as "hotdog" or "not hotdog"
func ClassifyIamge(image []byte)) (imageType string) {
    return hasHotdogColors(image) && hasHotDogShape(image)
}

func hasHotdogShape(image []byte) bool {
    // internal logic that the application doesn't need to know
    return true
}

func hasHotdogColors(image []byte) bool {
    // internal logic that the application doesn't need to know
    return true
}
```

We create an API by only exposing the function(s) that the application-level
needs to know about. All other logic is unexported to keep a clean separation of
concerns. The application doesn't need to know how to classify an image, just
the result of the classification.

#### 2. Don't Change APIs

The unexported functions within a package can and should change often for
testing, refactoring, and bug fixing.

A well-designed library will have a stable API so that users aren't receiving
breaking changes each time they update the package version. In Go, this means
not changing exported function's signatures.

#### 3. Don't Export Functions From The Main

A `main` package isn't a library, there's no need to export functions from it.

#### 4. Packages Shouldn't Know About Dependents

Perhaps one of the most important and most broken rules is that a package
shouldn't know anything about its dependents. In other words, a package should
never have specific knowledge about a particular application that uses it.
