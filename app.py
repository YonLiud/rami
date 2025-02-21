import sys
import os
from flask import Flask
from routes import main_bp
from services.excel_service import load_and_cache_excel

app = Flask(__name__)
app.config['DATABASE_FILE'] = None

@app.route('/favicon.ico')
def favicon():
    return Flask.send_from_directory(os.path.join(app.root_path, 'static'), 'favicon.ico', mimetype='image/vnd.microsoft.icon')

app.register_blueprint(main_bp, url_prefix='/')
if __name__ == '__main__':
    if not app.config['DATABASE_FILE']:
        database_file = sys.argv[1] if sys.argv[1:] else input("Enter the full path to the Excel file and press Enter: ").strip()
        if not database_file.endswith(".xlsx") or not os.path.isfile(database_file):
            print("Invalid file. Exiting.")
            sys.exit(1)
        app.config['DATABASE_FILE'] = database_file

    if not load_and_cache_excel(app.config['DATABASE_FILE']):
        print("Failed to load the Excel file.")
        sys.exit(1)

    app.run(debug=False)