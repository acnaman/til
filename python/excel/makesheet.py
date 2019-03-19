import openpyxl
wb = openpyxl.Workbook()

sheet = wb.active

sheet.title = 'test'

sheet['A1'] = '''長い\r改行のある\n文章はどうなる？
'''

wb.save('out/example_out.xlsx')

