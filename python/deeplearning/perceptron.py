def AND(x1, x2):
    w1 = 0.5
    w2 = 0.5

    y = x1 * w1 + x2 * w2
    if y > 0.7:
        return 1

    else:
        return 0

def OR(x1, x2):
    w1 = 0.5
    w2 = 0.5

    y = x1 * w1 + x2 * w2
    if y > 0:
        return 1

    else:
        return 0



