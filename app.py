import sys
import os
from time import sleep
from flask import Flask
from routes import main_bp  # Import the single blueprint
from services.excel_service import load_and_cache_excel

app = Flask(__name__)

app.register_blueprint(main_bp, url_prefix='/')

database_file = None

if __name__ == '__main__':
    if not database_file:
        if not sys.argv[1:]:
            database_file = input("Please drag and drop the Excel file to the terminal and press enter: ").strip()
        else:
            database_file = sys.argv[1]
        if not database_file.endswith(".xlsx"):
            print("Please provide a valid Excel file.")
            sleep(5)
            sys.exit(1)

        if not os.path.isfile(database_file):
            print(f"File not found: {database_file}")
            sleep(5)
            sys.exit(1)

    if not load_and_cache_excel(database_file):
        print("Failed to load and cache the Excel file.")
        sleep(5)
        sys.exit(1)

    app.run(debug=False)