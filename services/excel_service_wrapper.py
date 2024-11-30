from services.excel_service import *

def get_visitors_inside():
    """
    Get all visitors inside
    """

    visitors_inside = get_all_visitors_inside()

    return visitors_inside

def mark_entry(visitor_id: str):
    """
    Mark the entry of the visitor
    """
    try:
        mark_row_in_all_sheets(visitor_id, "YES")
        save_cached_data()
    except Exception as e:
        print(e)
        return False
    
def mark_exit(visitor_id: str):
    """
    Mark the exit of the visitor
    """
    try:
        mark_row_in_all_sheets(visitor_id, "NO")
        save_cached_data()
    except Exception as e:
        print(e)
        return False
    
    return True

def get_by_flag(flag: str):
    """
    Get the visitor by id
    """
    visitors = None

    visitors = search_value_in_data(flag)

    return visitors