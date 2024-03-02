class Point:
	def __init__(self, x, y, a, b):
		self.x = x
		self.y = y
		self.a = a
		self.b = b
		if self.x is None and self.y is None:
			return
		if self.y ** 2 != self.x ** 3 + a * x + b:
			raise ValueError('({}, {}) is not on the curve'.format(x, y))

	def __eq__(self, other):
		return self.x == other.x and self.y == other.y \
	and self.a == other.a and self.b == other.b

	def __ne__(self, other):
		return self.x != other.x or self.y != other.y \
	or self.a != other.a or self.b != other.b

	def __add__(self, other):
		if self.a != other.a or self.b != other.b:
			raise TypeError('Points {}, {} are not on the same curve'.format(self, other))

		if self.x == other.x and self.y != other.y:
			return self.__class__(None, None, self.a, self.b)

		if self.x is None:
			return other
		if other.x is None:
			return self