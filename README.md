# coup

To run this project:

1) Install Go. Homebrew is your friend :)
2) Clone this repo or run `go get github.com/andrewcopp/coup` to download
3) Create a virtualenv
4) Install TensorFlow in that virtualenv
5) Activate the virtualenv
3) Navigate to the top level directory of this project
4) Run `go cmd/trainer/trainer.go`

This will run the script that simulates games and trains and fits against the TensorFlow model.

The script is super slow unless you modify TensorFlow to run on a GPU and put it on an AWS EC2 instance. This is mostly to prove it works. 

All Python code can be found in `cmd/trainer/model`

This project was originally written in Go to help with concurrency. Concurrency was not
added in order to finish by the project deadline but will be added in my next iteration.
