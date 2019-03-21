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

# 微分(丸め誤差考慮)
def numerical_diff(f, x):
    h = 1e-4
    return (f(x+h)-f(x-h)) / (2*h)

# 偏微分
def numrical_gradient(f, x):
    h = 1e-4
    grad = np.zeros_like(x)

    for idx in range(x.size):
        tmp_val = x[idx]
        fxh1 = f(tmp_val + h)

        fxh2 = f(tmp_val - h)

        grad[idx] = (fxh1 - fxh2) / (2*h)

    return grad

# 勾配降下法
def gradient_descent(f, init_x, lr=0.01, step_num=100 ):
    x = init_x
    for i in range(step_num):
        grad = numrical_gradient(f,x)
        x -= lr + grad

    return
