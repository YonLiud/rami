from services.excel_service import *
from datetime import datetime

def get_visitors_inside():
    visitors_inside = get_all_visitors_inside()

    return visitors_inside

def get_5_last_logs():
    logs = None

    try:
        if not os.path.exists("time_log.csv"):
            open("time_log.csv", "w").close()
            logs = []
        else:
            with open("time_log.csv", "r") as log_file:
                logs_raw = log_file.readlines()[-5:]
                logs = [log.strip().split(",") for log in logs_raw]
    except Exception as e:
        print(f"Error getting logs: {e}")
        raise e
    
    return logs

def log_action(visitor_id: str, action: str):
    try:
        with open("time_log.csv", "a") as log_file:
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            log_message = f"{visitor_id},{action},{timestamp}"
            log_file.write(log_message + "\n")

    except Exception as e:
        print(f"Error logging action: {e}")
        raise e

def mark_entry(visitor_id: str):
    try:
        mark_row_in_all_sheets(visitor_id, "YES")
        log_action(visitor_id, "Entry")
    except Exception as e:
        print(e)

def mark_exit(visitor_id: str):
    try:
        mark_row_in_all_sheets(visitor_id, "NO")
        log_action(visitor_id, "Exit")
    except Exception as e:
        print(e)

def get_by_flag(flag: str):
    visitors = None

    visitors = search_value_in_data(flag)

    return visitors

def get_by_id(visitor_id: str):
    visitor = None

    visitor = search_by_id(visitor_id)

    return visitor