from services.excel_service import *

def get_visitors_inside():
    visitors_inside = get_all_visitors_inside()

    return visitors_inside

def mark_entry(visitor_id: str):
    try:
        mark_row_in_all_sheets(visitor_id, "YES")
    except Exception as e:
        print(e)
    
def mark_exit(visitor_id: str):
    try:
        mark_row_in_all_sheets(visitor_id, "NO")
    except Exception as e:
        print(e)

def get_by_flag(flag: str):
    visitors = None

    visitors = search_value_in_data(flag)

    return visitors