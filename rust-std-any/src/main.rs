use std::any::Any;
use std::fmt::Debug;

// Dynamic typing
//
// This is reflection. Useful when we want to do something differently depending
// on the type of the argument but we don't know the type until runtime.
//
// In this example we know we want something which implements [Debug] but we
// want to log a different kind of message depending on the concrete type,
// so [Any] allows us to look into the type, if it was just a generic
// `log<T: Debug>` we wouldn't have any type information at runtime.

fn log<T: Any + Debug>(val: &T) {
    let any = val as &dyn Any;

    if any.is::<String>() {
        println!("Value is a string, no downcast");
        // note that we cannot call `any.len()` here, we need to downcast first
    }
    match any.downcast_ref::<String>() {
        Some(string) => {
            println!(
                "Value is a string with downcast (len={}): {:?}",
                string.len(),
                string
            );
        }
        None => {
            println!("{val:?}");
        }
    }
}

fn main() {
    let x = "Lorem ipsum".to_string();
    log(&x);

    let y = 10.0;
    log(&y);
}
