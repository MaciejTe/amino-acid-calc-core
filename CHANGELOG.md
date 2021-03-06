<!---
#######################################
## Amino-acid calc
##
## Format: markdown (md)
## Latest versions should be placed as first
##
## Notation: 00.01.02
##      - 00: stable released version
##      - 01: new features
##      - 02: bug fixes and small changes
##
## Updating schema (mandatory):
##      <empty_line>
##      <version> (dd/mm/rrrr)
##      ----------------------
##      * <item>
##      * <item>
##      <empty_line>
##
## Useful tutorial: https://en.support.wordpress.com/markdown-quick-reference/
##
#######################################
-->
0.0.2 (18.01.2020)
---------------------
    - Added "calculate" subcommand CLI section (with recipe command)
    - Improved ingredient search & details commands
    - Added Calculator interface with implementing structure: Recipe
    - Created tables: recipes, amino_acids, nutrition_facts and ingredients (not used yet)
    - Added "Branded": true option for USDA DB client for easier ingredients matching; solution will be improved in future
    
0.0.1 (27.12.2019)
---------------------
    - Added ingredients subcommand CLI section (with details and search commands)
    - Added panic wrapper
    - Added basic postgresql DB client for future use
    - Added base for future DB migrations (done with https://github.com/golang-migrate/migrate)
    - Added configuration file (for now the only source of configuration for application)

0.0.0 (26.12.2019)
---------------------
    - Initialised repository, added LICENSE and .gitignore
