from flask import render_template, request, redirect, url_for, current_app
from datetime import datetime
from . import main_bp
from services.excel_service import load_and_cache_excel
from services.excel_service_wrapper import get_visitors_inside, get_5_last_logs

@main_bp.route('/')
def home():
    visitors_inside = get_visitors_inside()
    error_message = request.args.get('error_message')
    last_updated = request.args.get('last_updated')
    logs = get_5_last_logs()
    if not last_updated:
        last_updated = datetime.now()
    return render_template('home.html', visitors=visitors_inside, error_message=error_message, last_updated=last_updated, logs=logs)

@main_bp.route('/refresh/')
def route_refresh():
    load_and_cache_excel(current_app.config['DATABASE_FILE'])
    return redirect(url_for('main.home', last_updated=datetime.now()))
