from gedcom.element.individual import IndividualElement
from gedcom.element.file import FileElement
from gedcom.element.element import Element
from gedcom.parser import Parser
import os

# Path to your `.ged` file
file_path = r'FamilieNeu.ged'
base_path = r'C:\Users\Luis\Ahnenblatt'

def parse_gedcom():
    # Initialize the parser
    gedcom_parser = Parser()

    # Parse your file
    gedcom_parser.parse_file(file_path, False)

    elements = gedcom_parser.get_element_list()

    amount = 0

    content = ""
    for element in elements:
        if isinstance(element, FileElement):
            val = element.get_value()
            if ".ged" in val:
                continue

            content += val
            content += "\n"
            amount += 1

    with open("paths.txt", "w+") as file:
        file.write(content)


def parse_txt():
    with open(file_path, "r", encoding="utf-8") as file:
        txt = file.readlines()

    paths = ""


    for idx, line in enumerate(txt):
        if ".ged" in line:
            continue
        if "FILE " in line:
            path = line.split("FILE")[-1].strip()
            paths += f"{path}\n"

            filename = path.split('\\')[-1]

            next_line = txt[idx+1]
            if "ALTPATH " in next_line:
                line_split = next_line.split("ALTPATH")
                altpath = line_split[-1].strip()
                if altpath.startswith(".\\"):
                    altpath = altpath[2:]
                paths += f"{os.path.join(base_path, altpath, filename)}\n"
        
    with open("paths.txt", "w+") as file:
        file.write(paths)

if __name__ == "__main__":
    parse_txt()