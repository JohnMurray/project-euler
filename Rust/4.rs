// Written for Rust v0.3
fn main() {
  let mut largest_palindrome = 0;
  let mut right_hand_side = 999;

  while right_hand_side > 99 {
    let mut left_hand_side = 999;
    while left_hand_side > 99 {
      let possible_palindrome = right_hand_side * left_hand_side;

      if is_palindrome(possible_palindrome) && possible_palindrome > largest_palindrome {
        largest_palindrome = possible_palindrome;
      }
      left_hand_side -= 1;
    }

    right_hand_side -= 1;
  }

  io::println(int::str(largest_palindrome));
}

fn is_palindrome(num: int) -> bool {
  let mut rev = 0;
  let mut n = num;

  while n > 0 {
    rev = rev * 10 + (n % 10);
    n /= 10;
  }

  num == rev
}