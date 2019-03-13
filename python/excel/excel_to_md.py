import openpyxl

wb = openpyxl.load_workbook('example.xlsx')

sheet = wb.get_sheet_by_name('Sheet1')

md = ''

for n in range(sheet.max_row):
    row = list(sheet.rows)[n]

    for cell in row:
        md += '|'
        md += str(cell.value)
    

    md += '|\n'
            
    if n == 0:
        for cell in range(len(row)):
            md += '|---'

        md += '|\n'


print (md)

with open('example.md', encoding='utf-8', mode='w') as f:
    f.write(md)


