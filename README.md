#GMock [![Build Status](https://travis-ci.org/cfmobile/gmock.svg?branch=master)](https://travis-ci.org/cfmobile/gmock)
A mocking utility for Go.

This was created with unit testing in mind, to make it easier to isolate components.

##Installation

Just use go get

```
  go get github.com/cfmobile/gmock
```

##Usage

Import the header

```go
import (
  //...
  "github.com/cfmobile/gmock"
)
```

Use `GMock` to mock your variables with another value of the same type.

1. Call one of the available constructors with a **pointer** to the variable you'd like to mock. Retain the `*GMock` object returned.
2. Assign a mock value of the same type (if you haven't done so in the constructor) by passing it to the `Replace(_)` method in the `*GMock` object you retained.
3. Once you're done using your mocked value, restore the original value by calling `Restore()` in the `*GMock` object.

**Important: If you don't call `Restore()` and the GMock object is destroyed, you won't be able to restore the original value to your variable.**

####Example 1:
Using the constructor `CreateMockWithTarget(_)`

The default constructor doesn't change the value until `Replace(_)` is called.

```go
package main
import (
  "fmt"
  "github.com/cfmobile/gmock"
)

func main() {
  var someVariable = "my original value"

  fmt.Println("1:", someVariable)

  mock := gmock.CreateMockWithTarget(&someVariable)

  fmt.Println("2:", someVariable)

  mock.Replace("mocked value")

  fmt.Println("3:", someVariable)

  mock.Restore()

  fmt.Println("4:", someVariable)
}
```
Should print:
```
1: my original value
2: my original value
3: mocked value
4: my original value
```

####Example 2:
Using the constructor `MockTargetWithValue(_, _)`

The target variable is immediately mocked with the value specified in the second parameter

[Playground link](https://travis-ci.org/cfmobile/gmock.svg?branch=master)]

```go
package main
import (
  "fmt"
  "github.com/cfmobile/gmock"
)

func main() {
  var someVariable = "my original value"

  fmt.Println("1:", someVariable)

  mock := gmock.MockTargetWithValue(&someVariable, "mocked value")

  fmt.Println("2:", someVariable)

  mock.Restore()

  fmt.Println("3:", someVariable)
}
```
Should print:
```
1: my original value
2: mocked value
3: my original value
```

####Example 3:
Multiple calls to `Replace(_)`

Calling `Replace(_)` multiple times won't make `GMock` lose track of the original value.

```go
package main
import (
  "fmt"
  "github.com/cfmobile/gmock"
)

func main() {
  var someVariable = "my original value"

  fmt.Println("1:", someVariable)

  mock := gmock.MockTargetWithValue(&someVariable, "mocked value")

  fmt.Println("2:", someVariable)

  mock.Replace("another mock value")

  fmt.Println("3:", someVariable)

  mock.Restore()

  fmt.Println("4:", someVariable)
}
```
Should print:
```
1: my original value
2: mocked value
3: another mock value
4: my original value
```
