# Activity List - Example

This is a sample activity list for a Baseball Simulator project.

- May 19 09:00–10:00: Reviewed AI-generated EDA code (Andrew).
- May 19 10:00–11:00: Wrote missing test cases (Sam).
- May 20 13:00–14:00: Updated documentation (Priya).

## Backlog Checklist

### Open Items

#### Initialization & Data Loading
- [ ] **BL-001**: Initialize game state with teams, lineups, and starting pitchers
- [ ] **BL-002**: Load basic player stats for teams
- [ ] **BL-003**: Error handling for missing/corrupt player data

#### Simulation Engine
- [ ] **BL-004**: Simulate individual pitch outcomes based on player stats
- [ ] **BL-005**: Simulate full at-bats pitch-by-pitch
- [ ] **BL-006**: Update game state after every at-bat
- [ ] **BL-007**: Simulate full 9-inning games
- [ ] **BL-008**: Simulate multiple games in a series and aggregate outcomes
- [ ] **BL-009**: Save and restore simulation state

#### Logging & Analysis
- [ ] **BL-010**: Record detailed pitch logs for each simulation
- [ ] **BL-011**: Scrollable pitch log with play-by-play text

#### User Interface
- [ ] **BL-012**: Graphical baseball diamond with current base runners
- [ ] **BL-013**: Real-time scoreboard (inning, outs, score, balls, strikes)
- [ ] **BL-014**: Play, Pause, Fast Forward controls
- [ ] **BL-015**: Adjustable pitch delay speed with slider
- [ ] **BL-016**: Simulation pauses instantly when requested

#### Scheduling & Planning
- [ ] **BL-017**: Load upcoming two-week schedule for teams

#### Testing & Validation
- [ ] **BL-018**: Unit tests for pitch simulation and game state updates
- [ ] **BL-019**: Validation on user inputs and loaded data
- [ ] **BL-020**: Clear error notifications for schedule/data loading failures

#### Advanced Simulation Features
- [ ] **BL-021**: Batter vs pitcher splits for lineup recommendations
- [ ] **BL-022**: Model pitcher fatigue for rotation recommendations
- [ ] **BL-023**: Integrate salary and contract data for lineup/rotation constraints
- [ ] **BL-024**: Factor player injury histories/statuses into availability
- [ ] **BL-025**: Include pitch types and velocities in simulation
- [ ] **BL-026**: Pitch velocity influences pitch effectiveness

#### Project Infrastructure & Documentation
- [ ] **BL-027**: Set up Go module and workspace structure
- [ ] **BL-028**: Set up Hugo documentation site with code mounts
- [ ] **BL-029**: Add CI/CD pipeline for build and test automation
- [ ] **BL-030**: Write developer onboarding and contribution guide
- [ ] **BL-031**: Set up code review and issue tracking process

---

### Completed Items

#### Data Preparation & Normalization
- [x] **BL-032**: Implement data_prepper module for CSV cleaning/normalization
- [x] **BL-033**: Extract and standardize handedness and position columns
- [x] **BL-034**: Ensure all stats are numeric and player IDs are consistent
- [x] **BL-035**: Output clean, ready-to-load CSVs for simulation
- [x] **BL-036**: Add position flexibility metric and boolean columns for each position
- [x] **BL-038**: Remove stat family/sub-header rows from all cleaned and team totals CSVs

---