# Project Name
Code-and-share Backend

# Preparation Notes
## Define the problem
I need a tool that helps quickly modify datasets for the datafeed.
  
I need two things:
- Isolation to avoid using root user
- Quick visual ways to modify data without MySQL knowledge
# Solution Brainstorming
## Solution Chosen
* Backend written in HTML/CSS/JS(Using Bootstrap4), iterating with a dedicated Go Backend

# Related stakeholders
* Github user angelalonso
* Github user gamstc
# Competitors
TBD
# Objectives
## Specific
* I want to be able to quickly list objects, phases and paths :ok:
  * Main menu to access lists :ok:
  * Upper navigation bar used everywhere :ok:
  * List object :ok:
  * List phase :ok:
  * List path :ok:
* I want to be able to quickly add objects, phases and paths
  * Add possibility to save to DB a new object :ok:
  * Add object :ok:
  * Add phase :ok:
  * Force format -> [{"object": "forest001", "position": "1"}, {"object": "mountain001", "position": "2"}, {"object": "rain001", "position": "3"}, {"object": "beach001", "position": "4"}]
  * Get objects for phase from actual objects
  * Add Phase ID to paths_new :ok:
  * Add path :ok:
  * Get phases for path from actual phases
* I want to be able to quickly view objects, phases and paths
  * View object :ok:
  * View phase :ok:
  * Get objects for phase from actual objects, link to see them
  * View path :ok:
  * Get phases for path from actual phases, link to see them
* I want to be able to quickly edit objects, phases and paths
  * Get ids too, for saving later :ok:
  * Edit object :ok:
  * Edit phase :ok:
  * Get objects for phase from actual objects
  * Edit path :ok:
  * Get phases for path from actual phases
* I want to be able to quickly insert objects, phases and paths
  * Insert object :ok:
  * Make it possible to upload images to shared/open volume on the frontend side (or manage images from here)
  * Insert phase :ok:
  * Get objects for phase from actual objects
  * Insert path :ok:
  * Get phases for path from actual phases
* I want to be able to quickly delete objects, phases and paths
  * Delete object :ok:
  * Delete phase :ok:
  * Show objects that still exist but are no longer in phase
  * Delete path :ok:
  * Show phases that still exist but are no longer in path
* Build tests
* Containerize app

# Possible trade-offs between Quality – Time – Cost
TBD

# Estimate project resources
TBD






