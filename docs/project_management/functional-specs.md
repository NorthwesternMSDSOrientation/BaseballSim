# Baseball Simulator - Functional Specifications

## Overview

This document details the core functions for the baseball simulator project, divided into MVP iterations. It includes purpose, inputs, outputs, and key notes for each function to guide implementation.

---

## MVP 1 - Core Pitch-by-Pitch Simulation

### InitializeGame(teamA, teamB)

- **Purpose:** Initialize the game with teams, lineups, and starting pitchers.
- **Inputs:**  
  - `teamA`: Team object  
  - `teamB`: Team object
- **Outputs:**  
  - `gameState`: Initialized game state object
- **Notes:** Resets scores, inning, outs, and bases.

---

### LoadPlayerStats(team)

- **Purpose:** Load player stats required for simulation.
- **Inputs:**  
  - `team`: Team identifier or object
- **Outputs:**  
  - Player stats dictionary (batting and pitching)
- **Notes:** MVP1 uses basic stats (AVG, OBP, ERA, K%, BB%).

---

### SimulatePitch(pitcherStats, batterStats, count)

- **Purpose:** Simulate a single pitch outcome probabilistically.
- **Inputs:**  
  - `pitcherStats`  
  - `batterStats`  
  - `count`: Ball-strike count
- **Outputs:**  
  - `pitchOutcome` (e.g., strike, ball, foul, hit, out)
- **Notes:** Combines pitcher control and batter discipline.

---

### SimulateAtBat(pitcherStats, batterStats)

- **Purpose:** Simulate full at-bat until result (walk, out, hit).
- **Inputs:**  
  - `pitcherStats`  
  - `batterStats`
- **Outputs:**  
  - `atBatOutcome` (walk, strikeout, single, double, out, etc.)
- **Notes:** Loops pitch simulation until at-bat ends.

---

### UpdateGameState(gameState, atBatOutcome)

- **Purpose:** Update the game state based on at-bat result.
- **Inputs:**  
  - `gameState`  
  - `atBatOutcome`
- **Outputs:**  
  - Updated `gameState`
- **Notes:** Advances runners, updates outs, scores, innings.

---

### SimulateGame(teamA, teamB)

- **Purpose:** Run a full 9-inning game simulation.
- **Inputs:**  
  - `teamA`  
  - `teamB`
- **Outputs:**  
  - Final `gameState`
- **Notes:** Alternates half innings until game completion.

---

## MVP 2 - UI and Playback Controls

### DrawBaseballDiamond(gameState)

- **Purpose:** Render baseball diamond with bases and runners.
- **Inputs:**  
  - `gameState`
- **Outputs:**  
  - Visual display of diamond
- **Notes:** Updated after each pitch or at-bat.

---

### UpdateScoreboard(gameState)

- **Purpose:** Show inning, outs, score, count.
- **Inputs:**  
  - `gameState`
- **Outputs:**  
  - Scoreboard UI element

---

### PlaySimulation(pitchDelay)

- **Purpose:** Run simulation with delay between pitches.
- **Inputs:**  
  - `pitchDelay` (seconds)
- **Outputs:**  
  - Simulation running with UI updates.

---

### PauseSimulation()

- **Purpose:** Pause simulation execution.
- **Inputs:** None
- **Outputs:** Simulation halted.

---

### FastForwardSimulation()

- **Purpose:** Run simulation as fast as possible.
- **Inputs:** None
- **Outputs:** Rapid simulation run.

---

## MVP 3 - Advanced Modeling

### GetBatterVsPitcherSplits(batterID, pitcherID)

- **Purpose:** Retrieve batter vs. pitcher matchup stats.
- **Inputs:**  
  - `batterID`  
  - `pitcherID`
- **Outputs:**  
  - Matchup statistics

---

### CalculateFatigue(pitcherID, inningsPitched)

- **Purpose:** Model pitcher fatigue and performance drop.
- **Inputs:**  
  - `pitcherID`  
  - `inningsPitched`
- **Outputs:**  
  - Fatigue factor

---

### SimulatePitchSelection(pitcherStats, batterStats, count)

- **Purpose:** Choose pitch type and location based on stats.
- **Inputs:**  
  - `pitcherStats`  
  - `batterStats`  
  - `count`
- **Outputs:**  
  - Selected pitch and outcome probabilities.

---

### OptimizeLineup(schedule, salaries, stats)

- **Purpose:** Recommend optimal lineup given schedule and payroll.
- **Inputs:**  
  - Upcoming schedule  
  - Player salaries  
  - Player stats
- **Outputs:**  
  - Recommended batting order and lineup

---

## MVP 4 - Data Integration & Automation

- All cleaned and team totals CSVs must contain only the real column header and data rows.  
- Stat family/sub-header rows (e.g., `,,,Standard,Standard,...`) are dropped during data preparation.

---

### FetchLatestPlayerStats(teamName)

- **Purpose:** Automatically fetch updated stats from APIs.
- **Inputs:**  
  - `teamName`
- **Outputs:**  
  - Updated player stats dataset

---

### SaveSimulationResults(filename, data)

- **Purpose:** Save simulation outputs for analysis.
- **Inputs:**  
  - `filename`  
  - `data`
- **Outputs:**  
  - Saved file on disk or cloud.

---

### LoadSimulationScenario(filename)

- **Purpose:** Load saved simulation scenario.
- **Inputs:**  
  - `filename`
- **Outputs:**  
  - Loaded scenario for replay or modification.

---

# End of Functional Specifications
