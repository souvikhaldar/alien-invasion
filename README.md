# Alien-Invasion
## Situation
Alien-Invasion is a simulation of hypothetical situation (or maybe near future :P), where `N` number of aliens are unleashed on us. They kill each other if present in same city and destroy the city as well during the brutal fight. The final state of our planet is the map as stored in output file as defined in `config.json`. You can read the entire situation [here](https://github.com/souvikhaldar/alien-invasion/blob/main/problem_statement.md)  

## Video Demo
[Video demonstration of the simulation](https://youtu.be/vyg6K09PJzI)

## Steps to run the simulation
1. [Install golang](https://go.dev/doc/install)  
2. Clone this repository. `git clone git@github.com:souvikhaldar/alien-invasion.git`   
3. Get into the repository. `cd alien-invasion/`.  
4. Install the dependencies using `go mod tidy`.   
5. Run the simulation 5 aliens or any number you want `go run cmd/ai/main.go -N=5`  
6. By default it uses the credentials present in `config.json` present in the root. You can provide other configuration file by using the flag `-conf=<file-name`.  
7. The output file is `output_map.txt` in `testfiles` dir if not specified otherwise.  
