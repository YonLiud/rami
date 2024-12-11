import sys
import os
from time import sleep
from datetime import datetime
from flask import Flask, render_template, request, redirect, url_for
from services.excel_service import load_and_cache_excel, save_cached_data
from services.excel_service_wrapper import *

app = Flask(__name__)
database_file = None

@app.route('/')
def home():
    visitors_inside = get_visitors_inside()
    error_message = request.args.get('error_message')
    last_updated = request.args.get('last_updated')
    if not last_updated:
        last_updated = datetime.now()
    return render_template('home.html', visitors=visitors_inside, error_message=error_message, last_updated=last_updated)

@app.route('/search', methods=['POST'])
def search():
    search_id = request.form['search_id']
    search_results = search_value_in_data(search_id)
    return render_template('search_results.html', results=search_results)

@app.route('/visitor/<visitor_id>')
def visitor_details(visitor_id):
    visitor = get_by_id(visitor_id)
    if not visitor:
        return redirect(url_for('home', error_message="Visitor not found"))
    return render_template('visitor_details.html', visitor=visitor, visitor_id=visitor_id)

@app.route('/mark_entry/<visitor_id>')
def route_mark_entry(visitor_id):
    try:
        mark_entry(visitor_id)
        save_cached_data()
    except Exception as e:
        error_message = str(e)
        return redirect(url_for('home', error_message=error_message))
    return redirect(url_for('home'))

@app.route('/mark_exit/<visitor_id>')
def route_mark_exit(visitor_id):
    try:
        mark_exit(visitor_id)
        save_cached_data()
    except Exception as e:
        error_message = str(e)
        if isinstance(e, PermissionError):
            error_message = "Please close the Excel file and try again."
        return redirect(url_for('home', error_message=error_message))
    return redirect(url_for('home'))

@app.route('/refresh/')
def route_refresh():
    load_and_cache_excel(database_file)
    return redirect(url_for('home', last_updated=datetime.now()))

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