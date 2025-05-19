**Author:** Andrew D'Amico
**Email:** Andrew.Damico@u.northwestern.edu
# Data Prepper Module

This tool processes raw Baseball-Reference CSV files for the Cubs and White Sox, cleaning and normalizing the data for use in the simulation engine.

## Features

- Extracts and separates handedness from the `Player` or `B` column.
- Normalizes player names (removes handedness symbols).
- Normalizes position columns, adds position flexibility and boolean columns for each position.
- Handles roster and schedule files.
- Drops stat family/sub-header rows from all outputs.
- Outputs cleaned CSVs to the `data/prepared_data/` directory, using a consistent naming convention.

## Usage

From the project root, run:

```sh
go run ./internal_modules/data_prepper/data_prepper.go
```

## How It Works

- Reads each raw CSV (e.g., `standard_batting.csv`, `40_man_roster.csv`, `full_season_roster.csv`, `2025_schedule.csv`).
- For each row:
  - Extracts handedness (`L`, `R`, `S`) from the `Player` field or `B` column.
  - Removes handedness symbol from the player name.
  - Normalizes position columns and validates numeric columns.
  - Adds position flexibility and boolean columns where applicable.
  - **Drops any stat family/sub-header rows (e.g., `,,,Standard,Standard,...`). Only the real column header and data rows are retained in all outputs.**
- Writes a cleaned CSV to `data/clean/`, named like `cubs_standard_batting_clean.csv`.
- Writes a team totals CSV (if present) to `data/clean/`, named like `cubs_standard_batting_team_totals_clean.csv`, with only the real header and the team totals row.

## Output Example

| Player           | Handedness | ...other columns... | EligiblePositions | PositionFlexibility | Is1B | Is2B | ... |
|------------------|------------|--------------------|-------------------|---------------------|------|------|-----|
| Michael Busch    | L          | ...                | 1B,3B             | 2                   | 1    | 0    | ... |
| Ian Happ         | S          | ...                | LF                | 1                   | 0    | 0    | ... |
| Nico Hoerner     | R          | ...                | 2B                | 1                   | 0    | 1    | ... |

## Next Steps

- [ ] Process and normalize team schedule CSVs for simulation scheduling logic.
- [ ] Validate all cleaned files for Go struct mapping in the simulation engine.

## Requirements

- Go 1.18+
- No external dependencies for basic CSV processing.

---
