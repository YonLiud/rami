from flask import render_template, redirect, url_for
from . import main_bp
from services.excel_service_wrapper import get_by_id, mark_entry, mark_exit
from services.excel_service import save_cached_data

@main_bp.route('/visitor/<visitor_id>')
def visitor_details(visitor_id):
    visitor = get_by_id(visitor_id)
    if not visitor:
        return redirect(url_for('main.home', error_message="Visitor not found"))
    return render_template('visitor_details.html', visitor=visitor, visitor_id=visitor_id)

@main_bp.route('/mark_entry/<visitor_id>')
def route_mark_entry(visitor_id):
    try:
        mark_entry(visitor_id)
        save_cached_data()
    except Exception as e:
        error_message = str(e)
        return redirect(url_for('main.home', error_message=error_message))
    return redirect(url_for('main.home'))

@main_bp.route('/mark_exit/<visitor_id>')
def route_mark_exit(visitor_id):
    try:
        mark_exit(visitor_id)
        save_cached_data()
    except Exception as e:
        error_message = "Please close the Excel file and try again." if isinstance(e, PermissionError) else str(e)
        return redirect(url_for('main.home', error_message=error_message))
    return redirect(url_for('main.home'))
