#[derive(Debug)]
struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>,
}

#[derive(Debug)]
pub struct Linked_list<T> {
    head: Option<Box<Node<T>>>,
}

impl<T> Linked_list<T> {
    // Create an empty list
    pub fn new() -> Self {
        Linked_list { head: None }
    }

    // Push a value to the front of the list
    pub fn push(&mut self, value: T) {
        let new_node = Box::new(Node {
            value,
            next: self.head.take(),
        });
        self.head = Some(new_node);
    }

    // Pop the front value
    pub fn pop(&mut self) -> Option<T> {
        self.head.take().map(|node| {
            self.head = node.next;
            node.value
        })
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_ref().map(|node| &node.value)
    }
}
