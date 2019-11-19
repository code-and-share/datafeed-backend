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
  * Add possibility to get file from outside the container :ok:
  * Load images into backend :ok:
  * Change frontend to get images from this backend :ok:
* I want to be able to quickly list objects, phases and paths :ok:
  * Main menu to access lists :ok:
  * Upper navigation bar used everywhere :ok:
  * List object :ok:
  * List phase :ok:
  * List path :ok:
* I want to be able to quickly add objects, phases and paths :x:
  * Add possibility to save to DB a new object :ok:
  * Add object :ok:
  * Add phase :ok:
  * Force format -> [{"object": "forest001", "position": "1"}, {...}] :ok
  * Delete buttons should show hand mouse pointer :ok:
  * Get objects for phase from actual objects :ok:
   * Find a way for this to scale :x:
  * Add Phase ID to paths_new :ok:
  * Add path :ok:
  * Get phases for path from actual phases :ok:
   * Find a way for this to scale :x:
* I want to be able to quickly view objects, phases and paths :x:
  * View object :ok:
  * View phase :ok:
   * Get objects for phase from actual objects, link to see them :x:
  * View path :ok:
    * Get phases for path from actual phases, link to see them :x:
* I want to be able to quickly edit objects, phases and paths :x:
  * Get ids too, for saving later :ok:
  * Edit object :ok:
  * Edit phase :x:
  * Get objects for phase from actual objects :x:
  * Edit path :ok:
  * Get phases for path from actual phases :x:
* I want to be able to quickly insert objects, phases and paths :x:
  * Insert object :ok:
  * Make it possible to upload images to shared/open volume on the frontend side (or manage images from here) :x:
  * Insert phase :ok:
  * Get objects for phase from actual objects :x:
  * Insert path :ok:
  * Get phases for path from actual phases :x:
* I want to be able to quickly delete objects, phases and paths :x:
  * Delete object :ok:
  * Delete phase :ok:
  * Show objects that still exist but are no longer in phase :x:
  * Delete path :ok:
  * Show phases that still exist but are no longer in path :x:
* Build tests :x:
* Containerize app :x:

# Possible trade-offs between Quality – Time – Cost
TBD

# Estimate project resources
TBD






