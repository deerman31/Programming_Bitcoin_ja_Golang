use core::fmt;
use std::{
    intrinsics::mir::Field,
    ops::{Add, Mul, Sub},
};

struct FieldElement {
    num: i64,
    prime: i64,
}

impl FieldElement {
    fn new(num: i64, prime: i64) -> Self {
        if num >= prime || num < 0 {
            panic!("Num {} not in field range 0 to {}", num, prime - 1);
        }
        FieldElement { num, prime }
    }
    fn pow(&self, exponent: i64) -> Self {
        if exponent < 0 {
            let inv = self.pow(self.prime - 2 - exponent.abs());
            return inv;
        }

        let mut result = 1;
        let mut base = self.num % self.prime;
        let mut exp = exponent;
        let prime = self.prime;

        while exp > 0 {
            if exp % 2 == 1 {
                result = (result * base) % prime;
            }
            exp /= 2;
            base = (base * base) % prime;
        }
        FieldElement::new(result, prime)
    }

    fn mod_python_style(a: i64, b: i64) -> i64 {
        if a % b < 0 {
            a % b + b
        } else {
            a % b
        }
    }
}

impl PartialEq for FieldElement {
    fn eq(&self, other: &Self) -> bool {
        self.num == other.num && self.prime == other.prime
    }
}

impl fmt::Display for FieldElement {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "num: {} prime: {}", self.num, self.prime)
    }
}

impl Add for FieldElement {
    type Output = Self;
    fn add(self, other: Self) -> Self {
        if self.prime != other.prime {
            panic!("Cannot add two numbers in different Fields");
        }
        let num = Self::mod_python_style(self.num + other.num, self.prime);
        FieldElement::new(num, self.prime)
    }
}

impl Sub for FieldElement {
    type Output = Self;
    fn sub(self, other: Self) -> Self {
        if self.prime != other.prime {
            panic!("Cannot subtract two numbers in different Fields");
        }
        let num = Self::mod_python_style(self.num - other.num, self.prime);
        FieldElement::new(num, self.prime)
    }
}

impl Mul for FieldElement {
    type Output = Self;
    fn mul(self, other: Self) -> Self {
        if self.prime != other.prime {
            panic!("Cannot multiply two numbers in different Fields");
        }
        let onum: usize = other.num as usize;
        let inv = onum.pow((self.prime - 2) as u32) % self.prime as usize;
        let num = Self::mod_python_style(self.num * inv as i64, self.prime);
        FieldElement::new(num, self.prime)
    }
}
