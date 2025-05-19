# Backlog Management Process

When adding items to the backlog, follow this process:

1. **Write a User Story:**  
   - Use the format: *As a [persona], I would like [goal] so that [reason].*
2. **Link to Artifacts:**  
   - Reference related items in [functional-specs.md](functional-specs.md), [roadmap.md](roadmap.md), and [wbs.md](wbs.md).
3. **Fill in All Columns:**  
   - Include WBS, Sprint, Story Points, Persona, Progress, and the User Story.
4. **Check Off When Complete:**  
   - Use the checkbox at the start of each item. Mark as `[x]` when done.
5. **Keep Items Linked:**  
   - Each checklist item has a unique ID, which matches the detailed table below.

---

## Backlog Intro

## Backlog

---

## Detailed Backlog Table

| ID      | WBS    | Sprint | Story Points | Persona            | Progress | User Story                                                                                                    | Links to Artifacts                |
|---------|--------|--------|-------------|--------------------|----------|---------------------------------------------------------------------------------------------------------------|------------------------------------|
| BL-001  | 2.1.1  | 1      | 8           | Product Owner      | new      | As a Product Owner, I would like the system to initialize the game state with teams, lineups, and starting pitchers so that gameplay starts correctly. | #file:functional-specs, #file:roadmap, #file:wbs |
| BL-002  | 2.1.2  | 1      | 5           | Developer          | new      | As a Developer, I would like to load basic player stats (batting average, ERA, strikeout rate) for teams so that simulations use realistic data. | #file:functional-specs, #file:wbs |
| BL-003  | 2.1.2  | 1      | 5           | Developer          | new      | As a Developer, I would like error handling for missing or corrupt player data so that the simulator is robust. | #file:functional-specs, #file:wbs |
| BL-004  | 2.1.3  | 1      | 8           | Baseball Fan       | new      | As a Baseball Fan, I would like to simulate individual pitch outcomes based on player stats so that at-bats feel realistic. | #file:functional-specs, #file:wbs |
| BL-005  | 2.1.4  | 1      | 8           | Developer          | new      | As a Developer, I would like to simulate full at-bats pitch-by-pitch until a result is reached so that the game flow is realistic. | #file:functional-specs, #file:wbs |
| BL-006  | 2.1.5  | 1      | 13          | Product Owner      | new      | As a Product Owner, I would like the game state to update properly after every at-bat so that outs, runs, and base runners are tracked accurately. | #file:functional-specs, #file:wbs |
| BL-007  | 2.1.6  | 1      | 13          | Baseball Fan       | new      | As a Baseball Fan, I would like to simulate full 9-inning games between teams so that the results are realistic. | #file:functional-specs, #file:wbs |
| BL-008  | 2.2    | 1      | 8           | Product Owner      | new      | As a Product Owner, I would like to simulate multiple games in a series and aggregate outcomes so that I can predict series winners. | #file:functional-specs, #file:roadmap, #file:wbs |
| BL-009  | 2.2    | 2      | 8           | Developer          | new      | As a Developer, I would like to save and restore simulation state so that users can resume sessions. | #file:functional-specs, #file:roadmap, #file:wbs |
| BL-010  | 2.7.4  | 1      | 3           | Product Owner      | new      | As a Product Owner, I would like detailed pitch logs recorded for each simulation so that I can perform post-game analysis. | #file:functional-specs, #file:wbs |
| BL-011  | 2.3.4  | 2      | 8           | User               | new      | As a User, I would like a scrollable pitch log displaying detailed play-by-play text so that I can review each play. | #file:functional-specs, #file:wbs |
| BL-012  | 2.3.1  | 2      | 8           | User               | new      | As a User, I would like to view a graphical baseball diamond with current base runners so that I can easily track game progress. | #file:functional-specs, #file:wbs |
| BL-013  | 2.3.2  | 2      | 8           | User               | new      | As a User, I would like a real-time scoreboard displaying inning, outs, score, balls, and strikes so that I have full context. | #file:functional-specs, #file:wbs |
| BL-014  | 2.3.3  | 2      | 5           | User               | new      | As a User, I would like Play, Pause, and Fast Forward controls so that I can watch simulations at my preferred speed. | #file:functional-specs, #file:wbs |
| BL-015  | 2.3.3  | 2      | 5           | User               | new      | As a User, I would like to adjust the pitch delay speed with a slider so that I can customize playback pacing. | #file:functional-specs, #file:wbs |
| BL-016  | 2.3.3  | 2      | 8           | Product Owner      | new      | As a Product Owner, I would like simulation to pause instantly when requested so that users have control. | #file:functional-specs, #file:wbs |
| BL-017  | 2.6.2  | 2      | 8           | Developer          | new      | As a Developer, I would like to load the upcoming two-week schedule for teams so that planning features are enabled. | #file:functional-specs, #file:wbs |
| BL-018  | 2.7.1  | 1      | 3           | QA Tester          | new      | As a QA Tester, I would like unit tests covering pitch simulation and game state updates so that reliability is ensured. | #file:functional-specs, #file:wbs |
| BL-019  | 2.7.2  | 2      | 3           | Developer          | new      | As a Developer, I would like validation on user inputs and loaded data so that simulation errors are prevented. | #file:functional-specs, #file:wbs |
| BL-020  | 2.7.3  | 2      | 5           | Product Owner      | new      | As a Product Owner, I would like clear error notifications for schedule or data loading failures so that issues are obvious. | #file:functional-specs, #file:wbs |
| BL-021  | 2.4.1  | 3      | 13          | Manager            | new      | As a Manager, I would like batter vs pitcher splits incorporated so that lineup recommendations are improved. | #file:functional-specs, #file:wbs |
| BL-022  | 2.4.2  | 3      | 13          | Manager            | new      | As a Manager, I would like to model pitcher fatigue so that effective rotations are recommended and injury risk is reduced. | #file:functional-specs, #file:wbs |
| BL-023  | 2.6.3  | 3      | 13          | Business Analyst   | new      | As a Business Analyst, I would like salary and contract data integrated so that lineup and rotation recommendations are constrained. | #file:functional-specs, #file:wbs |
| BL-024  | 2.4.2  | 3      | 13          | Business Analyst   | new      | As a Business Analyst, I would like player injury histories and current injury statuses factored into availability so that lineup decisions are accurate. | #file:functional-specs, #file:wbs |
| BL-025  | 2.4.3  | 3      | 8           | User               | new      | As a User, I would like simulation to include pitch types and velocities so that pitching is more realistic. | #file:functional-specs, #file:wbs |
| BL-026  | 2.4.3  | 3      | 8           | User               | new      | As a User, I would like pitch velocity to influence pitch effectiveness dynamically so that outcomes are more realistic. | #file:functional-specs, #file:wbs |
| BL-027  | INF-1  | 0      | 3           | Developer          | new      | As a Developer, I would like to set up the Go module and workspace structure so that the project is organized. | #file:readme, #file:status        |
| BL-028  | INF-2  | 0      | 3           | Developer          | new      | As a Developer, I would like to set up a Hugo documentation site with code mounts so that documentation is easily accessible. | #file:readme, #file:status        |
| BL-029  | INF-3  | 0      | 5           | Developer          | new      | As a Developer, I would like to add a CI/CD pipeline for build and test automation so that code quality is maintained. | #file:readme, #file:status        |
| BL-030  | INF-4  | 0      | 3           | Developer          | new      | As a Developer, I would like to write a developer onboarding and contribution guide so that new contributors can get started quickly. | #file:readme, #file:status        |
| BL-031  | INF-5  | 0      | 3           | Developer          | new      | As a Developer, I would like to set up code review and issue tracking process so that collaboration is streamlined. | #file:readme, #file:status        |
| BL-032  | 2.6.1  | 1      | 8           | Developer          | done     | As a Developer, I would like to implement a data_prepper module to clean and normalize all player CSVs for simulation use so that the data is ready for the engine. | #file:encyclopedia, #file:WBS     |
| BL-033  | 2.6.1  | 1      | 5           | Developer          | done     | As a Developer, I would like to extract and standardize handedness and position columns in all player data tables so that the data is consistent. | #file:encyclopedia                |
| BL-034  | 2.6.1  | 1      | 5           | Developer          | done     | As a Developer, I would like to ensure all numeric/stat columns are properly typed and player IDs are consistent across tables so that data integrity is maintained. | #file:encyclopedia                |
| BL-035  | 2.6.1  | 1      | 3           | Developer          | done     | As a Developer, I would like to generate a clean, ready-to-load CSV for each player data type for the simulation engine so that loading is seamless. | #file:encyclopedia                |
| BL-036  | 2.6.1  | 1      | 5           | Developer          | done     | As a Developer, I would like to add a position flexibility metric and boolean columns for each position to support lineup logic so that the simulation can use this data. | #file:encyclopedia                |
| BL-037  | 2.6.1  | 2      | 5           | Developer          | new      | As a Developer, I would like to process and normalize team schedule CSVs so that the simulation engine can use real game schedules. | #file:encyclopedia                |
| BL-038  | 2.6.1  | 1      | 2           | Developer          | done     | As a Developer, I would like the data_prepper to remove stat family/sub-header rows from all cleaned and team totals CSVs so that only useful data remains. | #file:encyclopedia                |

---

> **Note:**  
> - IDs (e.g., BL-001) are used for cross-referencing between the checklist and the table.
> - Links to artifacts reference the relevant files for traceability.
> - Additional infrastructure and process items (BL-027 to BL-031) are included for project completeness.

# Product Backlog - Example

This is a sample backlog for a Baseball Simulator project.  
Replace these with your own user stories and priorities.

| ID      | Persona      | User Story                                                                 | Priority | Status |
|---------|--------------|----------------------------------------------------------------------------|----------|--------|
| BL-001  | User         | As a user, I want to upload a CSV so that the system can predict sales.    | High     | Open   |
| BL-002  | Developer    | As a developer, I want to validate input data so that errors are prevented.| Medium   | Open   |
| BL-003  | Product Owner| As a product owner, I want a dashboard so that I can view results easily.  | Low      | Open   |