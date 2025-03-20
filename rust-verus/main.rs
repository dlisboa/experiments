use vstd::prelude::*;

verus! {

#[derive(Debug)]
enum Beverage {
    Coffee { sugar: bool },
    Water { ice: bool },
}

fn f(i: i32) -> (x: i32)
  requires
    i > 10,
    i < 20,
  ensures
    x >= i,
{
    let y = i + 2;
    y
}

fn drink_coffee(bev: &Beverage) -> bool
    requires bev is Coffee
{
    true
}

fn binary_search(v: &Vec<u64>, k: u64) -> (r: usize)
    requires
        forall|i: int, j: int| 0 <= i <= j < v.len() ==> v[i] <= v[j],
        exists|i: int| 0 <= i < v.len() && k == v[i],
    ensures
        r < v.len(),
        k == v[r as int],
{
    // lower and upper bounds
    let mut lower: usize = 0;
    let mut upper: usize = v.len() - 1; // change to `upper = v.len()` to see verification error
    while lower != upper
        invariant
            upper < v.len(),
            exists|idx: int| lower <= idx <= upper && v[idx] == k,
            forall|low: int, up: int| 0 <= low <= up < v.len() ==> v[low] <= v[up],
    {
        let idx = lower + (upper - lower) / 2;
        if v[idx] < k {
            lower = idx + 1;
        } else {
            upper = idx;
        }
    }
    lower
}

mod tests {
    use super::*;

    fn prove_f() {
        f(19);
        //f(20); // failed precondition
    }

    fn prove_binary_search() {
        let mut v: Vec<u64> = vec![0, 10, 20, 30, 40];
        assert(v[3] == 30);  // needed to trigger exists|i: int| ... k == v[i]
        // change to 30 to 31 to see verification error
        let idx = binary_search(&v, 30);
        assert(idx == 3);
    }

    fn prove_drink_cofee() {
        // change to Beverage::Water to see verification error
        let coffee = Beverage::Coffee { sugar: true };
        drink_coffee(&coffee);
    }
}

}

// Code outside of versus!{} block will not be verified, but this allows us to use println! and
// other side effects.
fn main() {
    let a = f(20);

    let mut coffee = Beverage::Coffee { sugar: true };
    drink_coffee(&coffee);
    println!("{:?}", coffee);

    let v = vec![1, 4, 3, 5, 6, 7, 10, 2];
    let idx = binary_search(&v, 7);
    println!("{} {}", idx, v[idx]);
}
