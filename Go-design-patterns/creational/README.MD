Creational Patterns


Functional Optional Pattern VS Builder Pattern
- if conditional flag setting/ value assigning is there, then builder
- if functional options has to be provided, then functional optional pattern
        The builder works best if there are a bunch of conditionals:
        fooPrep := NewFoo()
        if conf.SetBar {
        fooPrep.SetBar()
        }
        foo := fooPrep.Build()

        If options are greater than just setting flags, then the options works well:
        foo := NewFoo(WithParsedValue(parser, value))


Singleton pattern
The Singleton pattern is implemented in Golang using a struct type and a private package-level variable that holds the only instance of the struct. The struct's constructor function is exported and is used to return the instance of the struct. 