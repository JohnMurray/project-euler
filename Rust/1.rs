use std;

// Written for Rust v0.3
fn main(args: ~[str]) {
  let mut sum = 0;
  let mut i = 0;
  while i < 1000 {
    if i % 5 == 0 || i % 3 == 0 {
      sum += i
    }
    i += 1;
  }

  io::println(int::str(sum));
}
