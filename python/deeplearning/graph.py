import numpy as np
import matplotlib.pylab as plt
import neuron

x = np.arange(-5.0, 5.0, 0.1)
y = neuron.sigmoid(x)

plt.plot(x,y)
plt.ylim(-0.1, 1.1)
plt.show()

