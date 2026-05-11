#![allow(dead_code)]
#![allow(unused_variables)]
use std::collections::HashMap;
use std::collections::HashSet;

fn is_ok(trail: &Vec<i32>, starting_node: i32) -> bool {
    let mut seen = HashSet::new();
    let mut start_count = 0;

    for &num in trail {
        if num == starting_node {
            start_count += 1;
            if start_count > 2 {
                return false;
            }
        } else if !seen.insert(num) {
            return false;
        }
    }

    true
}

fn sol(map: &HashMap<i32, Vec<i32>>, starting_node: i32, node: i32, trail: &mut Vec<i32>) -> bool {
    if node == starting_node && trail.len() == map.len() + 1 {
        return true;
    }

    if let Some(neighbours) = map.get(&node) {
        for &neighbour in neighbours {
            trail.push(neighbour);
            if is_ok(trail, starting_node) && sol(&map, starting_node, neighbour, trail) == true {
                return true;
            }
            trail.pop();
        }
    }

    false
}

fn main() {
    let mut map: HashMap<i32, Vec<i32>> = HashMap::new();

    map.insert(1, vec![5, 4, 2]);
    map.insert(2, vec![1, 3]);
    map.insert(3, vec![4, 2]);
    map.insert(4, vec![5, 1, 3]);
    map.insert(5, vec![1, 4]);

    let mut trail = vec![1];
    let has_hamiltonian: bool = sol(&map, 1, 1, &mut trail);
    print!("{has_hamiltonian}");
}
