struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>,
}

struct Linked_list<T> {
    head: Option<Box<Node<T>>>,
}

impl Linked_list<T> {
    fn new() -> Self {
        Linked_list { head: None }
    }
}
pub fn linked_list() {}
