#ifndef FIELD_ELEMENT_HPP
# define FIELD_ELEMENT_HPP

# include <stdexcept>
# include <string>
# include <iostream>

class FieldElement {
	private:
	const int		_num;
	const int		_prime;
	public:
	FieldElement(int, int);
	~FieldElement();
	int	getNum() const;
	int	getPrime() const;

	bool	operator==(const FieldElement &) const;
	bool	operator!=(const FieldElement &) const;

	FieldElement	operator+(const FieldElement &) const;
	FieldElement	operator-(const FieldElement &) const;
	FieldElement	operator*(const FieldElement &) const;
	FieldElement	operator/(const FieldElement &) const;

	FieldElement	pow(int) const;
};
std::ostream	&operator<<(std::ostream &, const FieldElement &);

#endif