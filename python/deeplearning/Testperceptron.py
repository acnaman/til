import unittest as ut
import perceptron as pt

class Testperceptron(ut.TestCase):
    def test_and(self):
        self.assertEqual(0, pt.AND(0, 0))
        self.assertEqual(0, pt.AND(0, 1))
        self.assertEqual(0, pt.AND(1, 0))
        self.assertEqual(1, pt.AND(1, 1))

    def test_or(self):
        self.assertEqual(0, pt.OR(0, 0))
        self.assertEqual(1, pt.OR(0, 1))
        self.assertEqual(1, pt.OR(1, 0))
        self.assertEqual(1, pt.OR(1, 1))

    def test_nand(self):
        self.assertEqual(1, pt.NAND(0, 0))
        self.assertEqual(1, pt.NAND(0, 1))
        self.assertEqual(1, pt.NAND(1, 0))
        self.assertEqual(0, pt.NAND(1, 1))


if __name__ == '__main__':
    ut.main()

