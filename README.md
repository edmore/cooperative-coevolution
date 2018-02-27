[![Build Status](https://travis-ci.com/edmore/cooperative-coevolution.svg?token=qCqiUCDFN1395pnZuyJY&branch=master)](https://magnum.travis-ci.com/edmore/cooperative-coevolution)

Cooperative Co-Evolution in Go(lang)

This work was done in fulfillment of a _Masters Degree in Computer Science_ at the University of Cape Town.

Dissertation: **_Accelerated Cooperative Co-Evolution on Multi-core Architectures_**.

Abstract:

``
The Cooperative Co-Evolution model has been used in Evolutionary Computation to optimize
the training of artificial neural networks (ANNs). This architecture has proven to be
a useful extension to domains such as Neuro-Evolution (NE), which is the training of ANNs
using concepts of natural evolution. However, there is a need for real-time systems and the
ability to solve more complex tasks which has prompted a further need to optimize these
Cooperative Co-Evolution methods. Cooperative Co-Evolution methods consist of a number
of phases, however the evaluation phase is still the most compute intensive phase, for some
complex tasks taking as long as weeks to complete. This study uses NE as a test case study
and we design a parallel Cooperative Co-Evolution processing framework and implement the
optimized serial and parallel versions using the Golang (Go) programming language. Go is a
multi-core programming language with first-class constructs, channels and goroutines, that
make it well suited to parallel programming. Our study focuses on Enforced Subpopulations
(ESP) for single-agent systems and Multi-Agent ESP for multi-agent systems. We evaluate
the parallel versions in the benchmark tasks; double pole balancing and prey-capture, for
single and multi-agent systems respectively, in tasks of increasing complexity. We observe a
maximum speed-up of 20x for the parallel Multi-Agent ESP implementation over our single
core optimized version in the prey-capture task and a maximum speedup of 16x for ESP in
the harder version of double pole balancing task. We also observe linear speed-ups for the
difficult versions of the tasks for a certain range of cores, indicating that the Go implementations
are efficient and that the parallel speed-ups are better for more complex tasks. We
find that in complex tasks, the Cooperative Co-Evolution Neuro-Evolution (CCNE) methods
are amenable to multi-core acceleration, which provides a basis for the study of even
more complex Cooperative Co-Evolution methods in a wider range of domains.
``

The author of this work is Edmore T. Moyo.

Each implementation is in its own branch, this repo has the following implementations:

- master - the master branch is the blue-print algorithm, based on the Enforced Subpopulations (ESP) algorithm. It can be used as a starting point for the parallelisation of other Cooperative Co-Evolution methods.
- **esp-serial** - the serial implementation of the ESP method in the double pole balancing task (supports Markov and non-Markov versions).
- **esp-parallel** - the _parallel_ implementation of the ESP method in the double pole balancing task  (supports Markov and non-Markov versions).
- **sane-serial** - the serial implementation of our neuron Symbiotic, Adaptive Neuro-Evolution (SANE) method in the double pole balancing task.
- **sane-parallel** - the _parallel_ implementation of our neuron SANE method in the double pole balancing task.
-  **multi-agent-esp-serial** - the serial implementation of the Multi-Agent ESP method in the prey-capture task (prey starts at one starting position [50,50]).
- **multi-agent-esp-serial-custom** - the serial implementation of the Multi-Agent ESP method in the prey-capture task (prey starts at nine different defined positions for each trial).
- **multi-agent-esp-parallel** - the _parallel_ implementation of the Multi-Agent ESP method in the prey-capture task (preys start at nine different defined positions for each trial).

Validation visualizations are available for download: https://www.dropbox.com/s/df5l30p1g9kz1aq/Validation_visualizations.zip
