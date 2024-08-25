Go does not support traditional class-based inheritance as seen in other object-oriented languages like Java or C++. Instead, Go promotes composition over inheritance, which means that instead of inheriting behavior from a base class, you compose behavior using smaller, reusable components.

Composition in Go
In Go, composition is achieved through struct embedding, where one struct is embedded within another. The embedded struct's fields and methods are "promoted" to the embedding struct, giving the appearance of inheritance.