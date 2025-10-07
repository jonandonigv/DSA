pub fn dynamic_array() {
    let mut dynamic_array: Vec<i32> = Vec::new();
    // add elements
    dynamic_array.push(1);
    dynamic_array.push(2);
    dynamic_array.push(3);

    // initialize with values
    let mut another_array = vec![4, 5, 6];

    // append elements from another vec
    dynamic_array.append(&mut another_array);

    // access elements
    println!("First element {}", dynamic_array[0]);
    println!("Array length {}", dynamic_array.len());

    // iterate over the array
    for iter in &dynamic_array {
        println!("Item: {}", iter);
    }

    if let Some(removed) = dynamic_array.pop() {
        println!("Removed: {}", removed);
    }

    // print the final array
    println!("Final array: {:?}", dynamic_array);
}
pub fn static_array() {}

pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
