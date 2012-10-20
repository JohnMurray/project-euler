// Written for Rust v0.3
fn main(args: ~[str]) {
  let big_o_number = 600851475143;

  let mut i = 2;
  let mut largest_prime = 1;

  while i < big_o_number / 2 + 1 {
    if big_o_number % i == 0 {
      if is_prime(i) {
        largest_prime = i;
      }
    }
    i += 1;
  }

  io::println(int::str(largest_prime));
}

fn is_prime(i: int) -> bool {
  if i % 2 == 0 {
    ret false;
  }
  let mut j = 3;
  while j < i / 2 + 1 {
    if i % j == 0 {
      ret false;
    }
    j += 1;
  }
  ret true;
}
