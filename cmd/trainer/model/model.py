import tensorflow as tf
import os

os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'

# The model is responsible for all of the TensorFlow logic.
class Model:

    def __init__(self):
        self.epsilon = 0.01

    def initialize(self, outfile):

        tf.reset_default_graph()

        n_inputs = 301
        n_outputs = 1
        n_hidden_layer = 256

        with tf.device('/cpu:0'):
            weights = {
                'hidden_layer': tf.Variable(tf.truncated_normal([n_inputs, n_hidden_layer])),
                'out': tf.Variable(tf.truncated_normal([n_hidden_layer, n_outputs]))
            }
            biases = {
                'hidden_layer': tf.Variable(tf.zeros([n_hidden_layer])),
                'out': tf.Variable(tf.zeros([n_outputs]))
            }

        init = tf.global_variables_initializer()

        saver = tf.train.Saver()

        with tf.Session() as sess:
            sess.run(init)

            saver.save(sess, outfile)

    def transfer(self, infile, outfile):

        n_inputs = 301
        n_outputs = 1
        n_hidden_layer = 256

        with tf.device('/cpu:0'):
            weights = {
                'hidden_layer': tf.Variable(tf.truncated_normal([n_inputs, n_hidden_layer])),
                'out': tf.Variable(tf.truncated_normal([n_hidden_layer, n_outputs]))
            }
            biases = {
                'hidden_layer': tf.Variable(tf.zeros([n_hidden_layer])),
                'out': tf.Variable(tf.zeros([n_outputs]))
            }

        saver = tf.train.Saver()

        with tf.Session() as sess:
            saver.restore(sess, infile)

            saver.save(sess, outfile)

    def train(self, infile, outfile, inputs, outputs):

        tf.reset_default_graph()

        learning_rate = 0.01
        n_inputs = 301
        n_outputs = 1
        n_hidden_layer = 256

        with tf.device('/cpu:0'):
            features = tf.placeholder(tf.float32, [None, n_inputs])
            labels = tf.placeholder(tf.float32, [None, n_outputs])

            weights = {
                'hidden_layer': tf.Variable(tf.truncated_normal([n_inputs, n_hidden_layer])),
                'out': tf.Variable(tf.truncated_normal([n_hidden_layer, n_outputs]))
            }
            biases = {
                'hidden_layer': tf.Variable(tf.zeros([n_hidden_layer])),
                'out': tf.Variable(tf.zeros([n_outputs]))
            }

            layer_1 = tf.add(tf.matmul(features, weights['hidden_layer']), biases['hidden_layer'])
            layer_1 = tf.nn.relu(layer_1)

            logits = tf.add(tf.matmul(layer_1, weights['out']), biases['out'])
            logits = tf.nn.tanh(logits)

            cost = tf.reduce_sum(tf.pow(logits-labels, 2))/(2*1)
            optimizer = tf.train.GradientDescentOptimizer(learning_rate=learning_rate).minimize(cost)

        saver = tf.train.Saver()

        with tf.Session() as sess:
            saver.restore(sess, infile)

            sess.run(optimizer, feed_dict={features: inputs, labels: outputs})

            saver.save(sess, outfile)

    def fit(self, infile, inputs):

        tf.reset_default_graph()

        n_inputs = 301
        n_outputs = 1
        n_hidden_layer = 256

        with tf.device('/cpu:0'):
            features = tf.placeholder(tf.float32, [None, n_inputs])
            labels = tf.placeholder(tf.float32, [None, n_outputs])

            weights = {
                'hidden_layer': tf.Variable(tf.truncated_normal([n_inputs, n_hidden_layer])),
                'out': tf.Variable(tf.truncated_normal([n_hidden_layer, n_outputs]))
            }
            biases = {
                'hidden_layer': tf.Variable(tf.zeros([n_hidden_layer])),
                'out': tf.Variable(tf.zeros([n_outputs]))
            }

            layer_1 = tf.add(tf.matmul(features, weights['hidden_layer']), biases['hidden_layer'])
            layer_1 = tf.nn.relu(layer_1)

            logits = tf.add(tf.matmul(layer_1, weights['out']), biases['out'])
            logits = tf.nn.tanh(logits)

        saver = tf.train.Saver()

        with tf.Session() as sess:
            saver.restore(sess, infile)

            prediction = sess.run(logits, feed_dict={features: inputs})
            return prediction


# layout - check
# communication
# saving - check
# data
# model
