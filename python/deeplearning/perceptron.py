def AND(x1, x2):
    w1 = 0.5
    w2 = 0.5
    theta = 0.7

    tmp = x1 * w1 + x2 * w2
    if tmp > theta:
        return 1

    else:
        return 0

def OR(x1, x2):
    w1 = 0.5
    w2 = 0.5
    theta = 0

    tmp = x1 * w1 + x2 * w2
    if tmp > theta:
        return 1

    else:
        return 0

def NAND(x1, x2):
    w1 = -0.5
    w2 = -0.5
    theta = -0.7

    tmp = x1 * w1 + x2 * w2
    if tmp > theta:
        return 1

    else:
        return 0



