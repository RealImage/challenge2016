Distributor Permission System
Overview

The Distributor Permission System is an open-source Go project designed to manage distributor permissions and distribution rights based on geographical regions (city, state, country). This project is a valuable addition to your portfolio, demonstrating your coding skills and problem-solving abilities.
Features
1. Permission Management

Easily configure permissions for distributors:

    Include: Specify regions (city, state, country) where a distributor is authorized to distribute.
    Exclude: Define regions to exclude from a distributor's distribution rights.

2. Hierarchical Structure

Create a hierarchical structure of distributors:

    Add sub-distributors to existing distributors, allowing for complex distribution networks.

3. Location-Based Authorization

The system checks authorization based on location data, providing distributors with precise distribution rights.
4. Interactive CLI

The interactive command-line interface (CLI) makes it simple to:

    Add new distributors.
    Configure permissions.
    Check distributions.
    Add sub-distributors.

Getting Started

Follow these steps to run the program and start using it:

    Ensure you have Go installed on your system.

    Clone the repository:

    bash

git clone https://github.com/RohithER12/challenge_Qube_cinimas.git
cd distributor-permission-system

Compile and run the program:

bash

    go run main.go

    Use the interactive menu to perform various actions and explore the features.

Usage
Adding a New Distributor

    Easily add new distributors by providing their names.
    Prevent duplicate entries with a built-in check.

Configuring Permissions

    Set precise permissions by including or excluding specific regions (city, state, country).
    The system checks parent distributor permissions for sub-distributors.

Checking Distributions

    Verify if a distributor is authorized to distribute in a particular region.
    Get detailed distribution information based on location.

Adding Sub-Distributors

    Expand your distribution network by adding sub-distributors to existing distributors.