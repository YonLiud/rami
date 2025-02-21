from flask import render_template, request
from . import main_bp
from services.excel_service_wrapper import search_value_in_data

@main_bp.route('/search', methods=['POST'])
def search():
    search_id = request.form['search_id']
    search_results = search_value_in_data(search_id)
    return render_template('search_results.html', results=search_results)
