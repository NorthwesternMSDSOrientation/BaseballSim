**Author:** Andrew D'Amico
**Email:** Andrew.Damico@u.northwestern.edu
# Baseball Simulation

This repository is a template and example for a baseball simulation project using real player, roster, and schedule data from [Baseball-Reference.com](https://www.baseball-reference.com/).  
It demonstrates how to organize your code, data, and project management artifacts for transparent, reproducible, and collaborative AI-assisted development.

---

## Assignment Prompt: AI Project Management & Collaboration

**Instructions:**  
Post the link to your group’s GitHub repo. Your repo should be organized as follows:

```
Root/  
  functional_code/  
  prepared_data/  
  documentation_and_project_management_artifacts/  
    Functional_Specs  
    Work_Breakdown_Structure  
    Product_Backlog  
    Status_Log  
    Activity_List  
    Roadmap  
  results/  
  supplementary_code/  
    raw_data_processing/  
```


### Why We Require AI Project-Management Artifacts

When you bring an AI “co-worker” into your project, you’re collaborating with a model whose decisions, suggestions, and outputs must be tracked, assessed, and iterated. These artifacts:

- **Ensure transparency:** See which parts the AI generated, what prompts you used, and how you reviewed or revised those outputs.
- **Facilitate reproducibility:** Future team members (or instructors!) can rerun prompts, validate results, and understand the evolution of your solution.
- **Support accountability:** Logging AI interactions helps you reflect on biases, errors, or creative breakthroughs introduced by the model.
- **Promote best practices:** Treating AI like a teammate encourages you to apply solid PM processes—requirements gathering, backlog grooming, sprint planning—even when the “developer” isn’t human.

### Building a Development Pipeline with an AI Co-Developer

- **Planning & Specification:**  
  Draft functional specs and user stories. Craft prompts to guide the AI in generating code, documentation, or tests.
- **Data Preparation:**  
  Collect and prepare datasets. Use AI to suggest cleaning steps, flag anomalies, or generate synthetic data.
- **Iterative Development:**  
  Break work into backlog items. For each: run prompt → review AI output → integrate/refine → commit.
- **Testing & Validation:**  
  Write tests. Use AI to generate test cases or mock data, then validate results manually.
- **Review & Retrospective:**  
  Log status updates and decisions. At the end of each sprint, review AI-generated code for quality and maintainability.

### Defining the Artifacts

| Artifact         | Description                                                                 | Example                                                                                           |
|------------------|-----------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| Functional Specs | Detailed features, user stories, and acceptance criteria                    | “As a user, I want to upload a CSV so that the system can predict sales. Acceptance: ...”         |
| WBS              | Hierarchical breakdown of all tasks                                         | 1. Data Ingestion<br>1.1 Download raw data<br>1.2 Validate schema                                 |
| Product Backlog  | Prioritized list of user stories, bugs, and technical tasks                 | 1. Implement CSV upload (priority: high)                                                          |
| Status Log       | Chronological record of progress, decisions, and issues                     | May 18: Prompted AI to write data-cleaning function; fixed edge-case in row filtering.            |
| Activity List    | Timestamped list of development activities with owners                      | “May 19 09:00–10:00: Reviewed AI-generated EDA code (Andrew).”                                   |
| Roadmap          | High-level timeline of milestones and deliverables                          | Q1 Week 1: Data pipeline complete. Q1 Week 2: Model MVP.                                          |

**Tip:** For each artifact, note explicitly which parts were produced by your AI co-developer (e.g., “Generated initial draft of Functional Specs via ChatGPT prompts on May 18”).

### Sample Repository Structure


GitHub Copilot
Root/  
  functional_code/  
  prepared_data/  
  documentation_and_project_management_artifacts/  
    Functional_Specs  
    Work_Breakdown_Structure  
    Product_Backlog  
    Status_Log  
    Activity_List  
    Roadmap  
  results/  
  supplementary_code/  
    raw_data_processing/


### Submission

- Post your GitHub repo link.
- In your submission, include a unique comment on working with an AI co-developer: what went well, what challenges you faced, and insights you gained from at least three other classmates’ simulations.

---

## Features

- **Data Pipeline:**  
  Cleans and normalizes raw Baseball-Reference CSVs for both teams, including batting, pitching, fielding, roster, and schedule data.  
  All cleaned and team totals CSVs contain only the real column header and data rows; stat family/sub-header rows are dropped.  
  All cleaned files include columns for handedness, eligible positions, position flexibility, and dummy variables for each position.

- **Simulation Engine:**  
  (In development) Uses cleaned data to simulate games, lineups, and season outcomes.

- **Project Management:**  
  Comprehensive backlog, WBS, and status tracking in the `/docs` folder.

- **Reproducibility:**  
  All data sources and processing steps are documented for transparency and academic use.

## Directory Structure

- `/data` — Raw and cleaned data files for both teams.
- `/docs` — Project management, specs, and presentations.
- `/internal_modules/data_prepper` — Go module for data cleaning and normalization.
- `/code/simulation` — (Planned) Simulation engine code.

## Data Source

All baseball statistics and schedules are sourced from [Baseball-Reference.com](https://www.baseball-reference.com/).

## Getting Started

1. Install Go 1.18+.
2. Run the data prepper:
   ```sh
   go run ./internal/data_prepper/data_prepper.go
   ```
3. See /docs for project management and /data/prepared_data for cleaned CSVs.

##License
For academic and educational use only.
See /docs for attribution and licensing details.


