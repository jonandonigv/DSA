use std::{
    cell::RefCell,
    rc::{Rc, Weak},
};

#[derive(Debug)]
struct Node<T> {
    value: T,
    next: Option<Rc<RefCell<Node<T>>>>,
    prev: Option<Weak<RefCell<Node<T>>>>,
}

#[derive(Debug)]
pub struct Doubly_linked_list<T> {
    head: Option<Rc<RefCell<Node<T>>>>,
    tail: Option<Rc<RefCell<Node<T>>>>,
}

impl<T: std::fmt::Debug> Doubly_linked_list<T> {
    pub fn new() -> Self {
        Doubly_linked_list {
            head: None,
            tail: None,
        }
    }

    pub fn push_front(&mut self, value: T) {
        let new_node = Rc::new(RefCell::new(Node {
            value,
            next: self.head.clone(),
            prev: None,
        }));

        if let Some(old_head) = self.head.take() {
            old_head.borrow_mut().prev = Some(Rc::downgrade(&new_node));
        } else {
            self.tail = Some(new_node.clone());
        }

        self.head = Some(new_node);
    }

    pub fn pop_front(&mut self) -> Option<T> {
        self.head.take().map(|old_head| {
            self.head = old_head.borrow().next.clone();

            if let Some(new_head) = &self.head {
                new_head.borrow_mut().prev = None;
            } else {
                self.tail = None;
            }

            Rc::try_unwrap(old_head).ok().unwrap().into_inner().value
        })
    }

    pub fn print_list(&self) {
        let mut current = self.head.clone();
        print!("List: ");
        while let Some(node) = current {
            print!("{:?} <-> ", node.borrow().value);
            current = node.borrow().next.clone();
        }

        println!("None");
    }
}
