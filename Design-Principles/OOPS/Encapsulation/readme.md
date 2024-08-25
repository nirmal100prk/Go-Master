Go doesn't have classes like traditional object-oriented languages, but it achieves encapsulation through the use of exported and unexported fields and methods.

Exported: A field or method is exported (public) if its name starts with an uppercase letter. This makes it accessible from outside the package.
Unexported: A field or method is unexported (private) if its name starts with a lowercase letter. This restricts access to within the package where it is defined.