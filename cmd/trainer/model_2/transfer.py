import tensorflow as tf
from model import Model
import sys

m = Model()
m.transfer(sys.argv[1], sys.argv[2])
