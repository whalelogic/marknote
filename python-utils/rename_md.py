import os
import re

def kebab_case(name):
    # Split by extension
    base, ext = os.path.splitext(name)
    # Replace non-alphanumeric with hyphens
    s = re.sub(r'[^a-zA-Z0-9]', '-', base)
    # Handle CamelCase/PascalCase by inserting hyphens
    s = re.sub(r'([a-z0-9])([A-Z])', r'\1-\2', s)
    # Lowercase
    s = s.lower()
    # Replace multiple hyphens with one
    s = re.sub(r'-+', '-', s)
    # Strip hyphens from ends
    s = s.strip('-')
    return s + ext

def rename_files(root_dir):
    for root, dirs, files in os.walk(root_dir):
        for filename in files:
            if filename.endswith('.md'):
                old_path = os.path.join(root, filename)
                new_filename = kebab_case(filename)
                new_path = os.path.join(root, new_filename)
                
                if old_path != new_path:
                    print(f"Renaming: {old_path} -> {new_path}")
                    os.rename(old_path, new_path)

if __name__ == "__main__":
    rename_files('static')
