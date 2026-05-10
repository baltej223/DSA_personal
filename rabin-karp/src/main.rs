#![allow(unused_variables)]
#![allow(dead_code)]

fn sol(t: Vec<i32>, p: Vec<i32>, d: i32, q: i32) -> i32 {
    let n = t.len();
    let m = p.len();

    if m > n {
        return -1;
    }

    let mut h = 1;
    for _ in 0..m - 1 {
        h = (h * d) % q;
    }

    let mut pattern_hash = 0;
    let mut text_hash = 0;

    for i in 0..m {
        pattern_hash = (d * pattern_hash + p[i]) % q;
        text_hash = (d * text_hash + t[i]) % q;
    }

    for i in 0..=n - m {
        if pattern_hash == text_hash {
            let mut matched = true;

            for j in 0..m {
                if p[j] != t[i + j] {
                    matched = false;
                    break;
                }
            }

            if matched {
                return i as i32;
            }
        }
        if i < n - m {
            text_hash = (d * (text_hash - t[i] * h) + t[i + m]) % q;

            if text_hash < 0 {
                text_hash += q;
            }
        }
    }
    -1
}

fn main() {
    let text: Vec<i32> = vec![2, 3, 5, 9, 0, 2, 3, 1, 4, 1, 5, 2, 6, 7];
    let pattern: Vec<i32> = vec![3, 1, 4, 1, 5];

    let pos = sol(text, pattern, 10, 13);

    println!("{pos}");
}

