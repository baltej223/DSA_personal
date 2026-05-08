fn main() {
    let mut v: Vec<i32> = Vec::new();
    v.push(1);
    v.push(3);
    v.push(2);

    InsertionSort(&mut v);
}

fn swap_if_bigger_than_prev(arr: &mut Vec<i32>, i: usize) {
    if i == 0 {
        return;
    }

    if arr[i] < arr[i - 1] {
        arr.swap(i, i - 1);
        swap_if_bigger_than_prev(arr, i - 1);
    }
}

fn InsertionSort(arr: &mut Vec<i32>) {
    for i in 1..arr.len() {
        swap_if_bigger_than_prev(arr, i)
    }
}
