# GO Project Setter

This is a Project setter built with go for go but can be customised for other languages. Minimal Setter with git initialisations, go initialisations, and few necessary files like .gitignore,readme,etc..

## Table of Contents

1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Future Work](#future-work)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Harichandra-Prasath/GO-Setter.git
cd GO-Setter
```

2. Run the application

```bash 
make run
```

3. Install it if needed

```bash
go install
```

4. You can now simply call GO-Setter (Go's bin should be in path)

```bash
GO-Setter
```

## Configuration

Configurable Parameters

1. **Name : String**  
Name of the Project  
2. **Root : String**  (Default - /home/$USER)  
Path string of the root where the Project has to be set up  
3. **Origin : String**  
Remote repo git link that should be set as Origin  
4. **Branch: String** (Default - main)  
Name of the branch to move/rename from default `master` branch  
5. **IsIgnore : bool**(Default - True)  
Bool representing the need of minimal .gitignore while Setting up the project  
6. **IsMake : bool**  (Default - True)   
Bool representing the need of minimal Makefile while Setting up the project  
7. **IsReadme : bool**(Default - True)  
Bool representing the need of minimal Readme while Setting up the project  
8. **Mod : String**   (Default will be extracted from origin)  
Module Path Name for the Go project to initialise  

**Moreover, One can change the templates for go for personalised file creation in set-up**  

## Future-Work

1.  More configurable parameters and support for various languages  