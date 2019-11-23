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
* I want to store the images on this container, make it available for the frontend to use :ok:
  * Add possibility to upload a file :ok:
  * Load images into backend :x:
  * Change frontend to get images from this backend :ok:
* I want to be able to quickly list objects, phases and paths :ok:
  * Main menu to access lists :ok:
  * Upper navigation bar used everywhere :ok:
  * List object :ok:
  * List phase :ok:
  * List path :ok:
* I want to be able to quickly add objects, phases and paths :x:
  * Add possibility to save to DB a new object :x:
  * Add object :x:
  * Add phase :x:
  * Force format -> [{"object": "forest001", "position": "1"}, {...}] :x:
  * Delete buttons should show hand mouse pointer :x:
  * Get objects for phase from actual objects :x:
   * Find a way for this to scale :x:
  * Add Phase ID to paths_new :x:
  * Add path :x:
  * Get phases for path from actual phases :x:
   * Find a way for this to scale :x:
* I want to be able to quickly view objects, phases and paths :x:
  * View object :x:
  * View phase :x:
   * Get objects for phase from actual objects, link to see them :x:
  * View path :x:
    * Get phases for path from actual phases, link to see them :x:
* I want to be able to quickly edit objects, phases and paths :x:
  * Get ids too, for saving later :x:
  * Edit object :x:
  * Edit phase :x:
  * Get objects for phase from actual objects :x:
  * Edit path :x:
  * Get phases for path from actual phases :x:
* I want to be able to quickly insert objects, phases and paths :x:
  * Insert object :x:
  * Make it possible to upload images to shared/open volume on the frontend side (or manage images from here) :x:
  * Insert phase :x:
  * Get objects for phase from actual objects :x:
  * Insert path :x:
  * Get phases for path from actual phases :x:
* I want to be able to quickly delete objects, phases and paths :x:
  * Delete object :x:
  * Delete phase :x:
  * Show objects that still exist but are no longer in phase :x:
  * Delete path :x:
  * Show phases that still exist but are no longer in path :x:
* Build tests :x:
* Containerize app :ok:

# Possible trade-offs between Quality – Time – Cost
TBD

# Estimate project resources
TBD






