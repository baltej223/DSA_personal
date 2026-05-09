fn is_ok(w: usize) -> bool {
    w < 8
}

fn sol(pw: &mut Vec<[i32; 2]>, object: usize, current_p: usize, current_w: usize) -> usize {
    if object == pw.len() {
        return current_p;
    }

    let new_w = current_w + pw[object][0] as usize;
    let new_p = current_p + pw[object][1] as usize;

    let skip = sol(pw, object + 1, current_p, current_w);

    let take: usize = if is_ok(new_w) {
        sol(pw, object + 1, new_p, new_w)
    } else {
        0
    };
    skip.max(take)
}

fn main() {
    let mut pw: Vec<[i32; 2]> = vec![[3, 2], [5, 3], [6, 4], [10, 5]];

    println!("{}", sol(&mut pw, 0, 0, 0));
}
