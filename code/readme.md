# Functional Code Directory

This folder contains the main simulation application code for your baseball simulation project.

## What to Include

- Main simulation scripts and modules  
- Model training and evaluation code (if part of the app)  
- Utility functions and helpers **used directly by the simulation app**  
- Unit and integration tests for the simulation code  
- Any other code used to run or manage the simulation engine

> **Note:**  
> Data preparation, cleaning scripts, and raw/processed data are stored in the `internal_modules/` and `data/` folders at the project root, not here.

## Suggested Structure

Organize your code in a way that makes it easy to understand and run. Here are some examples based on language:

### Python Example
```
code/
  simulation_engine.py
  model_training.py
  utils/
    __init__.py
    helpers.py
  tests/
    test_simulation.py
```

### R Example
```
code/
  simulation_engine.R
  model_training.R
  utils/
    helpers.R
  tests/
    test_simulation.R
```

### Go Example
```
code/
  main.go
  simulation/
    engine.go
    stats.go
  tests/
    simulation_test.go
```

**Notes**  
Use subfolders (e.g., utils/, tests/) to keep your code organized.  
Include a README or comments in your scripts to explain how to run your code and what each file does.

**Author:** Andrew D'Amico  
**Email:** Andrew.Damico@u.northwestern.edu

