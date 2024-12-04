# Rami - Entry and Exit logger

Rami is a web-based application designed to log the entry and exit of visitors. It allows users to search for visitors, mark their entry and exit, and view the current visitors inside. The application uses an Excel file to store and manage the visitor data.

## Features
- Search for visitors by ID
- Mark the entry and exit of visitors
- View a list of current visitors inside
- Data stored and managed in an Excel file

## Build instructions

### Prerequisites
- Python 3.12 or higher
- Pip

### Installation
1. Clone the repository
2. Run `pip install -r requirements.txt` to install the required packages
3. Run `pyinstaller ramiexe.spec` to build the executable
   The executable will be in the `dist` folder


>[!INFO]
> This will generate the executable based on the configuration specified in the ``ramiexe.spec`` file. The output will be placed in the ``dist`` directory.


## TODO

- Add logging functionality to log the entry and exit of visitors and ability to export the logs