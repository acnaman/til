import numpy as np

# 引数x：numpy.array()型
def step_function(x):
    y = x>0 
    return y.astype(np.int)

# 引数x：numpy.array()型
def sigmoid(x):
    return 1 / (1 + np.exp(-x))

def relu(x):
    return np.maximum(0,x)

# 他の実装と合わせるため、恒等関数を定義
def identity_function(x):
    return x

def softmax(a):
    c = np.max(a)

    exp_a = np.exp(a - c) # オーバーフロー対策。（softmax関数は全ての要素から同じ値を足し引きしても結果が変わらない）

    sum_exp_a = np.sum(exp_a)
    y = exp_a / sum_exp_a

    return y

# 2乗和誤差
def mean_squared_error(y, t):
    return 0.5 * np.sum((y-t)**2)

# 交差エントロピー誤差
def cross_entropy_error(y,t):
    if y.ndim == 1:
        t = t.reshape(1, t.size)
        y = y.reshape(1, y.size)

    delta = 1e-7
    batch_size = y.shape[0]
    return -np.sum(t * np.log(y - delta) / batch_size)


