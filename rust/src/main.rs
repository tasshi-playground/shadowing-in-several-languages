fn main() {
    println!("value of one() is {}", one()); // print 1

    let one: i32 = one();
    println!("value of one is {}", one); // print 1

    //let two = one() + 1;
    // Error: expected function, found `i32`

    let two = one + 1;
    println!("value of two is {}", two); // print 2
}

fn one() -> i32 {
    1
}
