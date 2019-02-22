import numpy as np

# 引数x：numpy.array()型
def step_function(x):
    y = x>0 
    return y.astype(np.int)

# 引数x：numpy.array()型
def sigmoid(x):
    return 1 / (1 + np.exp(-x))
