use dsa::{doubly_linked_list::Doubly_linked_list, linked_list::Linked_list};

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

    // Doubly linked list
    let mut d_list: Doubly_linked_list<i32> = dsa::doubly_linked_list::Doubly_linked_list::new();
    d_list.push_front(3);
    d_list.push_front(2);
    d_list.push_front(1);
    d_list.print_list();
    // Pop values
    println!("Popped: {:?}", d_list.pop_front());
    d_list.print_list();

    println!("Popped: {:?}", d_list.pop_front());
    d_list.print_list();
}
