import unittest as ut
import neuron
import numpy as np

class Testneuron(ut.TestCase):

    def assertArray(self, first, second):
        res = (np.array(first) == np.array(second)).all()            
        self.assertTrue(res)

    def test_step_function(self):
        self.assertArray(np.array([1,1,1]), neuron.step_function(np.array([1,1,1])))
        self.assertArray(np.array([1,0,1]), neuron.step_function(np.array([1,0,1])))
        self.assertArray(np.array([0,0,1]), neuron.step_function(np.array([-1,0,1.4])))


if __name__ == "__main__":
    ut.main()
