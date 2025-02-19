# Rami - Entry and Exit logger

Rami is a web-based application designed to log the entry and exit of visitors. It allows users to search for visitors, mark their entry and exit, and view the current visitors inside. The application uses an Excel file to store and manage the visitor data.

## Features
- Search for visitors by ID
- Mark the entry and exit of visitors
- View a list of current visitors inside
- Data stored and managed in an Excel file
- Log and Review all entries and exits with the ability of exporting the logs file

## Build instructions

### Prerequisites
- Python 3.12 or higher
- Pip

### Installation
1. Clone the repository
2. Run `pip install -r requirements.txt` to install the required packages
3. Run `pyinstaller ramiexe.spec` to build the executable
   The executable will be in the `dist` folder


> [!NOTE]
> This will generate the executable based on the configuration specified in the ``ramiexe.spec`` file. The output will be placed in the ``dist`` directory.


> [!WARN]
> Any excel file must have these 3 rows: ``הדועת רפסמ``, ``אלמ םש`` and ``םינפב``


## Routes Map

### 1. `/`
- **Method:** GET
- **Description:** Displays the home page with:
  - Visitors currently inside
  - Logs of the last actions
  - Error messages, if any

### 2. `/download_logs`
- **Method:** GET
- **Description:** Downloads the time logs as a CSV file (`time_log.csv`).

### 3. `/search`
- **Method:** POST
- **Description:** Searches for a visitor based on the `search_id` submitted via a form.

### 4. `/visitor/<visitor_id>`
- **Method:** GET
- **Description:** Displays the details of a specific visitor identified by `visitor_id`.

### 5. `/mark_entry/<visitor_id>`
- **Method:** GET
- **Description:** Marks the entry of a visitor with the given `visitor_id`.

### 6. `/mark_exit/<visitor_id>`
- **Method:** GET
- **Description:** Marks the exit of a visitor with the given `visitor_id`.

### 7. `/refresh/`
- **Method:** GET
- **Description:** Refreshes the data by reloading and caching the Excel file.
