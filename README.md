[![Build Status](https://travis-ci.com/edmore/cooperative-coevolution.svg?token=qCqiUCDFN1395pnZuyJY&branch=master)](https://magnum.travis-ci.com/edmore/cooperative-coevolution)

Cooperative Co-Evolution in Go(lang)

This work was done in fulfillment of a Masters Degree in Computer Science at the University of Cape Town.

The author of this work is Edmore T. Moyo

Each implementation is in its own branch, this repo has the following implementations:

- master - the master branch is the blue-print algorithm, based on the ESP algorithm
- esp-serial - the serial implementation of the ESP method in the double pole balancing task (supports Markov and non-Markov versions)
- esp-parallel - the parallel implementation of the ESP method in the double pole balancing task  (supports Markov and non-Markov versions)
- sane-serial - the serial implementation of our neuron SANE method in the double pole balancing task
- sane-parallel - the parallel implementation of our neuron SANE method in the double pole balancing task
-  multi-agent-esp-serial - the serial implementation of the Multi-Agent ESP method in the prey-capture task (prey starts at one starting position [50,50])
- multi-agent-esp-serial-custom - the serial implementation of the Multi-Agent ESP method in the prey-capture task (prey start at 9 different defined positions for each trial)
- multi-agent-esp-parallel - the parallel implementation of the Multi-Agent ESP method in the prey-capture task (prey start at 9 different defined positions for each trial)

Validation visualizations are available for download: https://www.dropbox.com/s/df5l30p1g9kz1aq/Validation_visualizations.zip
