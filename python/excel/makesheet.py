import openpyxl
wb = openpyxl.Workbook()

sheetnames = wb.get_sheet_names()

print(sheetnames)
