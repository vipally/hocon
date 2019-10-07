# hocon
hocon is a GO port of Human-Optimized Config Object Notation(HOCON).

hocon is a superset of [JSON]()
hocon main project(JAVA): [https://github.com/lightbend/config](https://github.com/lightbend/config)

# Features

- JSON
  - number
  - boolean
  - string
  - key-value
  - array

- HOCON add
  - Comments, with `#` or `//`
  - Allow omitting the `{}` around a root object
  - Allow `=` as a synonym for `:`
  - Allow omitting the `=` or `:` before a `{` so
    `foo { a : 42 }`
  - Allow omitting commas as long as there's a newline
  - Allow trailing commas after last element in objects and arrays
  - Allow unquoted strings for keys and values
  - Unquoted keys can use dot-notation for nested objects,
    `foo.bar=42` means `foo { bar : 42 }`
  - Duplicate keys are allowed; later values override earlier,
    except for object-valued keys where the two objects are merged
    recursively
  - `include` feature merges root object in another file into
    current object, so `foo { include "bar.json" }` merges keys in
    `bar.json` into the object `foo`
  - include with no file extension includes any of `.conf`,
    `.json`, `.properties`
  - you can include files, URLs, or classpath resources; use
    `include url("http://example.com")` or `file()` or
    `classpath()` syntax to force the type, or use just `include
    "whatever"` to have the library do what you probably mean
    (Note: `url()`/`file()`/`classpath()` syntax is not supported
    in Play/Akka 2.0, only in later releases.)
  - substitutions `foo : ${a.b}` sets key `foo` to the same value
    as the `b` field in the `a` object
  - substitutions concatenate into unquoted strings, `foo : the
    quick ${colors.fox} jumped`
  - substitutions fall back to environment variables if they don't
    resolve in the config itself, so `${HOME}` would work as you
    expect. Also, most configs have system properties merged in so
    you could use `${user.home}`.
  - substitutions normally cause an error if unresolved, but
    there is a syntax `${?a.b}` to permit them to be missing.
  - `+=` syntax to append elements to arrays, `path += "/bin"`
  - multi-line strings with triple quotes as in Python or Scala

### Examples of HOCON

All of these are valid HOCON.

Start with valid JSON:

    {
        "foo" : {
            "bar" : 10,
            "baz" : 12
        }
    }

Drop root braces:

    "foo" : {
        "bar" : 10,
        "baz" : 12
    }

Drop quotes:

    foo : {
        bar : 10,
        baz : 12
    }

Use `=` and omit it before `{`:

    foo {
        bar = 10,
        baz = 12
    }

Remove commas:

    foo {
        bar = 10
        baz = 12
    }

Use dotted notation for unquoted keys:

    foo.bar=10
    foo.baz=12

Put the dotted-notation fields on a single line:

    foo.bar=10, foo.baz=12
