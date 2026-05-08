fn is_ok(chessboard: &Vec<Vec<bool>>) -> bool {
    let mut positions: Vec<usize> = vec![];
    for i in 0..4 {
        for j in 0..4 {
            if chessboard[i][j] {
                positions.push(j);
            }
        }
    }

    for m in 0..positions.len() {
        // if the any of the quess shares j or diag
        // diag = (i - k, j - k), (i - k, j + k), (i + k, j - k), (i + k, j + k) where k starts
        // from 1 and ahead.
        for k in (m + 1)..positions.len() {
            if positions[m] == positions[k] {
                return false;
            }

            if m.abs_diff(k) == positions[m].abs_diff(positions[k]) {
                return false;
            }
        }
    }
    true
}

fn sol(chessboard: &mut Vec<Vec<bool>>, i: usize) {
    if i == 4 {
        for row in chessboard {
            for val in row {
                print!("{} ", val);
            }
            println!("");
        }
        println!("--------");
        return;
    }

    for j in 0..4 {
        chessboard[i][j] = true;

        if is_ok(&chessboard) {
            sol(chessboard, i + 1);
        }
        chessboard[i][j] = false;
    }
}

fn main() {
    let mut chessboard: Vec<Vec<bool>> = vec![vec![false; 4]; 4];
    sol(&mut chessboard, 0);
}
