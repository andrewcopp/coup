import tensorflow as tf
from model import Model
import sys

m = Model()

features = [e.split(',') for e in sys.argv[2].split('_')]

labels = m.fit(sys.argv[1], features)
sys.stdout.write(','.join(str(e[0]) for e in labels))
