import pandas as pd
import os
import json

cachced_data = {}
file_last_modified = None

def load_and_cache_excel(file_path: str):
    """
    Load an excel file and cache it in memory
    """
    global cachced_data, file_last_modified

    try:
        current_modified = os.path.getmtime(file_path)

        if current_modified == file_last_modified:
            return cachced_data
        
        excel_data = pd.read_excel(file_path, sheet_name=None)

        cachced_data = {sheet_name: df for sheet_name, df in excel_data.items()}
        file_last_modified = current_modified

        print(f"Loaded excel file: {file_path}")
        return cachced_data
    
    except Exception as e:
        print(f"Failed to load excel file: {file_path}")
        print(e)
        return None
    
def get_cached_data():
    """
    Get the cached data
    """
    return cachced_data

def data_to_json():
    """
    Convert the cached data to json
    """
    global cachced_data

    json_data = {}

    if not cachced_data:
        print("No cached data found")
        return
    
    for sheet_name, df in cachced_data.items():
        rows = df.to_dict(orient="records")
        json_data[sheet_name] = rows

    return json_data