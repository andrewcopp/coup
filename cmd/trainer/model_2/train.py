import tensorflow as tf
from model import Model
import sys

m = Model()

features = [e.split(',') for e in sys.argv[3].split('_')]
labels = [e.split(',') for e in sys.argv[4].split('_')]

m.train(sys.argv[1], sys.argv[2], features, labels)
