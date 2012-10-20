fn main(args: ~[str]) {
  let mut sum = 0;

  let mut i = 0;
  let mut j = 1;

  while j < 4000000 {
    if j % 2 == 0 {
      sum += j;
    }
    let temp = j;
    j = j + i;
    i = temp;
  }

  io::println(int::str(sum));
}
