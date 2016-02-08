use std::io;

fn main() {
	let mut s = String::new();
	io::stdin().read_line(&mut s).unwrap(); // read line into s
	s.clear(); // clear s
	io::stdin().read_line(&mut s).unwrap(); // read line into s
	let arr: Vec<u32> = s.split_whitespace(). // split per space
		map(|x| x.trim().parse().unwrap()). // each do (trim.parse.unwrap)
		rev(). // reverse
		collect(); // collection
	for x in arr {
		println!("{}", x);
	}
}
