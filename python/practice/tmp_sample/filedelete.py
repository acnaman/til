# !python3
# カレントディレクトリ以下の拡張子が.py以外のファイルを全て削除する

import send2trash, os, re

ignore_pattern = '.py'

for folßdername, subfolders, filenames in os.walk('.'):
    for filename in filenames:
        if filename.endswith(ignore_pattern) :
            print("{}は除外しないファイル".format(filename))
        
        else:
            print ("{}は除外する！！！！！".format(filename))





