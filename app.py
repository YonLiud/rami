from flask import Flask, render_template, request, redirect, url_for
from services.excel_service import load_and_cache_excel
from services.excel_service_wrapper import *

app = Flask(__name__)

@app.route('/')
def home():
    visitors_inside = get_visitors_inside()
    return render_template('home.html', visitors=visitors_inside)

@app.route('/search', methods=['POST'])
def search():
    search_id = request.form['search_id']
    search_results = search_value_in_data(search_id)
    return render_template('search_results.html', results=search_results)

@app.route('/mark_entry/<visitor_id>')
def route_mark_entry(visitor_id):
    mark_entry(visitor_id)
    save_cached_data()
    return redirect(url_for('home'))

@app.route('/mark_exit/<visitor_id>')
def route_mark_exit(visitor_id):
    mark_row_in_all_sheets(visitor_id, 'outside')
    save_cached_data()
    return redirect(url_for('home'))

@app.route('/refresh/')
def route_refresh():
    load_and_cache_excel("database.xlsx")
    return redirect(url_for('home'))

if __name__ == '__main__':
    load_and_cache_excel("database.xlsx")
    app.run(debug=True)