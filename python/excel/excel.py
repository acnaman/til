import openpyxl

wb = openpyxl.load_workbook('example.xlsx')

sheet = wb.get_sheet_by_name('Sheet1')

print(sheet['A1'].value)
print(sheet['B1'].value)
print(sheet['C1'].value)



