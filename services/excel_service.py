import pandas as pd
import os
import json

cachced_data = {}
file_last_modified = None
file_path = None

def load_and_cache_excel(file_path_i: str):
    """
    Load an excel file and cache it in memory
    """
    global cachced_data, file_last_modified, file_path

    if not os.path.exists(file_path_i):
        print(f"File not found: {file_path_i}")
        return
    
    file_path = file_path_i

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

def save_cached_data():
    global cachced_data, file_path

    if not cachced_data:
        print("No cached data found")
        return
    
    if not file_path:
        print("No file path found")
        return
    
    try:
        with pd.ExcelWriter(file_path) as writer:
            for sheet_name, df in cachced_data.items():
                df.to_excel(writer, sheet_name=sheet_name, index=False)
        
        print("Saved cached data to excel file")
    except Exception as e:
        print("Failed to save cached data to excel file")
        print(e)

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

def search_value_in_data(search_value):
    """
    Search for a value in the cached data, handling type mismatches.

    Parameters:
        search_value: The value to search for in the cached data.

    Returns:
        dict: A dictionary where keys are sheet names and values are lists of matching rows.
    """
    global cachced_data

    results = {}

    if not cachced_data:
        print("No cached data available.")
        return results

    search_value = str(search_value)

    try:
        for sheet_name, df in cachced_data.items():
            df = df.map(str)
            search_result = df[df.eq(search_value).any(axis=1)]
            if not search_result.empty:
                results[sheet_name] = search_result.to_dict(orient="records")

    except Exception as e:
        print("Failed to search value in data.")
        print(e)

    return results

def get_all_visitors_inside():
    """
    Get all visitors inside

    Returns:
        dict: A dictionary where keys are sheet names and values are lists of rows where 'בפנים' = 'YES'.
    """
    global cachced_data

    results = {}

    if not cachced_data:
        print("No cached data available.")
        return results

    try:
        for sheet_name, df in cachced_data.items():
            if "בפנים" not in df.columns:
                print(f"Column 'בפנים' not found in sheet '{sheet_name}'.")
                results[sheet_name] = []
                continue

            inside = df[df["בפנים"] == "YES"]
            results[sheet_name] = inside.to_dict(orient="records")

    except Exception as e:
        print("Failed to get visitors inside.")
        print(e)

    return results

def mark_row_in_all_sheets(search_value, mark_value):
    """
    Mark a specific row in the 'בפנים' column across all sheets based on a unique value 
    in the 'מספר תעודה' column.

    Parameters:
        search_value: The unique value to search for in the 'מספר תעודה' column.
        mark_value: The value to set in the 'בפנים' column.

    Returns:
        dict: A dictionary with sheet names as keys and success status as values.
    """
    global cachced_data

    if not cachced_data:
        print("No cached data available.")
        return {}

    results = {}

    try:
        for sheet_name, df in cachced_data.items():
            if "מספר תעודה" not in df.columns or "בפנים" not in df.columns:
                print(f"One or both columns 'מספר תעודה', 'בפנים' not found in sheet '{sheet_name}'.")
                results[sheet_name] = False
                continue

            match = df[df["מספר תעודה"].astype(str) == str(search_value)]

            if not match.empty:
                df.loc[df["מספר תעודה"].astype(str) == str(search_value), "בפנים"] = mark_value

                cachced_data[sheet_name] = df

                results[sheet_name] = True
                print(f"Updated sheet '{sheet_name}': Set 'בפנים' to '{mark_value}' where 'מספר תעודה' = '{search_value}'.")
            else:
                results[sheet_name] = False

    except Exception as e:
        print("Failed to update row in data.")
        print(e)

    return results