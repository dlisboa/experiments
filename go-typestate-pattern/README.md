# An exploration of the Typestate Pattern in Go

> The typestate pattern is an API design pattern that encodes
> information about an object’s run-time state in its compile-time type.
>
> — [The Typestate Pattern in Rust](https://cliffle.com/blog/rust-typestate/)

The typestate pattern allows one to write code with invariants that are checked
at compile time rather than run time. In this experiment we'll use the idea of a
content publishing software where Articles can be edited and then published.

These are the invariants we want to enforce:

* There is a type Article
* An Article can be in two states: Draft and Published
* Only a Draft Article is editable (can call Edit())
* Only a Published Article is linkable (can call Link())
* Both types of Article can call Content()
* Calling Publish() transitions a Draft Article to Published
* Calling Unpublish() transitions a Published Article back to Draft
* One cannot call Publish() on a Published Article
* One cannot call Unpublish() on a Draft Article

Invariants we wish we could enforce:

* One cannot use a variable of a Draft Article after it's been published

## Conclusions

All solutions will allow the "Use After Publish" problem where a variable with a
Draft article can still be used after it has transitioned to Published. This is
because Go does not have the same idea of "consume" as Rust (almost no language
has). The variable will remain valid in the scope.

### With manual monomorphization

Here we do the monomorphization (instantiation) by hand, creating a new named
type instantiated with the state. This works fine since each named type will
have its own set of methods and the compiler will check that we're not calling
the wrong method on each (like calling `Edit` on a Published Envelope). However
if we do implement a generic method it'll be callable in either the instantiated
type and the generic type. This might be a possible source of mistakes.

This is a good approach but we could argue it's no longer the Typestate
Pattern just normal subtyping with generics as a codegen tool.

