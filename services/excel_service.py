import pandas as pd
import os
import json

cachced_data = {}
file_path = None

def load_and_cache_excel(file_path_i: str):

    global cachced_data, file_path

    if not os.path.exists(file_path_i):
        print(f"File not found: {file_path_i}")
        return False
    
    file_path = file_path_i

    cachced_data = {}

    try:

        excel_data = pd.read_excel(file_path, sheet_name=None)

        cachced_data = {sheet_name: df for sheet_name, df in excel_data.items()}

        print(f"Loaded excel file: {file_path}")
        return cachced_data
    
    except Exception as e:
        print(f"Failed to load excel file: {file_path}")
        print(e)
        return False
    
    return True
    
def get_cached_data():
    return cachced_data

def save_cached_data():
    global cachced_data, file_path

    if not cachced_data:
        raise ValueError("No cached data found")

    if not file_path:
        raise ValueError("No file path found")

    with pd.ExcelWriter(file_path) as writer:
        for sheet_name, df in cachced_data.items():
            df.to_excel(writer, sheet_name=sheet_name, index=False)

    print("Saved cached data to excel file")

def data_to_json():
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

def search_by_id(search_value):
    global cachced_data

    results = {}

    if not cachced_data:
        print("No cached data available.")
        return results

    search_value = str(search_value)

    try:
        for sheet_name, df in cachced_data.items():
            if "מספר תעודה" not in df.columns:
                print(f"Column 'מספר תעודה' not found in sheet '{sheet_name}'.")
                results[sheet_name] = []
                continue

            match = df[df["מספר תעודה"].astype(str) == search_value]
            if not match.empty:
                results[sheet_name] = match.to_dict(orient="records")

    except Exception as e:
        print("Failed to search by id.")
        print(e)

    return results

def get_all_visitors_inside():
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