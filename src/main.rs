use dsa::{add, linked_list::Linked_list};

fn main() {
    // Linked list
    let mut list: Linked_list<i32> = dsa::linked_list::Linked_list::new();
    // Push some values
    list.push(1);
    list.push(2);
    list.push(3);
    println!("List: {:?}", list);
    // Peek at the head
    println!("Peek: {:?}", list.peek());
    // Pop values
    println!("Popped: {:?}", list.pop());
    println!("Popped: {:?}", list.pop());
    println!("List after pops: {:?}", list);

    let result = add(2, 2);
    println!("Hello world!");
    println!("{result}");
}
