#include "field_element.hpp"

FieldElement::FieldElement(int num, int prime)
: _num(num), _prime(prime) {
	if (num >= prime || num < 0) {
		throw std::invalid_argument("Num " + std::to_string(num) + " not in field range 0 to " + std::to_string(prime - 1));
	}
}

FieldElement::~FieldElement() {}

int	FieldElement::getNum() const { return this->_num; }

int	FieldElement::getPrime() const { return this->_prime; }

bool	FieldElement::operator==(const FieldElement &other) const {
	return this->getNum() == other.getNum() && this->getPrime() == other.getPrime();
}

bool	FieldElement::operator!=(const FieldElement &other) const {
	return (*this == other) == false;
}

static int mod_pythonic(int x, int prime) {
	return ((x % prime) + prime) % prime;
}

FieldElement	FieldElement::operator+(const FieldElement &other) const {
	if (this->getPrime() != other.getPrime()) {
		throw std::invalid_argument("Cannot add two numbers in different Fields");
	}
	int num = mod_pythonic(this->getNum() + other.getNum(), this->getPrime());
	return FieldElement(num, this->_prime);
}

FieldElement	FieldElement::operator-(const FieldElement &other) const {
	if (this->getPrime() != other.getPrime()) {
		throw std::invalid_argument("Cannot subtract two numbers in different Fields");
	}
	int num = mod_pythonic(this->getNum() - other.getNum(), this->getPrime());
	return FieldElement(num, this->_prime);
}

FieldElement	FieldElement::operator*(const FieldElement &other) const {
	if (this->getPrime() != other.getPrime()) {
		throw std::invalid_argument("Cannot multiply two numbers in different Fields");
	}
	int num = mod_pythonic(this->getNum() * other.getNum(), this->getPrime());
	return FieldElement(num, this->_prime);
}

FieldElement	FieldElement::operator/(const FieldElement &other) const {
	if (this->getPrime() != other.getPrime()) {
		throw std::invalid_argument("Cannot divide two numbers in different Fields");
	}
	int a = other.getPrime() - 2;
	int b = pow(other.getNum(), a);

	int num = mod_pythonic(this->getNum() * other.getNum(), this->getPrime());
	return FieldElement(num, this->_prime);
}

FieldElement	FieldElement::pow(int exponent) const {
}

std::ostream	&operator<<(std::ostream &os, const FieldElement &f) {
	os << "FieldElement_" << f.getNum() << "(" << f.getPrime() << ")";
	return os;
}
