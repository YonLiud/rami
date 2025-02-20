from flask import send_file, redirect, url_for
from . import main_bp

@main_bp.route('/download_logs')
def download_logs():
    try:
        return send_file("time_log.csv", as_attachment=True)
    except Exception as e:
        error_message = str(e)
        return redirect(url_for('main.home', error_message=error_message))
