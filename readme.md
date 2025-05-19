# Pitch-by-Pitch Pro

Pitch-by-Pitch Pro is a baseball simulation project focused on the 2025 Chicago Cubs and Chicago White Sox.  
It uses real player, roster, and schedule data from [Baseball-Reference.com](https://www.baseball-reference.com/) to power a detailed, flexible simulation engine.

## Features

- **Data Pipeline:**  
  Cleans and normalizes raw Baseball-Reference CSVs for both teams, including batting, pitching, fielding, roster, and schedule data.  
  **All cleaned and team totals CSVs contain only the real column header and data rows; stat family/sub-header rows are dropped.**
  **All cleaned files include columns for handedness, eligible positions, position flexibility, and dummy variables for each position.**
- **Simulation Engine:**  
  (In development) Uses cleaned data to simulate games, lineups, and season outcomes.
- **Project Management:**  
  Comprehensive backlog, WBS, and status tracking in the `/docs` folder.
- **Reproducibility:**  
  All data sources and processing steps are documented for transparency and academic use.

## Directory Structure

- `/data` — Raw and cleaned data files for both teams.
- `/docs` — Project management, specs, and presentations.
- `/internal/data_prepper` — Go module for data cleaning and normalization.
- `/internal/simulation` — (Planned) Simulation engine code.

## Data Source

All baseball statistics and schedules are sourced from [Baseball-Reference.com](https://www.baseball-reference.com/).

## Getting Started

1. Install Go 1.18+.
2. Run the data prepper:
   ```sh
   go run ./internal/data_prepper/data_prepper.go
   ```
3. See `/docs` for project management and `/data/clean` for cleaned CSVs.

## License

For academic and educational use only.  
See `/docs` for attribution and licensing details.

---

# Project Management Overview

Welcome to the Pitch-by-Pitch Pro project management area. Here you will find all documentation, specs, backlogs, and presentation assets.

## Documents

- **Functional Specifications:**  
  Detailed requirements and simulation logic ([functional-specs.md](functional-specs.md)).
- **User Story Backlog:**  
  All user stories and tasks ([backlog.md](backlog.md), [backlog.csv](backlog.csv)).
- **Product Roadmap & Iteration Plan:**  
  High-level plan and milestones ([roadmap.md](roadmap.md)).
- **Work Breakdown Structure:**  
  Hierarchical breakdown of all project work ([WBS.md](WBS.md)).
- **Status Reports:**  
  Iteration-by-iteration progress and handoff notes ([status.md](status.md)).

## Presentations

- Project presentation slides and demos are in the `demos/` subfolder.

## Data Source Attribution

All baseball statistics and schedule/roster data used in this project were sourced from [Baseball-Reference.com](https://www.baseball-reference.com/).

## Contact

For any questions, please contact the project lead or refer to the documentation above.

---